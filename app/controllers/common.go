package controllers

import (
    "strconv"
    "log"
    "math/rand"
    "encoding/hex"


    "crypto/sha256"
    // "crypto/aes"
    // "crypto/cipher"
    // "crypto/rand"
    "encoding/base64"
    // "errors"
    // "fmt"
    // "io"
)


const (
    _      = iota
    KB int = 1 << (10 * iota)
    MB
    GB
)

func parseUintOrDefault(intStr string, _default uint64) uint64 {
    if value, err := strconv.ParseUint(intStr, 0, 64); err != nil {
        return _default
    } else {
        return value
    }
}

func parseIntOrDefault(intStr string, _default int64) int64 {
    if value, err := strconv.ParseInt(intStr, 0, 64); err != nil {
        return _default
    } else {
        return value
    }
}

func checkErr(err error, msg string) {
    if err != nil {
        log.Println(msg, err)
    }
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func randString() string {
    randBytes := make([]byte, 16)
    rand.Read(randBytes)
    return hex.EncodeToString(randBytes)
}

// encrypt string to base64 crypto using SHA
func EncryptSHA256(text string) string {
    h := sha256.New()
    h.Write([]byte(text))
    s := base64.URLEncoding.EncodeToString(h.Sum(nil))
    return (s)
}
