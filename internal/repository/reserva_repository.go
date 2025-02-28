package repository

import (
	"github.com/IbzanJared/sistema-reservas/internal/domain/model"
	"gorm.io/gorm"
)

type ReservaRepository struct {
	db *gorm.DB
}

func NewReservaRepository(db *gorm.DB) *ReservaRepository {
	return &ReservaRepository{db}
}

func (r *ReservaRepository) Crear(reserva *model.Reserva) error {
	return r.db.Create(reserva).Error
}

func (r *ReservaRepository) ObtenerTodas() ([]model.Reserva, error) {
	var reservas []model.Reserva
	err := r.db.Find(&reservas).Error
	return reservas, err
}

func (r *ReservaRepository) ObtenerPorID(id uint) (*model.Reserva, error) {
	var reserva model.Reserva
	err := r.db.First(&reserva, id).Error
	return &reserva, err
}

func (r *ReservaRepository) Actualizar(reserva *model.Reserva) error {
	return r.db.Save(reserva).Error
}

func (r *ReservaRepository) Eliminar(id uint) error {
	return r.db.Delete(&model.Reserva{}, id).Error
}
