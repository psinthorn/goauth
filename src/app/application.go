package app

import (
	"github.com/gin-gonic/gin"
	"github.com/psinthorn/goauth/src/clients/cassandra"
	"github.com/psinthorn/goauth/src/domain/access_token"
	"github.com/psinthorn/goauth/src/http"
	"github.com/psinthorn/goauth/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {

	// Conntect to Cassandra just to make sure we can create session
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()

	dbRepository := db.NewRepository()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080")
}
