package http

import (
	"strconv"

	"../dbmodel"

	"github.com/go-openapi/strfmt"

	"github.com/gin-gonic/gin"
)

// swagger:operation GET /paymenttype PaymentType PaymentTypeGET
// Получение списка типов платежей
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
//     description: JSON массив типов платежей
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/PaymentType"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) PaymentTypeGET(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))

	shortName := c.Query("shortname")
	if shortName != "" {
		paymentType, err := http.Manager.PaymentTypeGetByShortName(shortName)
		if err != nil {
			c.JSON(400, err)
		}
		c.JSON(200, paymentType)
		return
	}

	PaymentTypes, err := http.Manager.PaymentTypeGet(size, page)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, PaymentTypes)
}

// swagger:operation GET /paymenttype/{id} PaymentType PaymentTypeGETByID
// Получение типа платежа по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор города
//   required: true
//   type: string <uuid4>
// responses:
//   '200':
//     description: JSON модель города
//     schema:
//         "$ref": "#/definitions/PaymentType"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) PaymentTypeGETByID(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	PaymentTypes, err := http.Manager.PaymentTypeGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, PaymentTypes)
}

// swagger:operation POST /paymenttype PaymentType PaymentTypePOST
// Создание типа платежа
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для создания типа платежа
//   required: true
//   schema:
//     "$ref": "#/definitions/PaymentType"
// responses:
//   '200':
//     description: JSON созданная модель типа платежа
//     schema:
//         "$ref": "#/definitions/PaymentType"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) PaymentTypePOST(c *gin.Context) {
	var PaymentType dbModel.PaymentType
	if c.Bind(&PaymentType) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	PaymentType, err := http.Manager.PaymentTypeCreate(PaymentType)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, PaymentType)
	return
}

// swagger:operation PUT /paymenttype PaymentType PaymentTypePUT
// Изменение типа платежа
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для изменения модели типа платежа
//   required: true
//   schema:
//     "$ref": "#/definitions/PaymentType"
// responses:
//   '200':
//     description: JSON ответ на изменение типа платежа
//     schema:
//         "$ref": "#/definitions/PaymentType"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) PaymentTypePUT(c *gin.Context) {
	var PaymentType dbModel.PaymentType
	if c.Bind(&PaymentType) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	if err := http.Manager.PaymentTypeUpdate(PaymentType); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, PaymentType)
	return
}

// swagger:operation DELETE /paymenttype/{id} PaymentType PaymentTypeDELETE
// Удаление типа платежа по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор типа платежа
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
func (http *HttpManager) PaymentTypeDELETE(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	PaymentType, err := http.Manager.PaymentTypeGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}

	if err := http.Manager.PaymentTypeDelete(PaymentType); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "")
	return
}

// swagger:operation GET /paymenttypes/count PaymentType PaymentTypeCount
// Получение количества типов платежей
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: JSON с количеством типов платежей
//     schema:
//         "$ref": "#/definitions/Count"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) PaymentTypeCount(c *gin.Context) {
	count, err := http.Manager.PaymentTypeCount()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, count)
}
