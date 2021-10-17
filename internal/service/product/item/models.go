package item

import "fmt"

type Item struct {
	Title string
}

func (item Item) String() string {
	return fmt.Sprintf("Item { title = %v }", item.Title)
}

func NewItem(title string) *Item {
	return &Item{
		Title: title,
	}
}
