package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"algo/webdemo/redpacket"
)

type apiResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// appHandler 统一返回数据和错误，由包装器负责写 JSON。
type appHandler func(r *http.Request) (interface{}, *appError)

type appError struct {
	Code int
	Msg  string
}

func main() {
	pool := redpacket.NewPool()

	mux := http.NewServeMux()
	mux.Handle("/redpacket/init", wrap(poolInitHandler(pool)))
	mux.Handle("/redpacket/grab", wrap(poolGrabHandler(pool)))

	srv := &http.Server{
		Addr:         ":8090",
		Handler:      loggingMiddleware(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("redpacket server listening on :8090")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}

func poolInitHandler(pool *redpacket.Pool) appHandler {
	return func(r *http.Request) (interface{}, *appError) {
		if r.Method != http.MethodPost {
			return nil, &appError{Code: http.StatusMethodNotAllowed, Msg: "method not allowed"}
		}
		var req struct {
			TotalAmount int64 `json:"total_amount"` // 单位：分
			Count       int   `json:"count"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return nil, &appError{Code: http.StatusBadRequest, Msg: "invalid json"}
		}
		if err := pool.Init(req.TotalAmount, req.Count); err != nil {
			return nil, &appError{Code: http.StatusBadRequest, Msg: err.Error()}
		}
		return map[string]any{"ok": true}, nil
	}
}

func poolGrabHandler(pool *redpacket.Pool) appHandler {
	return func(r *http.Request) (interface{}, *appError) {
		if r.Method != http.MethodPost {
			return nil, &appError{Code: http.StatusMethodNotAllowed, Msg: "method not allowed"}
		}
		amount, err := pool.Grab()
		if err != nil {
			return nil, &appError{Code: http.StatusBadRequest, Msg: err.Error()}
		}
		return map[string]any{"amount": amount}, nil
	}
}

// wrap 将业务处理函数统一包装为 JSON 响应，并做 panic 恢复。
func wrap(h appHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("panic: %v", rec)
				writeJSON(w, http.StatusInternalServerError, apiResponse{Code: 500, Msg: "internal error"})
			}
		}()

		data, appErr := h(r)
		if appErr != nil {
			writeJSON(w, appErr.Code, apiResponse{Code: appErr.Code, Msg: appErr.Msg})
			return
		}
		writeJSON(w, http.StatusOK, apiResponse{Code: 0, Msg: "ok", Data: data})
	})
}

func writeJSON(w http.ResponseWriter, status int, resp apiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}

// loggingMiddleware 作为最外层统一拦截器：记录方法、路径、耗时。
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

