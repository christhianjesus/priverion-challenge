package user

import (
	"net/http"
	"time"

	"github.com/christhianjesus/priverion-challenge/internal/domain/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userHandler struct {
	us user.UserService
}

type JwtCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewUserHandler(us user.UserService) *userHandler {
	return &userHandler{us}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	user := mongoUser{}

	if err := c.BindJSON(&user.e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentTime := time.Now()
	user.e.ID = primitive.NewObjectIDFromTimestamp(currentTime)
	user.e.CreatedAt = currentTime
	user.e.UpdatedAt = currentTime

	if err := h.us.CreateUser(ctx, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

type loginJSON struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *userHandler) AuthUser(c *gin.Context) {
	ctx := c.Request.Context()
	login := loginJSON{}

	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.us.AuthUser(ctx, login.Username, login.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	claims := &JwtCustomClaims{
		Username: user.ID(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": signedToken})
}

func (h *userHandler) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.Param("user")

	user, err := h.us.GetUser(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mongoUser := user.(*mongoUser)

	c.JSON(http.StatusOK, mongoUser.e)
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	user := mongoUser{}
	userID := c.Param("user")

	if err := c.BindJSON(&user.e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.e.ID, _ = primitive.ObjectIDFromHex(userID)
	currentTime := time.Now()
	user.e.UpdatedAt = currentTime

	if err := h.us.UpdateUser(ctx, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.Param("user")

	if err := h.us.DeleteUser(ctx, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}
