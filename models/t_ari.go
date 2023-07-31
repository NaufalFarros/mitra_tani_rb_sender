package models

import "time"

type TAri struct {
	NoSpat              string    `gorm:"column:no_spat;NOT NULL"` // no spat
	IDAri               int       `gorm:"column:id_ari;primary_key;AUTO_INCREMENT"`
	IDSpta              int       `gorm:"column:id_spta;NOT NULL"` // refrens t_spta id_spta
	PolTebu             float64   `gorm:"column:pol_tebu;default:0.00"`
	PersenBrixAri       float64   `gorm:"column:persen_brix_ari;default:0.00"` // ari
	PersenPolAri        float64   `gorm:"column:persen_pol_ari;default:0.00"`  // ari
	PhAri               float64   `gorm:"column:ph_ari;default:0.00"`          // ari
	Hk                  float64   `gorm:"column:hk;default:0.00"`
	FaktorPerah         float64   `gorm:"column:faktor_perah;default:0.0000"`   // faktor perah dari liquid
	FaktorRegresi       float64   `gorm:"column:faktor_regresi;default:0.0000"` // faktor regresi kedawung / faktor Kristal Rendemen
	NilaiNira           float64   `gorm:"column:nilai_nira;default:0.00"`
	FaktorRendemen      float64   `gorm:"column:faktor_rendemen;default:0.00"`   // jatiroto metod menjadi faktor perah
	RendemenAri         float64   `gorm:"column:rendemen_ari;default:0.00"`      // jatmed menjadi rendemen contoh
	FaktorKonversi      float64   `gorm:"column:faktor_konversi;default:0.00"`   // faktor konversi
	RendemenIndividu    float64   `gorm:"column:rendemen_individu;default:0.00"` // jatmed
	HablurAri           float64   `gorm:"column:hablur_ari;default:0.00"`
	GulaTotal           float64   `gorm:"column:gula_total;default:0.00"`
	TetesTotal          float64   `gorm:"column:tetes_total;default:0.00"`
	RendemenPtr         float64   `gorm:"column:rendemen_ptr;default:0.0000"`
	GulaPtr             float64   `gorm:"column:gula_ptr;default:0.00"`
	TetesPtr            float64   `gorm:"column:tetes_ptr;default:0.00"`
	GulaPg              float64   `gorm:"column:gula_pg;default:0.00"`
	TetesPg             float64   `gorm:"column:tetes_pg;default:0.00"`
	DitolakAri          int       `gorm:"column:ditolak_ari;default:0"` // 0 , 1 dotolak
	DitolakAlasan       string    `gorm:"column:ditolak_alasan"`
	TglAri              time.Time `gorm:"column:tgl_ari"`  // ari
	PtgsAri             string    `gorm:"column:ptgs_ari"` // ari
	SbhAriStatus        int       `gorm:"column:sbh_ari_status;default:0"`
	SbhAriTgl           time.Time `gorm:"column:sbh_ari_tgl"`
	SbhAriUser          string    `gorm:"column:sbh_ari_user"`
	PengolahanStatus    int       `gorm:"column:pengolahan_status;default:0"`
	PengolahanTgl       time.Time `gorm:"column:pengolahan_tgl"`
	PengolahanUser      string    `gorm:"column:pengolahan_user"`
	TanamanStatus       int       `gorm:"column:tanaman_status;default:0"`
	TanamanTgl          time.Time `gorm:"column:tanaman_tgl"`
	TanamanUser         string    `gorm:"column:tanaman_user"`
	AkuStatus           int       `gorm:"column:aku_status;default:0"`
	AkuTgl              time.Time `gorm:"column:aku_tgl"`
	AkuUser             string    `gorm:"column:aku_user"`
	SembilanpuluhPersen float64   `gorm:"column:sembilanpuluh_persen;default:0.00"`
	SepuluhPersen       float64   `gorm:"column:sepuluh_persen;default:0.00"`
	RSpg                float64   `gorm:"column:r_spg;default:0.0000"`
	KopensasiGula       float64   `gorm:"column:kopensasi_gula;default:0.0000"`
	Coresampler         int       `gorm:"column:coresampler;default:0"`
}
