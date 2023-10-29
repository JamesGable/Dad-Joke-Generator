package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type JokeResponse struct {
	Joke string `json:"joke"`
}

var jokes = []string{
	"Why don't skeletons fight each other? They don't have the guts!",
	"What do you call fake spaghetti? An impasta!",
	"Why couldn't the bicycle stand up by itself? It was two tired!",
	"Why don't we ever tell secrets on a farm? Because the potatoes have eyes and the corn has ears!",
	"Did you hear about the restaurant on the moon? Great food, no atmosphere.",
	"I used to play piano by ear, but now I use my hands.",
	"Why did the scarecrow win an award? Because he was outstanding in his field!",
	"Parallel lines have so much in common. It is a shame they will never meet.",
	"What's the best time to go to the dentist? Tooth-hurty!",
	"Did you hear about the guy who invented Lifesavers? They say he made a mint.",
	"What did the ocean say to the shore? Nothing, it just waved.",
	"Why did the bicycle fall over? Because it was two-tired!",
	"I told my wife she should embrace her mistakes. She gave me a hug.",
	"How do you organize a space party? You planet!",
	"What do you call an alligator in a vest? An investigator.",
	"Did you hear about the kidnapping at the park? They woke him up.",
	"What did one wall say to the other wall? I'll meet you at the corner.",
	"I would avoid the sushi if I was you. It's a little fishy.",
	// Add more jokes here
}

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	randomJoke := getRandomJoke()
	response := JokeResponse{Joke: randomJoke}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to generate JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func getRandomJoke() string {
	randomIndex := rand.Intn(len(jokes))
	return jokes[randomIndex]
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/get-joke", jokeHandler)

	handler := enableCORS(http.DefaultServeMux)
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	fmt.Println("Server listening on port 8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server failed to start")
	}
}
