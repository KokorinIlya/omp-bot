package item

import (
	"errors"
	"fmt"
	"math/rand"
)

type DummyService struct {
	indexById map[uint64]int
	items     []Item
} // We can use linked hash map to get faster removing

const (
	maxRetries    = 100
	noSuchItemIdx = -1
)

func NewDummyService() *DummyService {
	return &DummyService{
		items:     make([]Item, 0),
		indexById: make(map[uint64]int),
	}
}

func (service *DummyService) Describe(itemId uint64) (*Item, error) {
	idx, contains := service.indexById[itemId]
	if !contains || idx < 0 {
		return nil, errors.New(fmt.Sprintf(
			"No item with id %v", itemId,
		))
	}
	return &service.items[idx], nil
}

func (service *DummyService) List(cursor uint64, limit uint64) ([]Item, error) {
	totalCount := uint64(len(service.items))
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
	return service.items[cursor:right], nil
}

func (service *DummyService) Create(item Item) (uint64, error) {
	var newId uint64
	retries := 0
	for {
		newId = uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
		_, contains := service.indexById[newId]
		if !contains {
			break
		}
		if retries == maxRetries {
			return 0, errors.New("cannot allocate id, please try again later")
		}
		retries++
	}
	item.Id = newId
	service.items = append(service.items, item)
	service.indexById[newId] = len(service.items) - 1
	return newId, nil
}

func (service *DummyService) Update(itemId uint64, item Item) error {
	if itemId != item.Id {
		return errors.New("itemId != item.Id")
	}
	idx, contains := service.indexById[itemId]
	if !contains || idx < 0 {
		return errors.New(fmt.Sprintf(
			"No item with id %v", itemId,
		))
	}
	service.items[idx] = item
	return nil
}

func (service *DummyService) Remove(itemId uint64) error {
	idx, contains := service.indexById[itemId]
	if !contains || idx < 0 {
		return errors.New(fmt.Sprintf(
			"No item with id %v", itemId,
		))
	}
	for i := idx + 1; i < len(service.items); i++ {
		item := service.items[i]
		service.items[i-1] = item
		service.indexById[item.Id] = i - 1
	}

	// Ids of removed items will not be allocated for new items
	service.indexById[itemId] = noSuchItemIdx
	service.items = service.items[:len(service.items)-1]
	return nil
}

func (service *DummyService) ItemsCount() uint64 {
	return uint64(len(service.items))
}
