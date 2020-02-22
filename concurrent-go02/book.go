package main

import "fmt"

// Book description
type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	Book{
		ID:            1,
		Title:         "The Hitchhiker's Guide to the Galaxy",
		Author:        "Douglas Adams",
		YearPublished: 1979,
	},
	Book{
		ID:            2,
		Title:         "The Hobbit",
		Author:        "J.R.R. Tolkien",
		YearPublished: 1937,
	},
	Book{
		ID:            3,
		Title:         "A Tale of Two Cities",
		Author:        "Charles Dickens",
		YearPublished: 1859,
	},
	Book{
		ID:            4,
		Title:         "Book Title 4",
		Author:        "Book Author 4",
		YearPublished: 1934,
	},
	Book{
		ID:            5,
		Title:         "Book Title 5",
		Author:        "Book Author 5",
		YearPublished: 1955,
	},
	Book{
		ID:            6,
		Title:         "Book Title 6",
		Author:        "Book Author 6",
		YearPublished: 1966,
	},
	Book{
		ID:            7,
		Title:         "Book Title 7",
		Author:        "Book Author 7",
		YearPublished: 1977,
	},
	Book{
		ID:            8,
		Title:         "The Moon is a Harsh Mistress",
		Author:        "Robert A. Henlein",
		YearPublished: 1966,
	},
	Book{
		ID:            9,
		Title:         "On Basilisk Station",
		Author:        "David Weber",
		YearPublished: 1993,
	},
	Book{
		ID:            10,
		Title:         "The Android's Dream",
		Author:        "John Scalzi",
		YearPublished: 2006,
	},
}
