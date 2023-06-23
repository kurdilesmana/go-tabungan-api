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

	tabunganRoutes := app.Group("/tabungan")
	tabunganRoutes.Post("/daftar", s.api.Register)

	err := app.Listen(port)
	if err != nil {
		log.Fatal(err)
	}
}
