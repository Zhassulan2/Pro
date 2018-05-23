package model

type TicketlinesInput struct {
	Id                      TicketLinePK
	Product                 *string `json:"product"`
	Attributesetinstance_id *string `json:"attributesetinstance_id"`
	Units                   float64 `json:"units"`
	Price                   float64 `json:"price"`
	Taxid                   string  `json:"taxid"`
	Attributes              *string `json:"attributes"`
}

type TicketLinePK struct {
	Ticket string `json:"ticket"`
	Line   *int   `json:"line"`
}

type Ticketlines struct {
	Ticket                  string `gorm:"primary_key"`
	Line                    *int   `gorm:"primary_key"`
	Product                 *string
	Attributesetinstance_id *string
	Units                   float64
	Price                   float64
	Taxid                   string
	Attributes              *string
}
