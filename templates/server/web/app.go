package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"project-name/db"
)

type App struct {
	d db.DB
	r *mux.Router
}

func NewApp(d db.DB, cors bool) App {
	app := App{
		d: d,
		r: mux.NewRouter(),
	}
	handler := app.GetTechnologies
	if !cors {
		handler = disableCors(handler)
	}
	app.r.HandleFunc("/api/technologies", handler).Methods("GET")
	app.r.PathPrefix("/").Handler(http.FileServer(http.Dir("/webapp")))
	return app
}

func (a *App) Serve() error {
	return http.ListenAndServe(":8080", a.r)
}

func (a *App) GetTechnologies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	technologies, err := a.d.GetTechnologies()
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(technologies)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}
