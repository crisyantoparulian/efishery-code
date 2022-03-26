package fetchresourcewithprice

import "fetch-app/domain/repository"

// Outport of usecase
type Outport interface {
	repository.FindResourcesRepo
	repository.CurrencyConverter
}
