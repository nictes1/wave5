package main

import "github.com/gin-gonic/gin"

type request struct {
	ID       int
	Nombre   string
	Tipo     string
	Cantidad int
	Precio   float64
}

var products []request
var lastID int

func main() {
	server := gin.Default()

	pr := server.Group("/productos")
	pr.POST("/", Guardar())

	server.Run(":8080")
}

func Guardar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		req.ID = 4

		products = append(products, req)

		c.JSON(200, req)
	}
}

//lastID++
//req.ID = lastID

func GuardarWithAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		lastID++
		req.ID = lastID

		products = append(products, req)

		c.JSON(200, req)
	}

}
