package item

import (
	"errors"
	"fmt"
	"math/rand"
)

type DummyItemService struct {
	indexById map[uint64]int
	items     []Item
} // We can use linked hash map to get faster removing

const (
	maxRetries    = 100
	noSuchItemIdx = -1
)

func NewDummyItemService() *DummyItemService {
	return &DummyItemService{
		items:     make([]Item, 0),
		indexById: make(map[uint64]int),
	}
}

func (itemService *DummyItemService) Describe(itemId uint64) (*Item, error) {
	idx, contains := itemService.indexById[itemId]
	if !contains || idx < 0 {
		return nil, errors.New(fmt.Sprintf(
			"No item with id %v", itemId,
		))
	}
	return &itemService.items[idx], nil
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
	var newId uint64
	retries := 0
	for {
		newId = uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
		_, contains := itemService.indexById[newId]
		if !contains {
			break
		}
		if retries == maxRetries {
			return 0, errors.New("cannot allocate id, please try again later")
		}
		retries++
	}
	item.Id = newId
	itemService.items = append(itemService.items, item)
	itemService.indexById[newId] = len(itemService.items) - 1
	return newId, nil
}

func (itemService *DummyItemService) Update(itemId uint64, item Item) error {
	if itemId != item.Id {
		return errors.New("itemId != item.Id")
	}
	idx, contains := itemService.indexById[itemId]
	if !contains || idx < 0 {
		return errors.New(fmt.Sprintf(
			"No item with id %v", itemId,
		))
	}
	itemService.items[idx] = item
	return nil
}

func (itemService *DummyItemService) Remove(itemId uint64) error {
	idx, contains := itemService.indexById[itemId]
	if !contains || idx < 0 {
		return errors.New(fmt.Sprintf(
			"No item with id %v", itemId,
		))
	}
	for i := idx + 1; i < len(itemService.items); i++ {
		item := itemService.items[i]
		itemService.items[i-1] = item
		itemService.indexById[item.Id] = i - 1
	}

	// Ids of removed items will not be allocated for new items
	itemService.indexById[itemId] = noSuchItemIdx
	itemService.items = itemService.items[:len(itemService.items)-1]
	return nil
}

func (itemService *DummyItemService) ItemsCount() uint64 {
	return uint64(len(itemService.items))
}
