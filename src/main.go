package main

import (
	"encoding/json"
	"net/http"
)

type Data struct {
	Universidad                        string `json:"Universidad"`
	Curso                              string `json:"Curso"`
	Alumno                             string `json:"Alumno"`
	Periodo                            string `json:"Periodo"`
	Lenguaje_de_programacion_preferido string `json:"Lenguaje de programacion preferido"`
	Aspiracion_PostGraduacion          string `json:"Aspiración PostGraduación"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Data{"UTPL", "Procesos de Ingeniería de Software", "Jose Melkisedeh Abad Troya", "Abr/Ago 2021", "Python", "Emprendimiento"}
		json.NewEncoder(w).Encode(data)
	})
	http.ListenAndServe(":5656", nil)
}
