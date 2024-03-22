package education

import (
	"fmt"
)

type Cours struct {
	Title, Author string
	ID, Year      int
}

func NewCours(ID, year int, title, author string) Cours {
	return Cours{
		ID:     ID,
		Title:  title,
		Author: author,
		Year:   year,
	}
}

func (c *Cours) String() string {
	return fmt.Sprintf("ID: %v, Title: %s, Author: %v, Year: %v", c.ID, c.Title, c.Author, c.Year)
}
