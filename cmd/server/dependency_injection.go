package server

import (
	reservaHandler "github.com/IbzanJared/sistema-reservas/internal/handler"
	"github.com/IbzanJared/sistema-reservas/internal/repository"
	"github.com/IbzanJared/sistema-reservas/internal/service"
	"gorm.io/gorm"
)

type Handlers struct {
	ReservaHandler *reservaHandler.ReservaHandler
}

func GetHandlers(db *gorm.DB) Handlers {

	repoReserva := repository.NewReservaRepository(db)

	// - service
	serviceReserva := service.NewReservaService(repoReserva)

	// - handler
	handlerReserva := reservaHandler.NewReservaHandler(serviceReserva)

	return Handlers{
		ReservaHandler: handlerReserva,
	}
}
