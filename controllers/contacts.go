package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syamsitaufik/contact-address-be/models"
)

type CreateContactInput struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Address     string `json:"address"`
}

// GET /contacts
// Get all contacts
func FindContacts(c *gin.Context) {
	var contacts []models.Contact
	models.DB.Find(&contacts)

	c.JSON(http.StatusOK, gin.H{"data": contacts})
}

// POST /contact
// Create new contact
func CreateContact(c *gin.Context) {
	// Validate input
	var input CreateContactInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Contact
	contact := models.Contact{Name: input.Name, PhoneNumber: input.PhoneNumber, Address: input.Address}
	models.DB.Create(&contact)

	c.JSON(http.StatusOK, gin.H{"data": contact})
}
