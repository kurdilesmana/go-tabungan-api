package data

import "time"

type Rekening struct {
	Id         int     `db:"id" json:"id" gorm:"primaryKey"`
	Nama       string  `db:"nama" json:"nama"`
	Nik        string  `db:"nik" json:"nik"`
	NoHp       string  `db:"no_hp" json:"no_hp"`
	NoRekening string  `db:"no_rekening" json:"no_rekening"`
	Saldo      float64 `db:"saldo" json:"saldo"`
}

type Transaksi struct {
	Id            int       `db:"id" json:"id" gorm:"primaryKey"`
	Waktu         time.Time `db:"waktu" json:"waktu"`
	NoRekening    string    `db:"no_rekening" json:"no_rekening"`
	KodeTransaksi string    `db:"kode_transaksi" json:"kode_transaksi"`
	Nominal       float64   `db:"nominal" json:"nominal"`
}

type RegisterRekening struct {
	Nama string `json:"nama" validate:"required"`
	Nik  string `json:"nik" validate:"required"`
	NoHp string `json:"no_hp" validate:"required"`
}

type RegisterRekeningResponse struct {
	NoRekening string `json:"no_rekening" validate:"required"`
}

type Saving struct {
	NoRekening string  `json:"no_rekening" validate:"required"`
	Nominal    float64 `json:"nominal" validate:"required,numeric,min=1"`
}

type SavingResponse struct {
	Saldo float64 `json:"saldo"`
}

type CashWithdrawl struct {
	NoRekening string  `json:"no_rekening" validate:"required"`
	Nominal    float64 `json:"nominal" validate:"required,numeric,min=18"`
}

type CashWithdrawlResponse struct {
	Saldo float64 `json:"saldo"`
}

type Balance struct {
	NoRekening string `json:"no_rekening" validate:"required"`
}

type BalanceResponse struct {
	Saldo float64 `json:"saldo"`
}

type Mutation struct {
	NoRekening string `json:"no_rekening" validate:"required"`
}

type TransaksiMutasi struct {
	Waktu         time.Time `json:"waktu"`
	KodeTransaksi string    `json:"kode_transaksi"`
	Nominal       float64   `json:"nominal"`
}

type MutationResponse struct {
	Mutasi []TransaksiMutasi `json:"mutasi"`
}
