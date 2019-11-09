package app

import (
	"github.com/federicoleon/bookstore_oauth-api/src/domain/access_token"
	"github.com/federicoleon/bookstore_oauth-api/src/repository/db"
	"github.com/federicoleon/bookstore_oauth-api/src/http"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080")

}
