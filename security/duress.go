package security

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"os"
)

type AuthConfig struct {
	MasterHash string `json:"master_hash"`
	DuressHash string `json:"duress_hash"`
	Salt       string `json:"salt"`
}

type DuressManager struct {
	configPath string
	config     AuthConfig
}

func NewDuressManager(configPath string) *DuressManager {
	dm := &DuressManager{configPath: configPath}
	dm.loadConfig()
	return dm
}

func (dm *DuressManager) loadConfig() {
	file, err := os.ReadFile(dm.configPath)
	if err != nil {
		// Default fallback credentials if no config exists (for demo)
		// master password = "masterpassword", duress password = "duresspassword"
		// Salt = "localvault_salt_2025"
		dm.config = AuthConfig{
			MasterHash: dm.hashPassword("masterpassword", "localvault_salt_2025"),
			DuressHash: dm.hashPassword("duresspassword", "localvault_salt_2025"),
			Salt:       "localvault_salt_2025",
		}
		dm.saveConfig()
		return
	}
	json.Unmarshal(file, &dm.config)
}

func (dm *DuressManager) saveConfig() error {
	data, err := json.MarshalIndent(dm.config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dm.configPath, data, 0600)
}

func (dm *DuressManager) hashPassword(password, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password + salt))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Authenticate checks the password and returns:
// (isValid, isDuress)
func (dm *DuressManager) Authenticate(password string) (bool, bool) {
	inputHash := dm.hashPassword(password, dm.config.Salt)
	if subtle.ConstantTimeCompare([]byte(inputHash), []byte(dm.config.MasterHash)) == 1 {
		return true, false
	}
	if subtle.ConstantTimeCompare([]byte(inputHash), []byte(dm.config.DuressHash)) == 1 {
		return true, true
	}
	return false, false
}

// UpdateCredentials updates the hashes for master and duress passwords
func (dm *DuressManager) UpdateCredentials(masterPassword, duressPassword string) error {
	dm.config.MasterHash = dm.hashPassword(masterPassword, dm.config.Salt)
	dm.config.DuressHash = dm.hashPassword(duressPassword, dm.config.Salt)
	return dm.saveConfig()
}

// PopulateDecoyData inserts realistic fake items into the duress database store
func (dm *DuressManager) GetDecoyDocuments() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"provider_id": "rss",
			"source_id":   "rss_decoy_1",
			"doc_type":    "article",
			"title":       "Getting Started with Go and WebAssembly",
			"content":     "WebAssembly (Wasm) is a binary instruction format for a stack-based virtual machine. Wasm is designed as a portable compilation target for programming languages, enabling deployment on the web for client and server applications. In this article, we explore how to build a basic Go application, compile it to WebAssembly, and load it into a standard HTML dashboard dynamically.",
			"url":         "https://go.dev/blog/wasi",
			"author":      "Go Developer Blog",
		},
		{
			"provider_id": "github",
			"source_id":   "gh_decoy_2",
			"doc_type":    "code_commit",
			"title":       "Refactor config parser for clean environment loading",
			"content":     "Cleaned up the JSON parser routines in our utility loader. Replaced manual byte slicing with native json.NewDecoder to avoid extra memory allocations. Benchmark tests show a 12% improvement in startup memory footprint under heavy configurations. All integration tests are verified passing.",
			"url":         "https://github.com/LeeeeeeSheeeeep/utils",
			"author":      "LeeeeeeSheeeeep",
		},
		{
			"provider_id": "manual",
			"source_id":   "note_decoy_3",
			"doc_type":    "note",
			"title":       "Weekly Groceries and Shopping List",
			"content":     "1. Organic milk (2 cartons)\n2. Whole wheat sourdough bread\n3. Coffee beans (Ethiopian medium roast)\n4. Olive oil (Extra Virgin)\n5. Fresh basil and cherry tomatoes for caprese salad\n6. Greek yogurt (unsweetened)",
			"url":         "",
			"author":      "LeeeeeeSheeeeep",
		},
		{
			"provider_id": "rss",
			"source_id":   "rss_decoy_4",
			"doc_type":    "article",
			"title":       "Why Self-Hosting is the Future of Personal Computing",
			"content":     "In recent years, centralization has led to privacy breaches, ad-ridden platforms, and sudden shutdowns of services we rely on daily. Self-hosting your own RSS readers, storage repositories, and analytics allows you to regain absolute ownership of your digital life. Using Docker and SQLite, setting up a secure server takes less than an hour.",
			"url":         "https://news.ycombinator.com/item?id=3891023",
			"author":      "PrivacyAdvocate",
		},
	}
}
