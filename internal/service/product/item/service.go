package item

import (
	"errors"
	"fmt"
)

type DummyItemService struct {
	items []Item
}

func NewDummyItemService() *DummyItemService {
	return &DummyItemService{
		items: make([]Item, 0),
	}
}

func (itemService *DummyItemService) Describe(itemId uint64) (*Item, error) {
	if itemId < 0 || itemId >= uint64(len(itemService.items)) {
		return nil, errors.New(fmt.Sprintf(
			"Incorrect index %v, correct indexes are [0..%v]", itemId, len(itemService.items),
		))
	}
	return &itemService.items[itemId], nil
}

func (itemService *DummyItemService) List(cursor uint64, limit uint64) ([]Item, error) {
	totalCount := uint64(len(itemService.items))
	if cursor < 0 || cursor > totalCount {
		return nil, errors.New(fmt.Sprintf(
			"Incorrect cursor index %v, correct indexes are [0..%v]", cursor, len(itemService.items),
		))
	}
	right := cursor + limit
	if right > totalCount {
		right = totalCount
	}
	return itemService.items[cursor:right], nil
}

func (itemService *DummyItemService) Create(item Item) (uint64, error) {
	itemService.items = append(itemService.items, item)
	return uint64(len(itemService.items) - 1), nil
}

func (itemService *DummyItemService) Update(itemId uint64, item Item) error {
	if itemId < 0 || itemId >= uint64(len(itemService.items)) {
		return errors.New(fmt.Sprintf(
			"Incorrect index %v, correct indexes are [0..%v]", itemId, len(itemService.items),
		))
	}
	itemService.items[itemId] = item
	return nil
}

func (itemService *DummyItemService) Remove(itemId uint64) error {
	if itemId < 0 || itemId >= uint64(len(itemService.items)) {
		return errors.New(fmt.Sprintf(
			"Incorrect index %v, correct indexes are [0..%v]", itemId, len(itemService.items),
		))
	}
	itemService.items = append(itemService.items[:itemId], itemService.items[itemId+1:]...)
	return nil
}
