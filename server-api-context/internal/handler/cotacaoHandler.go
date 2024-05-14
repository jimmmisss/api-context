package handler

import (
	"encoding/json"
	"github.com/jimmmisss/server-api-context/internal/port"
	"log"
	"net/http"
)

type CotacaoHandler struct {
	Service port.CotacaoServiceInterface
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
