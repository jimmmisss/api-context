//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/jimmmisss/server-api-context/internal/db"
	"github.com/jimmmisss/server-api-context/internal/handler"
)

func Wire(db *db.ConnDB) *handler.CotacaoHandler {
	wire.Build(ProviderSet)
	return nil
}
