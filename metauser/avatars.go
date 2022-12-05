package metauser

import (
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type MetaAvatars struct {
	Avatars map[int]string
}

type MetaAvatarKeysJSON struct {
	Avatars []int `json:"avatars"`
}

func GetMetaAvatars() MetaAvatars {
	file, err := os.Open("./resources/avatars")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	names, err := file.Readdirnames(-1)
	if err != nil {
		panic(err)
	}

	avatars := make(map[int]string)
	for _, basename := range names {
		name := strings.TrimSuffix(basename, filepath.Ext(basename))
		nameInt, err := strconv.Atoi(name)
		if err != nil {
			continue
		}
		avatars[nameInt] = basename
	}
	
	return MetaAvatars{Avatars: avatars}
}

func GetMetaAvatarKeys() MetaAvatarKeysJSON {
	avatars := GetMetaAvatars().Avatars
	keys := make([]int, len(avatars))
	i := 0
	for k := range avatars {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return MetaAvatarKeysJSON{Avatars: keys}
}