package security

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

// DeadManConfig stores configuration for the Switch
type DeadManConfig struct {
	ThresholdDays int      `json:"threshold_days"`
	TotalShares   int      `json:"total_shares"`
	MinShares     int      `json:"min_shares"`
	Recipients    []string `json:"recipients"`
	LastHeartbeat time.Time `json:"last_heartbeat"`
}

type DeadManSwitch struct {
	mu         sync.RWMutex
	configPath string
	config     DeadManConfig
}

func NewDeadManSwitch(configPath string) *DeadManSwitch {
	s := &DeadManSwitch{
		configPath: configPath,
		config: DeadManConfig{
			ThresholdDays: 30,
			TotalShares:   5,
			MinShares:     3,
			Recipients:    []string{"partner@example.com", "trustee@example.com"},
			LastHeartbeat: time.Now(),
		},
	}
	s.loadConfig()
	return s
}

func (s *DeadManSwitch) loadConfig() {
	data, err := os.ReadFile(s.configPath)
	if err == nil {
		var cfg DeadManConfig
		if json.Unmarshal(data, &cfg) == nil {
			s.config = cfg
		}
	}
}

func (s *DeadManSwitch) SaveConfig() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := json.MarshalIndent(s.config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.configPath, data, 0600)
}

// RecordHeartbeat logs a user activity signal to prevent switch triggering
func (s *DeadManSwitch) RecordHeartbeat() {
	s.mu.Lock()
	s.config.LastHeartbeat = time.Now()
	s.mu.Unlock()
	s.SaveConfig()
}

// GetStatus returns the current dead man switch state
func (s *DeadManSwitch) GetStatus() (time.Time, int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	timeSinceLast := time.Since(s.config.LastHeartbeat)
	threshold := time.Duration(s.config.ThresholdDays) * 24 * time.Hour
	triggered := timeSinceLast > threshold

	return s.config.LastHeartbeat, s.config.ThresholdDays, triggered
}

// UpdateSettings updates configuration parameters
func (s *DeadManSwitch) UpdateSettings(thresholdDays, totalShares, minShares int, recipients []string) error {
	if thresholdDays <= 0 {
		return errors.New("threshold days must be positive")
	}
	if minShares < 2 {
		return errors.New("minimum shares to reconstruct must be at least 2")
	}
	if totalShares < minShares {
		return errors.New("total shares must be greater than or equal to minimum shares")
	}

	s.mu.Lock()
	s.config.ThresholdDays = thresholdDays
	s.config.TotalShares = totalShares
	s.config.MinShares = minShares
	s.config.Recipients = recipients
	s.mu.Unlock()

	return s.SaveConfig()
}

// ExecuteKeySplit splits the master encryption key based on SSS parameters
func (s *DeadManSwitch) ExecuteKeySplit(masterKey []byte) ([]ShamirShare, error) {
	s.mu.RLock()
	total := s.config.TotalShares
	min := s.config.MinShares
	s.mu.RUnlock()

	return SplitSecret(masterKey, total, min)
}

// VerifySecretReconstruction tries to rebuild the secret from shares to confirm math integrity
func (s *DeadManSwitch) VerifySecretReconstruction(shares []ShamirShare, original []byte) error {
	recovered, err := ReconstructSecret(shares)
	if err != nil {
		return fmt.Errorf("reconstruction failed: %w", err)
	}

	if len(recovered) != len(original) {
		return fmt.Errorf("size mismatch: recovered %d, original %d", len(recovered), len(original))
	}

	for i := range original {
		if recovered[i] != original[i] {
			return fmt.Errorf("content mismatch at byte index %d", i)
		}
	}

	return nil
}
