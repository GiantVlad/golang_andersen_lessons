package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"go_andr_less/23_25_sql_queue/http/domain"
	"go_andr_less/23_25_sql_queue/http/utils"
	"log"
	"sync"
	"time"
)

var db *sql.DB
var ctx context.Context

func init() {
	ctx = context.Background()
	dsn := "postgres://postgres:postgres@localhost:5432/calendar"
	var err error
	db, err = sql.Open("pgx", dsn) // *sql.DB
	if err != nil {
		log.Fatalf("failed to load driver: %v", err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	log.Println("...connected")
}

type Events struct {
}

var mu = sync.Mutex{}

func (events *Events) Add(e *domain.Event) error {
	query := `insert into events(content, start_date) values($1, $2)`
	result, err := db.ExecContext(ctx, query, e.Content, e.StartDate.UTC()) // sql.Result
	if err != nil {
		return fmt.Errorf("db error: %v", err)
	}
	// Auto-generated ID (SERIAL)
	id, err := result.LastInsertId()
	e.Id = int(id)

	return nil
}

func (events *Events) Get(id int) (domain.Event, error) {
	event := domain.Event{}
	query := `select id, content, start_date from events where id = $1`
	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return event, fmt.Errorf("db error: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var db_id int64
		var db_content string
		var db_start time.Time
		if err := rows.Scan(&db_id, &db_content, &db_start); err != nil {
			return event, fmt.Errorf("db scan error: %v", err)
		}

		event.Id = int(db_id)
		event.Content = db_content
		event.StartDate = utils.Date{db_start}
	}
	if err := rows.Err(); err != nil {
		return event, fmt.Errorf("db rows error: %v", err)
	}

	if event.Id > 0 {
		return event, nil
	}

	return domain.Event{}, fmt.Errorf("not found")
}

func (events *Events) Update(e *domain.Event) error {
	_, err := events.Get(e.Id)
	if err != nil {
		return fmt.Errorf("#%d not found", e.Id)
	}

	query := `update events set content=$1, start_date=$2 where id=$3`
	_, err = db.ExecContext(ctx, query, e.Content, e.StartDate.UTC(), e.Id)
	if err != nil {
		return fmt.Errorf("db error: %v", err)
	}

	return nil
}

func (events *Events) Delete(e *domain.Event) error {
	_, err := events.Get(e.Id)
	if err != nil {
		return fmt.Errorf("#%d not found", e.Id)
	}
	query := `delete from events where id=$1`
	_, err = db.ExecContext(ctx, query, e.Id)
	if err != nil {
		return fmt.Errorf("db error: %v", err)
	}

	return nil
}

//func (events *Events) List(dateFrom time.Time, dateTo time.Time) ([]domain.Event, error) {
//	mu.Lock()
//	defer mu.Unlock()
//	var filtered []domain.Event
//	var evs []domain.Event
//	evs = *events
//	for i := range evs {
//		if evs[i].StartDate.Time.Equal(dateFrom) || (evs[i].StartDate.Time.After(dateFrom) && evs[i].StartDate.Time.Before(dateTo)) {
//			filtered = append(filtered, evs[i])
//		}
//	}
//	return filtered, nil
//}
//
//func (events *Events) Count() int {
//	mu.Lock()
//	defer mu.Unlock()
//	return len(*events)
//}
