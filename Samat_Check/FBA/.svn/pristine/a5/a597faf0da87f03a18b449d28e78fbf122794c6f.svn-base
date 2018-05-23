package main

import (
	"encoding/json"
	"fmt"
	"openbravosync/db"
	"openbravosync/model"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var settings model.Settings
var dbs map[string]gorm.DB

func loadConfiguration(file string) model.Settings {
	dbs = make(map[string]gorm.DB)
	var config model.Settings
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	for _, dbconf := range config.DataBases {
		fmt.Println("DB config", dbconf)
		dbs[dbconf.Key], _ = DataBaseManager.GetDb(dbconf)
	}

	return config
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("key")
		db, ok := dbs[key]
		if !ok {
			fmt.Println("not key", key)
			c.JSON(404, "NO AUTH")
			c.Abort()
			return
		}
		c.Set("db", &db)
		c.Next()
		return
	}
}

func main() {
	settings = loadConfiguration("settings.json")
	applicationResource := DataBaseManager.ApplicationResource{}
	attributeResource := DataBaseManager.AttributeResource{}
	attributeinstanceResource := DataBaseManager.AttributeinstanceResource{}
	attributesetResource := DataBaseManager.AttributesetResource{}
	attributesetinstanceResource := DataBaseManager.AttributesetinstanceResource{}
	attributeuseResource := DataBaseManager.AttributeuseResource{}
	attributevalueResource := DataBaseManager.AttributevalueResource{}
	categoriesResource := DataBaseManager.CategoriesResource{}
	closedcashResource := DataBaseManager.ClosedcashResource{}
	customerResource := DataBaseManager.CustomerResource{}
	floorsResource := DataBaseManager.FloorsResource{}
	locationsResource := DataBaseManager.LocationsResource{}
	paymentsResource := DataBaseManager.PaymentsResource{}
	peopleResource := DataBaseManager.PeopleResource{}
	placesResource := DataBaseManager.PlacesResource{}
	productsCatResource := DataBaseManager.Products_CatResource{}
	productsComResource := DataBaseManager.Products_ComResource{}
	productResource := DataBaseManager.ProductResource{}
	receiptResource := DataBaseManager.ReceiptResource{}
	reservationCustomersResource := DataBaseManager.ReservationCustomersResource{}
	reservationsResource := DataBaseManager.ReservationsResource{}
	resourcesResource := DataBaseManager.ResourcesResource{}
	roleResource := DataBaseManager.RoleResource{}
	sharedticketsResource := DataBaseManager.SharedticketsResource{}
	stockcurrentResource := DataBaseManager.StockcurrentResource{}
	stockdiaryResource := DataBaseManager.StockdiaryResource{}
	stocklevelResource := DataBaseManager.StocklevelResource{}
	taxcategoriesResource := DataBaseManager.TaxcategoriesResource{}
	taxcustcategoriesResource := DataBaseManager.TaxcustcategoriesResource{}
	taxesResource := DataBaseManager.TaxesResource{}
	taxlinesResource := DataBaseManager.TaxlinesResource{}
	thirdpartiesResource := DataBaseManager.ThirdpartiesResource{}
	ticketResource := DataBaseManager.TicketResource{}
	ticketlinesResource := DataBaseManager.TicketLineResource{}

	teambonusconfighostsResource := DataBaseManager.TeambonusconfighostsResource{}
	teambonussessionmoneyResource := DataBaseManager.TeambonussessionmoneyResource{}
	teambonussessionpersonResource := DataBaseManager.TeambonussessionpersonResource{}
	teambonussessionResource := DataBaseManager.TeambonussessionResource{}

	r := gin.Default()

	base := r.Group("/", auth())

	base.POST("applications", applicationResource.Create)
	base.DELETE("applications/:id", applicationResource.Delete)

	base.POST("attribute", attributeResource.Create)
	base.DELETE("attribute/:id", attributeResource.Delete)

	base.POST("attributeinstance", attributeinstanceResource.Create)
	base.DELETE("attributeinstance/:id", attributeinstanceResource.Delete)

	base.POST("attributeset", attributesetResource.Create)
	base.DELETE("attributeset/:id", attributesetResource.Delete)

	base.POST("attributesetinstance", attributesetinstanceResource.Create)
	base.DELETE("attributesetinstance/:id", attributesetinstanceResource.Delete)

	base.POST("attributeuse", attributeuseResource.Create)
	base.DELETE("attributeuse/:id", attributeuseResource.Delete)

	base.POST("attributevalue", attributevalueResource.Create)
	base.DELETE("attributevalue/:id", attributevalueResource.Delete)

	base.POST("categories", categoriesResource.Create)
	base.DELETE("categories/:id", categoriesResource.Delete)
	base.GET("categories", categoriesResource.Get)

	base.POST("closedcash", closedcashResource.Create)
	base.DELETE("closedcash/:id", closedcashResource.Delete)

	base.POST("customers", customerResource.Create)
	base.DELETE("customers/:id", customerResource.Delete)

	base.POST("floors", floorsResource.Create)
	base.DELETE("floors/:id", floorsResource.Delete)

	base.POST("locations", locationsResource.Create)
	base.DELETE("locations/:id", locationsResource.Delete)

	base.POST("payments", paymentsResource.Create)
	base.DELETE("payments/:id", paymentsResource.Delete)

	base.POST("people", peopleResource.Create)
	base.DELETE("people/:id", peopleResource.Delete)

	base.POST("places", placesResource.Create)
	base.DELETE("places/:id", placesResource.Delete)

	base.POST("products_cat", productsCatResource.Create)
	base.DELETE("products_cat/:id", productsCatResource.Delete)
	base.GET("products_cat", productsCatResource.Get)

	base.POST("products_com", productsComResource.Create)
	base.DELETE("products_com/:id", productsComResource.Delete)
	base.GET("products_com", productsComResource.Get)

	base.POST("products", productResource.Create)
	base.DELETE("products/:id", productResource.Delete)
	base.GET("products", productResource.Get)

	base.POST("receipts", receiptResource.Create)
	base.DELETE("receipts/:id", receiptResource.Delete)

	base.POST("reservationcustomers", reservationCustomersResource.Create)
	base.DELETE("reservationcustomers/:id", reservationCustomersResource.Delete)

	base.POST("reservations", reservationsResource.Create)
	base.DELETE("reservations/:id", reservationsResource.Delete)

	base.POST("resources", resourcesResource.Create)
	base.DELETE("resources/:id", resourcesResource.Delete)

	base.POST("roles", roleResource.Create)
	base.DELETE("roles/:id", roleResource.Delete)

	base.POST("sharedtickets", sharedticketsResource.Create)
	base.DELETE("sharedtickets/:id", sharedticketsResource.Delete)

	base.POST("stockcurrent", stockcurrentResource.Create)
	base.DELETE("stockcurrent/:id", stockcurrentResource.Delete)

	base.POST("stockdiary", stockdiaryResource.Create)
	base.DELETE("stockdiary/:id", stockdiaryResource.Delete)
	base.GET("stockdeary", stockdiaryResource.Get)
	base.GET("stockdeary/:date", stockdiaryResource.Get)

	base.POST("stocklevel", stocklevelResource.Create)
	base.DELETE("stocklevel/:id", stocklevelResource.Delete)

	base.POST("taxcategories", taxcategoriesResource.Create)
	base.DELETE("taxcategories/:id", taxcategoriesResource.Delete)
	base.GET("taxcategories", taxcategoriesResource.Get)

	base.POST("taxcustcategories", taxcustcategoriesResource.Create)
	base.DELETE("taxcustcategories/:id", taxcustcategoriesResource.Delete)

	base.POST("taxes", taxesResource.Create)
	base.DELETE("taxes/:id", taxesResource.Delete)

	base.POST("taxlines", taxlinesResource.Create)
	base.DELETE("taxlines/:id", taxlinesResource.Delete)

	base.POST("thirdparties", thirdpartiesResource.Create)
	base.DELETE("thirdparties/:id", thirdpartiesResource.Delete)

	base.POST("tickets", ticketResource.Create)
	base.DELETE("tickets/:id", ticketResource.Delete)

	base.POST("ticketlines", ticketlinesResource.Create)
	base.DELETE("ticketlines/:id", ticketlinesResource.Delete)

	//BINUSES
	base.POST("teambonusconfighosts", teambonusconfighostsResource.Create)
	base.DELETE("teambonusconfighosts/:id", teambonusconfighostsResource.Delete)

	base.POST("teambonussessionmoney", teambonussessionmoneyResource.Create)
	base.DELETE("teambonussessionmoney/:id", teambonussessionmoneyResource.Delete)

	base.POST("teambonussessionperson", teambonussessionpersonResource.Create)
	base.DELETE("teambonussessionperson/:id", teambonussessionpersonResource.Delete)

	base.POST("teambonussession", teambonussessionResource.Create)
	base.DELETE("teambonussession/:id", teambonussessionResource.Delete)

	base.GET("gettables", func(c *gin.Context) {
		key := c.GetHeader("key")
		for _, db := range settings.DataBases {
			if db.Key == key {
				c.String(200, "%s", db.GetTables)
			}
		}
	})
	base.GET("settables", func(c *gin.Context) {
		key := c.GetHeader("key")
		for _, db := range settings.DataBases {
			if db.Key == key {
				//c.JSON(200, db.SetTables)
				c.String(200, "%s", db.SetTables)
			}
		}
	})

	r.Run(":" + strconv.Itoa(settings.Port))
}
