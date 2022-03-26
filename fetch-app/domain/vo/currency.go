package vo

type Currency string

type CurrencyRequest struct {
}

func NewCurrency(req CurrencyRequest) (Currency, error) {

	var obj Currency

	return obj, nil
}

const (
	CurrencyUSD Currency = "USD"
	CurrencyIDR Currency = "IDR"
)

func (r Currency) Validate() error {
	return nil
}

func (r Currency) String() string {
	return string(r)
}
