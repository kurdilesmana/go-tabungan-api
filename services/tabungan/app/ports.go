package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kurdilesmana/go-tabungan-api/services/tabungan/data"
	"gorm.io/gorm"
)

type TabunganServicePort interface {
	RegisterRekening(registerRekening data.RegisterRekening) (nomor_rekening string, err error)
	Saving(saving data.Saving) (saldo float64, err error)
}

type TabunganDatastorePort interface {
	Begin() (tx *gorm.DB)
	Rollback(tx *gorm.DB)
	Commit(tx *gorm.DB)
	RegisterRekening(rekening data.Rekening) (IDRekening int, err error)
	CheckRekeningExist(nik, noHp string) (exist bool, err error)
	CheckRekeningExistByNomorRekening(noRekening string) (exist bool, err error)
	RegisterNomorRekening(nomorRekening string, idRekening int) (err error)
	Saving(saving data.Transaksi) (err error)
	UpdateSavingSaldo(nomorRekening string, nominal float64) (err error)
}

type TabunganHandlerPort interface {
	Register(c *fiber.Ctx) error
	Saving(c *fiber.Ctx) error
}
