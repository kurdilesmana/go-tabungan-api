package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kurdilesmana/go-tabungan-api/pkg/logging"
	"github.com/kurdilesmana/go-tabungan-api/services/tabungan/api"
)

type Server struct {
	api api.TabunganAPI
	log *logging.Logger
}

func InitServer(api api.TabunganAPI, log *logging.Logger) *Server {
	return &Server{
		api: api,
		log: log,
	}
}

func (s *Server) Start(port string) {
	app := fiber.New()
	app.Post("/daftar", s.api.Register)
	app.Post("/tabung", s.api.Saving)
	app.Post("/tarik", s.api.CashWithdrawl)
	app.Get("/saldo/:no_rekening", s.api.Balance)
	app.Get("/mutasi/:no_rekening", s.api.Mutation)

	err := app.Listen(port)
	if err != nil {
		log.Fatal(err)
	}
}
