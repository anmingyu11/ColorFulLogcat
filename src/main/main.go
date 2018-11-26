package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "os/signal"
    "syscall"
)

//处理信号
func initSignalHandler() {
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGSTOP)

    go func() {
        sig := <-sigs
        fmt.Println("ReceivedSignal : ", sig.String())
        os.Exit(1);
    }()
}

func main() {
    //Todo : 命令行参数
    //Todo : 使用通道优化
    cmd := exec.Command("adb", "logcat")
    stdout, _ := cmd.StdoutPipe()
    r := bufio.NewReader(stdout)
    
    go func() {
        fmt.Println("-- waiting for devices --")
        for {
            line, _, _ := r.ReadLine()
            PrintColorfulLine(string(line))
        }
    }()

    cmd.Start()
    cmd.Wait()
}
