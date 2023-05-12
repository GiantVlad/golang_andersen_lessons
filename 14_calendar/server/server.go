package server

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type HttpCfg struct {
	Host string `config:"host"`
	Port uint32 `config:"port"`
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is the calendar!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello from Calendar app!\n")
}

func Start(httpCfg HttpCfg) {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(httpCfg.Host+":"+strconv.Itoa(int(httpCfg.Port)), nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
