package main

import (
	"fmt"
	"io"
	"local/backend/environment"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type requestAssertionHeader struct {
	key   string
	value string
}

type requestAssertion struct {
	text       string
	httpStatus int
	header     requestAssertionHeader
}

func assertRequest(request http.Request) *requestAssertion {
	switch {
	case request.Method != http.MethodPost:
		statusCode := http.StatusMethodNotAllowed
		return &requestAssertion{
			http.StatusText(statusCode),
			statusCode,
			requestAssertionHeader{"Allow", "POST"},
		}

	case !strings.Contains(
		request.Header.Get("Content-Type"),
		"application/x-www-form-urlencoded",
	):
		statusCode := http.StatusUnsupportedMediaType
		return &requestAssertion{
			http.StatusText(statusCode),
			statusCode,
			requestAssertionHeader{"Accept-Post", "application/x-www-form-urlencoded"},
		}
	}

	return nil
}

func main() {
	http.HandleFunc("/api/doctor", func(w http.ResponseWriter, r *http.Request) {
		if environment.CorsActive {
			w.Header().Add("Access-Control-Allow-Origin", "*")
		}

		assertionError := assertRequest(*r)

		if assertionError != nil {
			w.Header().Add(assertionError.header.key, assertionError.header.value)
			http.Error(w, assertionError.text, assertionError.httpStatus)
			return
		}

		answer := r.FormValue("answer")
		cmd := exec.Command("emacs", "--batch", "-Q", "--script", "doctor.el")

		stdin, error := cmd.StdinPipe()
		if error != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Println(error)
			return
		}

		go func() {
			defer stdin.Close()
			io.WriteString(stdin, answer)
		}()

		stdout, error := cmd.Output()
		if error != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Println(error)
			return
		}

		output := string(stdout)
		slicedOutput := output[113 : len(output)-2]
		fmt.Fprintf(w, "%s", slicedOutput)
	})

	http.Handle("/", http.FileServer(http.Dir("./public")))

	serverAddress := ":" + environment.Port
	log.Println("Starting server on", serverAddress)
	log.Fatalln(http.ListenAndServe(serverAddress, nil))
}
