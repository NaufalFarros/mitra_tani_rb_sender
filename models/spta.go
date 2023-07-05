package models

type TestSpta struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	SbhStatus     int32   `json:"sbh_status" `
	IdAri         uint    `json:"id_ari" `
	NoSpat        string  `json:"no_spat" `
	KodeKatLahan  string  `json:"kode_kat_lahan" `
	KodePlant     string  `json:"kode_plant" `
	KodeAffd      string  `json:"kode_affd" `
	KodeBlok      string  `json:"kode_blok" `
	TglSpta       string  `json:"tgl_spta" `
	TebangPg      int32   `json:"tebang_pg" `
	AngkutPg      int32   `json:"angkut_pg" `
	NamaPetani    string  `json:"nama_petani" `
	IdPetani      string  `json:"id_petani" `
	DeskripsiBlok string  `json:"deskripsi_blok" `
	LuasHa        float64 `json:"luas_ha" `
	HaTertebang   float64 `json:"ha_tertebang" `
	TglTebang     string  `json:"tgl_tebang" `
	BrixSel       int32   `json:"brix_sel" `
	PhSel         int32   `json:"ph_sel" `
	DitolakSel    int32   `json:"ditolak_sel" `
	DitolakAlasan string  `json:"ditolak_alasan" `
	CetakSptaTgl  string  `json:"cetak_spta_tgl" `
}
