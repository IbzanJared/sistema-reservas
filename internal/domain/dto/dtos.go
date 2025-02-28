package dto

type ReservaDTO struct {
	UsuarioID     uint   `json:"usuario_id"`
	RestauranteID uint   `json:"restaurante_id"`
	MesaID        uint   `json:"mesa_id"`
	Fecha         string `json:"fecha"`
	Hora          string `json:"hora"`
	Estado        string `json:"estado"`
	HorarioID     uint   `json:"horario_id"`
}
