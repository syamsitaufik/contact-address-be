package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syamsitaufik/contact-address-be/models"
)

// GET /contacts
// Get all contacts

func FindContacts(c *gin.Context) {
	var contacts []models.Contact
	models.DB.Find(&contacts)

	c.JSON(http.StatusOK, gin.H{"data": contacts})
}
