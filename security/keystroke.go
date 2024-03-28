package security

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

// TimingEntry represents keystroke timing metrics from a single password entry session
type TimingEntry struct {
	DwellTimes  []float64 `json:"dwell_times"`  // Time in milliseconds each key was held down
	FlightTimes []float64 `json:"flight_times"` // Time in milliseconds between keys (Key N release to Key N+1 press)
}

// TypingProfile represents the enrolled baseline profile with statistical averages and standard deviations
type TypingProfile struct {
	DwellMeans  []float64 `json:"dwell_means"`
	DwellDevs   []float64 `json:"dwell_devs"`
	FlightMeans []float64 `json:"flight_means"`
	FlightDevs  []float64 `json:"flight_devs"`
	SampleCount int       `json:"sample_count"`
}

type KeystrokeProfiler struct {
	profilePath string
	profile     *TypingProfile
}

func NewKeystrokeProfiler(profilePath string) *KeystrokeProfiler {
	kp := &KeystrokeProfiler{profilePath: profilePath}
	kp.loadProfile()
	return kp
}

func (kp *KeystrokeProfiler) loadProfile() {
	file, err := os.ReadFile(kp.profilePath)
	if err != nil {
		// Use a mock pre-enrolled profile representing master password "masterpassword" typing rhythms
		// configured for standard developer typing speed (dwell ~80ms, flight ~120ms)
		kp.profile = &TypingProfile{
			DwellMeans:  []float64{80, 85, 90, 75, 80, 85, 90, 80, 85, 90, 75, 80, 85, 90},
			DwellDevs:   []float64{15, 12, 18, 14, 15, 12, 18, 15, 12, 18, 14, 15, 12, 18},
			FlightMeans: []float64{120, 110, 130, 115, 120, 110, 130, 120, 110, 130, 115, 120, 110},
			FlightDevs:  []float64{25, 20, 30, 22, 25, 20, 30, 25, 20, 30, 22, 25, 20},
			SampleCount: 5,
		}
		kp.saveProfile()
		return
	}
	var prof TypingProfile
	if err := json.Unmarshal(file, &prof); err == nil {
		kp.profile = &prof
	}
}

func (kp *KeystrokeProfiler) saveProfile() error {
	data, err := json.MarshalIndent(kp.profile, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(kp.profilePath, data, 0644)
}

// EnrollTrain builds a baseline profile from multiple training attempts
func (kp *KeystrokeProfiler) Enroll(entries []TimingEntry) error {
	if len(entries) < 3 {
		return fmt.Errorf("need at least 3 typing attempts to enroll profile")
	}

	lenDwell := len(entries[0].DwellTimes)
	lenFlight := len(entries[0].FlightTimes)

	// Validate dimension consistency
	for idx, entry := range entries {
		if len(entry.DwellTimes) != lenDwell || len(entry.FlightTimes) != lenFlight {
			return fmt.Errorf("inconsistent keystroke dimensions in training sample index %d", idx)
		}
	}

	profile := &TypingProfile{
		DwellMeans:  make([]float64, lenDwell),
		DwellDevs:   make([]float64, lenDwell),
		FlightMeans: make([]float64, lenFlight),
		FlightDevs:  make([]float64, lenFlight),
		SampleCount: len(entries),
	}

	// 1. Calculate means
	for _, entry := range entries {
		for i := 0; i < lenDwell; i++ {
			profile.DwellMeans[i] += entry.DwellTimes[i]
		}
		for i := 0; i < lenFlight; i++ {
			profile.FlightMeans[i] += entry.FlightTimes[i]
		}
	}
	for i := 0; i < lenDwell; i++ {
		profile.DwellMeans[i] /= float64(len(entries))
	}
	for i := 0; i < lenFlight; i++ {
		profile.FlightMeans[i] /= float64(len(entries))
	}

	// 2. Calculate standard deviations
	for _, entry := range entries {
		for i := 0; i < lenDwell; i++ {
			diff := entry.DwellTimes[i] - profile.DwellMeans[i]
			profile.DwellDevs[i] += diff * diff
		}
		for i := 0; i < lenFlight; i++ {
			diff := entry.FlightTimes[i] - profile.FlightMeans[i]
			profile.FlightDevs[i] += diff * diff
		}
	}

	minDev := 8.0 // minimum standard deviation floor in ms to prevent division by zero
	for i := 0; i < lenDwell; i++ {
		profile.DwellDevs[i] = math.Sqrt(profile.DwellDevs[i] / float64(len(entries)))
		if profile.DwellDevs[i] < minDev {
			profile.DwellDevs[i] = minDev
		}
	}
	for i := 0; i < lenFlight; i++ {
		profile.FlightDevs[i] = math.Sqrt(profile.FlightDevs[i] / float64(len(entries)))
		if profile.FlightDevs[i] < minDev {
			profile.FlightDevs[i] = minDev
		}
	}

	kp.profile = profile
	return kp.saveProfile()
}

// Verify evaluates an attempt timing pattern against the profile baseline.
// Returns a distance score. A score < 2.0-2.5 generally represents a positive match.
func (kp *KeystrokeProfiler) Verify(entry TimingEntry) (float64, bool) {
	if kp.profile == nil || kp.profile.SampleCount == 0 {
		// No baseline registered; allow by default but warn
		return 0.0, true
	}

	lenDwell := len(entry.DwellTimes)
	lenFlight := len(entry.FlightTimes)

	// Inconsistent typing lengths (e.g. typing different password)
	if lenDwell != len(kp.profile.DwellMeans) || lenFlight != len(kp.profile.FlightMeans) {
		return 999.9, false
	}

	var sumZScore float64
	totalMeasurements := float64(lenDwell + lenFlight)

	// Calculate Mahalanobis-like Z-Score distances
	for i := 0; i < lenDwell; i++ {
		z := math.Abs(entry.DwellTimes[i] - kp.profile.DwellMeans[i]) / kp.profile.DwellDevs[i]
		sumZScore += z
	}

	for i := 0; i < lenFlight; i++ {
		z := math.Abs(entry.FlightTimes[i] - kp.profile.FlightMeans[i]) / kp.profile.FlightDevs[i]
		sumZScore += z
	}

	averageZScore := sumZScore / totalMeasurements

	// Threshold check: Accept if average timing deviation is within 2.5 standard deviations
	threshold := 2.5
	return averageZScore, averageZScore <= threshold
}
