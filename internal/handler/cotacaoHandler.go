package handler

import (
	"api-context/internal/api"
	"encoding/json"
	"net/http"
)

func CotacaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
	}

	dolar, error := api.GetCambioDolar()
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dolar)
}
