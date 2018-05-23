package http

import (
	"strconv"

	"../dbmodel"
	"../model"

	"github.com/go-openapi/strfmt"

	"github.com/gin-gonic/gin"
)

// swagger:operation GET /сategory Category CategoryGET
// Получение списка категорий
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
//         "$ref": "#/definitions/Category"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) CategoryGET(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	Categorys, err := http.Manager.CategoryGet(size, page)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, Categorys)
}

// swagger:operation GET /category/{id} Category CategoryGETByID
// Получение категории по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор категории
//   required: true
//   type: string <uuid4>
// responses:
//   '200':
//     description: JSON модель категории
//     schema:
//         "$ref": "#/definitions/Category"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) CategoryGETByID(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	Categorys, err := http.Manager.CategoryGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, Categorys)
}

// swagger:operation POST /category Category CategoryPOST
// Создание категории
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для создания категории
//   required: true
//   schema:
//     "$ref": "#/definitions/Category"
// responses:
//   '200':
//     description: JSON созданная модель категории
//     schema:
//         "$ref": "#/definitions/Category"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) CategoryPOST(c *gin.Context) {
	ti := c.MustGet("TokenInfo").(model.TokenInfo)
	var Category dbModel.Category
	if c.Bind(&Category) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	Category, err := http.Manager.CategoryCreate(Category, ti)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, Category)
	return
}

// swagger:operation PUT /category Category CategoryPUT
// Изменение категории
//
// ---
// produces:
// - application/json
// parameters:
// - name: bodyjson
//   in: body
//   description: JSON для изменения модели категории
//   required: true
//   schema:
//     "$ref": "#/definitions/Category"
// responses:
//   '200':
//     description: JSON ответ на изменение категории
//     schema:
//         "$ref": "#/definitions/Category"
//   '400':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) CategoryPUT(c *gin.Context) {
	var Category dbModel.Category
	if c.Bind(&Category) != nil {
		c.JSON(400, "problem decoding body")
		return
	}
	if err := http.Manager.CategoryUpdate(Category); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, Category)
	return
}

// swagger:operation DELETE /category/{id} Category CategoryDELETE
// Удаление категории по ID
//
// ---
// produces:
// - application/json
// parameters:
// - name: id
//   in: path
//   description: идентификатор категории
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
func (http *HttpManager) CategoryDELETE(c *gin.Context) {
	id := strfmt.UUID4(c.Param("id"))

	Category, err := http.Manager.CategoryGetByID(id)
	if err != nil {
		c.JSON(400, err)
	}

	if err := http.Manager.CategoryDelete(Category); err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "")
	return
}

// swagger:operation GET /categories/count Category CategoryCount
// Получение количества категорий
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: JSON с количеством категорий
//     schema:
//         "$ref": "#/definitions/Count"
//   '500':
//     description: JSON ответ об ошибке
//     schema:
//         "$ref": "#/definitions/Error"
func (http *HttpManager) CategoryCount(c *gin.Context) {
	count, err := http.Manager.CategoryCount()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, count)
}
