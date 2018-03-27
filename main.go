package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//IP:
func main() {
	http.HandleFunc("/swagger.json", swagger)
	changeHeaderThenServe := func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			h.ServeHTTP(w, r)
		}
	}

	http.Handle("/docker/", changeHeaderThenServe(http.StripPrefix("/docker", http.FileServer(http.Dir("/harry/docker/volume/swagger")))))
	// http.Handle("/docker/", http.StripPrefix("/docker", http.FileServer(http.Dir("/var/docker/volume/swagger"))))
	// http.Handle("/docker/", changeHeaderThenServe(http.StripPrefix("/docker", http.FileServer(http.Dir("C:\\docker\\gitlab\\volume\\swagger")))))
	// http.Handle("/docker2/", http.StripPrefix("/docker2", http.FileServer(http.Dir("C:\\docker\\gitlab\\volume\\swagger"))))

	//Gitlab Webhooks
	http.HandleFunc("/gitlab", gitlabMergeEvent)

	http.ListenAndServe(":8080", nil)
}

func swagger(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	http.ServeFile(w, req, "swagger.json")
}

type test_struct struct {
	Test string
}

type Data struct {
	Name  string
	Value string
}

func gitlabMergeEvent(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	total := []*Data{}

	for k, v := range req.Form {
		item := &Data{}
		item.Name = k
		item.Value = strings.Join(v, "")

		total = append(total, item)
	}

	if err := json.NewEncoder(w).Encode(total); err != nil {
		panic(err)
	}
}
