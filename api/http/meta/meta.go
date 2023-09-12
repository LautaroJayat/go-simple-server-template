package meta

import (
	"fmt"
	"log"
	"net/http"
)

type metaMux struct {
	l *log.Logger
}

func (m *metaMux) logError(e string) {
	m.l.Printf("error=%q", e)
}

func (m *metaMux) ready(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("ok"))

	if err != nil {
		m.logError(fmt.Sprintf("something happened while reporting readyness: %q", err))
	}

}

func (m *metaMux) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("ok"))

	if err != nil {
		m.logError(fmt.Sprintf("something happened while reporting status: %q", err))
	}

}

func NewMux(logger *log.Logger) *http.ServeMux {
	metaMux := &metaMux{logger}
	m := http.NewServeMux()
	m.HandleFunc("/ready", metaMux.ready)
	m.HandleFunc("/status", metaMux.health)
	return m
}
