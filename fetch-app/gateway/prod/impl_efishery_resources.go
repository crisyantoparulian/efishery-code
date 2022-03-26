package prod

import (
	"context"
	"errors"
	"fetch-app/domain/entity"
	"fetch-app/infrastructure/config"
	"fetch-app/infrastructure/log"

	jsoniter "github.com/json-iterator/go"
)

type EfisheryResourcesImpl struct {
	Config *config.Config
}

func (r *EfisheryResourcesImpl) FindResources(ctx context.Context) ([]*entity.Resources, error) {
	log.Info(ctx, "FindResources")

	result := make([]*entity.Resources, 0)

	req := newReqGeneralResty()
	resp, err := req.Get(r.Config.EfisheryApi.FetchResources)
	if err != nil {
		log.Error(ctx, "Error FindResources Get fail : ", err)
		return result, err
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		log.Error(ctx, "Error FindResources unmarshal fail : ", err)
		return result, err
	}

	if resp.StatusCode() != 200 {
		log.Error(ctx, "Error FindResources not 200 : ", resp)
		return result, errors.New("Error GeneralNotifWithBahasa not Ok")
	}

	return result, nil
}
