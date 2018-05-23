package model

type Products struct {
	Id              string `gorm:"primary_key"`
	Reference       string
	Code            string
	Codetype        *string
	Name            string
	Pricebuy        float64
	Pricesell       float64
	Category        string
	Taxcat          string
	Attributeset_Id *string
	Stockcost       *float64
	Stockvolume     *float64
	Image           *string `type:"bytea"`
	Iscom           string
	Isscale         string
	Attributes      *string `type:"bytea"`
}
