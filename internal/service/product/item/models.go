package item

import "fmt"

type Item struct { // TODO: add fields
	Id    uint64
	Title string
}

func (item Item) String() string {
	return fmt.Sprintf(
		"Item { id = %v; title = %v }",
		item.Id, item.Title,
	)
}

func NewItem(id uint64, title string) *Item {
	return &Item{
		Id:    id,
		Title: title,
	}
}
