package controllers

import (
    "strconv"
    "log"
    "math/rand"
    "encoding/hex"


    // "crypto/aes"
    // "crypto/cipher"
    // "crypto/rand"
    // "encoding/base64"
    // "errors"
    // //"fmt"
    // "io"
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

// func encrypt(key []byte, text string) (string, error) {

//     block, err := aes.NewCipher(key)
//     if err != nil {
//         return "FAIL", err
//     }
//     b := base64.StdEncoding.EncodeToString([]byte(text))
//     ciphertext := make([]byte, aes.BlockSize+len(b))
//     iv := ciphertext[:aes.BlockSize]
//     if _, err := io.ReadFull(rand.Reader, iv); err != nil {
//         return "FAIL", err
//     }
//     cfb := cipher.NewCFBEncrypter(block, iv)
//     cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
//     return string(ciphertext), nil
// }

// func decrypt(key []byte, text string) (string, error) {
//     block, err := aes.NewCipher(key)
//     if err != nil {
//         return "FAIL"   , err
//     }
//     if len(text) < aes.BlockSize {
//         return "FAIL", errors.New("ciphertext too short")
//     }
//     iv := []byte(text[:aes.BlockSize])
//     text = text[aes.BlockSize:]
//     cfb := cipher.NewCFBDecrypter(block, iv)
//     cfb.XORKeyStream([]byte(text), []byte(text))
//     data, err := base64.StdEncoding.DecodeString(string(text))
//     if err != nil {
//         return "FAIL", err
//     }
//     return string(data), nil
// }