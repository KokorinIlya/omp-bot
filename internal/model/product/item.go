package product

import "fmt"

type Item struct {
	Title string
	Id    uint64
}

func (item Item) String() string {
	return fmt.Sprintf("Item { title = %v; id = %v }", item.Title, item.Id)
}
