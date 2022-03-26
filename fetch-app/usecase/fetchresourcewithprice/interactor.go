package fetchresourcewithprice

import (
	"context"
	"fetch-app/domain/service"
	"fetch-app/domain/vo"
	"fetch-app/infrastructure/log"
	"fmt"
	"strconv"
	"time"
)

//go:generate mockery --name Outport -output mocks/

type fetchresourcewithpriceInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &fetchresourcewithpriceInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *fetchresourcewithpriceInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...

	resourcesObjs, err := r.outport.FindResources(ctx)
	if err != nil {
		return nil, err
	}

	currencyConvert := map[string]float64{}

	currencyKey := vo.CurrencyIDR.String() + "_" + vo.CurrencyUSD.String()
	cacheCurrency := service.GetCache(currencyKey)
	if cacheCurrency != nil {
		currencyConvert = cacheCurrency.(map[string]float64)
	}

	if _, ok := currencyConvert[currencyKey]; !ok {
		currencyConvert, err = r.outport.CurrencyConverter(ctx, vo.CurrencyIDR, vo.CurrencyUSD)
		if err != nil {
			return nil, err
		}
	}

	err = service.SetCache(vo.CurrencyIDR.String()+"_"+vo.CurrencyUSD.String(), currencyConvert, 1*time.Hour)
	if err != nil {
		log.Error(ctx, "Error SetCache ", err)
	}

	for _, obj := range resourcesObjs {
		var convertCurrency float64
		if d, ok := currencyConvert[currencyKey]; ok {
			convertCurrency = d
		}

		var strIdr, strUsd *string
		if obj.Price != nil {
			idrPrice, _ := strconv.ParseFloat(*obj.Price, 32)
			usdPrice := idrPrice * convertCurrency
			idr := fmt.Sprint(idrPrice)
			usd := fmt.Sprint(usdPrice)
			strIdr = &idr
			strUsd = &usd
		}

		res.Resources = append(res.Resources, &ResourceWithPrice{
			Uuid:      obj.Uuid,
			Komoditas: obj.Komoditas,
			Size:      obj.Size,
			TglParsed: obj.TglParsed,
			Timestamp: obj.Timestamp,
			Area: &Area{
				Provinsi: obj.AreaProvinsi,
				Kota:     obj.AreaKota,
			},
			Price: &Price{
				Idr: strIdr,
				Usd: strUsd,
			},
		})
	}

	return res, nil
}
