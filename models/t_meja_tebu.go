package models

import "time"

type TMejaTebu struct {
	NoSpat        string    `gorm:"column:no_spat;NOT NULL"` // no spat
	IDMejatebu    int       `gorm:"column:id_mejatebu;primary_key;AUTO_INCREMENT"`
	IDSpta        int       `gorm:"column:id_spta;NOT NULL"` // refrens t_spta id
	Daduk         int       `gorm:"column:daduk"`            // meja tebu
	Sogolan       int       `gorm:"column:sogolan"`          // meja tebu
	Pucuk         int       `gorm:"column:pucuk"`            // meja tebu
	AkarTanah     int       `gorm:"column:akar_tanah"`       // meja tebu
	NonTebu       int       `gorm:"column:non_tebu"`         // meja tebu
	Terbakar      int       `gorm:"column:terbakar"`         // meja tebu
	Cacahan       int       `gorm:"column:cacahan"`          // meja tebu
	Brondolan     int       `gorm:"column:brondolan"`        // meja tebu
	KondisiTebu   string    `gorm:"column:kondisi_tebu"`     // meja tebu
	PtgsMejaTebu  string    `gorm:"column:ptgs_meja_tebu"`   // meja tebu
	Gilingan      int       `gorm:"column:gilingan"`
	KodeMejaTebu  string    `gorm:"column:kode_meja_tebu"` // meja tebu
	WarnaMejaTebu string    `gorm:"column:warna_meja_tebu"`
	TglMejaTebu   time.Time `gorm:"column:tgl_meja_tebu"`            // meja tebu
	RafraksiAktif int       `gorm:"column:rafraksi_aktif;default:0"` // 0=rafaksi tidak aktif, 1=rafaksi aktif

}
