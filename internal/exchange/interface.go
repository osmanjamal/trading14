package exchange

type Exchange interface {
    PlaceOrder(symbol string, action string, price float64, amount float64) error
    GetBalance() (map[string]float64, error)
}