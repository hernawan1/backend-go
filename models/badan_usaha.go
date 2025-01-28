package models

import "gorm.io/gorm"

type BadanUsaha struct {
	IdBadanUsaha      uint   `gorm:"primary key;autoIncrement" json:"id_badan_usaha"`
	IdJenisBadanUsaha uint   `json:"id_jenis_badan_usah"`
	IdJabatan         uint   `json:"id_jabatan"`
	NIB               string `json:"nib"`
	NamaBadanUsaha    string `json:"nama_badan_usaha"`
	NoTelp            string `json:"no_telp"`
	Email             string `json:"email"`
	Fax               string `json:"fax"`
	Alamat            string `json:"alamat"`
	Website           string `json:"website"`
	SIUP              string `json:"siup"`
	NpwpBadanUsaha    string `json:"npwp_badan_usaha"`
	KodeNegara        string `json:"kode_negara"`
	BidangBadanUsaha  string `json:"bidang_badan_usaha"`
	KodeBadanUsaha    string `json:"kode_badan_usaha"`
	KodeDesa          string `json:"kode_desas"`
	RT                string `json:"rt"`
	RW                string `json:"rw"`
	IdNegara          string `json:"id_negara"`
	IdOss             string `json:"id_oss"`
	JenisPerseroan    string `json:"jenis_perseroan"`
	JenisPelakuUsaha  string `json:"jenis_pelaku_usaha"`
	JenisKawasanUsaha string `json:"jenis_kawasan_usaha"`
	TglPengajuanNib   string `json:"tgl_pengajuan_nib"`
	TglPenerbitanNib  string `json:"tgl_penerbitan_nib"`
	TglPerubahanNib   string `json:"tgl_perubahan_nib"`
	StatusBadanHukum  string `json:"status_badan_hukum"`
	Type              string `json:"type"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&BadanUsaha{})
	return err
}
