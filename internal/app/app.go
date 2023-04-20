package app

import (
	"avito_test_GO/config"
	v1 "avito_test_GO/internal/controller/http/v1"
	"avito_test_GO/internal/usecase"
	"avito_test_GO/internal/usecase/repo"
	"avito_test_GO/pkg/httpserver"
	"avito_test_GO/pkg/postgres"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	pg, err := postgres.New(cfg)
	if err != nil {
		log.Fatal("Cannot connect to Postgres")
	}

	us := usecase.NewUserUseCase(repo.NewUserRepo(pg))
	rp := usecase.NewReportUseCase(us)
	handler := gin.New()

	v1.NewRouter(handler,
		us,
		rp)

	serv := httpserver.New(handler, httpserver.Port(cfg.AppPort))
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interruption:
		log.Printf("signal: " + s.String())
	case err = <-serv.Notify():
		log.Printf("Notify from http server")
	}

	err = serv.Shutdown()
	if err != nil {
		log.Printf("Http server shutdown")
	}

}
