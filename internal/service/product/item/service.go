package product

import "github.com/ozonmp/omp-bot/internal/model/product"

type ItemService interface {
	Describe(itemId uint64) (*product.Item, error)
	List(cursor uint64, limit uint64) []product.Item
	Create(product.Item) (uint64, error)
	Update(itemId uint64, item product.Item) error
	Remove(itemId uint64) (bool, error)
}
