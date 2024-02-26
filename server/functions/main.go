package main

import (
    "context"
    "fmt"
    "net/http"
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
    client := createRedisClient(ctx)
    defer client.Close()

    // Set up HTTP handlers
    http.HandleFunc("/start", StartHandler(ctx,client))
    http.HandleFunc("/leaderboard", LeaderboardHandler(ctx, client))
    http.HandleFunc("/leaderboard-desc", leaderboardDescHandler(ctx, client))

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

func createRedisClient(ctx context.Context) *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr:     "redis-18033.c322.us-east-1-2.ec2.cloud.redislabs.com:18033",
        Password: "wkykaO0bL69zK9VpPRcijKS8QszZpwCo",
        DB:       0,
    })

    // Test Redis connection
    _, err := client.Ping(ctx).Result()
    if err != nil {
        fmt.Println("Error connecting to Redis:", err)
        // Handle error appropriately, maybe exit or return nil
    }

    return client
}
