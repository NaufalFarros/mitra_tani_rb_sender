package models

import "time"

type TSptaKuotaKkw struct {
	ID          int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	TglSpta     time.Time `gorm:"column:tgl_spta;NOT NULL" json:"tgl_spta"`
	IDAffd      int       `gorm:"column:id_affd;NOT NULL" json:"id_affd"`
	IDSptaKuota int       `gorm:"column:id_spta_kuota" json:"id_spta_kuota"`
	KodePlant   string    `gorm:"column:kode_plant;NOT NULL" json:"kode_plant"`
	KuotaSpta   int       `gorm:"column:kuota_spta" json:"kuota_spta"`
	TglInput    time.Time `gorm:"column:tgl_input" json:"tgl_input"`
	PtgsInput   string    `gorm:"column:ptgs_input" json:"ptgs_input"`
}
