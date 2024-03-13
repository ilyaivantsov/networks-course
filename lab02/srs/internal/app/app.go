package app

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	Engine *gin.Engine
}

func New() *App {
	return &App{Engine: gin.Default()}
}
