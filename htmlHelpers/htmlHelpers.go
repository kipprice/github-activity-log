package htmlHelpers

import (
	"fmt"
)

func Header(header string, level int) {
	fmt.Println("")
	fmt.Printf("<h%v>%v</h%v>", level, header, level)
}

func Span(innertext string, classname string) {
	fmt.Printf("<span class='%v'>%v</span>", classname, innertext)
}

func A(innertext string, link string, classname string) {
	fmt.Printf("<a href='%v' class='%v'>%v</a>", link, classname, innertext)
}

