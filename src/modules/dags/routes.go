package dags

import (
	"encoding/json"
	"io"
	"net/http"
)

func DagMux() http.Handler {
	useMux := http.NewServeMux()
	dc := ControllerConstructor()
	useMux.Handle("GET /" , http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dag, err := dc.GetDags()
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		json.NewEncoder(w).Encode(dag)
	}))

	useMux.Handle("POST /" , http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dag DagDTO
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 422)
		}
		json.Unmarshal(body, &dag)
		err = dc.CreateDag(dag)
		if err != nil {
			http.Error(w, err.Error(), 422)
		}
		json.NewEncoder(w).Encode(dag)
	}))

	useMux.Handle("PUT /" , http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}))

	useMux.Handle("DELETE /" , http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}))

	return useMux



}