package htmlHelpers

import (
	"fmt"
)

func Header(header string, level int) {
	fmt.Println("")
	fmt.Printf("<h%v>%v</h%v>", level, header, level)
}

func Span(innertext string, classname string, title string) {
	fmt.Printf("<span title='%v' class='%v'>%v</span>", title, classname, innertext)
}

func Div(innertext string, classname string, title string) {
	fmt.Printf("<div title='%v' class='%v'>%v</div>", title, classname, innertext)
}

func StartSpan(classname string) {
	fmt.Printf("<span class='%v'>", classname)
}

func EndSpan() {
	fmt.Printf("</span>")
}

func A(innertext string, link string, classname string) {
	fmt.Printf("<a href='%v' class='%v'>%v</a>", link, classname, innertext)
}

func StartPageHtml() {
	fmt.Printf("<!DOCTYPE><html><head><style>a { text-decoration :none !important; } body > * { width: 80vw; } .label { font-size: 0.8rem; text-transform: lowercase; } </style></head><body style='font-family: Courier New; display: flex; align-items: center; flex-direction: column;'>")
}

func EndPageHtml() {
	fmt.Printf("</body></html>")
}

func StartGrid(columns int) {
	fmt.Printf("<div style='display: grid; grid-template-columns: 1fr repeat(%v, 15%%); row-gap: 1rem; margin-bottom: 1rem;'>", columns)
}

func EndGrid() {
	fmt.Printf("</div>")
}
