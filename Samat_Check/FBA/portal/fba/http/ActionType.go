package http

import (
	"fmt"
	"strconv"

	"../dbmodel"
	"../model"

	"github.com/go-openapi/strfmt"

	"github.com/gin-gonic/gin"
)

// swagger:operation GET /actiontype ActionTypes ActionTypeGET
// Получение списка типов действий для журнала движения товаров
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
//     description: JSON массив записей
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/ActionType"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ActionTypeGET(c *gin.Context) {
	// работа с параметрами
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	params, err := http.ParseQuery(c.Request.URL.RawQuery)
	var parameters model.Parameters
	parameters.Params = params
	parameters.Controller = "ActionType"
	parameters.Config = http.SearchParameters

	//метод может вернуть массив или одну запись
	_, single := parameters.ToQuery()
	if single {
		fmt.Println("SINGLE")
		ActionType, err := http.Manager.ActionTypeSingeGet(size, page, parameters)
		if err != nil {
			c.JSON(400, err)
		}
		c.JSON(200, ActionType)
		return
	}
	fmt.Println("ARRAY")
	ActionTypes, err := http.Manager.ActionTypeGet(size, page, parameters)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, ActionTypes)
	return
}

// swagger:operation GET /actiontype/{id} ActionTypes ActionTypeGETByID
// Получение типа операциий в журнале движения товаров по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор типа операции
//   required: true
//   type: string <uuid4>
// responses:
//   '200':
//     description: JSON модель типа операции
//     schema:
//         "$ref": "#/definitions/ActionType"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ActionTypeGETByID(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	ActionTypes, err := http.Manager.ActionTypeGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, ActionTypes)
}

// swagger:operation POST /actiontype ActionTypes ActionTypePOST
// Создание города
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для создания типа операции
//   required: true
//   schema:
//     "$ref": "#/definitions/City"
// responses:
//   '200':
//     description: JSON созданная модель типа операции
//     schema:
//         "$ref": "#/definitions/ActionType"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ActionTypePOST(c *gin.Context) {
	var ActionType dbModel.ActionType
	if c.Bind(&ActionType) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	ActionType, err := http.Manager.ActionTypeCreate(ActionType)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, ActionType)
	return
}

// swagger:operation PUT /actiontype ActionTypes ActionTypePUT
// Изменение типа операции
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для изменения модели типа операции
//   required: true
//   schema:
//     "$ref": "#/definitions/ActionType"
// responses:
//   '200':
//     description: JSON ответ на изменение города
//     schema:
//         "$ref": "#/definitions/ActionType"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ActionTypePUT(c *gin.Context) {
	var ActionType dbModel.ActionType
	if c.Bind(&ActionType) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	if err := http.Manager.ActionTypeUpdate(ActionType); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, ActionType)
	return
}

// swagger:operation DELETE /actiontype/{id} ActionTypes ActionTypeDELETE
// Удаление типа операции по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор типа операции
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
func (http *HttpManager) ActionTypeDELETE(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	ActionType, err := http.Manager.ActionTypeGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}

	if err := http.Manager.ActionTypeDelete(ActionType); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "")
	return
}

// swagger:operation GET /actiontypes/count ActionTypes ActionTypeCount
// Получение количества типов операций
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: JSON с количеством типов операций
//     schema:
//         "$ref": "#/definitions/Count"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) ActionTypeCount(c *gin.Context) {
	count, err := http.Manager.ActionTypeCount()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, count)
}
