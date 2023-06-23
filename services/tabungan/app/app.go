package app

import (
	"fmt"
	"time"

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
	rekening, err := t.datastore.GetRekeningByNomorRekening(saving.NoRekening)
	if err != nil {
		err = fmt.Errorf("failed to check exist rekening by nomor rekening")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	if rekening == (data.Rekening{}) {
		err = fmt.Errorf("failed to create, rekening not exist")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	var transaksi data.Transaksi
	transaksi.NoRekening = saving.NoRekening
	transaksi.Nominal = saving.Nominal
	transaksi.KodeTransaksi = "C"
	transaksi.Waktu = time.Now()

	// do insert new rekening
	err = t.datastore.Saving(transaksi)
	if err != nil {
		err = fmt.Errorf("failed to saving")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	// do insert new rekening
	saldo, err = t.datastore.UpdateSavingSaldo(saving.NoRekening, saving.Nominal)
	if err != nil {
		err = fmt.Errorf("failed to update saldo saving")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	// commit transacton
	t.datastore.Commit(tx)

	return
}

func (t *TabunganApplication) CashWithdrawl(cashWithdrawl data.CashWithdrawl) (saldo float64, err error) {
	// Initiate transaction
	tx := t.datastore.Begin()

	// validate rekening
	rekening, err := t.datastore.GetRekeningByNomorRekening(cashWithdrawl.NoRekening)
	if err != nil {
		err = fmt.Errorf("failed to check exist rekening by nomor rekening")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	if rekening == (data.Rekening{}) {
		err = fmt.Errorf("failed to cashwithdrawl, rekening not found")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	// check saldo
	if cashWithdrawl.Nominal > rekening.Saldo {
		err = fmt.Errorf("failed to cashwithdrawl, saldo not enough")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	var transaksi data.Transaksi
	transaksi.NoRekening = cashWithdrawl.NoRekening
	transaksi.Nominal = cashWithdrawl.Nominal
	transaksi.KodeTransaksi = "D"
	transaksi.Waktu = time.Now()

	// do insert transaksi cashWithdrawl
	err = t.datastore.CashWithdrawl(transaksi)
	if err != nil {
		err = fmt.Errorf("failed to saving")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	// do update saldo cashWithdrawl
	saldo, err = t.datastore.UpdateCashWithdrawlSaldo(cashWithdrawl.NoRekening, cashWithdrawl.Nominal)
	if err != nil {
		err = fmt.Errorf("failed to update saldo cashwithdrawl")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	// commit transacton
	t.datastore.Commit(tx)

	return
}

func (t *TabunganApplication) Balance(balance data.Balance) (saldo float64, err error) {
	// Initiate transaction
	tx := t.datastore.Begin()

	// get rekening
	rekening, err := t.datastore.GetRekeningByNomorRekening(balance.NoRekening)
	if err != nil {
		err = fmt.Errorf("failed to check exist rekening by nomor rekening")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	if rekening == (data.Rekening{}) {
		err = fmt.Errorf("failed to cashwithdrawl, rekening not found")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	saldo = rekening.Saldo

	// commit transacton
	t.datastore.Commit(tx)

	return
}

func (t *TabunganApplication) Mutation(mutation data.Mutation) (transactions []data.Transaksi, err error) {
	// Initiate transaction
	tx := t.datastore.Begin()

	// get rekening
	rekening, err := t.datastore.GetRekeningByNomorRekening(mutation.NoRekening)
	if err != nil {
		err = fmt.Errorf("failed to check exist rekening by nomor rekening")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	if rekening == (data.Rekening{}) {
		err = fmt.Errorf("failed to mutation, rekening not found")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	transactions, err = t.datastore.Mutation(mutation.NoRekening)
	if err != nil {
		err = fmt.Errorf("failed get transactions by nomor rekening")
		t.log.Warn(logrus.Fields{}, nil, err.Error())
		t.datastore.Rollback(tx)
		return
	}

	// commit transacton
	t.datastore.Commit(tx)

	return
}

func InitApplication(datastore TabunganDatastorePort, log *logging.Logger) *TabunganApplication {
	return &TabunganApplication{
		datastore: datastore,
		log:       log,
	}
}
