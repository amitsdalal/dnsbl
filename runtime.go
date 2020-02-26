package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"gocloud.dev/server"
	"gocloud.dev/server/health"
	"gocloud.dev/server/requestlog"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"sync"
	"time"
)

type customHealthCheck struct {
	mu      sync.RWMutex
	healthy bool
}

func (h *customHealthCheck) CheckHealth() error {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if !h.healthy {
		return errors.New("not ready yet!")
	}
	return nil
}

func findIP(input string) string {
	numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	regexPattern := numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock
	regEx := regexp.MustCompile(regexPattern)
	return regEx.FindString(input)
}

func main() {
	scriptPtr := flag.String("script", "/bin/dnsbl", "Path to the script : ")
	portPtr := flag.String("port", "5000", "Port number : ")
	flag.Parse()
	healthCheck := new(customHealthCheck)
	time.AfterFunc(time.Second, func() {
		healthCheck.mu.Lock()
		defer healthCheck.mu.Unlock()
		healthCheck.healthy = true
	})

	srvOptions := &server.Options{
		RequestLogger: requestlog.NewNCSALogger(os.Stdout, func(e error) { fmt.Fprintln(os.Stderr, e) }),
		HealthChecks:  []health.Checker{healthCheck},
		Driver:        &server.DefaultDriver{},
	}

	surv := http.NewServeMux()
	srv := server.New(surv, srvOptions)

	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		for {
			<-interrupt
			srv.Shutdown(context.Background())
		}
	}()

	surv.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.CommandContext(r.Context(), "/bin/bash", *scriptPtr, findIP(r.URL.Path))
		cmd.Stderr = os.Stderr
		out, err := cmd.Output()
		if err != nil {
			w.WriteHeader(500)
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Write(out)
	})

	port := fmt.Sprintf(":%s", *portPtr)
	if err := srv.ListenAndServe(port); err != nil {
		log.Fatalf("%v", err)
	}

}
