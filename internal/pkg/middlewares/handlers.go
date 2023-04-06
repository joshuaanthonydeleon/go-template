package middlewares

import "github.com/gin-gonic/gin"

func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(405, gin.H{"message": "Metodo no permitido"})
	}
}

func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "The processing function of the request route was not found"})
	}
}
