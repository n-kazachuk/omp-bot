package education

import "fmt"

type Cours struct {
	Title, Author string
	Year          int
}

func (c *Cours) String() string {
	return fmt.Sprintf("Title: %s, Author: %v, Year: %v", c.Title, c.Author, c.Year)
}
