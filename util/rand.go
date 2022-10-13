package util

import "math/rand"

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberRunes = []rune("0123456789")
var allRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

//should we ensure the output is never explicit?
func RandChars(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}
func RandNums(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = numberRunes[rand.Intn(len(numberRunes))]
    }
    return string(b)
}
func RandAll(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = allRunes[rand.Intn(len(allRunes))]
    }
    return string(b)
}
//should we ensure the output is never explicit?
