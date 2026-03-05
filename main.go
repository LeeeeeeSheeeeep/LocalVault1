package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"localvault/ai"
	"localvault/connectors"
	"localvault/connectors/github"
	"localvault/connectors/rss"
	"localvault/security"
	"localvault/storage"
)

type App struct {
	store         *storage.Store
	duressStore   *storage.Store
	activeStore   *storage.Store
	connectors    map[string]connectors.Connector
	nlpEngine     *ai.NLP
	deadman       *security.DeadManSwitch
	duressManager *security.DuressManager
	keystroke     *security.KeystrokeProfiler
	stego         *security.Stego
}

func NewApp(dbPath, duressDBPath, schemaPath, deadmanPath, authPath, keystrokePath string) (*App, error) {
	store, err := storage.NewStore(dbPath, schemaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to init master store: %w", err)
	}

	duressStore, err := storage.NewStore(duressDBPath, schemaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to init duress store: %w", err)
	}

	app := &App{
		store:         store,
		duressStore:   duressStore,
		activeStore:   duressStore, // Starts in duress/decoy mode by default for safety!
		connectors:    make(map[string]connectors.Connector),
		nlpEngine:     ai.NewNLP(),
		deadman:       security.NewDeadManSwitch(deadmanPath),
		duressManager: security.NewDuressManager(authPath),
		keystroke:     security.NewKeystrokeProfiler(keystrokePath),
		stego:         security.NewStego(),
	}

	// Register connectors
	app.connectors["github"] = github.New()
	app.connectors["rss"] = rss.New()

	// Pre-populate decoy data if duress DB is empty
	ctx := context.Background()
	existing, err := duressStore.Search(ctx, "*", 1)
	if err != nil || len(existing) == 0 {
		decoys := app.duressManager.GetDecoyDocuments()
		for _, d := range decoys {
			doc := &storage.Document{
				ProviderID: d["provider_id"].(string),
				SourceID:   d["source_id"].(string),
				DocType:    d["doc_type"].(string),
				Title:      d["title"].(string),
				Content:    d["content"].(string),
				URL:        d["url"].(string),
				Author:     d["author"].(string),
				CreatedAt:  time.Now().AddDate(-1, -2, 0), // Created last year!
				UpdatedAt:  time.Now().AddDate(-1, -2, 0),
			}
			duressStore.SaveDocument(ctx, doc)
		}
	}

	return app, nil
}

func (a *App) handleUnlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	if r.Method == "OPTIONS" {
		return
	}

	var req struct {
		Password  string               `json:"password"`
		Keystroke security.TimingEntry `json:"keystroke"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	isValid, isDuress := a.duressManager.Authenticate(req.Password)
	if !isValid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"status": "fail", "reason": "invalid_credentials"})
		return
	}

	// Dynamic routing: if Duress Password was used
	if isDuress {
		a.activeStore = a.duressStore
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success", "vault": "decoy"})
		return
	}

	// Master Password was used - verify keystroke timing biometrics
	score, timingMatched := a.keystroke.Verify(req.Keystroke)
	
	// If timing doesn't match, silently route to the decoy vault!
	if !timingMatched {
		log.Printf("Keystroke timing profile mismatch (score: %.4f). Silently routing to decoy database.", score)
		a.activeStore = a.duressStore
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "success",
			"vault":  "decoy",
			"score":  score,
			"warn":   "Timing signature anomaly",
		})
		return
	}

	// Normal master unlock
	a.activeStore = a.store
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"vault":  "master",
		"score":  score,
	})
}

func (a *App) handleKeystrokeEnroll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	if r.Method == "OPTIONS" {
		return
	}

	var req struct {
		Entries []security.TimingEntry `json:"entries"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err := a.keystroke.Enroll(req.Entries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (a *App) handleStegoEncode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		return
	}

	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		http.Error(w, "failed to parse multipart form", http.StatusBadRequest)
		return
	}

	secret := r.FormValue("secret")
	if secret == "" {
		http.Error(w, "missing secret data to encode", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	var coverReader io.Reader
	if err != nil {
		// If no cover image uploaded, generate a default pattern image
		mockBytes, mockErr := a.stego.GenerateMockCoverImage()
		if mockErr != nil {
			http.Error(w, "failed to generate default cover: "+mockErr.Error(), http.StatusInternalServerError)
			return
		}
		coverReader = bytes.NewReader(mockBytes)
	} else {
		defer file.Close()
		coverReader = file
	}

	w.Header().Set("Content-Disposition", "attachment; filename=\"stego_vault.png\"")
	w.Header().Set("Content-Type", "image/png")

	err = a.stego.Encode(coverReader, []byte(secret), w)
	if err != nil {
		http.Error(w, "stego encoding failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) handleStegoDecode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		return
	}

	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		http.Error(w, "failed to parse multipart form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "missing stego image file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	decodedBytes, err := a.stego.Decode(file)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"status": "fail", "reason": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"secret": string(decodedBytes),
	})
}

func (a *App) handleSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	query := r.URL.Query().Get("q")
	if query == "" {
		json.NewEncoder(w).Encode([]storage.Document{})
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit := 50
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	results, err := a.activeStore.Search(r.Context(), query, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

type GraphNode struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Group int    `json:"group"`
	Size  int    `json:"size"`
}

type GraphEdge struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type GraphData struct {
	Nodes []GraphNode `json:"nodes"`
	Edges []GraphEdge `json:"edges"`
}

func (a *App) handleGraph(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	results, err := a.activeStore.Search(r.Context(), "*", 100)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data GraphData
	tagNodes := make(map[string]bool)

	for _, doc := range results {
		data.Nodes = append(data.Nodes, GraphNode{
			ID:    doc.ID,
			Label: doc.Title,
			Group: 1,
			Size:  15,
		})

		tags := a.nlpEngine.ExtractKeywords(doc.Content, 4)
		for _, tag := range tags {
			if !tagNodes[tag] {
				tagNodes[tag] = true
				data.Nodes = append(data.Nodes, GraphNode{
					ID:    tag,
					Label: tag,
					Group: 2,
					Size:  8,
				})
			}
			data.Edges = append(data.Edges, GraphEdge{
				Source: doc.ID,
				Target: tag,
			})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

type SyncRequest struct {
	ProviderID string            `json:"provider_id"`
	Config     map[string]string `json:"config"`
}

func (a *App) handleSync(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		return
	}

	var req SyncRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	connector, ok := a.connectors[req.ProviderID]
	if !ok {
		http.Error(w, "connector not found", http.StatusNotFound)
		return
	}

	ctx := context.Background()
	if err := connector.Auth(ctx, req.Config); err != nil {
		http.Error(w, fmt.Sprintf("auth failed: %v", err), http.StatusUnauthorized)
		return
	}

	docChan := make(chan *storage.Document, 100)
	errChan := make(chan error, 1)

	go func() {
		state := connectors.SyncState{}
		_, err := connector.Sync(ctx, state, docChan)
		if err != nil {
			log.Printf("Sync error for %s: %v", req.ProviderID, err)
		}
		close(docChan)
		errChan <- err
	}()

	count := 0
	for doc := range docChan {
		if err := a.activeStore.SaveDocument(ctx, doc); err != nil {
			log.Printf("Failed to save doc %s: %v", doc.SourceID, err)
		} else {
			count++
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"synced": count,
	})
}

func (a *App) handleWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		return
	}

	latStr := r.URL.Query().Get("lat")
	lonStr := r.URL.Query().Get("lon")
	dateStr := r.URL.Query().Get("date")

	temp := 22.0
	condition := "Clear"
	wind := 5.4
	solarAlt := 45.0

	if dateStr != "" {
		if date, err := time.Parse(time.RFC3339, dateStr); err == nil {
			hour := date.Hour()
			temp = 15.0 + 10.0*math.Sin(float64(hour-8)/24.0*2.0*math.Pi)
			solarAlt = 90.0 * math.Sin(float64(hour)/24.0*2.0*math.Pi)
			if solarAlt < 0 {
				solarAlt = 0
			}

			dayOfYear := date.YearDay()
			if (dayOfYear+hour)%7 == 0 {
				condition = "Rainy"
				temp -= 3.0
				wind = 12.5
			} else if (dayOfYear+hour)%9 == 0 {
				condition = "Windy"
				wind = 22.1
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"lat":            latStr,
		"lon":            lonStr,
		"temp":           temp,
		"condition":      condition,
		"wind_speed":     wind,
		"solar_altitude": solarAlt,
	})
}

func (a *App) handleDeadMan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	if r.Method == "OPTIONS" {
		return
	}

	a.deadman.RecordHeartbeat()

	if r.Method == "POST" {
		var req struct {
			Action     string   `json:"action"`
			Threshold  int      `json:"threshold"`
			Total      int      `json:"total"`
			Min        int      `json:"min"`
			Recipients []string `json:"recipients"`
			Key        string   `json:"key"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		switch req.Action {
		case "heartbeat":
			w.Write([]byte(`{"status":"success"}`))
		case "settings":
			err := a.deadman.UpdateSettings(req.Threshold, req.Total, req.Min, req.Recipients)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte(`{"status":"success"}`))
		case "split":
			if req.Key == "" {
				http.Error(w, "encryption key is required", http.StatusBadRequest)
				return
			}
			shares, err := a.deadman.ExecuteKeySplit([]byte(req.Key))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "success",
				"shares": shares,
			})
		default:
			http.Error(w, "invalid action", http.StatusBadRequest)
		}
		return
	}

	lastH, threshold, triggered := a.deadman.GetStatus()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"last_heartbeat": lastH,
		"threshold_days": threshold,
		"triggered":      triggered,
	})
}

func (a *App) handleBackup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method == "POST" {
		var segments []security.Segment
		if err := json.NewDecoder(r.Body).Decode(&segments); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		bm := security.NewQRBackupManager(450)
		recoveredBytes, err := bm.RestoreBackup(segments)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var docs []*storage.Document
		if err := json.Unmarshal(recoveredBytes, &docs); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		count := 0
		for _, doc := range docs {
			if err := a.activeStore.SaveDocument(ctx, doc); err == nil {
				count++
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":   "success",
			"restored": count,
		})
		return
	}

	docs, err := a.activeStore.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	docBytes, err := json.Marshal(docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bm := security.NewQRBackupManager(450)
	segments, err := bm.CreateBackup(docBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":   "success",
		"segments": segments,
	})
}

func main() {
	err := os.MkdirAll("data", 0755)
	if err != nil {
		log.Fatalf("Failed to create data dir: %v", err)
	}

	schemaPath := filepath.Join("storage", "schema.sql")
	dbPath := filepath.Join("data", "localvault.db")
	duressDBPath := filepath.Join("data", "localvault_duress.db")
	deadmanPath := filepath.Join("data", "deadman.json")
	authPath := filepath.Join("data", "auth.json")
	keystrokePath := filepath.Join("data", "keystroke.json")

	app, err := NewApp(dbPath, duressDBPath, schemaPath, deadmanPath, authPath, keystrokePath)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}
	defer app.store.Close()
	defer app.duressStore.Close()

	http.HandleFunc("/api/unlock", app.handleUnlock)
	http.HandleFunc("/api/keystroke/enroll", app.handleKeystrokeEnroll)
	http.HandleFunc("/api/stego/encode", app.handleStegoEncode)
	http.HandleFunc("/api/stego/decode", app.handleStegoDecode)

	http.HandleFunc("/api/search", app.handleSearch)
	http.HandleFunc("/api/graph", app.handleGraph)
	http.HandleFunc("/api/sync", app.handleSync)
	http.HandleFunc("/api/weather", app.handleWeather)
	http.HandleFunc("/api/deadman", app.handleDeadMan)
	http.HandleFunc("/api/backup", app.handleBackup)

	fmt.Println("LocalVault Backend running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
