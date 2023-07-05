package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type SptaResponse struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	SbhStatus     int32              `json:"sbh_status" bson:"sbh_status"`
	IdAri         uint               `json:"id_ari" bson:"id_ari"`
	NoSpat        string             `json:"no_spat" bson:"no_spat"`
	KodeKatLahan  string             `json:"kode_kat_lahan" bson:"kode_kat_lahan"`
	KodePlant     string             `json:"kode_plant" bson:"kode_plant"`
	KodeAffd      string             `json:"kode_affd" bson:"kode_affd"`
	KodeBlok      string             `json:"kode_blok" bson:"kode_blok"`
	TglSpta       string             `json:"tgl_spta" bson:"tgl_spta"`
	TebangPg      int32              `json:"tebang_pg" bson:"tebang_pg"`
	AngkutPg      int32              `json:"angkut_pg" bson:"angkut_pg"`
	NamaPetani    string             `json:"nama_petani" bson:"nama_petani"`
	IdPetani      string             `json:"id_petani" bson:"id_petani"`
	DeskripsiBlok string             `json:"deskripsi_blok" bson:"deskripsi_blok"`
	LuasHa        float64            `json:"luas_ha" bson:"luas_ha"`
	HaTertebang   float64            `json:"ha_tertebang" bson:"ha_tertebang"`
	TglTebang     string             `json:"tgl_tebang" bson:"tgl_tebang"`
	BrixSel       int32              `json:"brix_sel" bson:"brix_sel"`
	PhSel         int32              `json:"ph_sel" bson:"ph_sel"`
	DitolakSel    int32              `json:"ditolak_sel" bson:"ditolak_sel"`
	DitolakAlasan string             `json:"ditolak_alasan" bson:"ditolak_alasan"`
	CetakSptaTgl  string             `json:"cetak_spta_tgl" bson:"cetak_spta_tgl"`
}
