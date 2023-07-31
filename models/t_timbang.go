package models

import "time"

type TTimbang struct {
	NoSpat              string    `gorm:"column:no_spat;NOT NULL"` // no spat
	IDTimbangan         int       `gorm:"column:id_timbangan;primary_key;AUTO_INCREMENT"`
	IDSpat              int       `gorm:"column:id_spat;default:0;NOT NULL"`
	LokasiTimbang1      string    `gorm:"column:lokasi_timbang_1"` // id_timbangan
	LokasiTimbang2      string    `gorm:"column:lokasi_timbang_2"` // id_timbangan
	Bruto               int       `gorm:"column:bruto;default:0"`
	Tara                int       `gorm:"column:tara;default:0"`
	Netto               int       `gorm:"column:netto;default:0"`
	NettoFinal          int       `gorm:"column:netto_final;default:0"`         // nilai netto akhir setelah rafaksi
	NettoRafaksi        int       `gorm:"column:netto_rafaksi"`                 // nilai neto rafaksi
	RafaksiProsentis    float64   `gorm:"column:rafaksi_prosentis;default:0"`   // prosentis rafaksi
	TglBruto            time.Time `gorm:"column:tgl_bruto"`                     // waktu pengambilan bruto
	TglTara             time.Time `gorm:"column:tgl_tara"`                      // waktu pengambilan tara
	TglNetto            time.Time `gorm:"column:tgl_netto"`                     // waktu pengambilan netto
	TglRafaksi          time.Time `gorm:"column:tgl_rafaksi"`                   // waktu entry data rafaksi
	TransloadingStatus  int       `gorm:"column:transloading_status;default:0"` // 0=tidak transloading,1=transloading
	NoTransloading      string    `gorm:"column:no_transloading"`               // no lori, no kontainer, no emplasement
	PtgsTransloading    string    `gorm:"column:ptgs_transloading"`             // nama petugas transloading
	PtgsTimbang1        string    `gorm:"column:ptgs_timbang_1"`                // user penimbang
	PtgsTimbang2        string    `gorm:"column:ptgs_timbang_2"`                // user penimbang
	TglTransloading     time.Time `gorm:"column:tgl_transloading"`              // waktu tranloading
	MultiSling          string    `gorm:"column:multi_sling"`                   // text penyimpanan multisling
	NettoSebelumKoreksi int       `gorm:"column:netto_sebelum_koreksi"`         // netto sebelum terjadi koreksi timbangan
	KetKoreksiTimbangan string    `gorm:"column:ket_koreksi_timbangan"`         // keterangan penyebab koreksi
	TrainStat           string    `gorm:"column:train_stat"`                    // no trainstat
	NoLori              string    `gorm:"column:no_lori"`                       // no lori
	NoLoko              string    `gorm:"column:no_loko"`                       // no_loko
}
