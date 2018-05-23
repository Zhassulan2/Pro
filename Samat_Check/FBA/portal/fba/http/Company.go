package http

import (
	"strconv"

	"../dbmodel"
	"../model"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

// swagger:operation GET /company Company CompanyGET
// Получение списка компаний
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
//     description: JSON массив компаний
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Company"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) CompanyGET(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	Companys, err := http.Manager.CompanyGet(size, page)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, Companys)
}

// swagger:operation GET /company/{id} Company CompanyGETByID
// Получение компании по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор компании
//   required: true
//   type: string <uuid4>
// responses:
//   '200':
//     description: JSON модель компании
//     schema:
//         "$ref": "#/definitions/Company"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) CompanyGETByID(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	Companys, err := http.Manager.CompanyGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, Companys)
}

// swagger:operation POST /company Company CompanyPOST
// Создание компании
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для создания компании
//   required: true
//   schema:
//     "$ref": "#/definitions/Company"
// responses:
//   '200':
//     description: JSON созданная модель компании
//     schema:
//         "$ref": "#/definitions/Company"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) CompanyPOST(c *gin.Context) {
	ti := c.MustGet("TokenInfo").(model.TokenInfo)
	var Company dbModel.Company
	if c.Bind(&Company) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	Company, err := http.Manager.CompanyCreate(Company, ti)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, Company)
	return
}

// swagger:operation PUT /company Company CompanyPUT
// Изменение компании
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для изменения модели компании
//   required: true
//   schema:
//     "$ref": "#/definitions/Company"
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
func (http *HttpManager) CompanyPUT(c *gin.Context) {
	var Company dbModel.Company
	if c.Bind(&Company) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	if err := http.Manager.CompanyUpdate(Company); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, Company)
	return
}

// swagger:operation DELETE /company/{id} Company CompanyDELETE
// Удаление города по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор компании
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
func (http *HttpManager) CompanyDELETE(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	Company, err := http.Manager.CompanyGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}

	if err := http.Manager.CompanyDelete(Company); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "")
	return
}

// swagger:operation GET /companies/count Company CompanyCount
// Получение количества компаний
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: JSON с количеством компаний
//     schema:
//         "$ref": "#/definitions/Count"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) CompanyCount(c *gin.Context) {
	count, err := http.Manager.CompanyCount()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, count)
}
