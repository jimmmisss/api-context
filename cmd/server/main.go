package main

import (
	"api-context/internal/api"
	"api-context/internal/db"
	"api-context/internal/handler"
	"api-context/internal/service"
	"net/http"
)

func main() {

	conn, err := db.NewConnectDB()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	repository := db.NewCotacaoRepository(conn.DB)
	apiService := api.NewCotacaoAPI()
	svc := service.NewCotacaoService(repository, apiService)
	controller := handler.NewCotacaoHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /cotacao", controller.CotacaoHandler)

	err = http.ListenAndServe("127.0.0.1:8080", mux)
	if err != nil {
		panic(err)
	}
}
