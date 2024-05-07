package service

import (
	"api-context/internal/model"
)

type CotacaoService struct {
	repository model.CotacaoRepository
	apiService model.CotacaoAPI
}

func NewCotacaoService(repository model.CotacaoRepository, apiRepository model.CotacaoAPI) *CotacaoService {
	return &CotacaoService{
		repository: repository,
		apiService: apiRepository,
	}
}

func (c *CotacaoService) ObtemCotacaoESalva() (*model.BidResponse, error) {
	cotacaoDTO, err := c.apiService.GetCotacaoAPI()
	if err != nil {
		return nil, err
	}

	cotacao := &model.Cotacao{
		Code:       cotacaoDTO.Code,
		Codein:     cotacaoDTO.Codein,
		Name:       cotacaoDTO.Name,
		High:       cotacaoDTO.High,
		Low:        cotacaoDTO.Low,
		VarBid:     cotacaoDTO.VarBid,
		PctChange:  cotacaoDTO.PctChange,
		Bid:        cotacaoDTO.Bid,
		Ask:        cotacaoDTO.Ask,
		Timestamp:  cotacaoDTO.Timestamp,
		CreateDate: cotacaoDTO.CreateDate,
	}

	err = c.repository.Create(cotacao)
	if err != nil {
		return nil, err
	}

	bid := &model.BidResponse{
		Bid: cotacaoDTO.Bid,
	}

	return bid, nil
}
