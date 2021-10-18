package item

import (
	"errors"
	"fmt"
)

type DummyItemService struct { // TODO: id does not change, hashmap + list
	items []Item
}

func NewDummyItemService() *DummyItemService {
	return &DummyItemService{
		items: make([]Item, 0),
	}
}

func (itemService *DummyItemService) Describe(itemId uint64) (*Item, error) {
	if err := itemService.checkIndex(itemId); err != nil {
		return nil, err
	}
	return &itemService.items[itemId], nil
}

func (itemService *DummyItemService) List(cursor uint64, limit uint64) ([]Item, error) {
	totalCount := uint64(len(itemService.items))
	if cursor > totalCount {
		return nil, errors.New(fmt.Sprintf(
			"Incorrect cursor position %v, correct cursor positions are [0..%v]",
			cursor, totalCount,
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
	if err := itemService.checkIndex(itemId); err != nil {
		return err
	}
	itemService.items[itemId] = item
	return nil
}

func (itemService *DummyItemService) Remove(itemId uint64) error {
	if err := itemService.checkIndex(itemId); err != nil {
		return err
	}
	itemService.items = append(itemService.items[:itemId], itemService.items[itemId+1:]...)
	return nil
}

func (itemService *DummyItemService) ItemsCount() uint64 {
	return uint64(len(itemService.items))
}

func (itemService *DummyItemService) checkIndex(itemId uint64) error {
	if itemId >= uint64(len(itemService.items)) {
		if len(itemService.items) > 0 {
			return errors.New(fmt.Sprintf(
				"Incorrect index %v, correct indexes are [0..%v]", itemId, len(itemService.items)-1,
			))
		} else {
			return errors.New(fmt.Sprintf(
				"Incorrect index %v for empty item storage", itemId,
			))
		}
	} else {
		return nil
	}
}
