package handlers

import (
	"encoding/json"
	"esp32-server/internal/models"
	"esp32-server/internal/storage"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Handler struct {
	DB *storage.DB
}

func (h *Handler) IngestData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. Decode
	var raw models.IncomingPayload
	if err := json.NewDecoder(r.Body).Decode(&raw); err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}

	// 2. Transform
	telemetry := models.TelemetryData{
		CapturedAt:  time.Now().UTC(),
		SensorID:    raw.SensorID,
		MoisturePct: raw.Moisture,
		WakeTimeSec: float64(raw.DurationMs) / 1000.0,
	}

	// 3. Log (Console)
	printLog(telemetry)

	// 4. Save to DB (If there is one)
	if h.DB != nil {
		err := h.DB.Save(telemetry)
		if err != nil {
			fmt.Println("❌ DB Error:", err)
		} else {
			fmt.Println("💾 Saved to DB")
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ack"))
}

// Print to console
func printLog(t models.TelemetryData) {
	barLen := t.MoisturePct / 10
	if barLen > 10 { barLen = 10 }
	if barLen < 0 { barLen = 0 }
	bar := strings.Repeat("█", barLen) + strings.Repeat("░", 10-barLen)

	fmt.Printf("\n📦 [%s] PACKET: ID=%s | Time=%.2fs | Moisture=%s %d%%\n", 
		t.CapturedAt.Format("15:04:05"), t.SensorID, t.WakeTimeSec, bar, t.MoisturePct)
}