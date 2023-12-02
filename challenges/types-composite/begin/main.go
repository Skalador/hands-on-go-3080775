// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title  string
	author author
}

// define a library type to track a list of books

type library map[string][]book

// define addBook to add a book to the library
func (l library) addBook(b book) {
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(author string) []book {
	return l[author]
}

func main() {
	// create a new library
	lib := library{}

	// add 2 books to the library by the same author
	a1 := author{name: "Kevin"}
	b1 := book{title: "first book", author: a1}
	b2 := book{title: "second book", author: a1}
	lib.addBook(b1)
	lib.addBook(b2)

	// add 1 book to the library by a different author
	a2 := author{name: "Anita"}
	lib.addBook(book{title: "First book", author: a2})

	// dump the library with spew
	spew.Dump(lib)

	// lookup books by known author in the library
	fBook := lib.lookupByAuthor("Kevin")
	fmt.Println(fBook)

	// print out the first book's title and its author's name
	fmt.Printf("Book: %v by %v", fBook[0].title, fBook[0].author.name)

}
