package htmlHelpers

import (
	"fmt"
)

// Header creates a header element for HTML formatting
func Header(header string, level int) string {
	out := ""
	out += fmt.Sprintln("")
	out += fmt.Sprintf("<h%v>%v</h%v>", level, header, level)
	return out
}

// Span creates a span element for HTML formatting
func Span(innertext string, classname string, title string) string {
	return fmt.Sprintf("<span title='%v' class='%v'>%v</span>", title, classname, innertext)
}

// Div creates a div element for HTML formatting
func Div(innertext string, classname string, title string) string {
	return fmt.Sprintf("<div title='%v' class='%v'>%v</div>", title, classname, innertext)
}

// StartSpan generates the opening tag of a span element
func StartSpan(classname string) string {
	return fmt.Sprintf("<span class='%v'>", classname)
}

// EndSpan generates the closing tag of a span element
func EndSpan() string {
	return fmt.Sprintf("</span>")
}

// A generates a link tag for HTML formatting
func A(innertext string, link string, classname string) string {
	return fmt.Sprintf("<a href='%v' class='%v'>%v</a>", link, classname, innertext)
}

// StartGrid generates the starting tag of a css grid
func StartGrid(columns int) string {
	return fmt.Sprintf("<div style='display: grid; grid-template-columns: 1fr repeat(%v, 15%%); row-gap: 1rem; margin-bottom: 1rem;'>", columns)
}

// EndGrid generates the closing tag of a css grid
func EndGrid() string {
	return fmt.Sprintf("</div>")
}
