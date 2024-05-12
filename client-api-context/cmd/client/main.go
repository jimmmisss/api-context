package main

import (
	"github.com/jimmmisss/client-api-context/internal/api"
	"github.com/jimmmisss/client-api-context/internal/handler"
	"github.com/jimmmisss/client-api-context/internal/service"
	"net/http"
)

func main() {

	apiServer := api.NewCotacaoAPIServer()
	NewGeraArquivo := service.NewGeraArquivoTxtService()
	svc := service.NewCotacaoServerService(apiServer, NewGeraArquivo)

	controller := handler.NewCotacaoServerHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", controller.CotacaoServerHandler)

	err := http.ListenAndServe("127.0.0.1:8888", mux)
	if err != nil {
		panic(err)
	}
}
