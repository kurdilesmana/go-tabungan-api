package data

type Rekening struct {
	Id         int     `db:"id" json:"id" gorm:"primaryKey"`
	Nama       string  `db:"nama" json:"nama"`
	Nik        string  `db:"nik" json:"nik"`
	NoHp       string  `db:"no_hp" json:"no_hp"`
	NoRekening string  `db:"no_rekening" json:"no_rekening"`
	Saldo      float64 `db:"saldo" json:"saldo"`
}

type Transaksi struct {
	Id            int     `db:"id" json:"id" gorm:"primaryKey"`
	Waktu         string  `db:"waktu" json:"waktu"`
	NoRekening    string  `db:"no_rekening" json:"no_rekening"`
	KodeTransaksi string  `db:"kode_transaksi" json:"kode_transaksi"`
	Nominal       float64 `db:"nominal" json:"nominal"`
}

type RegisterRekening struct {
	Nama string `json:"nama"`
	Nik  string `json:"nik"`
	NoHp string `json:"no_hp"`
}

type RegisterRekeningResponse struct {
	NoRekening string `json:"no_rekening"`
}

type Saving struct {
	NoRekening string  `json:"no_rekening"`
	Nominal    float64 `json:"nominal"`
}

type SavingRespoonse struct {
	Saldo float64 `json:"nominal"`
}
