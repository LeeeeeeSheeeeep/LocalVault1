package tests

import (
	"bytes"
	"crypto/rand"
	"os"
	"path/filepath"
	"testing"

	"localvault/security"
)

func TestQRBackupAndRestore(t *testing.T) {
	// Generate random payload
	payload := make([]byte, 2048)
	if _, err := rand.Read(payload); err != nil {
		t.Fatalf("failed to generate random payload: %v", err)
	}

	bm := security.NewQRBackupManager(300)

	// Create backup segments
	segments, err := bm.CreateBackup(payload)
	if err != nil {
		t.Fatalf("CreateBackup failed: %v", err)
	}

	if len(segments) == 0 {
		t.Fatal("no segments created")
	}

	// Verify segment indices and totals
	expectedTotal := segments[0].Total
	for i, seg := range segments {
		if seg.Index != i+1 {
			t.Errorf("segment index mismatch: expected %d, got %d", i+1, seg.Index)
		}
		if seg.Total != expectedTotal {
			t.Errorf("segment total mismatch: expected %d, got %d", expectedTotal, seg.Total)
		}
	}

	// Restore backup
	restored, err := bm.RestoreBackup(segments)
	if err != nil {
		t.Fatalf("RestoreBackup failed: %v", err)
	}

	if !bytes.Equal(payload, restored) {
		t.Fatal("restored payload does not match original")
	}
}

func TestShamirSecretSharing(t *testing.T) {
	// 32-byte AES key
	secret := make([]byte, 32)
	if _, err := rand.Read(secret); err != nil {
		t.Fatalf("failed to generate random secret: %v", err)
	}

	n := 5
	k := 3

	// Split secret
	shares, err := security.SplitSecret(secret, n, k)
	if err != nil {
		t.Fatalf("SplitSecret failed: %v", err)
	}

	if len(shares) != n {
		t.Fatalf("expected %d shares, got %d", n, len(shares))
	}

	// 1. Reconstruct using shares [1, 2, 3] (Exact threshold K)
	subset1 := []security.ShamirShare{shares[0], shares[1], shares[2]}
	recovered1, err := security.ReconstructSecret(subset1)
	if err != nil {
		t.Fatalf("reconstruction with 3 shares failed: %v", err)
	}
	if !bytes.Equal(secret, recovered1) {
		t.Error("reconstructed secret does not match original (subset 1)")
	}

	// 2. Reconstruct using shares [2, 4, 5] (Different exact threshold K)
	subset2 := []security.ShamirShare{shares[1], shares[3], shares[4]}
	recovered2, err := security.ReconstructSecret(subset2)
	if err != nil {
		t.Fatalf("reconstruction with 3 shares failed: %v", err)
	}
	if !bytes.Equal(secret, recovered2) {
		t.Error("reconstructed secret does not match original (subset 2)")
	}

	// 3. Reconstruct using shares [1, 2, 3, 4] (More than threshold K)
	subset3 := []security.ShamirShare{shares[0], shares[1], shares[2], shares[3]}
	recovered3, err := security.ReconstructSecret(subset3)
	if err != nil {
		t.Fatalf("reconstruction with 4 shares failed: %v", err)
	}
	if !bytes.Equal(secret, recovered3) {
		t.Error("reconstructed secret does not match original (subset 3)")
	}

	// 4. Reconstruct using [1, 2] (Under threshold K)
	subset4 := []security.ShamirShare{shares[0], shares[1]}
	recovered4, err := security.ReconstructSecret(subset4)
	if err != nil {
		t.Fatalf("reconstruction with 2 shares returned error: %v", err)
	}
	if bytes.Equal(secret, recovered4) {
		t.Error("reconstruction with insufficient shares succeeded (it should fail/mismatch mathematically)")
	}
}

func TestDuressPlausibleDeniability(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "localvault_auth_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	authPath := filepath.Join(tempDir, "auth.json")
	dm := security.NewDuressManager(authPath)

	// Update credentials
	err = dm.UpdateCredentials("correct_master_123", "correct_duress_456")
	if err != nil {
		t.Fatalf("failed to update credentials: %v", err)
	}

	// 1. Correct master password
	isValid, isDuress := dm.Authenticate("correct_master_123")
	if !isValid || isDuress {
		t.Errorf("expected master authentic to succeed and not trigger duress: got valid=%v, duress=%v", isValid, isDuress)
	}

	// 2. Correct duress password
	isValid, isDuress = dm.Authenticate("correct_duress_456")
	if !isValid || !isDuress {
		t.Errorf("expected duress authentic to succeed and trigger duress: got valid=%v, duress=%v", isValid, isDuress)
	}

	// 3. Incorrect password
	isValid, isDuress = dm.Authenticate("wrong_password")
	if isValid {
		t.Errorf("expected authentication to fail for wrong password")
	}
}

func TestStegoEncodeDecode(t *testing.T) {
	stego := security.NewStego()

	// Generate cover image
	mockPNG, err := stego.GenerateMockCoverImage()
	if err != nil {
		t.Fatalf("failed to generate mock cover: %v", err)
	}

	// Payload data to hide
	secretData := []byte("LocalVault_RecoverySecret_Key_Shamir_2025_StegoTest")

	// Encode secret data into image
	var buf bytes.Buffer
	err = stego.Encode(bytes.NewReader(mockPNG), secretData, &buf)
	if err != nil {
		t.Fatalf("LSB Stego encoding failed: %v", err)
	}

	// Decode secret data from image
	decoded, err := stego.Decode(&buf)
	if err != nil {
		t.Fatalf("LSB Stego decoding failed: %v", err)
	}

	if !bytes.Equal(secretData, decoded) {
		t.Fatalf("decoded secret mismatch: expected %q, got %q", string(secretData), string(decoded))
	}
}

func TestKeystrokeDynamics(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "localvault_keystroke_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	keystrokePath := filepath.Join(tempDir, "keystroke.json")
	kp := security.NewKeystrokeProfiler(keystrokePath)

	// Simulate 3 training keystroke timing logs (similar cadence)
	trainingSamples := []security.TimingEntry{
		{
			DwellTimes:  []float64{80, 85, 90, 75, 80},
			FlightTimes: []float64{120, 110, 130, 115},
		},
		{
			DwellTimes:  []float64{82, 83, 89, 78, 81},
			FlightTimes: []float64{118, 112, 128, 118},
		},
		{
			DwellTimes:  []float64{78, 87, 91, 74, 79},
			FlightTimes: []float64{122, 108, 132, 112},
		},
	}

	err = kp.Enroll(trainingSamples)
	if err != nil {
		t.Fatalf("Enroll failed: %v", err)
	}

	// 1. Valid test sample (very close to training baseline)
	validSample := security.TimingEntry{
		DwellTimes:  []float64{81, 84, 90, 76, 80},
		FlightTimes: []float64{120, 111, 129, 116},
	}
	score, matched := kp.Verify(validSample)
	if !matched {
		t.Errorf("expected close timing sample to match baseline profile: score=%f", score)
	}

	// 2. Anomaly test sample (extremely slow, robotic, or typing jitter)
	anomalySample := security.TimingEntry{
		DwellTimes:  []float64{250, 300, 280, 220, 290},
		FlightTimes: []float64{500, 450, 600, 480},
	}
	score, matched = kp.Verify(anomalySample)
	if matched {
		t.Errorf("expected anomalous timing profile to fail baseline matching: score=%f", score)
	}
}
