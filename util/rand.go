package util

import (
	"encoding/hex"
	"math/rand"
	"os"
	"strconv"
)

var lobbyCodeRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
func LobbyCode() string {
    codeLength, err := strconv.Atoi(os.Getenv("LOBBY_CODE_LENGTH"))
    if err != nil {
        codeLength = 4
    }
    b := make([]rune, codeLength)
    for i := range b {
        b[i] = lobbyCodeRunes[rand.Intn(len(lobbyCodeRunes))]
    }
    return string(b)
}
//TODO should we ensure the output is never explicit?

func GenerateToken() string {
    b := make([]byte, 32)
    if _, err := rand.Read(b); err != nil {
        return ""
    }
    return hex.EncodeToString(b)
}
