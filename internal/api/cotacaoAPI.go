package api

import (
	"api-context/internal/model"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type CotacaoAPI struct {
}

func NewCotacaoAPI() *CotacaoAPI {
	return &CotacaoAPI{}
}

func (c *CotacaoAPI) GetCotacaoAPI() (*model.CotacaoDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	url := "https://economia.awesomeapi.com.br/json/USD-BRL"
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

	var cotacaoDTO []model.CotacaoDTO
	err = json.Unmarshal(body, &cotacaoDTO)
	if err != nil {
		panic(err)
	}

	return &cotacaoDTO[0], nil
}
