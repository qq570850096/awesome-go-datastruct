package main

import (
	"net/http"

	"algo/webdemo/minigin"
	"algo/webdemo/redpacket"
)

type apiResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func main() {
	pool := redpacket.NewPool()

	engine := minigin.New()
	engine.Use(minigin.Recovery(), minigin.Logger())

	engine.POST("/redpacket/init", func(c *minigin.Context) {
		var req struct {
			TotalAmount int64 `json:"total_amount"`
			Count       int   `json:"count"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, apiResponse{Code: http.StatusBadRequest, Msg: "invalid json"})
			return
		}
		if err := pool.Init(req.TotalAmount, req.Count); err != nil {
			c.JSON(http.StatusBadRequest, apiResponse{Code: http.StatusBadRequest, Msg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, apiResponse{Code: 0, Msg: "ok", Data: map[string]any{"ok": true}})
	})

	engine.POST("/redpacket/grab", func(c *minigin.Context) {
		amount, err := pool.Grab()
		if err != nil {
			c.JSON(http.StatusBadRequest, apiResponse{Code: http.StatusBadRequest, Msg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, apiResponse{Code: 0, Msg: "ok", Data: map[string]any{"amount": amount}})
	})

	http.ListenAndServe(":8092", engine)
}

