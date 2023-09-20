package advertisement

import (
	"net/http"
	"time"

	"github.com/christhianjesus/priverion-challenge/internal/domain/advertisement"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type advertisementHandler struct {
	ar advertisement.AdvertisementRepository
}

func NewAdvertisementHandler(ar advertisement.AdvertisementRepository) *advertisementHandler {
	return &advertisementHandler{ar}
}

func (h *advertisementHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	advertisements, err := h.ar.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mongoAdvertisements := make([]*encapsulatedMongoAdvertisement, 0, len(advertisements))
	for _, advertisement := range advertisements {
		mongoAdvertisements = append(mongoAdvertisements, &advertisement.(*mongoAdvertisement).e)
	}

	c.JSON(http.StatusOK, mongoAdvertisements)
}

func (h *advertisementHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	advertisement := mongoAdvertisement{}

	if err := c.BindJSON(&advertisement.e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentTime := time.Now()
	advertisement.e.ID = primitive.NewObjectIDFromTimestamp(currentTime)
	advertisement.e.CreatedAt = currentTime
	advertisement.e.UpdatedAt = currentTime
	advertisement.e.UserID = c.MustGet(gin.AuthUserKey).(string)

	if err := h.ar.Create(ctx, &advertisement); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (h *advertisementHandler) GetOne(c *gin.Context) {
	ctx := c.Request.Context()
	advertisementID := c.Param("advertisement")

	advertisement, err := h.ar.GetOne(ctx, advertisementID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mongoAdvertisement := advertisement.(*mongoAdvertisement)

	c.JSON(http.StatusOK, mongoAdvertisement.e)
}

func (h *advertisementHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()
	advertisement := mongoAdvertisement{}
	advertisementID := c.Param("advertisement")

	if err := c.BindJSON(&advertisement.e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	advertisement.e.ID, _ = primitive.ObjectIDFromHex(advertisementID)
	currentTime := time.Now()
	advertisement.e.UpdatedAt = currentTime

	if err := h.ar.Update(ctx, &advertisement); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (h *advertisementHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	advertisementID := c.Param("advertisement")

	if err := h.ar.Delete(ctx, advertisementID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}
