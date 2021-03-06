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

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

//Get all movies controller
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conten-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

//Delete movie by id controller
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conten-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies) //getting all d movies after deletion
}

//Get a single movie by id controller
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conten-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item) //getting all d items
			return
		}
	}
}

//Create a movie controller
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conten-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie) //new movie inside movie
	json.NewEncoder(w).Encode(movie)
}

//Update a movie controller
func updateMovie(w http.ResponseWriter, r *http.Request) {

	//set json content type
	w.Header().Set("Conten-Type", "application/json")
	//params
	params := mux.Vars(r)
	//loops over movies,range
	//delete the movies wih the id that you have sent
	//add a new movie - the movie that we have sent in the body of postman
	for index, item := range movies {
		if item.ID == params["id"] { // we found the id here
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(100000000))
			movies = append(movies, movie) //new movie inside movie
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "1915", Title: "Movie One", Director: &Director{Firstname: "Aamir", Lastname: "Waris"}})
	movies = append(movies, Movie{ID: "2", Isbn: "42377", Title: "Movie Two", Director: &Director{Firstname: "Ayana", Lastname: "Rukasar"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at por 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
