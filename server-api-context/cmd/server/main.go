package main

import (
	"github.com/jimmmisss/server-api-context/internal/api"
	"github.com/jimmmisss/server-api-context/internal/db"
	"github.com/jimmmisss/server-api-context/internal/handler"
	"github.com/jimmmisss/server-api-context/internal/repository"
	"github.com/jimmmisss/server-api-context/internal/service"
	"net/http"
)

func main() {
	conn, err := db.NewConnectDB()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	repository := repository.NewCotacaoRepository(conn.DB)
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
