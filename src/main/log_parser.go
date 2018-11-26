package main

import (
    "fmt"
    "regexp"
    "strings"
)

// 解析对应Log行的Tag
func parseLine(line string) string {
    regex, err := regexp.Compile("\\s[VDIWE]\\s");
    panicErr(err)

    tag := regex.FindString(line);
    return strings.Trim(tag, " ")
}

func PrintColorfulLine(line string) {

    tag := parseLine(line)
    switch tag {
    case "V":
        Linux_print_verbose(line)
    case "D":
        Linux_print_debug(line)
    case "I":
        Linux_print_info(line)
    case "W":
        Linux_print_warn(line)
    case "E":
        Linux_print_err(line)
    default:
        fmt.Println(line)
    }
}

func TestPrintColorfulLine() {
    PrintColorfulLine(" E fuck")
    PrintColorfulLine(" W fuck")
    PrintColorfulLine(" I fuck")
    PrintColorfulLine(" D fuck")
    PrintColorfulLine(" V fuck")
}
