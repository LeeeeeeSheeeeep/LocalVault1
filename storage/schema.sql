-- Unified Database Schema for LocalVault

-- 1. Source Providers (e.g., GitHub, Gmail, Notion)
CREATE TABLE IF NOT EXISTS providers (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    auth_status TEXT DEFAULT 'disconnected',
    last_sync DATETIME
);

-- 2. Unified Documents Table (Core Data)
CREATE TABLE IF NOT EXISTS documents (
    id TEXT PRIMARY KEY,
    provider_id TEXT NOT NULL,
    source_id TEXT NOT NULL,      -- ID from the original platform (e.g., tweet ID, email message-id)
    doc_type TEXT NOT NULL,       -- e.g., 'email', 'issue', 'tweet', 'note', 'article'
    title TEXT,
    content TEXT,                 -- Markdown or plain text
    raw_data JSON,                -- Store the raw API response for future-proofing
    url TEXT,
    author TEXT,
    created_at DATETIME NOT NULL,
    updated_at DATETIME,
    FOREIGN KEY (provider_id) REFERENCES providers(id)
);

-- 3. SQLite FTS5 Virtual Table for Full-Text Search
CREATE VIRTUAL TABLE IF NOT EXISTS documents_fts USING fts5(
    title,
    content,
    author,
    doc_type UNINDEXED,
    id UNINDEXED,
    content='documents',
    content_rowid='rowid'
);

-- Triggers to keep FTS index updated
CREATE TRIGGER IF NOT EXISTS docs_ai AFTER INSERT ON documents BEGIN
  INSERT INTO documents_fts(rowid, title, content, author, doc_type, id)
  VALUES (new.rowid, new.title, new.content, new.author, new.doc_type, new.id);
END;

CREATE TRIGGER IF NOT EXISTS docs_ad AFTER DELETE ON documents BEGIN
  INSERT INTO documents_fts(documents_fts, rowid, title, content, author, doc_type, id)
  VALUES('delete', old.rowid, old.title, old.content, old.author, old.doc_type, old.id);
END;

CREATE TRIGGER IF NOT EXISTS docs_au AFTER UPDATE ON documents BEGIN
  INSERT INTO documents_fts(documents_fts, rowid, title, content, author, doc_type, id)
  VALUES('delete', old.rowid, old.title, old.content, old.author, old.doc_type, old.id);
  INSERT INTO documents_fts(rowid, title, content, author, doc_type, id)
  VALUES (new.rowid, new.title, new.content, new.author, new.doc_type, new.id);
END;

-- 4. Tags / Labels
CREATE TABLE IF NOT EXISTS tags (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS document_tags (
    document_id TEXT NOT NULL,
    tag_id INTEGER NOT NULL,
    PRIMARY KEY (document_id, tag_id),
    FOREIGN KEY (document_id) REFERENCES documents(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id)
);
