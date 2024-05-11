package service

import (
	"github.com/jimmmisss/client-api-context/internal/api"
	"github.com/jimmmisss/client-api-context/internal/model"
)

type CotacaoServerService struct {
	api     api.CotacaoAPIServer
	geraTxt GeraArquivoTxtService
}

func NewCotacaoServerService(api api.CotacaoAPIServer, geraTxt GeraArquivoTxtService) *CotacaoServerService {
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
