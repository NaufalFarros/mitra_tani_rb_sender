package models

import "time"

type TSptaKuotaTot struct {
	ID             int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	TglSpta        time.Time `gorm:"column:tgl_spta;NOT NULL" json:"tgl_spta"`
	KodeBlok       string    `gorm:"column:kode_blok;NOT NULL" json:"kode_blok"`
	KoutaTot       int       `gorm:"column:kouta_tot" json:"kouta_tot"`
	TglInput       time.Time `gorm:"column:tgl_input" json:"tgl_input"`
	PtgsInput      string    `gorm:"column:ptgs_input" json:"ptgs_input"`
	IDSptaKuota    int       `gorm:"column:id_spta_kuota" json:"id_spta_kuota"`
	IDSptaKuotaKkw int       `gorm:"column:id_spta_kuota_kkw" json:"id_spta_kuota_kkw"`
}
