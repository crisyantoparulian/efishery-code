package repository

import (
	"context"
	"fetch-app/domain/entity"
	"fetch-app/domain/vo"
)

type FindResourcesRepo interface {
	FindResources(ctx context.Context) ([]*entity.Resources, error)
}

type CurrencyConverter interface {
	CurrencyConverter(ctx context.Context, currentCurrency, targetCurrency vo.Currency) (map[string]float64, error)
}
