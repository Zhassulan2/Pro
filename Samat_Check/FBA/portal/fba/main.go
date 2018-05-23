//FBA POS API
//     Schemes: http, https
//     Host: FBAappuv01.fortebank.com:8082
//     Version: 1.0
//     Consumes:
//     - application/json
//     Produces:
//     - application/json
//     Security:
//     - api_key:
// swagger:meta
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	fbaHTTP "./http"
	"./model"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var httpManager fbaHTTP.HttpManager
var settings model.Settings
var searchSettings []model.ConfigParameter

func main() {
	settings, err := loadConfiguration("settings.json")
	searchSettings, err = loadSearchConfiguration("search.json")
	db, err := getDb(settings)
	check(err)
	httpManager.Manager.DB = db
	httpManager.SearchParameters = searchSettings
	httpManager.Manager.Init()
	//httpManager.Manager.InitTable()
	r := gin.Default()
	r.Use(setCORSMiddleware())

	base := r.Group("/", auth(settings), checkPage(settings))

	//===========  ACTION TYPE =========================================
	base.GET("actiontype", httpManager.ActionTypeGET)
	base.GET("actiontype/:id", httpManager.ActionTypeGETByID)
	base.POST("actiontype", httpManager.ActionTypePOST)
	base.PUT("actiontype", httpManager.ActionTypePUT)
	base.DELETE("actiontype/:id", httpManager.ActionTypeDELETE)
	base.GET("actiontypes/count", httpManager.ActionTypeCount)
	//===========  Category  ============================================
	base.GET("category", httpManager.CategoryGET)
	base.GET("category/:id", httpManager.CategoryGETByID)
	base.POST("category", httpManager.CategoryPOST)
	base.PUT("category", httpManager.CategoryPUT)
	base.DELETE("category/:id", httpManager.CategoryDELETE)
	base.GET("categories/count", httpManager.CategoryCount)
	//===========  city  ============================================
	base.GET("city", httpManager.CityGET)
	base.GET("city/:id", httpManager.CityGETByID)
	base.POST("city", httpManager.CityPOST)
	base.PUT("city", httpManager.CityPUT)
	base.DELETE("city/:id", httpManager.CityDELETE)
	base.GET("cities/count", httpManager.CityCount)
	//===========  Company  ============================================
	base.GET("company", httpManager.CompanyGET)
	base.GET("company/:id", httpManager.CompanyGETByID)
	base.POST("company", httpManager.CompanyPOST)
	base.PUT("company", httpManager.CompanyPUT)
	base.DELETE("company/:id", httpManager.CompanyDELETE)
	base.GET("companies/count", httpManager.CompanyCount)
	//===========  Device  =============================================
	base.GET("device", httpManager.DeviceGET)
	base.GET("device/:id", httpManager.DeviceGETByID)
	base.POST("device", httpManager.DevicePOST)
	base.PUT("device", httpManager.DevicePUT)
	base.DELETE("device/:id", httpManager.DeviceDELETE)
	base.GET("devices/count", httpManager.DeviceCount)
	//===========  PaymentType ==========================================
	base.GET("paymenttype", httpManager.PaymentTypeGET)
	base.GET("paymenttype/:id", httpManager.PaymentTypeGETByID)
	base.POST("paymenttype", httpManager.PaymentTypePOST)
	base.PUT("paymenttype", httpManager.PaymentTypePUT)
	base.DELETE("paymenttype/:id", httpManager.PaymentTypeDELETE)
	base.GET("paymenttypes/count", httpManager.PaymentTypeCount)
	//===========  POINT  ===============================================
	base.GET("/point/:id", httpManager.PointGETByID)
	base.GET("/point/:id/authorize", httpManager.PointGetOauthClient)
	base.GET("/points", httpManager.PointGET)
	base.GET("/points/count", httpManager.PointCount)
	base.POST("/point", httpManager.PointPOST)
	base.PUT("/point", httpManager.PointPUT)
	base.DELETE("/point/:id", httpManager.PointDELETE)
	//===========  Product ==========================================
	base.GET("product", httpManager.ProductGET)
	base.GET("product/:id", httpManager.ProductGETByID)
	base.GET("product/:id/averageprice/:pointid", httpManager.ProductAveragePrice)
	base.GET("product/:id/lastprice/:pointid", httpManager.ProductLastPrice)

	base.POST("product", httpManager.ProductPOST)
	base.PUT("product", httpManager.ProductPUT)
	base.DELETE("product/:id", httpManager.ProductDELETE)
	base.GET("products/count", httpManager.ProductCount)

	//===========  productaction ==========================================
	base.GET("productaction", httpManager.ProductActionGET)
	base.GET("productaction/:id", httpManager.ProductActionGETByID)
	//base.GET("productaction/:shortname", httpManager.ProductActionDetailGET)
	base.POST("productaction", httpManager.ProductActionPOST)
	base.PUT("productaction", httpManager.ProductActionPUT)
	base.DELETE("productaction/:id", httpManager.ProductActionDELETE)
	base.GET("productactions/count", httpManager.ProductActionCount)
	//===========  ProductActionDetail ==========================================
	base.GET("productactiondetail", httpManager.ProductActionDetailGET)
	base.GET("productactiondetail/:id", httpManager.ProductActionDetailGETByID)
	base.POST("productactiondetail", httpManager.ProductActionDetailPOST)
	base.POST("productactiondetails", httpManager.ProductActionDetailsPOST)
	base.PUT("productactiondetail", httpManager.ProductActionDetailPUT)
	base.DELETE("productactiondetail/:id", httpManager.ProductActionDetailDELETE)
	base.GET("productactiondetails/count", httpManager.ProductActionDetailCount)
	base.GET("productactiondetails/avg", httpManager.ProductActionDetailCount)

	//===========  productstock ==========================================
	base.GET("productstock", httpManager.ProductStockGET)
	base.GET("productstock/:pointID", httpManager.ProductStockSearchGET)
	base.POST("productstock", httpManager.ProductStockPOST)
	base.PUT("productstock", httpManager.ProductStockPUT)
	base.DELETE("productstock/:id", httpManager.ProductStockDELETE)
	base.GET("productstocks/count", httpManager.ProductStockCount)
	//===========  role ==========================================
	base.GET("role", httpManager.RoleGET)
	base.GET("role/:id", httpManager.RoleGETByID)
	base.POST("role", httpManager.RolePOST)
	base.PUT("role", httpManager.RolePUT)
	base.DELETE("role/:id", httpManager.RoleDELETE)
	base.GET("roles/count", httpManager.RoleCount)
	//===========  Staff ==========================================
	base.GET("staff", httpManager.StaffGET)
	base.GET("staff/:id", httpManager.StaffGETByID)
	base.POST("staff", httpManager.StaffPOST)
	base.PUT("staff", httpManager.StaffPUT)
	base.DELETE("staff/:id", httpManager.StaffDELETE)
	base.GET("staffs/count", httpManager.StaffCount)
	//===========  StaffPoint ==========================================
	base.GET("staffpoint", httpManager.StaffPointGET)
	base.GET("staffpoint/:id", httpManager.StaffPointGETByID)
	base.POST("staffpoint", httpManager.StaffPointPOST)
	base.PUT("staffpoint", httpManager.StaffPointPUT)
	base.DELETE("staffpoint/:id", httpManager.StaffPointDELETE)
	base.GET("staffpoints/count", httpManager.StaffPointCount)
	//===========  Supplier ==========================================
	base.GET("supplier", httpManager.SupplierGET)
	base.GET("supplier/:id", httpManager.SupplierGETByID)
	base.POST("supplier", httpManager.SupplierPOST)
	base.PUT("supplier", httpManager.SupplierPUT)
	base.DELETE("supplier/:id", httpManager.SupplierDELETE)
	base.GET("suppliers/count", httpManager.SupplierCount)
	//===========  Tax ==========================================
	base.GET("tax", httpManager.TaxGET)
	base.GET("tax/:id", httpManager.TaxGETByID)
	base.POST("tax", httpManager.TaxPOST)
	base.PUT("tax", httpManager.TaxPUT)
	base.DELETE("tax/:id", httpManager.TaxDELETE)
	base.GET("taxs/count", httpManager.TaxCount)
	//===========  WorkSession ==========================================
	base.GET("worksession", httpManager.WorkSessionGET)
	base.GET("worksession/:id", httpManager.WorkSessionGETByID)
	base.POST("worksession", httpManager.WorkSessionPOST)
	base.PUT("worksession", httpManager.WorkSessionPUT)
	base.DELETE("worksession/:id", httpManager.WorkSessionDELETE)
	base.GET("worksessions/count", httpManager.WorkSessionCount)

	//===========  ADD ILYA  ============================================
	//===========  Customer  ============================================
	base.GET("customer", httpManager.CustomerGET)
	base.GET("customer/:id", httpManager.CustomerGETByID)
	base.POST("customer", httpManager.CustomerPOST)
	base.PUT("customer", httpManager.CustomerPUT)
	base.DELETE("customer", httpManager.CustomerDELETE)
	base.GET("customers/count", httpManager.CustomerCount)
	//===========  Discount  ============================================
	base.GET("discount", httpManager.DiscountGET)
	base.GET("discount/:id", httpManager.DiscountGETByID)
	base.POST("discount", httpManager.DiscountPOST)
	base.PUT("discount", httpManager.DiscountPUT)
	base.DELETE("discount/:id", httpManager.DiscountDELETE)
	base.GET("discounts/count", httpManager.DiscountCount)
	//===========  DiscountType  ============================================
	base.GET("discounttype", httpManager.DiscountTypeGET)
	base.GET("discounttype/:id", httpManager.DiscountTypeGETByID)
	base.POST("discounttype", httpManager.DiscountTypePOST)
	base.PUT("discounttype", httpManager.DiscountTypePUT)
	base.DELETE("discounttype", httpManager.DiscountTypeDELETE)
	base.GET("discounttypes/count", httpManager.DiscountTypeCount)
	//===========  Inventory  ==================================================
	base.GET("inventory", httpManager.InventoryGET)
	base.GET("inventory/:id", httpManager.InventoryGETByID)
	base.POST("inventory", httpManager.InventoryPOST)
	base.PUT("inventory", httpManager.InventoryPUT)
	base.DELETE("inventory/:id", httpManager.InventoryDELETE)
	base.GET("inventories/count", httpManager.InventoryCount)
	//===========  InventoryDetail  ============================================
	base.GET("inventorydetail", httpManager.InventoryDetailGET)
	base.GET("inventorydetail/:id", httpManager.InventoryDetailGETByID)
	base.POST("inventorydetail", httpManager.InventoryDetailPOST)
	base.PUT("inventorydetail", httpManager.InventoryDetailPUT)
	base.DELETE("inventorydetail/:id", httpManager.InventoryDetailDELETE)
	base.GET("inventorydetails/count", httpManager.InventoryDetailCount)
	//===========  InventoryDetail  ============================================
	base.GET("receipt", httpManager.ReceiptGET)
	base.GET("receipt/:id", httpManager.ReceiptGETByID)
	base.POST("receipt", httpManager.ReceiptPOST)
	base.PUT("receipt", httpManager.ReceiptPUT)
	base.DELETE("receipt", httpManager.ReceiptDELETE)
	base.GET("receipts/count", httpManager.ReceiptCount)
	//=========== IMPORT =======================================================
	base.POST("import/product/:pointid", httpManager.ImportProductPOST)
	//=========== REPORT =======================================================
	base.GET("report/remainder", httpManager.ReportRemainderGET)
	base.GET("report/incomesales/:companyid", httpManager.ReportIncomeSalesGET)

	r.Run(":" + strconv.Itoa(settings.HTTPPort))
}

func auth(settings model.Settings) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			token := c.Request.Header.Get("Authorization")
			var tokenInfo model.TokenInfo
			client := &http.Client{}
			req, err := http.NewRequest("GET", settings.OAuth+"/check", nil)
			req.Header.Add("Authorization", token)
			resp, err := client.Do(req)
			if err != nil {
				c.AbortWithError(400, err)
			}
			if resp.StatusCode == 200 {
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					c.AbortWithStatusJSON(400, model.NewError("Problem with client information (body)"))
				}
				if err := json.Unmarshal(body, &tokenInfo); err != nil {
					c.AbortWithStatusJSON(400, model.NewError("Problem with client information (parsing)"))
				}
				tokenInfo.Token = token
				c.Set("TokenInfo", tokenInfo)
				c.Next()
			} else if resp.StatusCode == 401 {
				c.AbortWithStatus(401)
			} else {
				c.AbortWithStatusJSON(400, model.NewError("Problem with authorization"))
			}
		}
	}
}

func checkPage(settings model.Settings) gin.HandlerFunc {
	return func(c *gin.Context) {
		size := c.Query("size")
		if size == "" {
			c.Next()
			return
		}
		sizeN, err := strconv.Atoi(size)
		if err != nil {
			c.AbortWithStatusJSON(400, err)
		}
		if sizeN > settings.MaxPageSize {
			c.AbortWithStatusJSON(400, model.NewError("Problem with page size, max = 50, size = "+size))
		}
	}
}

func check(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err)
		panic(err)
	}
}

func loadConfiguration(file string) (config model.Settings, err error) {
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return
}

func loadSearchConfiguration(file string) (config []model.ConfigParameter, err error) {
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return
}

func getDb(settings model.Settings) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%v port=%v user=%v dbname=%v sslmode=disable password=%v", settings.Host, settings.Port, settings.User, settings.DBName, settings.Password)
	fmt.Println(connectionString)
	db, err := gorm.Open("postgres", connectionString)
	db.LogMode(true)
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(80)
	db.DB().SetMaxIdleConns(9)
	return db, err
}

func setCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "authorization, content-type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusOK)
		}
	}
}
