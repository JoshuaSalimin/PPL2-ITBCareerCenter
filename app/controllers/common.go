package controllers

import (
    "strconv"
    "log"
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