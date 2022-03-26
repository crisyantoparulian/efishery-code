package prod

import (
	"context"
	"errors"
	"fetch-app/domain/vo"
	"fetch-app/infrastructure/config"
	"fetch-app/infrastructure/log"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type CurrencyImpl struct {
	Config *config.Config
}

func (r *CurrencyImpl) CurrencyConverter(ctx context.Context, currentCurrency, targetCurrency vo.Currency) (map[string]float64, error) {
	log.Info(ctx, "CurrencyConverter")

	q := currentCurrency + "_" + targetCurrency
	url := fmt.Sprintf(r.Config.CurrencyConverter.Url, q, r.Config.CurrencyConverter.ApiKey)

	result := map[string]float64{}

	req := newReqGeneralResty()
	resp, err := req.Get(url)
	if err != nil {
		log.Error(ctx, "Error CurrencyConverter Get fail : ", err)
		return result, err
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		log.Error(ctx, "Error CurrencyConverter unmarshal fail : ", err)
		return result, err
	}

	if resp.StatusCode() != 200 {
		log.Error(ctx, "Error CurrencyConverter not 200 : ", resp)
		return result, errors.New("Error CurrencyConverter not Ok")
	}

	return result, nil
}
