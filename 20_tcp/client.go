package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func readRoutine(ctx context.Context, cancel context.CancelFunc, conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for {
		select {
		case <-ctx.Done():
			break
		default:
			if !scanner.Scan() {
				log.Printf("CANNOT SCAN")
				cancel()
				break
			}
			text := scanner.Text()
			log.Printf("From server: %s", text)
		}
	}
	log.Printf("Finished readRoutine")
}

func writeRoutine(ctx context.Context, conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			break
		default:
			if !scanner.Scan() {
				break
			}
			str := scanner.Text()
			log.Printf("To server %v\n", str)

			conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}

	}
	log.Printf("Finished writeRoutine")
}

var host string
var port string
var timeout int

func init() {
	println("start")
	flag.IntVar(&timeout, "timeout", 10, "timeout")
}

func main() {
	flag.Parse()
	host = flag.Arg(0)
	port = flag.Arg(1)
	if host == "" || port == "" {
		log.Fatal("Host and Port arguments are required. Example go run server.go my_site.com 80 --timeout=20")
	}
	dialer := &net.Dialer{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	conn, err := dialer.DialContext(ctx, "tcp", host+":"+port)
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		readRoutine(ctx, cancel, conn)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		writeRoutine(ctx, conn)
		wg.Done()
	}()

	wg.Wait()
	conn.Close()
}
