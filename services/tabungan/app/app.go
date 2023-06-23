package app

import (
	"fmt"

	"github.com/kurdilesmana/go-tabungan-api/pkg/logging"
	"github.com/kurdilesmana/go-tabungan-api/services/tabungan/data"
	"github.com/sirupsen/logrus"
)

type TabunganApplication struct {
	datastore TabunganDatastorePort
	log       *logging.Logger
}

func (t *TabunganApplication) RegisterRekening(registerRekening data.RegisterRekening) (nomor_rekening string, err error) {
	// Initiate transaction
	tx := t.datastore.Begin()

	// validate rekening
	isExist, err := t.datastore.CheckRekeningExist(registerRekening.Nik, registerRekening.NoHp)
	if err != nil {
		err = fmt.Errorf("failed to check exist rekening")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	if isExist {
		err = fmt.Errorf("failed to create, rekening exist")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	var rekening data.Rekening
	rekening.Nama = registerRekening.Nama
	rekening.Nik = registerRekening.Nik
	rekening.NoHp = registerRekening.NoHp

	// do insert new rekening
	IDRekening, err := t.datastore.RegisterRekening(rekening)
	if err != nil {
		err = fmt.Errorf("failed to create rekening")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	// do update nomor_rekening
	nomor_rekening = fmt.Sprintf("%05d", IDRekening)
	err = t.datastore.RegisterNomorRekening(nomor_rekening, IDRekening)
	if err != nil {
		err = fmt.Errorf("failed to create rekening")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	// commit transacton
	t.datastore.Commit(tx)

	return
}

func (t *TabunganApplication) Saving(saving data.Saving) (saldo float64, err error) {
	// Initiate transaction
	tx := t.datastore.Begin()

	// validate rekening
	isExist, err := t.datastore.CheckRekeningExistByNomorRekening(saving.NoRekening)
	if err != nil {
		err = fmt.Errorf("failed to check exist rekening by nomor rekening")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	if isExist {
		err = fmt.Errorf("failed to create, rekening exist")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	var transaksi data.Transaksi
	transaksi.NoRekening = saving.NoRekening
	transaksi.Nominal = saving.Nominal
	transaksi.KodeTransaksi = "C"

	// do insert new rekening
	err = t.datastore.Saving(transaksi)
	if err != nil {
		err = fmt.Errorf("failed to saving")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	return
}

func InitApplication(datastore TabunganDatastorePort, log *logging.Logger) *TabunganApplication {
	return &TabunganApplication{
		datastore: datastore,
		log:       log,
	}
}
