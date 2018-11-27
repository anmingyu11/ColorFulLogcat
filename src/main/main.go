package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "os/signal"
    "printer"
    "syscall"
)

const (
    CMD_ADB    = "adb"
    CMD_LOGCAT = "logcat"
)

//处理信号
func init_signal_handler() {
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGSTOP)

    go func() {
        sig := <-sigs
        fmt.Println("ReceivedSignal : ", sig.String(), " : ", sig.Signal)
        os.Exit(1);
    }()
}

func construct_printer() (p printer.Color_printer_i) {
    return printer.NewPrinter(printer.Color_printer_params{1})
}

func run_logcat_colorful_printer(p printer.Color_printer_i) {
    cmd := exec.Command(CMD_ADB, CMD_LOGCAT)
    stdout, _ := cmd.StdoutPipe()
    r := bufio.NewReader(stdout)

    go func() {
        fmt.Println("-- waiting for devices --")
        for {
            bytes, _, _ := r.ReadLine()
            line := string(bytes)
            p.PrintLog(line)
        }
    }()

    cmd.Start()
    cmd.Wait()
}

func main() {
    //初始化信号处理程序
    init_signal_handler()
    //构造打印机
    c_p := construct_printer()
    //执行主程序
    run_logcat_colorful_printer(c_p)
}
