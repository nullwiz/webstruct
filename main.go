package main

import (
	"webstruct/api/handler"
	mock "webstruct/repository/mockdatabase"
	"webstruct/usecase/base"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	srv *base.Service
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := log.With().Logger()
	repo := mock.NewMockDatabase()
	srv = base.LoadService(repo, &logger)

	gin.SetMode(gin.ReleaseMode)
	r := handler.NewGinHandler(srv)
	// start gin
	r.Run("127.0.0.1:8000")
}
