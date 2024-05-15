package handler

import (
	"encoding/json"
	"github.com/jimmmisss/server-api-context/internal/model"
	"log"
	"net/http"
)

type CotacaoServiceInterface interface {
	ObtemCotacaoViaApiSalva() (*model.BidResponse, error)
}

type CotacaoHandler struct {
	Service CotacaoServiceInterface
}

func NewCotacaoHandler(service CotacaoServiceInterface) *CotacaoHandler {
	return &CotacaoHandler{Service: service}
}

func (c *CotacaoHandler) CotacaoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bid, err := c.Service.ObtemCotacaoViaApiSalva()
	if err != nil {
		log.Printf("================== ERRO NA API - CotacaoHandler: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bid)
}
