package main

import "runtime"

func panicErr(err error) {
    if err != nil {
        panic(err)
    }
}

func getOS() string {
    return runtime.GOOS
}
