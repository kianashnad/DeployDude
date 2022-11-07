package main

import (
	"DeployDude/controller"
	"DeployDude/foundation"
	"github.com/gin-gonic/gin"
)

func main() {
	foundation.LoadENVs()
	foundation.GenerateDBSchema()

	r := gin.Default()
	r.POST("/deploy/", controller.DeployProject)
	r.POST("/signal/", controller.AddProject)
	r.DELETE("/signal/", controller.RemoveProject)

	err := r.Run("0.0.0.0:" + foundation.GetENV("DEDU_PORT"))
	if err != nil {
		panic(err)
	}
}
