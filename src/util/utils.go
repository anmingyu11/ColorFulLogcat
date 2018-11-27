package util

import "runtime"

func PanicErr(err error) {
    if err != nil {
        panic(err)
    }
}

func GetOS() string {
    return runtime.GOOS
}
