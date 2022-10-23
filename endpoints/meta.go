package endpoints

import (
	"edu/letu/wan/database"

	"github.com/gin-gonic/gin"
)

func MetaNames(c *gin.Context) {
	if database.IsAuthorized(c) {
		return
	}
	
	c.JSON(200, gin.H{
		"names": []string{
			"John",
			"Paul",
			"George",
			"Ringo",
		},
	})
}

func MetaPictures(c *gin.Context) {
	if database.IsAuthorized(c) {
		return
	}

	c.JSON(200, gin.H{
		"pictures": []string{
			"https://dummyimage.com/100x100/000/fff",
			"https://dummyimage.com/100x100/000/ddd",
			"https://dummyimage.com/100x100/000/bbb",
			"https://dummyimage.com/100x100/000/999",
		},
	})
}