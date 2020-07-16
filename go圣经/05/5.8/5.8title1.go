package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func title(url string) error {
	resp , err := http.Get(url)
	if err != nil {
		return err
	}
	// Check Content-Type is HTML (e.g., "text/html;charset=utf-8").
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct,"text/html;") {
		resp.Body.Close()
		//return fmt.Errorf("%s has type %s, not text/html",url, ct)
		return log.Fatalf("%s has type %s, not text/html",url, ct)
	}
}