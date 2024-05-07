package handler

import (
	"api-context/internal/service"
	"encoding/json"
	"net/http"
)

type CotacaoHandler struct {
	service *service.CotacaoService
}

func NewCotacaoHandler(service *service.CotacaoService) *CotacaoHandler {
	return &CotacaoHandler{service: service}
}

func (c *CotacaoHandler) CotacaoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bid, err := c.service.ObtemCotacaoESalva()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bid)
}
