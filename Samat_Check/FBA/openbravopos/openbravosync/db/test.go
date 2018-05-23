package DataBaseManager

import (
	"fmt"
	"openbravosync/model"
	"reflect"

	"github.com/gin-gonic/gin"
)

//Create записывает любой объект в БД

func Create(c *gin.Context, typeObject reflect.Type) {

	connection := c.MustGet("key").(model.DataBaseConnect)

	reflect.TypeOf(model.Attributeuse{})
	db, err := GetDb(connection)
	if err != nil {
		panic(err)
	}

	value := reflect.New(typeObject).Interface()

	if c.Bind(value) != nil {
		c.JSON(400, model.NewError("problem decoding body"))
		return
	}
	//ListFields(value)
	if err := db.Create(value).Error; err != nil {
		c.JSON(404, err)
		return
	}
	c.JSON(201, value)
	db.Close()
}

func SomeHandler(typeObject reflect.Type) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		connection := c.MustGet("key").(model.DataBaseConnect)

		db, err := GetDb(connection)
		if err != nil {
			panic(err)
		}

		value := reflect.New(typeObject).Interface()

		if c.Bind(value) != nil {
			c.JSON(400, model.NewError("problem decoding body"))
			return
		}

		fmt.Println("VALUE = ", value)

		ListFields(value)

		fmt.Println("VALUE = ", value)

		if err := db.Save(value).Error; err != nil {
			c.JSON(404, err)
			return
		}
		c.JSON(201, value)
		db.Close()
	}

	return gin.HandlerFunc(fn)
}

func ListFields(a interface{}) {
	v := reflect.ValueOf(a).Elem()
	for j := 0; j < v.NumField(); j++ {
		f := v.Field(j)
		//n := v.Type().Field(j).Name
		//t := f.Type().String()
		fmt.Println("TYPE = ", v.Type().Field(j).Tag.Get("type"))

		if v.Type().Field(j).Tag.Get("type") == "bytea" {
			fmt.Println("bytea")
			if f.String() != "" {
				input := f.String()
				x := ToString(&input)
				newValue := reflect.ValueOf(x)
				f.Set(newValue)
			}
		}

		//fmt.Printf("Name: %s  Kind: %s  Type: %s\n", n, f.Kind(), t)
	}
}
