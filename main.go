package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	appUser "github.com/christhianjesus/priverion-challenge/internal/application/user"
	"github.com/christhianjesus/priverion-challenge/internal/infrastructure/advertisement"
	"github.com/christhianjesus/priverion-challenge/internal/infrastructure/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	authorized := r.Group("/")

	authorized.Use(JWT())
	{
		authorized.GET("/user/:user", userHandler.GetUser)
		authorized.PATCH("/user/:user", userHandler.UpdateUser)
		authorized.DELETE("/user/:user", userHandler.DeleteUser)

		authorized.GET("/advertisement/", advertisementHandler.GetAll)
		authorized.POST("/advertisement/", advertisementHandler.Create)
		authorized.GET("/advertisement/:advertisement", advertisementHandler.GetOne)
		authorized.PATCH("/advertisement/:advertisement", advertisementHandler.Update)
		authorized.DELETE("/advertisement/:advertisement", advertisementHandler.Delete)
	}

	r.Run()
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if strings.HasPrefix(h, "Bearer ") {
			tokenString := strings.Replace(h, "Bearer ", "", 1)

			token, err := jwt.ParseWithClaims(tokenString, &user.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			if claims, ok := token.Claims.(*user.JwtCustomClaims); ok && token.Valid {
				// c.MustGet(gin.AuthUserKey).
				c.Set(gin.AuthUserKey, claims.Username)
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		return
	}
}
