package data

// a model for my data
type BookList interface {
}

// Book struct
type TextBook struct {
	Title  string `json:"title"`
	Series int    `json:"series"`
	// Author string `json:"author"`
}

type ShortStory struct {
	Title    string `json:"title"`
	Chapters int    `json:"chapters"`
	// Author   string `json:"author"`
}

type Author struct {
	Name   string     `json:"name"`
	Age    int        `json:"age"`
	Gendor int        `json:gendor`
	Books  []BookList `json:"books"`
}

var BookStore []BookList
var AuthorList []Author

func init() {

	var JKBookList = []BookList{
		&TextBook{"Harry Potter", 7},
		&ShortStory{"Harry Potter Prequel", 8},
		&TextBook{"Fantastic Beasts and Where to Find Them", 1},
	}
	JK := Author{Name: "J. K. Rowling", Age: 53, Gendor: 0, Books: JKBookList}

	var MBookList = []BookList{
		&TextBook{"A Game of Thrones", 8},
		&TextBook{"The Hedge Knight", 1},
		&TextBook{"The Princess and the Queen", 1},
	}
	Martin := Author{Name: "George R R Martin", Age: 69, Gendor: 1, Books: MBookList}

	var BBookList = []BookList{
		&TextBook{"Elantris", 4},
		&TextBook{"Mistborn", 1},
		&ShortStory{"A Song for Lya", 10},
	}
	Brandon := Author{Name: "Brandon Sanderson", Age: 42, Gendor: 1, Books: BBookList}

	AuthorList = append(AuthorList, JK, Martin, Brandon)
	BookStore = append(BookStore, JKBookList...)
	BookStore = append(BookStore, MBookList...)
	BookStore = append(BookStore, BBookList...)
}
