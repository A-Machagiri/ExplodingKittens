package main

import (
    "context"
    "encoding/json"
    "net/http"
	"github.com/go-redis/redis/v8"
    "sort"
    "strconv"
)

func leaderboardDescHandler(ctx context.Context, client *redis.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
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
    }
}
