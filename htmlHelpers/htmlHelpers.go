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

func StartPageHtml() {
	fmt.Printf("<!DOCTYPE><html><head><style>a { text-decoration :none !important; } body > * { width: 80vw; } </style></head><body style='font-family: Courier New; display: flex; align-items: center; flex-direction: column;'>")
}

func EndPageHtml() {
	fmt.Printf("</body></html>")
}

func StartGrid(columns int) {
	fmt.Printf("<div style='display: grid; grid-template-columns: 1fr repeat(%v, 15%%);'>", columns)
}

func EndGrid() {
	fmt.Printf("</div>")
}
