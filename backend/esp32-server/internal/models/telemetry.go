package models

import "time"

// 1. The raw data recieved from esp32
type IncomingPayload struct {
	SensorID   string `json:"sensor_id"`
	Moisture   int    `json:"moisture"`
	DurationMs int    `json:"duration_ms"`
}

// 2. the data that will be saved to DB
type TelemetryData struct {
	CapturedAt  time.Time `json:"captured_at"`
	SensorID    string    `json:"sensor_id"`
	MoisturePct int       `json:"moisture_percent"`
	WakeTimeSec float64   `json:"wake_time_seconds"`
}