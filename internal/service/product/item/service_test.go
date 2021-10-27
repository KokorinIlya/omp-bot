package item

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestEmptyServiceItemsCount(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	assert.Equal(t, uint64(0), service.ItemsCount())
}

func TestGetNonExistingItem(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	_, err := service.Describe(42)
	assert.NotNil(t, err)
}

func TestCreateItem(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	item := NewItem(0, 10, 10, "title")
	_, err := service.Create(*item)
	assert.Nil(t, err)
}

func TestRequestCreatedItem(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	item := NewItem(0, 10, 10, "title")
	uid, err := service.Create(*item)
	assert.Nil(t, err)
	reqItem, err := service.Describe(uid)
	assert.Nil(t, err)
	assert.Equal(t, "title", reqItem.Title)
	assert.Equal(t, uid, reqItem.Id)
}

func TestCountCreates(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	assert.Equal(t, uint64(0), service.ItemsCount())
	var i uint64
	for i = 1; i < 20; i++ {
		_, _ = service.Create(*NewItem(0, 10, 10, fmt.Sprintf("title_%v", i)))
		assert.Equal(t, i, service.ItemsCount())
	}
}

func TestUpdateTheOnlyItem(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	item := NewItem(0, 10, 10, "title")
	uid, _ := service.Create(*item)
	newItem := NewItem(uid, 10, 10,"new title")
	err := service.Update(uid, *newItem)
	assert.Nil(t, err)
	reqItem, err := service.Describe(uid)
	assert.Nil(t, err)
	assert.Equal(t, "new title", reqItem.Title)
	assert.Equal(t, uid, reqItem.Id)
}

func TestUpdateOneOfManyItems(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	uid1, _ := service.Create(*NewItem(0, 10, 10, "title_1"))
	uid2, _ := service.Create(*NewItem(0, 10, 10, "title_2"))
	uid3, _ := service.Create(*NewItem(0, 10, 10, "title_3"))
	uid4, _ := service.Create(*NewItem(0, 10, 10, "title_4"))
	newItem := NewItem(uid3, 10, 10,"new title_3")
	err := service.Update(uid3, *newItem)
	assert.Nil(t, err)
	reqItem, err := service.Describe(uid3)
	assert.Nil(t, err)
	assert.Equal(t, "new title_3", reqItem.Title)
	assert.Equal(t, uid3, reqItem.Id)
	for idx, uid := range [...]uint64{uid1, uid2, uid4} {
		if idx == 2 {
			idx++
		}
		reqItem, err = service.Describe(uid)
		assert.Nil(t, err)
		assert.Equal(t, fmt.Sprintf("title_%v", idx+1), reqItem.Title)
		assert.Equal(t, uid, reqItem.Id)
	}
}

func TestUpdateChangeId(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	item := NewItem(0, 10, 10,"title")
	uid, _ := service.Create(*item)
	newItem := NewItem(uid+1, 10, 10, "new title")
	err := service.Update(uid, *newItem)
	assert.NotNil(t, err)
}

func TestUpdateNonExistingItem(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	newItem := NewItem(42, 10, 10,"new title")
	err := service.Update(42, *newItem)
	assert.NotNil(t, err)
}

func TestListEmptyService(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	lst, err := service.List(0, 10)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(lst))
}

func TestListEmptyServiceBigIndex(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	_, err := service.List(1, 10)
	assert.NotNil(t, err)
}

func assertItemsList(t *testing.T, items []Item, titles []string) {
	assert.Equal(t, len(titles), len(items))
	for i := 0; i < len(items); i++ {
		assert.Equal(t, titles[i], items[i].Title)
	}
}

func TestListNonEmptyService(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	var i uint64
	for i = 0; i < 20; i++ {
		_, _ = service.Create(*NewItem(0, 10, 10, fmt.Sprintf("title_%v", i)))
	}
	lst, err := service.List(0, 3)
	assert.Nil(t, err)
	assertItemsList(t, lst, []string{"title_0", "title_1", "title_2"})
	lst, err = service.List(10, 5)
	assert.Nil(t, err)
	assertItemsList(t, lst, []string{"title_10", "title_11", "title_12", "title_13", "title_14"})
}

func TestListNearRightBound(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	var i uint64
	for i = 0; i < 20; i++ {
		_, _ = service.Create(*NewItem(0, 10, 10, fmt.Sprintf("title_%v", i)))
	}
	lst, err := service.List(18, 100)
	assert.Nil(t, err)
	assertItemsList(t, lst, []string{"title_18", "title_19"})
}

func TestListRightBound(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	var i uint64
	for i = 0; i < 20; i++ {
		_, _ = service.Create(*NewItem(0, 10, 10, fmt.Sprintf("title_%v", i)))
	}
	lst, err := service.List(20, 100)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(lst))
}

func TestListBeyondRightBound(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	var i uint64
	for i = 0; i < 20; i++ {
		_, _ = service.Create(*NewItem(0, 10, 10, fmt.Sprintf("title_%v", i)))
	}
	_, err := service.List(21, 100)
	assert.NotNil(t, err)
}

func TestRemoveNonExisting(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	err := service.Remove(42)
	assert.NotNil(t, err)
}

func TestRemoveTheOnlyItem(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	uid, _ := service.Create(*NewItem(0, 10, 10, "title"))
	err := service.Remove(uid)
	assert.Nil(t, err)
	_, err = service.Describe(uid)
	assert.NotNil(t, err)
}

func TestRemoveItemsCount(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	//goland:noinspection SpellCheckingInspection
	uids := make([]uint64, 0)
	for i := 0; i < 20; i++ {
		uid, _ := service.Create(*NewItem(0, 10, 10, fmt.Sprintf("title_%v", i)))
		uids = append(uids, uid)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(uids), func(i, j int) { uids[i], uids[j] = uids[j], uids[i] })
	for i, uid := range uids {
		err := service.Remove(uid)
		assert.Nil(t, err)
		assert.Equal(t, uint64(20 - i - 1), service.ItemsCount())
	}
}

func TestRemoveLastItem(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	uid1, _ := service.Create(*NewItem(0, 10, 10, "title_1"))
	uid2, _ := service.Create(*NewItem(0, 10, 10, "title_2"))
	uid3, _ := service.Create(*NewItem(0, 10, 10, "title_3"))

	err := service.Remove(uid3)
	assert.Nil(t, err)
	_, err = service.Describe(uid3)
	assert.NotNil(t, err)

	req, err := service.Describe(uid1)
	assert.Nil(t, err)
	assert.Equal(t, "title_1", req.Title)

	req, err = service.Describe(uid2)
	assert.Nil(t, err)
	assert.Equal(t, "title_2", req.Title)

	lst, err := service.List(0, 200)
	assert.Nil(t, err)
	assertItemsList(t, lst, []string{"title_1", "title_2"})
}

func TestRemoveMiddleItem(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	uid1, _ := service.Create(*NewItem(0, 10, 10, "title_1"))
	uid2, _ := service.Create(*NewItem(0, 10, 10, "title_2"))
	uid3, _ := service.Create(*NewItem(0, 10, 10, "title_3"))

	err := service.Remove(uid2)
	assert.Nil(t, err)
	_, err = service.Describe(uid2)
	assert.NotNil(t, err)

	req, err := service.Describe(uid1)
	assert.Nil(t, err)
	assert.Equal(t, "title_1", req.Title)

	req, err = service.Describe(uid3)
	assert.Nil(t, err)
	assert.Equal(t, "title_3", req.Title)

	lst, err := service.List(0, 200)
	assert.Nil(t, err)
	assertItemsList(t, lst, []string{"title_1", "title_3"})
}

func TestRemoveFirstItem(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	uid1, _ := service.Create(*NewItem(0, 10, 10, "title_1"))
	uid2, _ := service.Create(*NewItem(0, 10, 10, "title_2"))
	uid3, _ := service.Create(*NewItem(0, 10, 10, "title_3"))

	err := service.Remove(uid1)
	assert.Nil(t, err)
	_, err = service.Describe(uid1)
	assert.NotNil(t, err)

	req, err := service.Describe(uid2)
	assert.Nil(t, err)
	assert.Equal(t, "title_2", req.Title)

	req, err = service.Describe(uid3)
	assert.Nil(t, err)
	assert.Equal(t, "title_3", req.Title)

	lst, err := service.List(0, 200)
	assert.Nil(t, err)
	assertItemsList(t, lst, []string{"title_2", "title_3"})
}

func TestRemoveAllItems(t *testing.T) {
	t.Parallel()
	service := NewDummyService()
	items := make([]Item, 0)
	for i := 0; i < 20; i++ {
		uid, _ := service.Create(*NewItem(0, 10, 10, fmt.Sprintf("title_%v", i)))
		item, _ := service.Describe(uid)
		items = append(items, *item)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })
	for i, item := range items {
		err := service.Remove(item.Id)
		assert.Nil(t, err)
		for j := i + 1; j < len(items); j++ {
			curItem := items[j]
			reqItem, err := service.Describe(curItem.Id)
			assert.Nil(t, err)
			assert.Equal(t, curItem, *reqItem)
		}
	}
}
