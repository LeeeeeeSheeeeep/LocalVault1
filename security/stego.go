package security

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
)

// Stego represents the image steganography module
type Stego struct{}

func NewStego() *Stego {
	return &Stego{}
}

// Encode embeds data into a cover image (PNG) and writes the result to output
func (s *Stego) Encode(coverReader io.Reader, data []byte, outputWriter io.Writer) error {
	img, err := png.Decode(coverReader)
	if err != nil {
		return fmt.Errorf("failed to decode cover image: %w", err)
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Convert to RGBA for writing
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	// Combine length of data (4 bytes) and data payload
	dataLen := uint32(len(data))
	header := []byte{
		byte(dataLen >> 24),
		byte(dataLen >> 16),
		byte(dataLen >> 8),
		byte(dataLen),
	}
	payload := append(header, data...)

	// Total bits to write
	totalBits := len(payload) * 8
	maxBits := width * height * 3 // 3 channels (R, G, B) per pixel
	if totalBits > maxBits {
		return fmt.Errorf("image too small: payload requires %d bits, image fits %d bits max", totalBits, maxBits)
	}

	bitIdx := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := rgba.RGBAAt(x, y)

			// R channel
			if bitIdx < totalBits {
				bit := getBit(payload, bitIdx)
				c.R = setLSB(c.R, bit)
				bitIdx++
			}
			// G channel
			if bitIdx < totalBits {
				bit := getBit(payload, bitIdx)
				c.G = setLSB(c.G, bit)
				bitIdx++
			}
			// B channel
			if bitIdx < totalBits {
				bit := getBit(payload, bitIdx)
				c.B = setLSB(c.B, bit)
				bitIdx++
			}

			rgba.SetRGBA(x, y, c)

			if bitIdx >= totalBits {
				break
			}
		}
		if bitIdx >= totalBits {
			break
		}
	}

	// Write out the modified PNG
	err = png.Encode(outputWriter, rgba)
	if err != nil {
		return fmt.Errorf("failed to encode output image: %w", err)
	}

	return nil
}

// Decode extracts hidden data from an encoded image (PNG)
func (s *Stego) Decode(stegoReader io.Reader) ([]byte, error) {
	img, err := png.Decode(stegoReader)
	if err != nil {
		return nil, fmt.Errorf("failed to decode stego image: %w", err)
	}

	bounds := img.Bounds()
	rgba, ok := img.(*image.RGBA)
	if !ok {
		// If not *image.RGBA, convert it to draw
		rgba = image.NewRGBA(bounds)
		draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	}

	// 1. Read header (4 bytes = 32 bits) to find payload length
	header := make([]byte, 4)
	headerBits := 32
	bitIdx := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := rgba.RGBAAt(x, y)

			// R channel
			if bitIdx < headerBits {
				setBit(header, bitIdx, getLSB(c.R))
				bitIdx++
			}
			// G channel
			if bitIdx < headerBits {
				setBit(header, bitIdx, getLSB(c.G))
				bitIdx++
			}
			// B channel
			if bitIdx < headerBits {
				setBit(header, bitIdx, getLSB(c.B))
				bitIdx++
			}

			if bitIdx >= headerBits {
				break
			}
		}
		if bitIdx >= headerBits {
			break
		}
	}

	// Parse payload length
	dataLen := (uint32(header[0]) << 24) |
		(uint32(header[1]) << 16) |
		(uint32(header[2]) << 8) |
		uint32(header[3])

	// Sanity limit check: limit max extraction size to 10MB to avoid memory overflow on corruption
	if dataLen > 10*1024*1024 {
		return nil, fmt.Errorf("invalid payload length extracted: %d (corrupted image or no stego data)", dataLen)
	}

	// 2. Read actual payload data bits
	totalBits := int(dataLen) * 8
	payload := make([]byte, dataLen)
	payloadBitIdx := 0
	bitIdx = 0 // Reset bit index from start of pixel scan, skipping header bits

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := rgba.RGBAAt(x, y)

			// R channel
			if bitIdx >= headerBits {
				if payloadBitIdx < totalBits {
					setBit(payload, payloadBitIdx, getLSB(c.R))
					payloadBitIdx++
				}
			}
			bitIdx++

			// G channel
			if bitIdx >= headerBits {
				if payloadBitIdx < totalBits {
					setBit(payload, payloadBitIdx, getLSB(c.G))
					payloadBitIdx++
				}
			}
			bitIdx++

			// B channel
			if bitIdx >= headerBits {
				if payloadBitIdx < totalBits {
					setBit(payload, payloadBitIdx, getLSB(c.B))
					payloadBitIdx++
				}
			}
			bitIdx++

			if payloadBitIdx >= totalBits {
				break
			}
		}
		if payloadBitIdx >= totalBits {
			break
		}
	}

	return payload, nil
}

// GenerateMockCoverImage creates a simple 100x100 PNG to act as a default template if none is uploaded
func (s *Stego) GenerateMockCoverImage() ([]byte, error) {
	bounds := image.Rect(0, 0, 150, 150)
	rgba := image.NewRGBA(bounds)
	
	// Create a colorful gradient grid cover image
	for y := 0; y < 150; y++ {
		for x := 0; x < 150; x++ {
			rgba.SetRGBA(x, y, color.RGBA{
				R: uint8(x % 256),
				G: uint8(y % 256),
				B: uint8((x + y) % 256),
				A: 255,
			})
		}
	}
	
	var buf bytes.Buffer
	err := png.Encode(&buf, rgba)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Helper: check if bit at index is 1 or 0
func getBit(data []byte, bitIdx int) byte {
	byteIdx := bitIdx / 8
	localBitIdx := 7 - (bitIdx % 8)
	return (data[byteIdx] >> localBitIdx) & 1
}

// Helper: set bit at index in byte slice
func setBit(data []byte, bitIdx int, val byte) {
	byteIdx := bitIdx / 8
	localBitIdx := 7 - (bitIdx % 8)
	if val == 1 {
		data[byteIdx] |= (1 << localBitIdx)
	} else {
		data[byteIdx] &= ^(1 << localBitIdx)
	}
}

// Helper: set the LSB (bit 0) of a channel byte
func setLSB(channel byte, bit byte) byte {
	if bit == 1 {
		return channel | 1
	}
	return channel &^ 1
}

// Helper: read the LSB of a channel byte
func getLSB(channel byte) byte {
	return channel & 1
}
