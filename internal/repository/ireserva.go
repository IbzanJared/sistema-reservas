package repository

import "github.com/IbzanJared/sistema-reservas/internal/domain/model"

type ReservaRepositoryInterface interface {
	Crear(reserva *model.Reserva) error
	ObtenerTodas() ([]model.Reserva, error)
	ObtenerPorID(id uint) (*model.Reserva, error)
	Actualizar(reserva *model.Reserva) error
	Eliminar(id uint) error
}
