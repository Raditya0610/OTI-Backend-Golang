package models

import "time"

type Enrollment struct {
	ID        uint      `json : "pendaftaran_id"`
	EventID   uint      `json : "event_id"`
	UserID    uint      `json : "user_id"`
	CreatedAt time.Time `json : "tanggal_pendaftaran"`
	Informasi string    `json : "informasi_tambahan"`
}
