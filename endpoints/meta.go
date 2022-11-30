package endpoints

import (
	"edu/letu/wan/database"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type MetaNamesJSON struct {
	Adjectives map[int]string `json:"adjectives"`
	Nouns map[int]string `json:"nouns"`
}

type MetaAvatarsJSON struct {
	Avatars map[int]string `json:"avatars"`
}

func MetaNames(c *gin.Context) {
	if !database.IsAuthorized(c) {
		return
	}

	content, err := ioutil.ReadFile("./resources/names.json")
    if err != nil {
        c.AbortWithStatus(500)
    }
 
	// validate json before sending it
    var payload MetaNamesJSON
    err = json.Unmarshal(content, &payload)
    if err != nil {
        c.AbortWithStatus(500)
    }

	c.JSON(200, payload)
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