package security

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"errors"
	"fmt"
	"hash/adler32"
	"io"
)

// Segment represents a single slice of the backup data intended for a single QR code
type Segment struct {
	Index     int    `json:"i"` // 1-based index
	Total     int    `json:"t"` // Total segments
	Payload   string `json:"p"` // Base64 encoded zlib payload chunk
	Checksum  uint32 `json:"c"` // Adler-32 checksum of this segment's payload
}

// QRBackupManager handles slicing, compressing, and validating backup chunks
type QRBackupManager struct {
	MaxSegmentBytes int
}

func NewQRBackupManager(maxBytes int) *QRBackupManager {
	if maxBytes <= 0 {
		maxBytes = 500 // Default safe size for standard QR code capacity (V20-V25)
	}
	return &QRBackupManager{MaxSegmentBytes: maxBytes}
}

// CreateBackup compresses the input data and splits it into QR-friendly segments
func (m *QRBackupManager) CreateBackup(data []byte) ([]Segment, error) {
	if len(data) == 0 {
		return nil, errors.New("no data to backup")
	}

	// 1. Compress the data using zlib
	var buf bytes.Buffer
	writer := zlib.NewWriter(&buf)
	if _, err := writer.Write(data); err != nil {
		return nil, fmt.Errorf("zlib compression failed: %w", err)
	}
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("zlib writer close failed: %w", err)
	}

	compressedBytes := buf.Bytes()
	compressedBase64 := base64.StdEncoding.EncodeToString(compressedBytes)

	// 2. Slice base64 into segments
	base64Bytes := []byte(compressedBase64)
	totalLength := len(base64Bytes)
	numSegments := (totalLength + m.MaxSegmentBytes - 1) / m.MaxSegmentBytes

	segments := make([]Segment, 0, numSegments)

	for i := 0; i < numSegments; i++ {
		start := i * m.MaxSegmentBytes
		end := start + m.MaxSegmentBytes
		if end > totalLength {
			end = totalLength
		}

		chunk := string(base64Bytes[start:end])
		checksum := adler32.Checksum([]byte(chunk))

		segments = append(segments, Segment{
			Index:     i + 1,
			Total:     numSegments,
			Payload:   chunk,
			Checksum:  checksum,
		})
	}

	return segments, nil
}

// RestoreBackup merges, validates, and decompresses segments back into original bytes
func (m *QRBackupManager) RestoreBackup(segments []Segment) ([]byte, error) {
	if len(segments) == 0 {
		return nil, errors.New("no segments provided for restore")
	}

	total := segments[0].Total
	if len(segments) != total {
		return nil, fmt.Errorf("missing segments: got %d, expected %d", len(segments), total)
	}

	// Order segments by index (1-based)
	ordered := make([]string, total)
	for _, seg := range segments {
		if seg.Index < 1 || seg.Index > total {
			return nil, fmt.Errorf("invalid segment index %d", seg.Index)
		}
		
		// Validate checksum
		calcChecksum := adler32.Checksum([]byte(seg.Payload))
		if calcChecksum != seg.Checksum {
			return nil, fmt.Errorf("checksum mismatch for segment %d: expected %x, got %x", seg.Index, seg.Checksum, calcChecksum)
		}
		
		ordered[seg.Index-1] = seg.Payload
	}

	// Join payloads
	var base64Buf bytes.Buffer
	for i, chunk := range ordered {
		if chunk == "" {
			return nil, fmt.Errorf("missing segment at index %d", i+1)
		}
		base64Buf.WriteString(chunk)
	}

	// Decode Base64
	compressedBytes, err := base64.StdEncoding.DecodeString(base64Buf.String())
	if err != nil {
		return nil, fmt.Errorf("base64 decoding failed: %w", err)
	}

	// Decompress zlib
	reader, err := zlib.NewReader(bytes.NewReader(compressedBytes))
	if err != nil {
		return nil, fmt.Errorf("zlib reader creation failed: %w", err)
	}
	defer reader.Close()

	var decompressedBuf bytes.Buffer
	// Cap decompressed output at 50MB to prevent zlib bomb attacks
	limitedReader := io.LimitReader(reader, 50*1024*1024)
	if _, err := io.Copy(&decompressedBuf, limitedReader); err != nil {
		return nil, fmt.Errorf("zlib decompression failed: %w", err)
	}

	return decompressedBuf.Bytes(), nil
}
