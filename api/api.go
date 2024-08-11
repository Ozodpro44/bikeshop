package api

import "github.com/gorilla/mux"

func Api() {
	r := mux.NewRouter()
    r.HandleFunc("/", api.HomeHandler).Methods("GET")
    r.HandleFunc("/admin", api.AdminHandler).Methods("GET")
    r.HandleFunc("/admin/add", api.AddBikeHandler).Methods("POST")
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}