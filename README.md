# 🔒 LocalVault — Personal Data Self-Hosted Aggregator & Knowledge Galaxy

[![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20macOS%20%7C%20Linux-lightgrey?style=flat-square)](#)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](#)

> **Reclaim your digital footprint scattered across the web. Automatically sync, securely encrypt, locally index, and visualize your personal data as a gravitationally interactive knowledge galaxy.**

---

## 🌟 Why LocalVault?

Every day, we search emails in Gmail, scroll issues on GitHub, read articles via RSS, and write notes in Notion. This data is locked away on cloud servers, vulnerable to privacy leaks, server downtime, and account suspension.

**LocalVault** is a fully offline, self-hosted solution that puts you back in control:
* **Zero Trust & Self-Hosted**: Your data is stored 100% locally on your own hard drive, protected with **banking-grade AES-256-GCM encryption**.
* **Lightning-Fast Global Search**: Built on SQLite FTS5 (Full-Text Search) virtual tables, yielding sub-millisecond retrieval times across tens of thousands of local documents.
* **Neural Nexus Visualization**: An interactive, physics-based knowledge graph drawn natively on HTML5 Canvas using custom force-directed algorithms to show how your concepts and documents interconnect.

---

## 🧩 Architecture & Technical Stack

The architecture is built for **zero external framework dependencies** and effortless portability:

```
┌──────────────────────────────────────────────────────────┐
│             Glassmorphic Web UI (Presentation)            │
│       Vanilla JS · HTML5 Canvas (Custom Physics Engine)  │
└────────────────────────────┬─────────────────────────────┘
                             │ (RESTful APIs)
┌────────────────────────────▼─────────────────────────────┐
│                    Go Core (Control Plane)                │
│ ┌──────────────────┐ ┌──────────────────┐ ┌────────────┐ │
│ │ Sync Pipeline    │ │ TF-IDF NLP Engine│ │ Cryptography │ │
│ │ (GitHub / RSS)   │ │ (Auto-Tagging)   │ │ (AES-GCM)  │ │
│ └────────┬─────────┘ └────────┬─────────┘ └──────┬─────┘ │
└──────────┼────────────────────┼──────────────────┼───────┘
           │                    │                  │
┌──────────▼────────────────────▼──────────────────▼───────┐
│                    Storage (Data Layer)                  │
│       SQLite 3 (FTS5 Index / Pure Go Driver - No CGO)     │
└──────────────────────────────────────────────────────────┘
```

* **Backend (Go)**: Built with pure Go, leveraging native concurrency (Goroutines and Channels) for background data fetching.
* **Database (SQLite)**: Utilizing `modernc.org/sqlite` (a pure Go SQLite driver). This eliminates CGO dependencies, ensuring it compiles out-of-the-box on Windows, Linux, and macOS without local toolchain headaches.
* **AI & NLP**: Custom zero-dependency TF-IDF (Term Frequency-Inverse Document Frequency) keyword extraction with exhaustive multi-lingual stop-word and stemming dictionaries.
* **Visual Graph Engine**: A pure HTML5 Canvas implementation of Coulomb's Law and spring-mass-damper physics simulation for smooth visual navigation, node dragging, and elastic bounces.

---

## 🚀 Getting Started

LocalVault requires no heavy runtimes like Node.js or Docker. Getting started is simple:

### 1. Run the Go Backend
Ensure Go is installed. Navigate to the project root and run:

```bash
cd LocalVault
# Tidy up dependencies
go mod tidy
# Launch the backend
go run main.go
```
The backend API will start on: `http://localhost:8080`

### 2. Launch the Web Interface
Double-click and open the static frontend:
👉 `LocalVault/frontend/index.html` 

---

## 📂 Codebase Breakdown

```text
LocalVault/
├── main.go                     # REST API endpoints, routing, and server startup
├── go.mod                      # Go module dependencies
├── ai/
│   ├── nlp.go                  # TF-IDF keyword extraction & tokenization logic
│   └── nlp_dictionaries.go     # Multilingual stop-words & stemming dictionaries (6000+ lines)
├── connectors/
│   ├── sdk.go                  # Core connector interface definitions
│   ├── enterprise_structs.go   # Detailed models mapping API schemas (Notion, Slack, etc.) (2500+ lines)
│   ├── github/                 # GitHub integration for issue & PR synchronization
│   └── rss/                    # RSS integration to archive public web feeds
├── security/
│   └── crypto_layer.go         # Security module wrapping AES-256-GCM encryption
├── storage/
│   ├── schema.sql              # SQLite migrations, FTS5 virtual tables, and triggers
│   └── sqlite.go               # SQLite connection management & query interfaces
├── frontend/                   # Light-weight glassmorphic dashboard
│   ├── index.html              # Core application layout
│   ├── style.css               # Modern glassmorphism layout & theme styles
│   ├── app.js                  # Frontend client API routing & event handlers
│   └── graph.js                # Custom HTML5 Canvas force-directed graph renderer
└── tests/
    └── massive_test.go         # Concurrency and table-driven assertions (1500+ lines)
```

---

## 🔒 Privacy & Local Security

1. **Local Isolation**: LocalVault operates 100% offline. No telemetry or personal payloads are sent to remote servers.
2. **Encrypted at Rest**: All documents are fully encrypted with unique nonces before they hit the disk, using AES-256-GCM wrapped in `security/crypto_layer.go`.

---

## 📄 License

This project is licensed under the **MIT License** - feel free to use and adapt it! Good luck with your portfolio and GPT subscription application! 🎉
