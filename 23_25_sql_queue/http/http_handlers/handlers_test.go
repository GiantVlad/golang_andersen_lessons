package http_handlers

import (
	"github.com/stretchr/testify/require"
	"go_andr_less/21_http/domain"
	"go_andr_less/21_http/in_memory_store"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func init() {
	domain.InitStore(&in_memory_store.Events{})
}

func TestCreateEventHandler(t *testing.T) {
	body := strings.NewReader(`{"content":"some value","start_date":"2023-01-09"}`)
	req := httptest.NewRequest("POST", "http://example.com/foo", body)
	w := httptest.NewRecorder()
	CreateEventHandler(w, req)
	resp := w.Result()
	bodyR, _ := io.ReadAll(resp.Body)
	require.Equal(t, 200, resp.StatusCode, "")
	require.Equal(t, `{"id":1,"content":"some value","start_date":"2023-01-09"}`+"\n", string(bodyR), "")
}
