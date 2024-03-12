package education

import "fmt"

type Cours struct {
	Title, Author string
	Year          int
}

func NewCours(title, author string, year int) Cours {
	return Cours{
		Title:  title,
		Author: author,
		Year:   year,
	}
}

func (c *Cours) String() string {
	return fmt.Sprintf("Title: %s, Author: %v, Year: %v", c.Title, c.Author, c.Year)
}
