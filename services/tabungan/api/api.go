package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/kurdilesmana/go-tabungan-api/pkg/logging"
	"github.com/kurdilesmana/go-tabungan-api/services/tabungan/app"
	"github.com/kurdilesmana/go-tabungan-api/services/tabungan/data"
	"github.com/sirupsen/logrus"
)

type TabunganAPI struct {
	app       app.TabunganServicePort
	log       *logging.Logger
	validator *validator.Validate
}

func (t *TabunganAPI) Register(c *fiber.Ctx) (err error) {
	registerRekening := data.RegisterRekening{}

	// parse body, attach to request struct
	if err = c.BodyParser(&registerRekening); err != nil {
		remark := "failed to parse request to register rekening"
		t.log.Error(logrus.Fields{"error": err.Error()}, registerRekening, remark)
		err = fiber.NewError(fiber.StatusBadRequest, err.Error())
		return
	}

	// validate data request
	if err = t.validator.Struct(registerRekening); err != nil {
		remark := "failed to validate request register rekening"
		t.log.Error(logrus.Fields{"error": err.Error()}, registerRekening, remark)
		err = c.Status(fiber.StatusBadRequest).JSON(map[string]string{"remark": err.Error()})
		return
	}

	// do register rekening
	response := data.RegisterRekeningResponse{}
	response.NoRekening, err = t.app.RegisterRekening(registerRekening)
	if err != nil {
		remark := "failed to register rekening"
		t.log.Error(logrus.Fields{"error": err.Error()}, registerRekening, remark)
		err = c.Status(fiber.StatusBadRequest).JSON(map[string]string{"remark": err.Error()})
		return
	}

	return c.Status(fiber.StatusOK).JSON(&response)
}

func (t *TabunganAPI) Saving(c *fiber.Ctx) (err error) {
	saving := data.Saving{}

	// parse body, attach to request struct
	if err = c.BodyParser(&saving); err != nil {
		remark := "failed to parse request to saving"
		t.log.Error(logrus.Fields{"error": err.Error()}, saving, remark)
		err = c.Status(fiber.StatusBadRequest).JSON(map[string]string{"remark": err.Error()})
		return
	}

	// validate data request
	if err = t.validator.Struct(saving); err != nil {
		remark := "failed to validate request saving rekening"
		t.log.Error(logrus.Fields{"error": err.Error()}, saving, remark)
		err = c.Status(fiber.StatusBadRequest).JSON(map[string]string{"remark": err.Error()})
		return
	}

	// do register rekening
	response := data.SavingResponse{}
	response.Saldo, err = t.app.Saving(saving)
	if err != nil {
		remark := "failed to register rekening"
		t.log.Error(logrus.Fields{"error": err.Error()}, saving, remark)
		err = c.Status(fiber.StatusBadRequest).JSON(map[string]string{"remark": err.Error()})
		return
	}

	return c.Status(fiber.StatusOK).JSON(&response)
}

func (t *TabunganAPI) CashWithdrawl(c *fiber.Ctx) (err error) {
	cashWithdrawl := data.CashWithdrawl{}

	// parse body, attach to request struct
	if err = c.BodyParser(&cashWithdrawl); err != nil {
		remark := "failed to parse request to cashWithdrawl"
		t.log.Error(logrus.Fields{"error": err.Error()}, cashWithdrawl, remark)
		err = c.Status(fiber.StatusBadRequest).JSON(map[string]string{"remark": err.Error()})
		return
	}

	// validate data request
	if err = t.validator.Struct(cashWithdrawl); err != nil {
		remark := "failed to validate request cashWithdrawl rekening"
		t.log.Error(logrus.Fields{"error": err.Error()}, cashWithdrawl, remark)
		err = c.Status(fiber.StatusBadRequest).JSON(map[string]string{"remark": err.Error()})
		return
	}

	// do register rekening
	response := data.CashWithdrawlResponse{}
	response.Saldo, err = t.app.CashWithdrawl(cashWithdrawl)
	if err != nil {
		remark := "failed to cashWithdrawl rekening"
		t.log.Error(logrus.Fields{"error": err.Error()}, cashWithdrawl, remark)
		err = c.Status(fiber.StatusBadRequest).JSON(map[string]string{"remark": err.Error()})
		return
	}

	return c.Status(fiber.StatusOK).JSON(&response)
}

func (t *TabunganAPI) Balance(c *fiber.Ctx) (err error) {
	balance := data.Balance{
		NoRekening: c.Params("no_rekening"),
	}

	// validate data request
	if err = t.validator.Struct(balance); err != nil {
		remark := "failed to validate request balance rekening"
		t.log.Error(logrus.Fields{"error": err.Error()}, balance, remark)
		err = c.Status(fiber.StatusBadRequest).JSON(map[string]string{"remark": err.Error()})
		return
	}

	// do register rekening
	response := data.BalanceResponse{}
	response.Saldo, err = t.app.Balance(balance)
	if err != nil {
		remark := "failed to get balance rekening"
		t.log.Error(logrus.Fields{"error": err.Error()}, balance, remark)
		err = c.Status(fiber.StatusBadRequest).JSON(map[string]string{"remark": err.Error()})
		return
	}

	return c.Status(fiber.StatusOK).JSON(&response)
}

func (t *TabunganAPI) Mutation(c *fiber.Ctx) (err error) {
	mutation := data.Mutation{
		NoRekening: c.Params("no_rekening"),
	}

	// validate data request
	if err = t.validator.Struct(mutation); err != nil {
		remark := "failed to validate request register rekening"
		t.log.Error(logrus.Fields{"error": err.Error()}, mutation, remark)
		err = c.Status(fiber.StatusBadRequest).JSON(map[string]string{"remark": err.Error()})
		return
	}

	// do register rekening
	transactions, err := t.app.Mutation(mutation)
	if err != nil {
		remark := "failed to get mutation rekening"
		t.log.Error(logrus.Fields{"error": err.Error()}, mutation, remark)
		err = c.Status(fiber.StatusBadRequest).JSON(map[string]string{"remark": err.Error()})
		return
	}

	response := data.MutationResponse{}
	response.Mutasi = []data.TransaksiMutasi{}
	err = copier.Copy(&response.Mutasi, transactions)
	if err != nil {
		remark := "failed to copy response mutation rekening"
		t.log.Error(logrus.Fields{"error": err.Error()}, mutation, remark)
		err = c.Status(fiber.StatusBadRequest).JSON(map[string]string{"remark": err.Error()})
		return
	}

	return c.Status(fiber.StatusOK).JSON(&response)
}

func InitTabunganAPI(app app.TabunganServicePort, log *logging.Logger) *TabunganAPI {
	return &TabunganAPI{
		app:       app,
		log:       log,
		validator: validator.New(),
	}
}
