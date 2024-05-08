package api

import (
	"api-context/internal/model"
	"api-context/internal/port"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CotacaoAPI struct {
	api *port.CotacaoAPIService
}

func NewCotacaoAPI(api *port.CotacaoAPIService) *CotacaoAPI {
	return &CotacaoAPI{api: api}
}

func (c *CotacaoAPI) GetCotacaoAPI() (*model.CotacaoDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	cotacao, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(cotacao)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var cotacaoDTO model.CotacaoDTO
	err = json.Unmarshal(body, &cotacaoDTO)
	if err != nil {
		panic(err)
	}
	fmt.Println(cotacaoDTO)

	return &cotacaoDTO, nil
}
