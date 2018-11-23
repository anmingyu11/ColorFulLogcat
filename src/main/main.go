package main

import (
    "bufio"
    "os/exec"
)

//Todo 信号处理程序
func main() {
    //Todo 使用通道优化
    cmd := exec.Command("adb", "logcat")
    stdout, _ := cmd.StdoutPipe()
    r := bufio.NewReader(stdout)
    go func() {
        for {
            line, _, _ := r.ReadLine()
            printColorfulLine(string(line))
        }
    }()

    cmd.Start()
    cmd.Wait()
}
