package util

import (
	"encoding/hex"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var lobbyCodeRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
func LobbyCode() string {
    rand.Seed(time.Now().UnixNano())
    codeLength, err := strconv.Atoi(os.Getenv("LOBBY_CODE_LENGTH"))
    if err != nil {
        codeLength = 4
    }
    b := make([]rune, codeLength)
    for i := range b {
        b[i] = lobbyCodeRunes[rand.Intn(len(lobbyCodeRunes))]
    }
    code := string(b)
    if IsProfane(code) {
        return LobbyCode()
    }
    return code
}

func GenerateToken() string {
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, 32)
    if _, err := rand.Read(b); err != nil {
        return ""
    }
    return hex.EncodeToString(b)
}

func RandomKeyFromMap(m map[int]string) int {
    rand.Seed(time.Now().UnixNano())
    var keys []int
    for k := range m {
        keys = append(keys, k)
    }
    return keys[rand.Intn(len(keys))]
}

func ValidateKeyFromMap(m map[int]string, key int) int {
    _, ok := m[key]
    if ok {
        return key
    }
    return RandomKeyFromMap(m)
}

func RemoveFromSlice[S ~[]E, E any](s S, index int) S {
    if index < len(s) - 1 {
        return append(s[:index], s[index+1:]...)
    } else {
        return s[:index]
    }
}
