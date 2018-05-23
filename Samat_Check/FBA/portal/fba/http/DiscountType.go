package http

import (
	"strconv"

	"../dbmodel"

	"github.com/go-openapi/strfmt"

	"github.com/gin-gonic/gin"
)

// swagger:operation GET /discounttype DiscountType DiscountTypeGET
// Получение списка типов скидок
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
// - name: shortname
//   in: query
//   description: мнемоника типа скидки
//   required: false
//   type: string
// responses:
//   '200':
//     description: JSON массив типов скидок
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/DiscountType"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) DiscountTypeGET(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))

	shortName := c.Query("shortname")
	if shortName != "" {
		actionType, err := http.Manager.DiscountTypeGetByShortName(shortName)
		if err != nil {
			c.JSON(400, err)
		}
		c.JSON(200, actionType)
		return
	}

	DiscountTypes, err := http.Manager.DiscountTypeGet(size, page)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, DiscountTypes)
}

// swagger:operation GET /discounttype/{id} DiscountType DiscountTypeGETByID
// Получение типов скидок по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор типов скидок
//   required: true
//   type: string <uuid4>
// responses:
//   '200':
//     description: JSON модель типов скидок
//     schema:
//         "$ref": "#/definitions/DiscountType"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) DiscountTypeGETByID(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	DiscountTypes, err := http.Manager.DiscountTypeGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, DiscountTypes)
}

// swagger:operation POST /discounttype DiscountType DiscountTypePOST
// Создание типа скидки
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для создания типа скидки
//   required: true
//   schema:
//     "$ref": "#/definitions/DiscountType"
// responses:
//   '200':
//     description: JSON созданная модель типа скидки
//     schema:
//         "$ref": "#/definitions/DiscountType"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) DiscountTypePOST(c *gin.Context) {
	var DiscountType dbModel.DiscountType
	if c.Bind(&DiscountType) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	DiscountType, err := http.Manager.DiscountTypeCreate(DiscountType)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, DiscountType)
	return
}

// swagger:operation PUT /discounttype DiscountType DiscountTypePUT
// Изменение типа скидки
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для изменения модели типа скидки
//   required: true
//   schema:
//     "$ref": "#/definitions/DiscountType"
// responses:
//   '200':
//     description: JSON ответ на изменение типа скидки
//     schema:
//         "$ref": "#/definitions/DiscountType"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) DiscountTypePUT(c *gin.Context) {
	var DiscountType dbModel.DiscountType
	if c.Bind(&DiscountType) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	if err := http.Manager.DiscountTypeUpdate(DiscountType); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, DiscountType)
	return
}

// swagger:operation DELETE /discounttype/{id} DiscountType DiscountTypeDELETE
// Удаление типа скидки по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор типа скидки
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
func (http *HttpManager) DiscountTypeDELETE(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	DiscountType, err := http.Manager.DiscountTypeGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}

	if err := http.Manager.DiscountTypeDelete(DiscountType); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "")
	return
}

// swagger:operation GET /discounttypes/count DiscountType DiscountTypeCount
// Получение количества типов скидок
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: JSON с количеством типов скидок
//     schema:
//         "$ref": "#/definitions/Count"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) DiscountTypeCount(c *gin.Context) {
	count, err := http.Manager.DiscountTypeCount()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, count)
}
