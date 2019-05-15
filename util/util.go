package util

import (
    "strings"
)

func Strip(target string) string {
    return strings.Trim(target, "\n")
}
