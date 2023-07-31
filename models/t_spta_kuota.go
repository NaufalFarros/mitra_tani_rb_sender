package models

import "time"

type TSptaKuota struct {
	ID            int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	TahunGiling   string    `gorm:"column:tahun_giling" json:"tahun_giling"`
	TglSpta       time.Time `gorm:"column:tgl_spta" json:"tgl_spta"`
	KuotaHarian   int       `gorm:"column:kuota_harian;default:0" json:"kuota_harian"`
	KuotaTambahan int       `gorm:"column:kuota_tambahan;default:0" json:"kuota_tambahan"`
	Simpan        int       `gorm:"column:simpan;default:0" json:"simpan"`
	Retur         int       `gorm:"column:retur;default:0" json:"retur"`
	Cetak         int       `gorm:"column:cetak;default:0" json:"cetak"`
	Analisa       int       `gorm:"column:analisa;default:0" json:"analisa"`
	PtgsInput     string    `gorm:"column:ptgs_input" json:"ptgs_input"`
	TglInput      time.Time `gorm:"column:tgl_input" json:"tgl_input"`
	KodePlant     string    `gorm:"column:kode_plant" json:"kode_plant"`
}
