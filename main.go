package main

import (
	"context"
	"log"
	"os"
	"time"

	appUser "github.com/christhianjesus/priverion-challenge/internal/application/user"
	"github.com/christhianjesus/priverion-challenge/internal/infrastructure/advertisement"
	"github.com/christhianjesus/priverion-challenge/internal/infrastructure/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURL := os.Getenv("ME_CONFIG_MONGODB_URL")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	userRepo := user.NewMongoUserRepository(client.Database("default"))
	userService := appUser.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	advertisementRepo := advertisement.NewMongoAdvertisementRepository(client.Database("default"))
	advertisementHandler := advertisement.NewAdvertisementHandler(advertisementRepo)

	r := gin.Default()
	r.POST("/login", userHandler.AuthUser)
	r.POST("/register", userHandler.CreateUser)

	r.GET("/user/:user", userHandler.GetUser)
	r.PATCH("/user/:user", userHandler.UpdateUser)
	r.DELETE("/user/:user", userHandler.DeleteUser)

	r.GET("/advertisement/", advertisementHandler.GetAll)
	r.POST("/advertisement/", advertisementHandler.Create)
	r.GET("/advertisement/:advertisement", advertisementHandler.GetOne)
	r.PATCH("/advertisement/:advertisement", advertisementHandler.Update)
	r.DELETE("/advertisement/:advertisement", advertisementHandler.Delete)
	r.Run()
}
