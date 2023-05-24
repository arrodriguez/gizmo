package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	s := &APISimulator{Engine: r}

	r.PUT("/simulations/:id", s.updateSimulation)
	r.POST("/simulations", s.createSimulation)
	r.DELETE("/simulations/:id", s.deleteSimulation)

	r.Run(":8080")

}
