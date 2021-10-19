package item

import "fmt"

type Item struct {
	Id        uint64
	OwnerId   uint64
	ProductId uint64
	Title     string
}

func (item Item) String() string {
	return fmt.Sprintf(
		"Item { id = %v; owner_id =%v, product_id = %v; title = %v }",
		item.Id, item.OwnerId, item.ProductId, item.Title,
	)
}

func NewItem(id uint64, ownerId uint64, productId uint64, title string) *Item {
	return &Item{
		Id:        id,
		OwnerId:   ownerId,
		ProductId: productId,
		Title:     title,
	}
}
