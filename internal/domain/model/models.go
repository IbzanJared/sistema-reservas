package model

type Reserva struct {
	ID            uint   `gorm:"primaryKey"`
	UsuarioID     uint   `gorm:"not null"`
	RestauranteID uint   `gorm:"not null"`
	MesaID        uint   `gorm:"not null"`
	Fecha         string `gorm:"type:date;not null"`
	Hora          string `gorm:"type:time;not null"`
	Estado        string `gorm:"type:enum('pendiente', 'confirmada', 'cancelada', 'completada');default:'pendiente'"`
	HorarioID     uint   `gorm:"not null"`
}
