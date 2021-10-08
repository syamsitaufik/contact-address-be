// Contact address
// 	- Create contact
// 		* Name, phone number, address
// 	- Read contact
// 	- Update contact
// 		* Name, phone number, address
// 	- Delete contact
// Bonus:
// 	- Categorise contacts - Family, Friend, Work
// 	- Search contact using
// 		* Full characters
// 		* partial characters - before, mid, after
// 	- Add to favourite
// 	- List data starting with this hierarchical order
// 		* Recent searched
// 		* Favourite
// 		* alphabetical
// 	- Add contact image

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syamsitaufik/contact-address-be/controllers"
	"github.com/syamsitaufik/contact-address-be/models"
)

func main() {
	router := gin.Default()

	models.ConnectDataBase()

	router.GET("/", getHelloWorld)

	router.GET("/contacts", controllers.FindContacts)
	router.POST("/contact", controllers.CreateContact)

	router.Run("localhost:3000")
}

func getHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Hello Wordl!"})
}
