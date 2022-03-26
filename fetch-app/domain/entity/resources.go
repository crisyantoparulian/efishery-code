package entity

type Resources struct {
	Uuid         *string `json:"uuid"`
	Komoditas    *string `json:"komoditas"`
	AreaProvinsi *string `json:"area_provinsi"`
	AreaKota     *string `json:"area_kota"`
	Size         *string `json:"size"`
	Price        *string `json:"price"`
	TglParsed    *string `json:"tgl_parsed"`
	Timestamp    *string `json:"timestamp"`
}

type ResourcesRequest struct {
}

func NewResources(req ResourcesRequest) (*Resources, error) {

	var obj Resources

	// assign value here

	err := obj.Validate()
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (r *Resources) Validate() error {
	return nil
}
