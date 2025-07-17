package security

import (
	"crypto/rand"
	"errors"
	"fmt"
)

// ShamirShare represents a single share of a split secret key
type ShamirShare struct {
	Index byte   `json:"index"` // X-coordinate
	Value []byte `json:"value"` // Y-coordinate bytes
}

// Galois Field GF(2^8) tables
var (
	gfExp [512]byte
	gfLog [256]byte
)

func init() {
	// Initialize exp and log tables for GF(2^8) multiplication
	// Irreducible polynomial: x^8 + x^4 + x^3 + x^2 + 1 (0x11d)
	var x byte = 1
	for i := 0; i < 255; i++ {
		gfExp[i] = x
		gfExp[i+255] = x
		gfLog[x] = byte(i)
		
		// Multiply x by 2 in GF(2^8)
		next := int(x) << 1
		if next&0x100 != 0 {
			next ^= 0x11d
		}
		x = byte(next)
	}
}

// add adds two elements in GF(2^8) (XOR operation)
func gfAdd(a, b byte) byte {
	return a ^ b
}

// mul multiplies two elements in GF(2^8) using log tables
func gfMul(a, b byte) byte {
	if a == 0 || b == 0 {
		return 0
	}
	return gfExp[int(gfLog[a])+int(gfLog[b])]
}

// div divides two elements in GF(2^8)
func gfDiv(a, b byte) byte {
	if b == 0 {
		panic("division by zero in GF(2^8)")
	}
	if a == 0 {
		return 0
	}
	logA := int(gfLog[a])
	logB := int(gfLog[b])
	diff := logA - logB
	if diff < 0 {
		diff += 255
	}
	return gfExp[diff]
}

// evaluatePolynomial evaluates a polynomial in GF(2^8) at x
// coefficients[0] is the constant term (the secret)
func evaluatePolynomial(coefficients []byte, x byte) byte {
	if x == 0 {
		return coefficients[0]
	}
	var result byte = 0
	// Horner's method
	for i := len(coefficients) - 1; i >= 0; i-- {
		result = gfAdd(gfMul(result, x), coefficients[i])
	}
	return result
}

// SplitSecret splits a secret byte array into N shares, requiring at least K shares to reconstruct
func SplitSecret(secret []byte, n, k int) ([]ShamirShare, error) {
	if k < 2 {
		return nil, errors.New("threshold K must be at least 2")
	}
	if n < k {
		return nil, errors.New("total shares N must be greater than or equal to threshold K")
	}
	if n > 255 {
		return nil, errors.New("total shares N cannot exceed 255")
	}

	shares := make([]ShamirShare, n)
	for i := 0; i < n; i++ {
		shares[i] = ShamirShare{
			Index: byte(i + 1),
			Value: make([]byte, len(secret)),
		}
	}

	// For each byte of the secret, create a polynomial of degree k-1
	for byteIndex, secretByte := range secret {
		coefficients := make([]byte, k)
		coefficients[0] = secretByte

		// Generate random coefficients for terms x^1 to x^(k-1)
		if _, err := rand.Read(coefficients[1:]); err != nil {
			return nil, fmt.Errorf("failed to generate random coefficients: %w", err)
		}

		// Evaluate polynomial for each share
		for i := 0; i < n; i++ {
			x := shares[i].Index
			shares[i].Value[byteIndex] = evaluatePolynomial(coefficients, x)
		}
	}

	return shares, nil
}

// ReconstructSecret recovers the secret from K shares
func ReconstructSecret(shares []ShamirShare) ([]byte, error) {
	if len(shares) == 0 {
		return nil, errors.New("no shares provided")
	}

	secretLen := len(shares[0].Value)
	numShares := len(shares)

	// Check that all shares have matching sizes
	for _, share := range shares {
		if len(share.Value) != secretLen {
			return nil, errors.New("all shares must have the same length")
		}
	}

	secret := make([]byte, secretLen)

	// Interpolate for each byte index
	for byteIndex := 0; byteIndex < secretLen; byteIndex++ {
		var secretByte byte = 0

		// Lagrange interpolation at x = 0
		for i := 0; i < numShares; i++ {
			xi := shares[i].Index
			yi := shares[i].Value[byteIndex]

			// Compute Lagrange coefficient li(0)
			var li byte = 1
			for j := 0; j < numShares; j++ {
				if i == j {
					continue
				}
				xj := shares[j].Index
				
				// li(0) *= xj / (xj - xi)
				num := xj
				den := gfAdd(xj, xi) // Subtract is same as add in GF(2^8)
				li = gfMul(li, gfDiv(num, den))
			}

			secretByte = gfAdd(secretByte, gfMul(yi, li))
		}
		secret[byteIndex] = secretByte
	}

	return secret, nil
}
