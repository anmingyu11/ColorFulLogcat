package main

import "fmt"

const (
    // linux终端显示转义控制字符esc
    LINUX_ESCAPE = 0X1B

    // 文字显示模式
    LINUX_DISPLAY_TYPE_NORMAL  = 0
    LINUX_DISPLAY_TYPE_HILIGHT = 1

    // 文字前景
    LINUX_FOREGROUND_RED       = 31
    LINUX_FOREGROUND_GREEN     = 32
    LINUX_FOREGROUND_YELLOW    = 33
    LINUX_FOREGROUND_BLUE      = 34
    LINUX_FOREGROUND_BLUEGREEN = 36
    LINUX_FOREGROUND_WHITE     = 37

    // 级别
    LINUX_FOREGROUND_LEVEL_ERR     = LINUX_FOREGROUND_RED
    LINUX_FOREGROUND_LEVEL_WARN    = LINUX_FOREGROUND_YELLOW
    LINUX_FOREGROUND_LEVEL_INFO    = LINUX_FOREGROUND_GREEN
    LINUX_FOREGROUND_LEVEL_DEBUG   = LINUX_FOREGROUND_BLUE
    LINUX_FOREGROUND_LEVEL_VERBOSE = LINUX_FOREGROUND_WHITE
)

//Todo 结构化，接口化
type linux_printer struct {
    display_type int
}

func Linux_print_err(s string) {
    linux_print(LINUX_FOREGROUND_LEVEL_ERR, LINUX_DISPLAY_TYPE_HILIGHT, s)
}

func Linux_print_warn(s string) {
    linux_print(LINUX_FOREGROUND_LEVEL_WARN, LINUX_DISPLAY_TYPE_HILIGHT, s)
}

func Linux_print_info(s string) {
    linux_print(LINUX_FOREGROUND_LEVEL_INFO, LINUX_DISPLAY_TYPE_HILIGHT, s)
}

func Linux_print_debug(s string) {
    linux_print(LINUX_FOREGROUND_LEVEL_DEBUG, LINUX_DISPLAY_TYPE_HILIGHT, s)
}

func Linux_print_verbose(s string) {
    linux_print(LINUX_FOREGROUND_LEVEL_VERBOSE, LINUX_DISPLAY_TYPE_HILIGHT, s)
}

func linux_print(foreground int, display int, s string) {
    const gap = " "
    fmt.Printf("%s%c[%d;%dm%s%c[0m \n", gap, LINUX_ESCAPE, display, foreground, s, LINUX_ESCAPE)
}

func Test_color_printer_linux() {
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
