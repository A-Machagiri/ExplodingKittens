package main

import (
    "context"
    "encoding/json"
	"github.com/go-redis/redis/v8"
    "fmt"
    "net/http"
)

func StartHandler(ctx context.Context, client *redis.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
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
    }
}
