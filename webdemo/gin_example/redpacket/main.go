package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"algo/webdemo/redpacket"
)

type apiResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func main() {
	r := gin.Default()
	pool := redpacket.NewPool()

	r.POST("/redpacket/init", func(c *gin.Context) {
		var req struct {
			TotalAmount int64 `json:"total_amount"`
			Count       int   `json:"count"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, apiResponse{Code: http.StatusBadRequest, Msg: "invalid json"})
			return
		}
		if err := pool.Init(req.TotalAmount, req.Count); err != nil {
			c.JSON(http.StatusBadRequest, apiResponse{Code: http.StatusBadRequest, Msg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, apiResponse{Code: 0, Msg: "ok", Data: gin.H{"ok": true}})
	})

	r.POST("/redpacket/grab", func(c *gin.Context) {
		amount, err := pool.Grab()
		if err != nil {
			c.JSON(http.StatusBadRequest, apiResponse{Code: http.StatusBadRequest, Msg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, apiResponse{Code: 0, Msg: "ok", Data: gin.H{"amount": amount}})
	})

	r.Run(":8093")
}
