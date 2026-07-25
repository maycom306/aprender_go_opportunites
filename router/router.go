package router

import "github.com/gin-gonic/gin"

func Inicializar() {
	//Inicializar Rotas
	router := gin.Default()

	Incializarotas(router)
	//rodandos servidor
	router.Run("localhost:8080")
}
