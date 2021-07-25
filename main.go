package main

//Recomendaciones
//Para el correcto funcionamiento de la api debe ser instalada la dependencia gorilla/mux
// go get -u github.com/gorilla/mux
//La api corre en el puerto 8080
// Puede acceder a la solucion en la siguiente direccion http://localhost:8080/sort
// La peticion post debe tener la siguiente estructura:
// { "unsorted" : [1,1,1,7,7,3,8,5,8,8,10,1,2] }

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//Se define la estructura del JSON
type List struct {
	Unsorted []int `json:"unsorted"`
	Sorted []int `json:"sorted"`

}

//Funcion "Ordenamiento" Recibe dos parametros "Unsort" que hace referencia a una lista desordenada
// y n a la capacidad de la lista
func Ordenamiento(Unsort []int, n int) []int {

	aux := make([]int, n)

	var sort []int

	for i := 0; i < len(Unsort); i++ {
		aux[Unsort[i]]++
	}

	for i := 0; i < n; i++ {
		if aux[i] != 0 {
			sort = append(sort, i)
			aux[i]--
		}
	}

	//los números duplicados se mueven al final de la lista ordenada
	for j := 0; j < len(Unsort); {
		if aux[j] != 0 {
			sort = append(sort, j)
			aux[j]--
		} else {
			j++
		}
	}

	return sort
}

//Se define la funcion post que recibe un JSON con un array desordenado con key "Unsorted"
// y retorna un JSON con el arry ordenado.
func Postsort(w http.ResponseWriter, r *http.Request)  {
	//fmt.Fprint(w, "Metodo POST")
	var Response List

	reqBody, err := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &Response)

	Response.Sorted = Ordenamiento(Response.Unsorted, 100)

	w.Header().Set("Contet-Type", "application/json")
	j, err := json.Marshal(Response)
	if err != nil {
		panic("error")
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(j)

}


func main()  {
	// r Hace referencia al enrutador
	r := mux.NewRouter().StrictSlash(false)

	r.HandleFunc("/sort", Postsort).Methods("POST")

	//Configuración del servidor
	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening...")
	log.Fatal(server.ListenAndServe())
}
