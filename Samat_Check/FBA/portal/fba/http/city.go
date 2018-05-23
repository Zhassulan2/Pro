package http

import (
	"fmt"
	"strconv"

	"github.com/go-openapi/strfmt"

	"../dbmodel"
	"../model"

	"github.com/gin-gonic/gin"
)

// swagger:operation GET /city City CityGET
// Получение списка городов
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
//     description: JSON массив городов
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/City"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) CityGET(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	Citys, count, err := http.Manager.CityGet(size, page)
	if err != nil {
		c.JSON(400, err)
	}
	fmt.Println(count)
	c.Header("count", strconv.Itoa(count))
	c.JSON(200, Citys)
}

// swagger:operation GET /city/{id} City CityGETByID
// Получение города по ID
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
//         "$ref": "#/definitions/City"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) CityGETByID(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))
	Citys, err := http.Manager.CityGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, Citys)
}

// swagger:operation POST /city City CityPOST
// Создание города
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для создания города
//   required: true
//   schema:
//     "$ref": "#/definitions/City"
// responses:
//   '200':
//     description: JSON созданная модель города
//     schema:
//         "$ref": "#/definitions/City"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) CityPOST(c *gin.Context) {
	var City dbModel.City
	if c.Bind(&City) != nil {
		c.JSON(400, model.NewError("problem decoding body"))
		return
	}
	City, err := http.Manager.CityCreate(City)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, City)
	return
}

// swagger:operation PUT /city City CityPUT
// Изменение города
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для изменения модели города
//   required: true
//   schema:
//     "$ref": "#/definitions/City"
// responses:
//   '200':
//     description: JSON ответ на изменение города
//     schema:
//         "$ref": "#/definitions/City"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) CityPUT(c *gin.Context) {
	var City dbModel.City
	if c.Bind(&City) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	if err := http.Manager.CityUpdate(City); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, City)
	return
}

// swagger:operation DELETE /city/{id} City CityDELETE
// Удаление города по ID
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
func (http *HttpManager) CityDELETE(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))
	City, err := http.Manager.CityGetByID(id)
	if err != nil {
		c.JSON(400, err)
		return
	}

	if err := http.Manager.CityDelete(City); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "")
	return
}

// swagger:operation GET /cities/count City CityCount
// Получение количества городов
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
func (http *HttpManager) CityCount(c *gin.Context) {
	count, err := http.Manager.CityCount()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, count)
}
