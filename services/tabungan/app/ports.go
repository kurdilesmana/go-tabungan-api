package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kurdilesmana/go-tabungan-api/services/tabungan/data"
	"gorm.io/gorm"
)

type TabunganServicePort interface {
	RegisterRekening(registerRekening data.RegisterRekening) (nomor_rekening string, err error)
	Saving(saving data.Saving) (saldo float64, err error)
	CashWithdrawl(cashWithdrawl data.CashWithdrawl) (saldo float64, err error)
	Balance(balance data.Balance) (saldo float64, err error)
	Mutation(mutation data.Mutation) (transaksi []data.Transaksi, err error)
}

type TabunganDatastorePort interface {
	Begin() (tx *gorm.DB)
	Rollback(tx *gorm.DB)
	Commit(tx *gorm.DB)
	RegisterRekening(rekening data.Rekening) (IDRekening int, err error)
	CheckRekeningExist(nik, noHp string) (exist bool, err error)
	RegisterNomorRekening(nomorRekening string, idRekening int) (err error)
	GetRekeningByNomorRekening(noRekening string) (rekening data.Rekening, err error)
	Saving(saving data.Transaksi) (err error)
	UpdateSavingSaldo(nomorRekening string, nominal float64) (saldo float64, err error)
	CashWithdrawl(cashWithdrawl data.Transaksi) (err error)
	UpdateCashWithdrawlSaldo(nomorRekening string, nominal float64) (saldo float64, err error)
	Mutation(nomorRekening string) (transaksi []data.Transaksi, err error)
}

type TabunganHandlerPort interface {
	Register(c *fiber.Ctx) error
	Saving(c *fiber.Ctx) error
	CashWithdrawl(c *fiber.Ctx) error
	Balance(c *fiber.Ctx) error
	Mutation(c *fiber.Ctx) error
}
