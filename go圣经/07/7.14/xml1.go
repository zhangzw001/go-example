package main

type Name struct {
	Local string // e.g., "Title" or "id"
}

type Attr struct { // e.g., name="value"
	Name  Name
	Value string
}
