package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func sort(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Metodo POST")
}


func main()  {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/sort", sort).Methods("POST")

	server := &http.Server{
		Addr:           ":8080",          // port
		Handler:        router,           //
		ReadTimeout:    10 * time.Second, // tiempo de lectura
		WriteTimeout:   10 * time.Second, // tiempo de escritura
		MaxHeaderBytes: 1 << 20,          // 1mega en bits
	}
	log.Println("Escuchando....")
	log.Fatal(server.ListenAndServe())
}
