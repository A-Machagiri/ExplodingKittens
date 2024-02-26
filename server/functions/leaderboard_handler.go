package main

import (
    "context"
    "encoding/json"
    "fmt"
	"github.com/go-redis/redis/v8"
    "net/http"
    "strconv"
)

func LeaderboardHandler(ctx context.Context, client *redis.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            // Your existing /leaderboard GET handler logic
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
    }
}
