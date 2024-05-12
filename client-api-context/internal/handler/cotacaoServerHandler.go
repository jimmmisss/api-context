package handler

import (
	"encoding/json"
	"github.com/jimmmisss/client-api-context/internal/service"
	"net/http"
)

type CotacaoServerHandler struct {
	service *service.CotacaoServerService
}

func NewCotacaoServerHandler(service *service.CotacaoServerService) *CotacaoServerHandler {
	return &CotacaoServerHandler{service: service}
}

func (c *CotacaoServerHandler) CotacaoServerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bid, err := c.service.CotacaoUSDBRLService()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bid)
}
