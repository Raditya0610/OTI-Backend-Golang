package controller

import (
	"net/http"
	"tugas-oti/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type enrollmentInput struct {
	EventID   uint   `json : "event_id"`
	UserID    uint   `json : "user_id"`
	Informasi string `json : "informasi_tambahan"`
}

// GetAllEnrollment godoc
// @Summary Get all Enrollment.
// @Description Get a list of Enrollment.
// @Tags Enrollment
// @Produce json
// @Success 200 {object} []models.Enrollment
// @Router /enrollment [get]
func GetAllEnrollment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var enrollment []models.Enrollment
	db.Preload("Event").Find(&enrollment)

	c.JSON(http.StatusOK, gin.H{"data": enrollment})
}

// CreateEnrollment godoc
// @Summary Create New Enrollment.
// @Description Creating a new Enrollment.
// @Tags Enrollment
// @Param Body body enrollmentInput true "the body to create a new enrollment"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Enrollment
// @Router /enrollment [post]
func CreateEnrollment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input enrollmentInput
	var event models.Event
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.EventID).First(&event).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "eventID not found!"})
		return
	}

	var count int64
	db.Model(&models.Enrollment{}).Where("event_id = ?", input.EventID).Count(&count)

	if count >= event.Capacity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event sudah penuh"})
		return
	}

	enrollment := models.Enrollment{
		EventID:   input.EventID,
		UserID:    input.UserID,
		Informasi: input.Informasi}
	db.Create(&enrollment)

	c.JSON(http.StatusOK, gin.H{"data": enrollment})
}

// DeleteEnrollment godoc
// @Summary Delete an enrollment
// @Tags Enrollment
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "enrollment id"
// @Success 200 {object} map[string]boolean
// @Router /enrollment/{id} [delete]
func DeleteEnrollment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var enrollment models.Enrollment
	if err := db.Where("id = ?", c.Param("id")).First(&enrollment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&enrollment)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
