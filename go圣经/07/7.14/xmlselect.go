package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)
//
//type Name struct {
//	Local string // e.g., "Title" or "id"
//}
//
//type Attr struct { // e.g., name="value"
//	Name  Name
//	Value string
//}
//
//// A Token includes StartElement, EndElement, CharData,
//// and Comment, plus a few esoteric types (not shown).
//type Token interface{}
//type StartElement struct { // e.g., <name>
//	Name Name
//	Attr []Attr
//}
//type EndElement struct { Name Name } // e.g., </name>
//type CharData []byte                 // e.g., <p>CharData</p>
//type Comment []byte                  // e.g., <!-- Comment -->
//
//type Decoder struct{ /* ... */ }
//func NewDecoder(io.Reader) *Decoder
//func (*Decoder) Token() (Token, error) // returns next Token in sequence
//
// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}

		if x[0] == y[0] {
			y = y[1:]
		}

		x = x[1:]
	}

	return false
}


func main() {
	resp , err := http.Get("http://www.w3.org/TR/2006/REC-xml11-20060816")
	if err != nil {
		log.Fatalf("http.get err : %v\n",err)
	}
	defer resp.Body.Close()
	dec := xml.NewDecoder(resp.Body)
	var stack []string // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err )
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack , " "), tok)
			}
		}
	}
}
