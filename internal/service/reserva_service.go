package service

import (
	"github.com/IbzanJared/sistema-reservas/internal/domain/model"
	"github.com/IbzanJared/sistema-reservas/internal/repository"
)

type ReservaService struct {
	repo repository.ReservaRepositoryInterface
}

func NewReservaService(repo repository.ReservaRepositoryInterface) *ReservaService {
	return &ReservaService{repo}
}

func (s *ReservaService) CrearReserva(reserva *model.Reserva) error {
	return s.repo.Crear(reserva)
}

func (s *ReservaService) ListarReservas() ([]model.Reserva, error) {
	return s.repo.ObtenerTodas()
}

func (s *ReservaService) BuscarReserva(id uint) (*model.Reserva, error) {
	return s.repo.ObtenerPorID(id)
}

func (s *ReservaService) ModificarReserva(reserva *model.Reserva) error {
	return s.repo.Actualizar(reserva)
}

func (s *ReservaService) EliminarReserva(id uint) error {
	return s.repo.Eliminar(id)
}
