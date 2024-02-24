package controller

import (
	"net/http"
	"tugas-oti/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type eventInput struct {
	Nama           string `json : "nama_event"`
	Deskripsi      string `json : "deskripsi_event"`
	TanggalMulai   string `json : "tanggal_mulai"`
	TanggalSelesai string `json : "tanggal_selesai"`
	Lokasi         string `json : "lokasi"`
	Capacity       int64  `json : "kapasitas_peserta"`
}

// GetAllEvent godoc
// @Summary Get All Event.
// @Description Get a list of Event.
// @Tags Event
// @Produce json
// @Success 200 {object} []models.Event
// @Router /event [get]
func GetAllEvent(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var event []models.Event
	db.Preload("Enrollment").Find(&event)

	c.JSON(http.StatusOK, gin.H{"data": event})
}

// CreateEvent godoc
// @Summary Create New Event.
// @Description Creating a new event.
// @Tags Event
// @Param Body body eventInput true "the body to create a new event"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Event
// @Router /event [post]
func CreateEvent(c *gin.Context) {
	var input eventInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(input.Nama) > 190 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama event tidak boleh lebih dari 190 Character"})
	}

	event := models.Event{Nama: input.Nama,
		Deskripsi:      input.Deskripsi,
		TanggalMulai:   input.TanggalMulai,
		TanggalSelesai: input.TanggalSelesai,
		Lokasi:         input.Lokasi,
		Capacity:       input.Capacity}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&event)

	c.JSON(http.StatusOK, gin.H{"data": event})
}

// GetEventById godoc
// @Summary Get Event.
// @Description Get an Event by id.
// @Tags Event
// @Produce json
// @Param id path string true "event id"
// @Success 200 {object} models.Event
// @Router /event/{id} [get]
func GetEventByID(c *gin.Context) {
	var event models.Event

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": event})
}

// UpdateEvent godoc
// @Summary Update Event.
// @Description Update Event by id.
// @Tags Event
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "event id"
// @Param Body body eventInput true "the body to update event"
// @Success 200 {object} models.Event
// @Router /event/{id} [put]
func UpdateEvent(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var event models.Event
	if err := db.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input eventInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Event
	updatedInput.Nama = input.Nama
	updatedInput.Deskripsi = input.Deskripsi
	updatedInput.TanggalMulai = input.TanggalMulai
	updatedInput.Lokasi = input.Lokasi
	updatedInput.Capacity = input.Capacity

	if len(updatedInput.Nama) > 190 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama event tidak boleh lebih dari 190 Character"})
	}

	db.Model(&event).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": event})
}

// DeleteEvent godoc
// @Summary Delete an Event
// @Tags Event
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "event id"
// @Success 200 {object} map[string]boolean
// @Router /event/{id} [delete]
func DeleteEvent(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var event models.Event
	if err := db.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&event)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
