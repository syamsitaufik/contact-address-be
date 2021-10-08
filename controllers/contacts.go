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

type UpdateContactInput struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

// GET /contacts
// Get all contacts
func FindContacts(c *gin.Context) {
	var contacts []models.Contact
	models.DB.Find(&contacts)

	c.JSON(http.StatusOK, gin.H{"data": contacts})
}

// POST /contacts
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
	// var contact models.Contact
	// c.BindJSON(&input)
	models.DB.Create(&contact)

	c.JSON(http.StatusOK, gin.H{"data": contact})
}

// GET /contacts/:id
// Find a contact
func FindContact(c *gin.Context) {
	var contact models.Contact

	if err := models.DB.Where("id = ?", c.Param("id")).First(&contact).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Contact not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contact})
}

// PATCH /contacts/:id
// Update a contact
func UpdateContact(c *gin.Context) {
	// Get model if exist
	var contact models.Contact

	if err := models.DB.Where("id = ?", c.Param("id")).First(&contact).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Contact not found!"})
		return
	}

	//Validate input
	var input UpdateContactInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.BindJSON(&contact)

	// models.DB.Model(&contact).Updates(input)
	models.DB.Save(&contact)

	c.JSON(http.StatusOK, gin.H{"data": contact})
}

// DELETE /contacts/:id
// Delete a contact

func DeleteContact(c *gin.Context) {
	// Get mode if exist
	var contact models.Contact
	if err := models.DB.Where("id = ?", c.Param("id")).First(&contact).Error; err != nil {
		return
	}

	models.DB.Delete(&contact)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
