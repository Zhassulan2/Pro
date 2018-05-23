package http

import (
	"strconv"

	"../dbmodel"
	"../model"

	"github.com/go-openapi/strfmt"

	"github.com/gin-gonic/gin"
)

// swagger:operation GET /inventorydetail InventoryDetail InventoryDetailGET
// Получение списка детализаций инвентаризаций
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
//     description: JSON массив детализаций инвентаризаций
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/InventoryDetail"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) InventoryDetailGET(c *gin.Context) {
	iID, ok := c.GetQuery("inventoryID")
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))

	var InventoryDetails []dbModel.InventoryDetail
	if ok {
		inventoryID := strfmt.UUID4(iID)
		InventoryDetails, err = http.Manager.InventoryDetailGetByInventory(inventoryID, size, page)
		if err != nil {
			c.JSON(400, err)
		}
		c.JSON(200, InventoryDetails)
	} else {
		InventoryDetails, err = http.Manager.InventoryDetailGet(size, page)
		if err != nil {
			c.JSON(400, err)
		}
		c.JSON(200, InventoryDetails)
	}
}

// swagger:operation GET /inventorydetail/{id} InventoryDetail InventoryDetailGETByID
// Получение детализации инвентаризации по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор детализации инвентаризации
//   required: true
//   type: string <uuid4>
// responses:
//   '200':
//     description: JSON модель детализации инвентаризации
//     schema:
//         "$ref": "#/definitions/InventoryDetail"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) InventoryDetailGETByID(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	InventoryDetails, err := http.Manager.InventoryDetailGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, InventoryDetails)
}

// swagger:operation POST /inventorydetail InventoryDetail InventoryDetailPOST
// Создание детализации инвентаризации
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для создания детализации инвентаризации
//   required: true
//   schema:
//     "$ref": "#/definitions/InventoryDetail"
// responses:
//   '200':
//     description: JSON созданная модель детализации инвентаризации
//     schema:
//         "$ref": "#/definitions/InventoryDetail"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) InventoryDetailPOST(c *gin.Context) {
	ti := c.MustGet("TokenInfo").(model.TokenInfo)
	var InventoryDetail dbModel.InventoryDetail
	if c.Bind(&InventoryDetail) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	InventoryDetail, err := http.Manager.InventoryDetailCreate(InventoryDetail, ti)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, InventoryDetail)
	return
}

// swagger:operation PUT /inventorydetail InventoryDetail InventoryDetailPUT
// Изменение детализации инвентаризации
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для изменения модели детализации инвентаризации
//   required: true
//   schema:
//     "$ref": "#/definitions/InventoryDetail"
// responses:
//   '200':
//     description: JSON ответ на изменение детализации инвентаризации
//     schema:
//         "$ref": "#/definitions/InventoryDetail"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) InventoryDetailPUT(c *gin.Context) {
	var InventoryDetail dbModel.InventoryDetail
	if c.Bind(&InventoryDetail) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	if err := http.Manager.InventoryDetailUpdate(InventoryDetail); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, InventoryDetail)
	return
}

// swagger:operation DELETE /inventorydetail/{id} InventoryDetail InventoryDetailDELETE
// Удаление детализации инвентаризации по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор детализации инвентаризации
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
func (http *HttpManager) InventoryDetailDELETE(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	InventoryDetail, err := http.Manager.InventoryDetailGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}

	if err := http.Manager.InventoryDetailDelete(InventoryDetail); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "")
	return
}

// swagger:operation GET /inventorydetails/count InventoryDetail InventoryDetailCount
// Получение количества детализаций инвентаризаций
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: JSON с количеством детализаций инвентаризаций
//     schema:
//         "$ref": "#/definitions/Count"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) InventoryDetailCount(c *gin.Context) {
	count, err := http.Manager.InventoryDetailCount()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, count)
}
