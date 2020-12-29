package app

import (
	"github.com/gin-gonic/gin"
	"github.com/psinthorn/goauth/src/domain/access_token"
	"github.com/psinthorn/goauth/src/http"
	"github.com/psinthorn/goauth/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication(port string) {

	dbRepository := db.NewRepository()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8083")
}
