package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/rs/cors"
)

// UserData represents the JSON structure for user data
type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LeaderboardEntry represents the JSON structure for leaderboard entry
type LeaderboardEntry struct {
	Username string `json:"username"`
	Wins     int    `json:"wins"`
	Losses   int    `json:"losses"`
}

func main() {
	ctx := context.Background()

	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-18033.c322.us-east-1-2.ec2.cloud.redislabs.com:18033",
		Password: "wkykaO0bL69zK9VpPRcijKS8QszZpwCo",
		DB:       0,
	})
	defer client.Close()
	
	// Test Redis connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	fmt.Println("Connected to Redis")

	// Define an HTTP handler for login
	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		// Decode JSON request body
		var userData UserData
		if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		// Store user data in Redis
		userKey := fmt.Sprintf("user:%s", userData.Username)
		if err := client.HSet(ctx, userKey, "password", userData.Password).Err(); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Respond with success
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User %s logged in successfully", userData.Username)
	})
	// Define an HTTP handler for leaderboard
	// Define an HTTP handler for leaderboard
	// Define an HTTP handler for leaderboard
	// Define an HTTP handler for leaderboard
	http.HandleFunc("/leaderboard", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Get leaderboard data from Redis
			leaderboard := make([]LeaderboardEntry, 0)
			keys, err := client.Keys(ctx, "leaderboard:*").Result()
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			for _, key := range keys {
				fields, err := client.HMGet(ctx, key, "wins", "losses").Result()
				if err != nil {
					http.Error(w, "Internal server error", http.StatusInternalServerError)
					return
				}
				wins, _ := strconv.Atoi(fields[0].(string))
				losses, _ := strconv.Atoi(fields[1].(string))
				username := key[len("leaderboard:"):]
				leaderboard = append(leaderboard, LeaderboardEntry{
					Username: username,
					Wins:     wins,
					Losses:   losses,
				})
			}

			// Respond with leaderboard data
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(leaderboard)
		case http.MethodPost:
			// Decode JSON request body
			var leaderboardData struct {
				Username string `json:"username"`
				GameWon  int    `json:"gameWon"`
				LostGame int    `json:"lostGame"`
			}
			if err := json.NewDecoder(r.Body).Decode(&leaderboardData); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}

			// Check if the user already exists
			leaderboardKey := fmt.Sprintf("leaderboard:%s", leaderboardData.Username)
			wins, _ := client.HGet(ctx, leaderboardKey, "wins").Int()
			losses, _ := client.HGet(ctx, leaderboardKey, "losses").Int()

			if wins == 0 && losses == 0 {
				// New user, add all entries
				if err := client.HSet(ctx, leaderboardKey, "wins", leaderboardData.GameWon, "losses", leaderboardData.LostGame).Err(); err != nil {
					http.Error(w, "Internal server error", http.StatusInternalServerError)
					return
				}
				// Respond with success
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "New user %s added to the leaderboard", leaderboardData.Username)
			} else {
				// Existing user, update win/loss counts
				wins += leaderboardData.GameWon
				losses += leaderboardData.LostGame
				if err := client.HSet(ctx, leaderboardKey, "wins", wins, "losses", losses).Err(); err != nil {
					http.Error(w, "Internal server error", http.StatusInternalServerError)
					return
				}
				// Respond with success
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "Leaderboard updated for user %s", leaderboardData.Username)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	// Define an HTTP handler for leaderboard in descending order
	http.HandleFunc("/leaderboard-desc", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Get leaderboard data from Redis
			leaderboard := make([]LeaderboardEntry, 0)
			keys, err := client.Keys(ctx, "leaderboard:*").Result()
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			for _, key := range keys {
				fields, err := client.HMGet(ctx, key, "wins", "losses").Result()
				if err != nil {
					http.Error(w, "Internal server error", http.StatusInternalServerError)
					return
				}
				wins, _ := strconv.Atoi(fields[0].(string))
				losses, _ := strconv.Atoi(fields[1].(string))
				username := key[len("leaderboard:"):]
				leaderboard = append(leaderboard, LeaderboardEntry{
					Username: username,
					Wins:     wins,
					Losses:   losses,
				})
			}

			// Sort leaderboard entries in descending order based on wins
			sort.Slice(leaderboard, func(i, j int) bool {
				return leaderboard[i].Wins > leaderboard[j].Wins
			})

			// Respond with leaderboard data
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(leaderboard)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Set up CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
	})

	// Use the CORS middleware
	handler := c.Handler(http.DefaultServeMux)

	// Start the HTTP server with CORS support
	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", handler)
}
