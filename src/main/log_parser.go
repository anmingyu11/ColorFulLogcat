package main

import (
    "fmt"
    "regexp"
    "strings"
)
//Todo 折行

// 解析对应Log行的Tag
func parseLine(line string) string {
    regex, err := regexp.Compile("\\s[VDIWE]\\s");
    panicErr(err)

    tag := regex.FindString(line);
    return strings.Trim(tag, " ")
}

func printColorfulLine(line string) {
    
    tag := parseLine(line)
    switch tag {
    case "V":
        linux_print_verbose(line)
    case "D":
        linux_print_debug(line)
    case "I":
        linux_print_info(line)
    case "W":
        linux_print_warn(line)
    case "E":
        linux_print_err(line)
    default:
        fmt.Println(line)
    }
}

func testPrintColorfulLine() {
    printColorfulLine(" E fuck")
    printColorfulLine(" W fuck")
    printColorfulLine(" I fuck")
    printColorfulLine(" D fuck")
    printColorfulLine(" V fuck")
}
