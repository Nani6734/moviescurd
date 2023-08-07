package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			break
		}

	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content type", "application/json")
	var movie movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content type", "Application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["Id"] {
			movies = append(movies[:index], movies[:index+1]...)

			var movie movie

			_ = json.NewDecoder(r.Body).Decode(&movie)

			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)

			return
		}
	}

}

func main() {
	r := mux.NewRouter()
	movies = append(movies, movie{ID: "1", Isbn: "1234", Title: "Firstmovie", Director: &Director{Firstname: "vinith", Lastname: "kumar"}})
	movies = append(movies, movie{ID: "2", Isbn: "2345", Title: "secondmovie", Director: &Director{Firstname: "john", Lastname: "cena"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
