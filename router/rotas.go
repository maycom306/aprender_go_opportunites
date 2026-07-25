package router

import (
	
	"github.com/gin-gonic/gin"
	"github.com/maycom306/aprender_go_opportunites/handler"
)

func Incializarotas(router *gin.Engine){
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler) //UT é igual a update 

		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}