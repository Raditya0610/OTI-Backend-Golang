package models

type Event struct {
	ID             uint         `json : "event_id" gorm : "primary_key"`
	Nama           string       `json : "nama_event"`
	Deskripsi      string       `json : "deskripsi_event"`
	TanggalMulai   string       `json : "tanggal_mulai"`
	TanggalSelesai string       `json : "tanggal_selesai"`
	Lokasi         string       `json : "lokasi"`
	Capacity       int64        `json : "kapasitas_peserta"`
	Enrollment     []Enrollment `json : "pendaftar"`
}
