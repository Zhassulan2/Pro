package http

import (
	"strconv"

	"../dbmodel"
	"../model"

	"github.com/go-openapi/strfmt"

	"github.com/gin-gonic/gin"
)

// swagger:operation GET /supplier Supplier SupplierGET
// Получение списка поставщиков
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
//     description: JSON массив поставщиков
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Supplier"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) SupplierGET(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	Suppliers, err := http.Manager.SupplierGet(size, page)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, Suppliers)
}

// swagger:operation GET /supplier/{id} Supplier SupplierGETByID
// Получение поставщика по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор поставщика
//   required: true
//   type: string <uuid4>
// responses:
//   '200':
//     description: JSON модель поставщика
//     schema:
//         "$ref": "#/definitions/Supplier"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) SupplierGETByID(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))
	Suppliers, err := http.Manager.SupplierGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, Suppliers)
}

// swagger:operation POST /supplier Supplier SupplierPOST
// Создание поставщика
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для создания поставщика
//   required: true
//   schema:
//     "$ref": "#/definitions/Supplier"
// responses:
//   '200':
//     description: JSON созданная модель поставщика
//     schema:
//         "$ref": "#/definitions/Supplier"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) SupplierPOST(c *gin.Context) {
	ti := c.MustGet("TokenInfo").(model.TokenInfo)
	var Supplier dbModel.Supplier
	if c.Bind(&Supplier) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	Supplier, err := http.Manager.SupplierCreate(Supplier, ti)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, Supplier)
	return
}

// swagger:operation PUT /city Supplier SupplierPUT
// Изменение поставщика
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для изменения модели поставщика
//   required: true
//   schema:
//     "$ref": "#/definitions/Supplier"
// responses:
//   '200':
//     description: JSON ответ на изменение поставщика
//     schema:
//         "$ref": "#/definitions/Supplier"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) SupplierPUT(c *gin.Context) {
	var Supplier dbModel.Supplier
	if c.Bind(&Supplier) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	if err := http.Manager.SupplierUpdate(Supplier); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, Supplier)
	return
}

// swagger:operation DELETE /supplier/{id} Supplier SupplierDELETE
// Удаление поставщика по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор поставщика
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
func (http *HttpManager) SupplierDELETE(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	Supplier, err := http.Manager.SupplierGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}

	if err := http.Manager.SupplierDelete(Supplier); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "")
	return
}

// swagger:operation GET /suppliers/count Supplier SupplierCount
// Получение количества поставщиков
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: JSON с количеством поставщиков
//     schema:
//         "$ref": "#/definitions/Count"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) SupplierCount(c *gin.Context) {
	count, err := http.Manager.SupplierCount()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, count)
}
