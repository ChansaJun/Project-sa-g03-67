package controller

import (
	"net/http"

	"errors"
	"time"

	"example.com/project-sa-g03/config"
	"example.com/project-sa-g03/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// POST /users
func CreatePayment(c *gin.Context) {
	var payments entity.Payment

	// Bind JSON payload to Payment struct
	if err := c.ShouldBindJSON(&payments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	// Create new payment entry
	p := entity.Payment{
		Name:       payments.Name,
		Date:       payments.Date,
		TotalPrice: payments.TotalPrice,
		Slip:       payments.Slip,
		ReserveID:  payments.ReserveID, // Ensure ReserveID is correctly bound
	}

	// Save to database
	if err := db.Create(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created successfully", "data": p})
}

// GET /users
func ListPayment(c *gin.Context) {
	// Get the ID parameter from the URL
	ID := c.Param("id")

	// Define a struct to hold the result set
	var payments []struct {
		Name      string    `json:"name"`
		Date      time.Time `json:"date"`
		TotalPice float32   `json:"total_price"`
		Slip      string    `json:"slipt"`
	}

	// Get the database connection
	db := config.DB()

	// Query to join reserves and shops tables and select the required fields
	results := db.Table("payments").
		Select("payments.date, payments.total_price").
		Joins(" join shops on payments.reserve_id = shops.id").
		Joins(" join reserves on payments.reserve_id = reserves.id").
		Where("reserves.id = ?", ID). // Filter by shop ID matching the parameter received
		Scan(&payments)               // Scan the results into the reserves struct

	// Check for errors in the query
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	// Return the results as JSON
	c.JSON(http.StatusOK, payments)
}

func GetReserveById(c *gin.Context) {
	ID := c.Param("id")
	var user entity.Reserve

	db := config.DB()

	// Query the user by ID
	results := db.Where("id = ?", ID).First(&user)
	if results.Error != nil {
		if errors.Is(results.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": results.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetReserveDetails(c *gin.Context) {
	var reservesDetails []entity.ReserveDetails
	db := config.DB()
	db.Find(&reservesDetails)
	c.JSON(http.StatusOK, &reservesDetails)
}
