package http

import (
	"strconv"

	"../dbmodel"
	"../model"

	"github.com/go-openapi/strfmt"

	"github.com/gin-gonic/gin"
)

// swagger:operation GET /discount Discount DiscountGET
// Получение списка скидок
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
//     description: JSON массив скидок
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Discount"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) DiscountGET(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	Discounts, err := http.Manager.DiscountGet(size, page)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, Discounts)
}

// swagger:operation GET /discount/{id} Discount DiscountGETByID
// Получение скидки по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор скидки
//   required: true
//   type: string <uuid4>
// responses:
//   '200':
//     description: JSON модель скидки
//     schema:
//         "$ref": "#/definitions/Discount"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) DiscountGETByID(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	Discounts, err := http.Manager.DiscountGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, Discounts)
}

// swagger:operation POST /discount Discount DiscountPOST
// Создание скидки
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для создания скидки
//   required: true
//   schema:
//     "$ref": "#/definitions/Discount"
// responses:
//   '200':
//     description: JSON созданная модель скидки
//     schema:
//         "$ref": "#/definitions/Discount"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) DiscountPOST(c *gin.Context) {
	ti := c.MustGet("TokenInfo").(model.TokenInfo)
	var Discount dbModel.Discount
	if c.Bind(&Discount) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	Discount, err := http.Manager.DiscountCreate(Discount, ti)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, Discount)
	return
}

// swagger:operation PUT /discount Discount DiscountPUT
// Изменение скидки
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для изменения модели скидки
//   required: true
//   schema:
//     "$ref": "#/definitions/Discount"
// responses:
//   '200':
//     description: JSON ответ на изменение скидки
//     schema:
//         "$ref": "#/definitions/Discount"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) DiscountPUT(c *gin.Context) {
	var Discount dbModel.Discount
	if c.Bind(&Discount) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	if err := http.Manager.DiscountUpdate(Discount); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, Discount)
	return
}

// swagger:operation DELETE /discount/{id} Discount DiscountDELETE
// Удаление скидки по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор скидки
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
func (http *HttpManager) DiscountDELETE(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	Discount, err := http.Manager.DiscountGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}

	if err := http.Manager.DiscountDelete(Discount); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "")
	return
}

// swagger:operation GET /discounts/count Discount DiscountCount
// Получение количества скидок
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: JSON с количеством скидок
//     schema:
//         "$ref": "#/definitions/Count"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) DiscountCount(c *gin.Context) {
	count, err := http.Manager.DiscountCount()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, count)
}
