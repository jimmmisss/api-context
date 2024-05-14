package main

import (
	"github.com/jimmmisss/server-api-context/internal/config"
	"github.com/jimmmisss/server-api-context/internal/db"
	"net/http"
)

func main() {
	conn, err := db.NewConnectDB()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	handler := config.Wire(nil)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /cotacao", handler.CotacaoHandler)

	err = http.ListenAndServe("127.0.0.1:8080", mux)
	if err != nil {
		panic(err)
	}
}
