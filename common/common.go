package common

import (
    "bytes"
)

func CheckError(e error) {
    if e != nil {
        panic(e)
    }
}

func Concat(s1, s2 string) string {
    var b bytes.Buffer
    b.WriteString(s1)
    b.WriteString(s2)
    return b.String()
}
