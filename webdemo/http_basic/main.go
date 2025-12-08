package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

// Todo 是一个最小可用的任务模型，用于演示原生 net/http 的 REST 风格接口。
type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type todoStore struct {
	mu    sync.Mutex
	next  int
	todos []Todo
}

func newTodoStore() *todoStore {
	return &todoStore{next: 1}
}

func (s *todoStore) list() []Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Todo, len(s.todos))
	copy(out, s.todos)
	return out
}

func (s *todoStore) add(text string) Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	t := Todo{ID: s.next, Text: text}
	s.next++
	s.todos = append(s.todos, t)
	return t
}

func main() {
	store := newTodoStore()

	mux := http.NewServeMux()

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleList(store, w, r)
		case http.MethodPost:
			handleCreate(store, w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      loggingMiddleware(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("http_basic server listening on :8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}

func handleList(store *todoStore, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(store.list())
}

func handleCreate(store *todoStore, w http.ResponseWriter, r *http.Request) {
	var req struct {
		Text string `json:"text"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if req.Text == "" {
		http.Error(w, "text required", http.StatusBadRequest)
		return
	}
	todo := store.add(req.Text)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(todo)
}

// loggingMiddleware 是最简单的统一拦截器示例：打印方法、路径和耗时。
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

