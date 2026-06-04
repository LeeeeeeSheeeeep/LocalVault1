package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"

	"github.com/cloudflare/circl/kem"
	"github.com/cloudflare/circl/kem/kyber/kyber1024"
)

// PQCManager handles Post-Quantum Key Encapsulation Mechanisms (KEM)
type PQCManager struct {
	scheme kem.Scheme
	pk     kem.PublicKey
	sk     kem.PrivateKey
}

// NewPQCManager generates a new Kyber1024 keypair for the vault
func NewPQCManager() (*PQCManager, error) {
	scheme := kyber1024.Scheme()
	pk, sk, err := scheme.GenerateKeyPair()
	if err != nil {
		return nil, fmt.Errorf("failed to generate Kyber keypair: %w", err)
	}

	return &PQCManager{
		scheme: scheme,
		pk:     pk,
		sk:     sk,
	}, nil
}

// SealVaultKey takes a standard 32-byte AES master key and encapsulates it
// using the Post-Quantum Kyber algorithm, returning the KEM ciphertext and the wrapped key bytes.
func (p *PQCManager) SealVaultKey(masterKey []byte) (kemCiphertext []byte, wrappedKey []byte, err error) {
	if len(masterKey) != 32 {
		return nil, nil, errors.New("master key must be 32 bytes for AES-256")
	}

	// Kyber encapsulation generates a shared secret (ss) and the ciphertext (ct)
	ct, ss, err := p.scheme.Encapsulate(p.pk)
	if err != nil {
		return nil, nil, err
	}

	// In a real hybrid system, we use the shared secret (ss) to encrypt the masterKey
	// using AES-GCM. We use `ss` as the wrapping key.
	block, err := aes.NewCipher(ss[:32]) // Kyber shared secret is 32 bytes
	if err != nil {
		return nil, nil, err
	}
	
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	wrapped := gcm.Seal(nonce, nonce, masterKey, nil)

	return ct, wrapped, nil
}

// OpenVaultKey takes the KEM ciphertext and wrapped master key, decapsulates the shared secret,
// and unwraps the master key.
func (p *PQCManager) OpenVaultKey(kemCiphertext []byte, wrappedKey []byte) ([]byte, error) {
	// Kyber decapsulation recovers the shared secret (ss)
	ss, err := p.scheme.Decapsulate(p.sk, kemCiphertext)
	if err != nil {
		return nil, fmt.Errorf("PQC decapsulation failed: %w", err)
	}

	// Unwrap the master key using the recovered shared secret
	block, err := aes.NewCipher(ss[:32])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(wrappedKey) < nonceSize {
		return nil, errors.New("wrapped key too short")
	}

	nonce, ciphertext := wrappedKey[:nonceSize], wrappedKey[nonceSize:]
	masterKey, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to unwrap master key: %w", err)
	}

	return masterKey, nil
}
