package http

import (
	"fmt"
	"strconv"

	"../dbmodel"

	"github.com/go-openapi/strfmt"

	"github.com/gin-gonic/gin"
)

// swagger:operation GET /productactiondetail ProductActionDetail ProductActionDetailGET
// Получение получение детализации по записи журнала движения товаров
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
// - name: productActionID
//   in: query
//   description: идентификатор записи в журнале движения товаров
//   required: false
//   type: string <uuid4>
// responses:
//   '200':
//     description: JSON массив детализации журнала движения товаров
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/ProductActionDetail"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ProductActionDetailGET(c *gin.Context) {
	paID, ok := c.GetQuery("productActionID")
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		c.JSON(400, err)
	}
	var ProductActionDetails []dbModel.ProductActionDetail

	if ok {
		productActionID := strfmt.UUID4(paID)
		if err != nil {
			c.JSON(400, err)
		}
		ProductActionDetails, err = http.Manager.ProductActionDetailGetByProductAction(productActionID, size, page)
	} else {
		ProductActionDetails, err = http.Manager.ProductActionDetailGet(size, page)
		if err != nil {
			c.JSON(400, err)
		}
	}

	c.JSON(200, ProductActionDetails)
}

// swagger:operation GET /productactiondetail/{id} ProductActionDetail ProductActionDetailGETByID
// Получение детализации журнала движения товаров по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор детализации журнала движения товаров
//   required: true
//   type: string <uuid4>
// responses:
//   '200':
//     description: JSON модель детализации журнала движения товаров
//     schema:
//         "$ref": "#/definitions/ProductActionDetail"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ProductActionDetailGETByID(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	ProductActionDetails, err := http.Manager.ProductActionDetailGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, ProductActionDetails)
}

/*
func (http *HttpManager) ProductActionDetailGETByProductAtionID(c *gin.Context) {
	id := strfmt.UUID4(c.Query("productActionID"))
	if err != nil {
		c.JSON(400, err)
		return
	}
	ProductActionDetails, err := http.Manager.ProductActionDetailGetByProductAction(id)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, ProductActionDetails)
}*/

// swagger:operation POST /productactiondetail ProductActionDetail ProductActionDetailPOST
// Создание записи детализации журнала движения товаров
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для создания записи детализации журнала движения товаров
//   required: true
//   schema:
//     "$ref": "#/definitions/ProductActionDetail"
// responses:
//   '200':
//     description: JSON созданная модель записи детализации журнала движения товаров
//     schema:
//         "$ref": "#/definitions/ProductActionDetail"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ProductActionDetailPOST(c *gin.Context) {
	var ProductActionDetail dbModel.ProductActionDetail
	if c.Bind(&ProductActionDetail) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	ProductActionDetail, err := http.Manager.ProductActionDetailCreate(ProductActionDetail)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, ProductActionDetail)
	return
}

// swagger:operation POST /productactiondetails ProductActionDetail ProductActionDetailsPOST
// Создание списка записей детализации журнала движения товаров
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для создания списка записей детализации журнала движения товаров
//   required: true
//   schema:
//     type: array
//     items:
//       "$ref": "#/definitions/ProductActionDetail"
// responses:
//   '200':
//     description: JSON ответ
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/ProductActionDetail"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ProductActionDetailsPOST(c *gin.Context) {
	fmt.Println("start ProductActionDetailsPOST")
	var ProductActionDetails []dbModel.ProductActionDetail
	if c.Bind(&ProductActionDetails) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	ProductActionDetailsRet, err := http.Manager.ProductActionDetailsCreate(ProductActionDetails)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, ProductActionDetailsRet)
	return
}

// swagger:operation PUT /city ProductActionDetail ProductActionDetailPUT
// Изменение записи детализации журнала движения товаров
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для изменения модели записи детализации журнала движения товаров
//   required: true
//   schema:
//     "$ref": "#/definitions/ProductActionDetail"
// responses:
//   '200':
//     description: JSON ответ на изменение записи детализации журнала движения товаров
//     schema:
//         "$ref": "#/definitions/ProductActionDetail"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ProductActionDetailPUT(c *gin.Context) {
	var ProductActionDetail dbModel.ProductActionDetail
	if c.Bind(&ProductActionDetail) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	if err := http.Manager.ProductActionDetailUpdate(ProductActionDetail); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, ProductActionDetail)
	return
}

// swagger:operation DELETE /productactiondetail/{id} ProductActionDetail ProductActionDetailDELETE
// Удаление записи детализации журнала движения товаров по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор записи детализации журнала движения товаров
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
func (http *HttpManager) ProductActionDetailDELETE(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	ProductActionDetail, err := http.Manager.ProductActionDetailGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}

	if err := http.Manager.ProductActionDetailDelete(ProductActionDetail); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "")
	return
}

// swagger:operation GET /productactiondetails/count ProductActionDetail ProductActionDetailCount
// Получение количества записей детализации журнала движения товаров
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: JSON с количеством записей детализации журнала движения товаров
//     schema:
//         "$ref": "#/definitions/Count"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ProductActionDetailCount(c *gin.Context) {
	count, err := http.Manager.ProductActionDetailCount()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, count)
}
