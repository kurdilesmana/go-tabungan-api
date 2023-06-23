package datastore

import (
	"fmt"

	"github.com/kurdilesmana/go-tabungan-api/pkg/logging"
	"github.com/kurdilesmana/go-tabungan-api/services/tabungan/data"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func (t *TabunganDatabase) CheckRekeningExistByNomorRekening(noRekening string) (exist bool, err error) {
	// check rekening exist
	var rekenings []data.Rekening
	res := t.db.Where(data.Rekening{NoRekening: noRekening}).Find(&rekenings)
	if res.Error != nil {
		remark := "failed to get rekening by nomor rekening"
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
	res := t.db.Save(&data.Rekening{Id: idRekening, NoRekening: nomorRekening})
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

func (t *TabunganDatabase) UpdateSavingSaldo(nomorRekening string, idRekening int) (err error) {
	// register nomor rekening
	res := t.db.Save(&data.Rekening{Id: idRekening, NoRekening: nomorRekening})
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
