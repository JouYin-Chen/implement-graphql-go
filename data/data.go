package data

// a model for my data
// Book struct
type TextBook struct {
	Title  string `json:"title"`
	Series int    `json:"series"`
	Author string `json:"author"`
}
type Author struct {
	Name   string     `json:"name"`
	Age    int        `json:"age"`
	Gendor int        `json:gendor`
	Books  []TextBook `json:"books"`
}

var TextBookList []TextBook
var AuthorList []Author

func init() {

	JKBookList := []TextBook{
		TextBook{
			"Harry Potter",
			7,
			"J. K. Rowling",
		},
		TextBook{
			"Fantastic Beasts and Where to Find Them",
			1,
			"J. K. Rowling",
		},
	}

	JK := Author{Name: "J. K. Rowling", Age: 53, Gendor: 0, Books: JKBookList}

	MBookList := []TextBook{
		TextBook{
			"A Game of Thrones",
			8,
			"George R R Martin",
		},
		TextBook{
			"The Hedge Knight",
			1,
			"George R R Martin",
		},
		TextBook{
			"The Princess and the Queen",
			1,
			"George R R Martin",
		},
	}

	Martin := Author{Name: "George R R Martin", Age: 69, Gendor: 1, Books: MBookList}

	BBookList := []TextBook{
		TextBook{
			"Elantris",
			4,
			"Brandon Sanderson",
		},
		TextBook{
			"Mistborn",
			1,
			"Brandon Sanderson",
		},
	}

	Brandon := Author{Name: "Brandon Sanderson", Age: 42, Gendor: 1, Books: BBookList}

	AuthorList = append(AuthorList, JK, Martin, Brandon)
	TextBookList = append(TextBookList, JKBookList...)
	TextBookList = append(TextBookList, MBookList...)
	TextBookList = append(TextBookList, BBookList...)

}
