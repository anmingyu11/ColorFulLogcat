package printer

import (
    "regexp"
    "strings"
    "util"
)

// 解析对应Log行的Tag
func parseLine(line string) (tag string, l string) {
    regex, err := regexp.Compile("\\s[VDIWE]\\s");
    util.PanicErr(err)

    tag = regex.FindString(line);
    return strings.Trim(tag, " "), line
}
