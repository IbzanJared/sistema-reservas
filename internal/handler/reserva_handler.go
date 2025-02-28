package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/IbzanJared/sistema-reservas/internal/domain/dto"
	"github.com/IbzanJared/sistema-reservas/internal/domain/model"
	"github.com/IbzanJared/sistema-reservas/internal/service"
	"github.com/go-chi/chi/v5"
)

type ReservaHandler struct {
	service service.ReservaServiceInterface
}

func NewReservaHandler(service service.ReservaServiceInterface) *ReservaHandler {
	return &ReservaHandler{service}
}

func (h *ReservaHandler) CrearReserva(w http.ResponseWriter, r *http.Request) {
	var reservaDTO dto.ReservaDTO

	if err := json.NewDecoder(r.Body).Decode(&reservaDTO); err != nil {
		http.Error(w, "Error en el formato de la solicitud", http.StatusBadRequest)
		return
	}

	if reservaDTO.UsuarioID == 0 || reservaDTO.RestauranteID == 0 || reservaDTO.MesaID == 0 || reservaDTO.HorarioID == 0 {
		http.Error(w, "IDs no pueden ser 0", http.StatusBadRequest)
		return
	}

	reserva := model.Reserva{
		UsuarioID:     reservaDTO.UsuarioID,
		RestauranteID: reservaDTO.RestauranteID,
		MesaID:        reservaDTO.MesaID,
		Fecha:         reservaDTO.Fecha,
		Hora:          reservaDTO.Hora,
		Estado:        reservaDTO.Estado,
		HorarioID:     reservaDTO.HorarioID,
	}

	if err := h.service.CrearReserva(&reserva); err != nil {
		http.Error(w, "No se pudo crear la reserva", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reserva)
}

func (h *ReservaHandler) ObtenerReservas(w http.ResponseWriter, r *http.Request) {
	reservas, err := h.service.ListarReservas()
	if err != nil {
		http.Error(w, "Error obteniendo las reservas", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(reservas)
}

func (h *ReservaHandler) ObtenerReservaPorID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	reserva, err := h.service.BuscarReserva(uint(id))
	if err != nil {
		http.Error(w, "Reserva no encontrada", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(reserva)
}

func (h *ReservaHandler) ModificarReserva(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var reservaDTO dto.ReservaDTO
	if err := json.NewDecoder(r.Body).Decode(&reservaDTO); err != nil {
		http.Error(w, "Error en el formato de la solicitud", http.StatusBadRequest)
		return
	}

	if reservaDTO.UsuarioID == 0 || reservaDTO.RestauranteID == 0 || reservaDTO.MesaID == 0 || reservaDTO.HorarioID == 0 {
		http.Error(w, "IDs no pueden ser 0", http.StatusBadRequest)
		return
	}

	reserva := model.Reserva{
		ID:            uint(id),
		UsuarioID:     reservaDTO.UsuarioID,
		RestauranteID: reservaDTO.RestauranteID,
		MesaID:        reservaDTO.MesaID,
		Fecha:         reservaDTO.Fecha,
		Hora:          reservaDTO.Hora,
		Estado:        reservaDTO.Estado,
		HorarioID:     reservaDTO.HorarioID,
	}

	if err := h.service.ModificarReserva(&reserva); err != nil {
		http.Error(w, "No se pudo actualizar la reserva", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(reserva)
}

func (h *ReservaHandler) EliminarReserva(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	if err := h.service.EliminarReserva(uint(id)); err != nil {
		http.Error(w, "No se pudo eliminar la reserva", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
