package http_handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_andr_less/21_http/domain"
	"log"
	"net/http"
	"strconv"
	"time"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	event := domain.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Printf("err: %v", err)
		addError(err.Error(), http.StatusBadRequest, w)
		return
	}
	err = domain.Add(&event)
	if err != nil {
		log.Printf("err1: %v", err)
		addError("da9b216c-024a-4020-8ec9-8b75399a0aab", http.StatusHTTPVersionNotSupported, w)
		return
	}
	err = json.NewEncoder(w).Encode(event)
	if err != nil {
		addError("7ab8a66b-b46f-4bb5-bc90-f5a01e88c0ad", http.StatusInternalServerError, w)
		log.Printf("err: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	event := domain.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		addError(err.Error(), http.StatusBadRequest, w)
		return
	}
	err = domain.Update(&event)
	if err != nil {
		addError("b514aa7d-f773-47b7-8b54-03f6fcc2c889", http.StatusInternalServerError, w)
		return
	}
	err = json.NewEncoder(w).Encode(event)
	if err != nil {
		addError("c3011d4f-97a2-4182-9438-44726b87bc7b", http.StatusInternalServerError, w)
		log.Printf("err: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idParam := vars["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		addError(err.Error(), http.StatusBadRequest, w)
		log.Printf("Invalid event ID")
		return
	}
	event, err := domain.Get(id)
	if err != nil {
		addError(err.Error(), http.StatusNotFound, w)
		log.Printf("Event Not found")
		return
	}
	err = domain.Delete(&event)
	if err != nil {
		addError("418facda-49d0-4346-8567-04abc6863859", http.StatusInternalServerError, w)
		log.Printf("err: %v", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetDayEventsHandler(w http.ResponseWriter, r *http.Request) {
	eventStore := r.Context().Value("store").(*domain.EventStore)
	date := r.URL.Query().Get("date")
	if date == "" {
		addError("The date query param is required", http.StatusBadRequest, w)
		log.Printf("The date query param is required")
		return
	}
	dateTime, err := time.Parse(time.DateOnly, date)
	if err != nil {
		addError("The date query param is invalid", http.StatusBadRequest, w)
		log.Printf("error: %v", err)
		return
	}
	list, err := (*eventStore).List(dateTime, dateTime)
	if err != nil {
		addError("4f581e19-7a63-4349-b976-d634def1b2b6", http.StatusInternalServerError, w)
		log.Printf("error: %v", err)
		return
	}
	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetWeekEventsHandler(w http.ResponseWriter, r *http.Request) {
	eventStore := r.Context().Value("store").(*domain.EventStore)
	date := r.URL.Query().Get("date")
	if date == "" {
		addError("The date query param is required", http.StatusBadRequest, w)
		log.Printf("The date query param is required")
		return
	}
	dateTime, err := time.Parse(time.DateOnly, date)
	if err != nil {
		addError("The date query param is invalid", http.StatusBadRequest, w)
		log.Printf("error: %v", err)
		return
	}
	weekday := time.Duration(dateTime.Weekday())
	dateFrom := dateTime.Add(-(24 * time.Hour * (7 - weekday)))
	dateTo := dateTime.Add(24 * time.Hour * (7 - weekday))
	list, err := (*eventStore).List(dateFrom, dateTo)
	if err != nil {
		addError("42d9d00f-200f-414e-a257-e10d6a1c0e18", http.StatusInternalServerError, w)
		log.Printf("error: %v", err)
		return
	}
	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		addError("3bbc078f-b0e7-4ad4-911f-868abd08d2f7", http.StatusInternalServerError, w)
		log.Printf("error: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetMonthEventsHandler(w http.ResponseWriter, r *http.Request) {
	eventStore := r.Context().Value("store").(*domain.EventStore)
	date := r.URL.Query().Get("date")
	if date == "" {
		addError("The date query param is required", http.StatusBadRequest, w)
		log.Printf("The date query param is required")
		return
	}
	dateTime, err := time.Parse(time.DateOnly, date)
	if err != nil {
		addError("The date query param is invalid", http.StatusBadRequest, w)
		log.Printf("error: %v", err)
		return
	}
	dateFrom := dateTime.AddDate(0, 0, -dateTime.Day()+1)
	dateTo := dateTime.AddDate(0, 1, -dateTime.Day())
	list, err := (*eventStore).List(dateFrom, dateTo)
	if err != nil {
		addError("6e339ac4-8209-4b18-a31d-e4f7e3759eb6", http.StatusInternalServerError, w)
		log.Printf("error: %v", err)
		return
	}
	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		addError("4fbde102-1a79-4c65-aba2-f81e7add311a", http.StatusInternalServerError, w)
		log.Printf("error: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
