package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

// Document represents the unified data structure
type Document struct {
	ID         string
	ProviderID string
	SourceID   string
	DocType    string
	Title      string
	Content    string
	RawData    map[string]interface{}
	URL        string
	Author     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Store struct {
	db *sql.DB
}

// NewStore initializes a SQLite store
func NewStore(dbPath string, schemaPath string) (*Store, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	// Read and execute schema
	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read schema: %w", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return nil, fmt.Errorf("failed to init schema: %w", err)
	}

	return &Store{db: db}, nil
}

// SaveDocument inserts or updates a document
func (s *Store) SaveDocument(ctx context.Context, doc *Document) error {
	if doc.ID == "" {
		doc.ID = uuid.New().String()
	}

	rawJSON, err := json.Marshal(doc.RawData)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO documents (id, provider_id, source_id, doc_type, title, content, raw_data, url, author, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			title=excluded.title,
			content=excluded.content,
			raw_data=excluded.raw_data,
			updated_at=excluded.updated_at
	`

	_, err = s.db.ExecContext(ctx, query,
		doc.ID, doc.ProviderID, doc.SourceID, doc.DocType,
		doc.Title, doc.Content, string(rawJSON), doc.URL, doc.Author,
		doc.CreatedAt, doc.UpdatedAt,
	)
	return err
}

// Search queries the FTS5 virtual table
func (s *Store) Search(ctx context.Context, query string, limit int) ([]Document, error) {
	// Sanitize query for FTS5 (escape quotes and wrap in literal quotes)
	safeQuery := `"` + strings.ReplaceAll(query, `"`, `""`) + `"`

	// Simple match syntax for FTS5
	sqlQuery := `
		SELECT id, doc_type, title, content, author
		FROM documents_fts
		WHERE documents_fts MATCH ?
		ORDER BY rank
		LIMIT ?
	`
	rows, err := s.db.QueryContext(ctx, sqlQuery, safeQuery, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Document
	for rows.Next() {
		var doc Document
		if err := rows.Scan(&doc.ID, &doc.DocType, &doc.Title, &doc.Content, &doc.Author); err != nil {
			return nil, err
		}
		results = append(results, doc)
	}
	return results, nil
}

// GetAll retrieves all documents stored in the database
func (s *Store) GetAll(ctx context.Context) ([]Document, error) {
	query := `
		SELECT id, provider_id, source_id, doc_type, title, content, raw_data, url, author, created_at, updated_at
		FROM documents
	`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Document
	for rows.Next() {
		var doc Document
		var rawJSON string
		if err := rows.Scan(
			&doc.ID, &doc.ProviderID, &doc.SourceID, &doc.DocType,
			&doc.Title, &doc.Content, &rawJSON, &doc.URL, &doc.Author,
			&doc.CreatedAt, &doc.UpdatedAt,
		); err != nil {
			return nil, err
		}
		
		if rawJSON != "" {
			json.Unmarshal([]byte(rawJSON), &doc.RawData)
		}
		results = append(results, doc)
	}
	return results, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

