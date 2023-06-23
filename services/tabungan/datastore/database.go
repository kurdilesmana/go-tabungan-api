package datastore

import (
	"fmt"

	"github.com/kurdilesmana/go-tabungan-api/pkg/logging"
	"github.com/kurdilesmana/go-tabungan-api/services/tabungan/data"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TabunganDatabase struct {
	db  *gorm.DB
	log *logging.Logger
}

func (t *TabunganDatabase) Begin() (tx *gorm.DB) {
	return t.db.Begin()
}

func (t *TabunganDatabase) Rollback(tx *gorm.DB) {
	tx.Rollback()
}

func (t *TabunganDatabase) Commit(tx *gorm.DB) {
	tx.Commit()
}

func (t *TabunganDatabase) RegisterRekening(rekening data.Rekening) (IDRekening int, err error) {
	// insert new rekening
	res := t.db.Create(&rekening)
	if res.Error != nil {
		remark := "failed to create rekening"
		t.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, nil, remark)
		err = fmt.Errorf(remark)
		return
	}

	// set value IDRekening
	IDRekening = rekening.Id

	return
}

func (t *TabunganDatabase) CheckRekeningExist(nik, noHp string) (exist bool, err error) {
	// check rekening exist
	var rekenings []data.Rekening
	res := t.db.Where(data.Rekening{Nik: nik}).Or(data.Rekening{NoHp: noHp}).Find(&rekenings)
	if res.Error != nil {
		remark := "failed to get rekening by nik nohp"
		t.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, nil, remark)
		err = fmt.Errorf(remark)
		return
	}

	if res.RowsAffected > 0 {
		exist = true
	}
	return
}

func (t *TabunganDatabase) RegisterNomorRekening(nomorRekening string, idRekening int) (err error) {
	// register nomor rekening
	res := t.db.Model(data.Rekening{}).Where("id = ?", idRekening).Updates(data.Rekening{NoRekening: nomorRekening})
	if res.Error != nil {
		remark := "failed to register nomor rekening"
		t.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, nil, remark)
		err = fmt.Errorf(remark)
		return
	}

	return
}

func (t *TabunganDatabase) GetRekeningByNomorRekening(noRekening string) (rekening data.Rekening, err error) {
	// check rekening exist
	res := t.db.Where(data.Rekening{NoRekening: noRekening}).Find(&rekening)
	if res.Error != nil {
		remark := "failed to get rekening by nomor rekening"
		t.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, nil, remark)
		err = fmt.Errorf(remark)
		return
	}
	return
}

func (t *TabunganDatabase) Saving(saving data.Transaksi) (err error) {
	// create saving transaksi
	res := t.db.Create(&saving)
	if res.Error != nil {
		remark := "failed to saving transaksi"
		t.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, nil, remark)
		err = fmt.Errorf(remark)
		return
	}
	return
}

func (t *TabunganDatabase) UpdateSavingSaldo(nomorRekening string, nominal float64) (saldo float64, err error) {
	var rekenings data.Rekening
	res := t.db.Model(&rekenings).Clauses(clause.Returning{Columns: []clause.Column{{Name: "saldo"}}}).Where(
		"no_rekening = ?", nomorRekening).UpdateColumn("saldo", gorm.Expr("saldo + ?", nominal))
	if res.Error != nil {
		remark := "failed to updata saldo saving rekening"
		t.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, nil, remark)
		err = fmt.Errorf(remark)
		return
	}
	saldo = rekenings.Saldo

	return
}

func (t *TabunganDatabase) CashWithdrawl(cashWithdrawl data.Transaksi) (err error) {
	// create cashWithdrawl transaksi
	res := t.db.Create(&cashWithdrawl)
	if res.Error != nil {
		remark := "failed to CashWithdrawl transaksi"
		t.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, nil, remark)
		err = fmt.Errorf(remark)
		return
	}
	return
}

func (t *TabunganDatabase) UpdateCashWithdrawlSaldo(nomorRekening string, nominal float64) (saldo float64, err error) {
	var rekenings data.Rekening
	res := t.db.Model(&rekenings).Clauses(clause.Returning{Columns: []clause.Column{{Name: "saldo"}}}).Where(
		"no_rekening = ?", nomorRekening).UpdateColumn("saldo", gorm.Expr("saldo - ?", nominal))
	if res.Error != nil {
		remark := "failed to update saldo cashwithdrawl rekening"
		t.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, nil, remark)
		err = fmt.Errorf(remark)
		return
	}
	saldo = rekenings.Saldo

	return
}

func (t *TabunganDatabase) Mutation(noRekening string) (transactions []data.Transaksi, err error) {
	// get transaksi rekening
	res := t.db.Where(data.Transaksi{NoRekening: noRekening}).Order("waktu desc").Find(&transactions)
	if res.Error != nil {
		remark := "failed to get transactions by nomor rekening"
		t.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, nil, remark)
		err = fmt.Errorf(remark)
		return
	}
	return
}

func Init(DBUser, DBPass, DBName, DBHost string, DBPort int, log *logging.Logger) *TabunganDatabase {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", DBUser, DBPass, DBHost, DBPort, DBName)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// init table database
	db.AutoMigrate(&data.Rekening{})
	db.AutoMigrate(&data.Transaksi{})

	return &TabunganDatabase{
		db:  db,
		log: log,
	}
}
