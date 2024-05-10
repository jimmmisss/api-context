package api

import (
	"context"
	"encoding/json"
	"github.com/jimmmisss/client-api-context/internal/model"
	"io"
	"log"
	"net/http"
	"time"
)

type CotacaoAPIServer struct {
}

func NewCotacaoAPIServer() *CotacaoAPIServer {
	return &CotacaoAPIServer{}
}

func (c *CotacaoAPIServer) GetCotacaoAPI() (*model.Bid, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	url := "http://localhost:8080/cotacao"
	cotacao, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Printf("Error: %v", err)
		panic(err)
	}

	res, err := http.DefaultClient.Do(cotacao)
	if err != nil {
		log.Printf("Error: %v", err)
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error: %v", err)
		panic(err)
	}

	var bid model.Bid
	err = json.Unmarshal(body, &bid)
	if err != nil {
		log.Printf("Error: %v", err)
		panic(err)
	}

	return &bid, nil
}
