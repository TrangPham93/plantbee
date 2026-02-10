package storage

import (
	"database/sql"
	"esp32-server/internal/models"
	"fmt"
	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

// Connect the database
func New(connectionString string) (*DB, error) {
	if connectionString == "" {
		return nil, fmt.Errorf("connection string is empty")
	}

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{conn: db}, nil
}

// Save the standard data to db 
func (d *DB) Save(t models.TelemetryData) error {
	query := `
		INSERT INTO sensor_readings (captured_at, sensor_id, moisture_percent, wake_time_sec) 
		VALUES ($1, $2, $3, $4)`
	
	_, err := d.conn.Exec(query, t.CapturedAt, t.SensorID, t.MoisturePct, t.WakeTimeSec)
	return err
}