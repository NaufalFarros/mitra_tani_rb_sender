package models

type LogDeteksi struct {
	ID               int  `gorm:"primaryKey;autoIncrement"`
	OriginalDataID   int  // Kolom untuk menyimpan ID dari data asli yang dimigrasi
	MigrateToMongoDb bool // Kolom untuk status migrasi ke MongoDB
}
