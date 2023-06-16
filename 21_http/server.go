package main

import (
	"context"
	"github.com/gorilla/mux"
	"go_andr_less/21_http/domain"
	"go_andr_less/21_http/http_handlers"
	"go_andr_less/21_http/in_memory_store"
	"log"
	"net/http"
	"time"
)

type StoreMW struct {
	store *domain.EventStore
}

func (storeMW *StoreMW) AddStoreMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "store", storeMW.store)
		r = r.WithContext(ctx)
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	storeMW := StoreMW{domain.InitStore(&in_memory_store.Events{})}

	router := mux.NewRouter()
	router.HandleFunc("/create_event", http_handlers.CreateEventHandler).Methods("POST")
	router.HandleFunc("/update_event/{id:[0-9]+}", http_handlers.UpdateEventHandler).Methods("PUT")
	router.HandleFunc("/delete_event/{id:[0-9]+}", http_handlers.DeleteEventHandler).Methods("DELETE")
	router.HandleFunc("/events_for_day", http_handlers.GetDayEventsHandler).Methods("GET")
	router.HandleFunc("/events_for_week", http_handlers.GetWeekEventsHandler).Methods("GET")
	router.HandleFunc("/events_for_month", http_handlers.GetMonthEventsHandler).Methods("GET")
	router.Use(storeMW.AddStoreMiddleware)
	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	// http.Handle("/", r)
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8082",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal("Can't start server")
	}
}
