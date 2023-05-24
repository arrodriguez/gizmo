package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/myesui/uuid"
)

type APISimulator struct {
	Engine *gin.Engine
}

type APISimulation struct {
	Method           string  `json:"method"`
	APIResource      string  `json:"api_resource"`
	Latency          string  `json:"latency"`
	UseAllCores      bool    `json:"use_all_cores"`     //  If the hardware or Virtualization or container has more cpus cores, percentage will applyed to all cores.
	CPUPercentage    float32 `json:"cpu_percentage"`    // Such as CPU intensive laod. E.g encryption, Math operations, etc
	IOWaitPercentage float32 `json:"iowait_percentage"` // Such as networking or Disk thus CPU became Idle
}

func (s *APISimulator) updateSimulation(c *gin.Context) {
}

func (s *APISimulator) createSimulation(c *gin.Context) {
	var as APISimulation
	if err := c.BindJSON(&as); err != nil {
		fmt.Printf("error parsing create simulation body %v", err)
		return
	}

	s.Engine.Handle(as.Method, as.APIResource, func(c *gin.Context) {
		d, err := time.ParseDuration(as.Latency)

		if err != nil {
			c.Error(err)
			return
		}

		bl := BusyLoad{Load: float64(as.CPUPercentage), Duration: d, TickFrecuency: 1 * time.Millisecond}
		bl.Run()

		c.IndentedJSON(http.StatusOK, gin.H{"msg": "executed"})

	})

	c.IndentedJSON(http.StatusOK, gin.H{"id": uuid.NewV4()})
}

func (s *APISimulator) deleteSimulation(c *gin.Context) {
}
