package service

import "github.com/IbzanJared/sistema-reservas/internal/domain/model"

type ReservaServiceInterface interface {
	CrearReserva(reserva *model.Reserva) error
	ListarReservas() ([]model.Reserva, error)
	BuscarReserva(id uint) (*model.Reserva, error)
	ModificarReserva(reserva *model.Reserva) error
	EliminarReserva(id uint) error
}
