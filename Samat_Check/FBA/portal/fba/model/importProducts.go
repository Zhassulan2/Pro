package model

//modelRowsCount структура для количества записей под swagger
//swagger:model Count
type ImportProduct struct {
	Name         string
	BarCode      string
	PriceBuy     float64
	PriceSell    float64
	Count        int64
	CategoryName string
}
