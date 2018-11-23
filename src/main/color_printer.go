package main

import "fmt"

const (
    // linux终端显示转义控制字符esc
    linux_escape = 0x1B

    // 文字显示模式
    linux_display_type_normal  = 0
    linux_display_type_hilight = 1

    // 文字前景
    linux_foreground_red       = 31
    linux_foreground_green     = 32
    linux_foreground_yellow    = 33
    linux_foreground_blue      = 34
    linux_foreground_bluegreen = 36
    linux_foreground_white     = 37

    // 级别
    linux_foreground_level_err     = linux_foreground_red
    linux_foreground_level_warn    = linux_foreground_yellow
    linux_foreground_level_info    = linux_foreground_green
    linux_foreground_level_debug   = linux_foreground_blue
    linux_foreground_level_verbose = linux_foreground_white
)

func linux_print(foreground int, display int, s string) {
    const gap = " "
    fmt.Printf("%s%c[%d;%dm%s%c[0m \n", gap, linux_escape, display, foreground, s, linux_escape)
}

func linux_print_err(s string) {
    linux_print(linux_foreground_level_err, linux_display_type_hilight, s)
}

func linux_print_warn(s string) {
    linux_print(linux_foreground_level_warn, linux_display_type_normal, s)
}

func linux_print_info(s string) {
    linux_print(linux_foreground_level_info, linux_display_type_normal, s)
}

func linux_print_debug(s string) {
    linux_print(linux_foreground_level_debug, linux_display_type_normal, s)
}

func linux_print_verbose(s string) {
    linux_print(linux_foreground_level_verbose, linux_display_type_normal, s)
}

func test_color_printer_linux() {
    // 前景 背景 颜色
    // ---------------------------------------
    // 30  40  黑色
    // 31  41  红色
    // 32  42  绿色
    // 33  43  黄色
    // 34  44  蓝色
    // 35  45  紫红色
    // 36  46  青蓝色
    // 37  47  白色
    //
    // 代码 意义
    // -------------------------
    //  0  终端默认设置
    //  1  高亮显示
    //  4  使用下划线
    //  5  闪烁
    //  7  反白显示
    //  8  不可见

    drange := []int{0, 1, 4, 5, 7, 8}
    for b := 40; b <= 47; b++ { // 背景色彩 = 40-47
        for f := 30; f <= 37; f++ { // 前景色彩 = 30-37
            for _, d := range drange { // 显示方式 = 0,1,4,5,7,8
                fmt.Printf(" %c[%d;%d;%dm%s(f=%d,b=%d,d=%d)%c[0m ", 0x1B, d, b, f, "", f, b, d, 0x1B)
            }
            fmt.Println("")
        }
    }
    fmt.Println("")

}
