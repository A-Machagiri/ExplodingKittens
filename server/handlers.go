// package main

// import (
//     "context"
//     "fmt"
//     "github.com/go-redis/redis/v8"
// )

// func main() {
//     // Connect to Redis
// 	client:=redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 		Password: "",
// 		DB:0,
// 	})
// 	ctx:=context.Background()
// 	defer client.Close()
// 	userKey:="Key:"
// 	fieldsAndValues:=map[string]interface{}{
// 		"username2":"Giri",
// 		"password2":"#123",
// 	}
// 	err:=client.HMSet(ctx,userKey,fieldsAndValues).Err()
// 	if(err!=nil){
// 		fmt.Println("Error setting fields",err)
// 		return
// 	}
// 	fmt.Println("Fields set successfully")
// 	// err := client.HDel(context.Background(), userKey, "password2").Err()
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// fmt.Println("Removed model and brand from key:1")
// 	field:="username1"
// 	username,err:=client.HGet(ctx,userKey,field).Result()
// 	if err!=nil{
// 		fmt.Println("Error getting value:", err)
//         return
// 	}
// 	fmt.Printf("Value of field %s for key %s: %s\n", field, userKey, username)
// }