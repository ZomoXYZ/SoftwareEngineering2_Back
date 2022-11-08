package endpoints

import (
	"edu/letu/wan/database"

	"github.com/gin-gonic/gin"
)

func MetaNames(c *gin.Context) {
	if !database.IsAuthorized(c) {
		return
	}
	
	// TODO make it a map instead of a list
	// {"names": {"0": "name", "1": "name", "2": "name"}}
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
	if !database.IsAuthorized(c) {
		return
	}

	// TODO make it a map instead of a list
	// {"pictures": {"0": "url", "1": "url", "2": "url"}}
	c.JSON(200, gin.H{
		"pictures": []string{
			"https://dummyimage.com/100x100/000/fff",
			"https://dummyimage.com/100x100/000/ddd",
			"https://dummyimage.com/100x100/000/bbb",
			"https://dummyimage.com/100x100/000/999",
		},
	})
}