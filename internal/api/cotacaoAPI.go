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
	repository model.CotacaoRepository
}

func NewCotacaoAPI(repository model.CotacaoRepository) *CotacaoAPI {
	return &CotacaoAPI{repository: repository}
}

func (s *CotacaoAPI) GetCotacaoAPI() (*model.CotacaoDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	cotacao, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/all/USD-BRL", nil)
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

	var model model.CotacaoDTO
	err = json.Unmarshal(body, &model)
	if err != nil {
		panic(err)
	}

	return &model, nil
}
