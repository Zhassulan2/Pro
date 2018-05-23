package http

import (
	"../dbmodel"
	"../model"
	"fmt"

	"github.com/go-openapi/strfmt"

	"github.com/gin-gonic/gin"
)

//********* TODO - добавить входные параметры
// swagger:operation GET /product Product ProductGET
// Получение списка товаров
//
// ---
// produces:
// - application/json
// parameters:
// - name: page
//   in: query
//   description: номер страницы пагинатора
//   required: false
//   type: int
// - name: size
//   in: query
//   description: количество записей на странице
//   required: false
//   type: int
// responses:
//   '200':
//     description: JSON массив товаров
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Product"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ProductGET(c *gin.Context) {
	//ti := c.MustGet("TokenInfo").(model.TokenInfo)
	params := http.Params(c)
	Products, err := http.Manager.ProductGet(params)
	if err != nil {
		c.JSON(400, err)
		fmt.Println(err)
		return
	}
	c.JSON(200, Products)
}

// swagger:operation GET /product/{id} Product ProductGETByID
// Получение товара по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор товара
//   required: true
//   type: string <uuid4>
// responses:
//   '200':
//     description: JSON модель товара
//     schema:
//         "$ref": "#/definitions/Product"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ProductGETByID(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	Products, err := http.Manager.ProductGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, Products)
}

func (http *HttpManager) ProductAveragePrice(c *gin.Context) {
	productID := strfmt.UUID4(c.Param("id"))
	pointID := strfmt.UUID4(c.Param("pointid"))

	averagePrice := http.Manager.ProductAveragePrice(productID, pointID)
	c.JSON(200, averagePrice)
}

func (http *HttpManager) ProductLastPrice(c *gin.Context) {
	productID := strfmt.UUID4(c.Param("id"))
	pointID := strfmt.UUID4(c.Param("pointid"))

	lastPrice := http.Manager.ProductLastPrice(productID, pointID)

	c.JSON(200, lastPrice)
}

// swagger:operation POST /product Product ProductPOST
// Создание товара
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для создания товара
//   required: true
//   schema:
//     "$ref": "#/definitions/Product"
// responses:
//   '200':
//     description: JSON созданная модель товара
//     schema:
//         "$ref": "#/definitions/Product"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ProductPOST(c *gin.Context) {
	ti := c.MustGet("TokenInfo").(model.TokenInfo)
	var Product dbModel.Product
	if c.Bind(&Product) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	Product, err := http.Manager.ProductCreate(Product, ti)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, Product)
	return
}

// swagger:operation PUT /product Product ProductPUT
// Изменение товара
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для изменения модели товара
//   required: true
//   schema:
//     "$ref": "#/definitions/Product"
// responses:
//   '200':
//     description: JSON ответ на изменение товара
//     schema:
//         "$ref": "#/definitions/Product"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ProductPUT(c *gin.Context) {
	var Product dbModel.Product
	if c.Bind(&Product) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	if err := http.Manager.ProductUpdate(Product); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, Product)
	return
}

// swagger:operation DELETE /product/{id} Product ProductDELETE
// Удаление товара по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор товара
//   required: true
//   type: string <uuid4>
// responses:
//   '200':
//     description: Успешное удаление
//     schema:
//         type: string
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ProductDELETE(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	Product, err := http.Manager.ProductGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}

	if err := http.Manager.ProductDelete(Product); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "")
	return
}

// swagger:operation GET /products/count Product ProductCount
// Получение количества товаров
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: JSON с количеством товаров
//     schema:
//         "$ref": "#/definitions/Count"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ProductCount(c *gin.Context) {
	count, err := http.Manager.ProductCount()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, count)
}
