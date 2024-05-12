package service

import (
	"github.com/jimmmisss/client-api-context/internal/model"
	"github.com/jimmmisss/client-api-context/internal/port"
)

type CotacaoServerService struct {
	api     port.CotacaoAPIInterface
	geraTxt port.CotacaoCriaArquivoInterface
}

func NewCotacaoServerService(
	api port.CotacaoAPIInterface,
	geraTxt port.CotacaoCriaArquivoInterface) *CotacaoServerService {
	return &CotacaoServerService{api: api, geraTxt: geraTxt}
}

func (c *CotacaoServerService) CotacaoUSDBRLService() (*model.Bid, error) {
	bid, err := c.api.GetCotacaoAPI()
	if err != nil {
		return nil, err
	}

	err = c.geraTxt.GeraArquivoTxt(bid)
	if err != nil {
		return nil, err
	}

	return bid, nil
}
