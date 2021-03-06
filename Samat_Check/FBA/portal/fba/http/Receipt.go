package http

import (
	"strconv"

	"../dbmodel"

	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

// swagger:operation GET /receipt Receipt ReceiptGET
// Получение списка чеков
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
//     description: JSON массив чеков
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Receipt"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ReceiptGET(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	receipts, err := http.Manager.ReceiptGet(size, page)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, receipts)
}

// swagger:operation GET /receipt/{id} Receipt ReceiptGETByID
// Получение чека по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор чека
//   required: true
//   type: string <uuid4>
// responses:
//   '200':
//     description: JSON модель чека
//     schema:
//         "$ref": "#/definitions/Receipt"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ReceiptGETByID(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.JSON(400, err)
		return
	}
	receipts, err := http.Manager.ReceiptGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, receipts)
}

// swagger:operation POST /receipt Receipt ReceiptPOST
// Создание чека
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для создания чека
//   required: true
//   schema:
//     "$ref": "#/definitions/Receipt"
// responses:
//   '200':
//     description: JSON созданная модель чека
//     schema:
//         "$ref": "#/definitions/Receipt"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ReceiptPOST(c *gin.Context) {
	//ti := c.MustGet("TokenInfo").(model.TokenInfo)
	var receipt dbModel.Receipt
	if c.Bind(&receipt) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	receipt, err := http.Manager.ReceiptCreate(receipt)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, receipt)
	return
}

// swagger:operation PUT /receipt Receipt ReceiptPUT
// Изменение чека
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для изменения модели чека
//   required: true
//   schema:
//     "$ref": "#/definitions/Receipt"
// responses:
//   '200':
//     description: JSON ответ на изменение чека
//     schema:
//         "$ref": "#/definitions/Receipt"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ReceiptPUT(c *gin.Context) {
	var receipt dbModel.Receipt
	if c.Bind(&receipt) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	if err := http.Manager.ReceiptUpdate(receipt); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, receipt)
	return
}

// swagger:operation DELETE /receipt/{id} Receipt ReceiptDELETE
// Удаление чека по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор чека
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
func (http *HttpManager) ReceiptDELETE(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.JSON(400, err)
		return
	}

	receipt, err := http.Manager.ReceiptGetByID(id)
	if err != nil {
		c.JSON(400, err)
		return
	}

	if err := http.Manager.ReceiptDelete(receipt); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "")
	return
}

// swagger:operation GET /receipts/count Receipt ReceiptCount
// Получение количества чеков
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: JSON с количеством городов
//     schema:
//         "$ref": "#/definitions/Count"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ReceiptCount(c *gin.Context) {
	count, err := http.Manager.ReceiptCount()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, count)
}
