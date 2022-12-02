package metauser

import (
	"encoding/json"
	"os"
)

type MetaNamesJSON struct {
	Adjectives map[int]string `json:"adjectives"`
	Nouns map[int]string `json:"nouns"`
}

type MetaAvatarsJSON struct {
	Avatars map[int]string `json:"avatars"`
}

func GetMetaNames() MetaNamesJSON {
	content, err := os.ReadFile("./resources/names.json")
    if err != nil {
		panic(err)
    }
    var payload MetaNamesJSON
    err = json.Unmarshal(content, &payload)
	if err != nil {
		panic(err)
	}
	return payload
}