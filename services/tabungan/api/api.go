package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
		err = fiber.NewError(fiber.StatusBadRequest, err.Error())
		return
	}

	// do register rekening
	response := data.RegisterRekeningResponse{}
	response.NoRekening, err = t.app.RegisterRekening(registerRekening)
	if err != nil {
		remark := "failed to register rekening"
		t.log.Error(logrus.Fields{"error": err.Error()}, registerRekening, remark)
		err = fiber.NewError(fiber.StatusNotFound, err.Error())
		return
	}

	return c.Status(fiber.StatusCreated).JSON(&response)
}

func (t *TabunganAPI) Saving(c *fiber.Ctx) (err error) {
	saving := data.Saving{}

	// parse body, attach to request struct
	if err = c.BodyParser(&saving); err != nil {
		remark := "failed to parse request to saving"
		t.log.Error(logrus.Fields{"error": err.Error()}, saving, remark)
		err = fiber.NewError(fiber.StatusBadRequest, err.Error())
		return
	}

	// validate data request
	if err = t.validator.Struct(saving); err != nil {
		remark := "failed to validate request register rekening"
		t.log.Error(logrus.Fields{"error": err.Error()}, saving, remark)
		err = fiber.NewError(fiber.StatusBadRequest, err.Error())
		return
	}

	// do register rekening
	response := data.SavingRespoonse{}
	response.Saldo, err = t.app.Saving(saving)
	if err != nil {
		remark := "failed to register rekening"
		t.log.Error(logrus.Fields{"error": err.Error()}, saving, remark)
		err = fiber.NewError(fiber.StatusNotFound, err.Error())
		return
	}

	return c.Status(fiber.StatusCreated).JSON(&response)
}

func InitTabunganAPI(app app.TabunganServicePort, log *logging.Logger) *TabunganAPI {
	return &TabunganAPI{
		app:       app,
		log:       log,
		validator: validator.New(),
	}
}
