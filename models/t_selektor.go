package models

import "time"

type TSelektor struct {
	NoSpat          string    `gorm:"column:no_spat"`
	IDSelektor      int       `gorm:"column:id_selektor;primary_key;AUTO_INCREMENT"` //
	IDSpta          int       `gorm:"column:id_spta;NOT NULL"`                       // refrens t_spta id
	PersnoMandorTma string    `gorm:"column:persno_mandor_tma"`                      // refrens ke sap_m_karyawaman
	TglTebang       time.Time `gorm:"column:tgl_tebang"`                             // tanggal tebang yang di isikan di spta
	NoAngkutan      string    `gorm:"column:no_angkutan"`                            // no polisi/nomor angkutan traktor
	PtgsAngkutan    string    `gorm:"column:ptgs_angkutan"`                          // supir/operator traktor
	HaTertebang     float64   `gorm:"column:ha_tertebang"`                           // ha tertebang inputan selektor
	TerbakarSel     int       `gorm:"column:terbakar_sel"`                           // 0=tidak terbakar, 1=terbakar
	DitolakSel      int       `gorm:"column:ditolak_sel"`                            // 0=diterima, 1=ditolak
	DitolakAlasan   string    `gorm:"column:ditolak_alasan"`                         // free text
	BrixSel         float64   `gorm:"column:brix_sel"`                               // penilaian brix di selektor
	PhSel           float64   `gorm:"column:ph_sel"`                                 // penilaian ph di selektor
	TglPintumasuk   time.Time `gorm:"column:tgl_pintumasuk"`                         // tgl jam peluncuran
	PtgsPintumasuk  string    `gorm:"column:ptgs_pintumasuk"`                        // user petugas luncuran
	NoUrutTimbang   int       `gorm:"column:no_urut_timbang"`                        // no urut penimbangan, di reset ke 0 pada jam 6 pagi
	NoTrainstat     string    `gorm:"column:no_trainstat"`                           // no trainstat
	NoHv            string    `gorm:"column:no_hv"`
	OpHv            string    `gorm:"column:op_hv"`
	NoStipping      string    `gorm:"column:no_stipping"`
	OpStipping      string    `gorm:"column:op_stipping"`
	NoGl            string    `gorm:"column:no_gl"`
	OpGl            string    `gorm:"column:op_gl"`
	PtgsSelektor    string    `gorm:"column:ptgs_selektor"`            // user yang melakukan selektor
	TglSelektor     time.Time `gorm:"column:tgl_selektor"`             // tgl saat melakukan selektor
	TanamanStatus   int       `gorm:"column:tanaman_status;default:0"` // 0. belum verif 1. sudah verif
	TanamanUser     string    `gorm:"column:tanaman_user"`
	TanamanAct      time.Time `gorm:"column:tanaman_act"`
	NoUrut          int       `gorm:"column:no_urut"`
	TglUrut         time.Time `gorm:"column:tgl_urut"`
	RfidCard        string    `gorm:"column:rfid_card"`
	OpenGate        time.Time `gorm:"column:open_gate"`
}
