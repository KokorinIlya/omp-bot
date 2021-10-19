package item

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestEmptyServiceItemsCount(t *testing.T) {
	service := NewDummyItemService()
	if cnt := service.ItemsCount(); cnt != 0 {
		t.Errorf("Expected items coutn = 0 on empty service, but %v received", cnt)
	}
}

func TestGetNonExistingItem(t *testing.T) {
	service := NewDummyItemService()
	item, err := service.Describe(42)
	if err == nil {
		t.Errorf("Expected error when requesting non-existing item, but %v received", item)
	}
}

func TestCreateItem(t *testing.T) {
	service := NewDummyItemService()
	item := NewItem(0, "title")
	_, err := service.Create(*item)
	if err != nil {
		t.Errorf("Expected successful creation, but %v received", err)
	}
}

func TestRequestCreatedItem(t *testing.T) {
	service := NewDummyItemService()
	item := NewItem(0, "title")
	uid, err := service.Create(*item)
	if err != nil {
		t.Errorf("Expected successful creation, but %v received", err)
	}
	reqItem, err := service.Describe(uid)
	if err != nil {
		t.Errorf("Expected previously created item, but %v received", err)
	}
	if reqItem.Title != "title" || reqItem.Id != uid {
		t.Errorf("Expected previously created item with returned id, but %v received", reqItem)
	}
}

func TestCountCreates(t *testing.T) {
	service := NewDummyItemService()
	if cnt := service.ItemsCount(); cnt != 0 {
		t.Errorf("Expected 0 items at empty service, but received %v", cnt)
	}
	var i uint64
	for i = 1; i < 20; i++ {
		_, _ = service.Create(*NewItem(0, fmt.Sprintf("title_%v", i)))
		if cnt := service.ItemsCount(); cnt != i {
			t.Errorf("Expected %v items at empty service, but received %v", i, cnt)
		}
	}
}

func TestUpdateTheOnlyItem(t *testing.T) {
	service := NewDummyItemService()
	item := NewItem(0, "title")
	uid, _ := service.Create(*item)
	newItem := NewItem(uid, "new title")
	err := service.Update(uid, *newItem)
	if err != nil {
		t.Errorf("Expected successful update, but %v received", err)
	}
	reqItem, err := service.Describe(uid)
	if err != nil {
		t.Errorf("Expected previously updated item, but %v received", err)
	}
	if reqItem.Title != "new title" || reqItem.Id != uid {
		t.Errorf("Expected previously created item with returned id, but %v received", reqItem)
	}
}

func TestUpdateOneOfManyItems(t *testing.T) {
	service := NewDummyItemService()
	uid1, _ := service.Create(*NewItem(0, "title_1"))
	uid2, _ := service.Create(*NewItem(0, "title_2"))
	uid3, _ := service.Create(*NewItem(0, "title_3"))
	uid4, _ := service.Create(*NewItem(0, "title_4"))
	newItem := NewItem(uid3, "new title_3")
	err := service.Update(uid3, *newItem)
	if err != nil {
		t.Errorf("Expected successful update, but %v received", err)
	}
	reqItem, err := service.Describe(uid3)
	if err != nil {
		t.Errorf("Expected previously updated item, but %v received", err)
	}
	if reqItem.Title != "new title_3" || reqItem.Id != uid3 {
		t.Errorf("Expected previously created item with returned id, but %v received", reqItem)
	}
	for idx, uid := range [...]uint64{uid1, uid2, uid4} {
		if idx == 2 {
			idx++
		}
		reqItem, err = service.Describe(uid)
		if err != nil {
			t.Errorf("Expected item, but %v received", err)
		}
		if reqItem.Title != fmt.Sprintf("title_%v", idx+1) || reqItem.Id != uid {
			t.Errorf("Expected unchanged item with returned id, but %v received", reqItem)
		}
	}
}

func TestUpdateChangeId(t *testing.T) {
	service := NewDummyItemService()
	item := NewItem(0, "title")
	uid, _ := service.Create(*item)
	newItem := NewItem(uid+1, "new title")
	err := service.Update(uid, *newItem)
	if err == nil {
		t.Error("Expected update fail, but update succeed")
	}
}

func TestUpdateNonExistingItem(t *testing.T) {
	service := NewDummyItemService()
	newItem := NewItem(42, "new title")
	err := service.Update(42, *newItem)
	if err == nil {
		t.Error("Expected update fail, but update succeed")
	}
}

func TestListEmptyService(t *testing.T) {
	service := NewDummyItemService()
	lst, err := service.List(0, 10)
	if err != nil {
		t.Errorf("Expected empty slice, received %v", err)
	}
	if len(lst) != 0 {
		t.Errorf("Expected empty slice, received %v", lst)
	}
}

func TestListEmptyServiceBigIndex(t *testing.T) {
	service := NewDummyItemService()
	lst, err := service.List(1, 10)
	if err == nil {
		t.Errorf("Expected error, but received %v", lst)
	}
}

func assertItemsList(t *testing.T, items []Item, titles []string) {
	if len(items) != len(titles) {
		t.Errorf("Expected items slice of length %v, but %v received", len(titles), items)
	}
	for i := 0; i < len(items); i++ {
		if items[i].Title != titles[i] {
			t.Errorf("Expected title %v at %v-th item, but %v received", titles[i], i, items[i])
		}
	}
}

func TestListNonEmptyService(t *testing.T) {
	service := NewDummyItemService()
	var i uint64
	for i = 0; i < 20; i++ {
		_, _ = service.Create(*NewItem(0, fmt.Sprintf("title_%v", i)))
	}
	lst, err := service.List(0, 3)
	if err != nil {
		t.Errorf("Expected successful list, but %v received", err)
	}
	assertItemsList(t, lst, []string{"title_0", "title_1", "title_2"})
	lst, err = service.List(10, 5)
	if err != nil {
		t.Errorf("Expected successful list, but %v received", err)
	}
	assertItemsList(t, lst, []string{"title_10", "title_11", "title_12", "title_13", "title_14"})
}

func TestListNearRightBound(t *testing.T) {
	service := NewDummyItemService()
	var i uint64
	for i = 0; i < 20; i++ {
		_, _ = service.Create(*NewItem(0, fmt.Sprintf("title_%v", i)))
	}
	lst, err := service.List(18, 100)
	if err != nil {
		t.Errorf("Expected successful list, but %v received", err)
	}
	assertItemsList(t, lst, []string{"title_18", "title_19"})
}

func TestListRightBound(t *testing.T) {
	service := NewDummyItemService()
	var i uint64
	for i = 0; i < 20; i++ {
		_, _ = service.Create(*NewItem(0, fmt.Sprintf("title_%v", i)))
	}
	lst, err := service.List(20, 100)
	if err != nil {
		t.Errorf("Expected successful list, but %v received", err)
	}
	if len(lst) != 0 {
		t.Errorf("Expected empty list, but %v received", lst)
	}
}

func TestListBeyondRightBound(t *testing.T) {
	service := NewDummyItemService()
	var i uint64
	for i = 0; i < 20; i++ {
		_, _ = service.Create(*NewItem(0, fmt.Sprintf("title_%v", i)))
	}
	lst, err := service.List(21, 100)
	if err == nil {
		t.Errorf("Expected error, but %v received", lst)
	}
}

func TestRemoveNonExisting(t *testing.T) {
	service := NewDummyItemService()
	err := service.Remove(42)
	if err == nil {
		t.Errorf("Expected error on remove, but remove was successful")
	}
}

func TestRemoveTheOnlyItem(t *testing.T) {
	service := NewDummyItemService()
	uid, _ := service.Create(*NewItem(0, "title"))
	err := service.Remove(uid)
	if err != nil {
		t.Errorf("Expected successful remove, but %v received", err)
	}
	res, err := service.Describe(uid)
	if err == nil {
		t.Errorf("Expected error on trying to retrieve removed item, but %v received", res)
	}
}

func TestRemoveItemsCount(t *testing.T) {
	service := NewDummyItemService()
	//goland:noinspection SpellCheckingInspection
	uids := make([]uint64, 0)
	for i := 0; i < 20; i++ {
		uid, _ := service.Create(*NewItem(0, fmt.Sprintf("title_%v", i)))
		uids = append(uids, uid)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(uids), func(i, j int) { uids[i], uids[j] = uids[j], uids[i] })
	for i, uid := range uids {
		err := service.Remove(uid)
		if err != nil {
			t.Errorf("Expected successful deletion, but %v received", err)
		}
		expCnt := uint64(20 - i - 1)
		if cnt := service.ItemsCount(); cnt != expCnt {
			t.Errorf("Expected %v items, but %v received", expCnt, cnt)
		}
	}
}

func TestRemoveLastItem(t *testing.T) {
	service := NewDummyItemService()
	uid1, _ := service.Create(*NewItem(0, "title_1"))
	uid2, _ := service.Create(*NewItem(0, "title_2"))
	uid3, _ := service.Create(*NewItem(0, "title_3"))
	err := service.Remove(uid3)
	if err != nil {
		t.Errorf("Expected successful removal, but %v received", err)
	}
	req, err := service.Describe(uid3)
	if err == nil {
		t.Errorf("Expected error on trying to retrieve removed item, but %v received", *req)
	}

	req, err = service.Describe(uid1)
	if err != nil {
		t.Errorf("Expected successful retrieval, but %v received", err)
	}
	if req.Title != "title_1" {
		t.Errorf("Expected item with title_1, but %v received", *req)
	}

	req, err = service.Describe(uid2)
	if err != nil {
		t.Errorf("Expected successful retrieval, but %v received", err)
	}
	if req.Title != "title_2" {
		t.Errorf("Expected item with title_2, but %v received", *req)
	}

	lst, err := service.List(0, 200)
	if err != nil {
		t.Errorf("Expected successful list, but %v received", err)
	}
	assertItemsList(t, lst, []string{"title_1", "title_2"})
}

func TestRemoveMiddleItem(t *testing.T) {
	service := NewDummyItemService()
	uid1, _ := service.Create(*NewItem(0, "title_1"))
	uid2, _ := service.Create(*NewItem(0, "title_2"))
	uid3, _ := service.Create(*NewItem(0, "title_3"))
	err := service.Remove(uid2)
	if err != nil {
		t.Errorf("Expected successful removal, but %v received", err)
	}
	req, err := service.Describe(uid2)
	if err == nil {
		t.Errorf("Expected error on trying to retrieve removed item, but %v received", *req)
	}

	req, err = service.Describe(uid1)
	if err != nil {
		t.Errorf("Expected successful retrieval, but %v received", err)
	}
	if req.Title != "title_1" {
		t.Errorf("Expected item with title_1, but %v received", *req)
	}

	req, err = service.Describe(uid3)
	if err != nil {
		t.Errorf("Expected successful retrieval, but %v received", err)
	}
	if req.Title != "title_3" {
		t.Errorf("Expected item with title_3, but %v received", *req)
	}

	lst, err := service.List(0, 200)
	if err != nil {
		t.Errorf("Expected successful list, but %v received", err)
	}
	assertItemsList(t, lst, []string{"title_1", "title_3"})
}

func TestRemoveFirstItem(t *testing.T) {
	service := NewDummyItemService()
	uid1, _ := service.Create(*NewItem(0, "title_1"))
	uid2, _ := service.Create(*NewItem(0, "title_2"))
	uid3, _ := service.Create(*NewItem(0, "title_3"))
	err := service.Remove(uid1)
	if err != nil {
		t.Errorf("Expected successful removal, but %v received", err)
	}
	req, err := service.Describe(uid1)
	if err == nil {
		t.Errorf("Expected error on trying to retrieve removed item, but %v received", *req)
	}

	req, err = service.Describe(uid2)
	if err != nil {
		t.Errorf("Expected successful retrieval, but %v received", err)
	}
	if req.Title != "title_2" {
		t.Errorf("Expected item with title_1, but %v received", *req)
	}

	req, err = service.Describe(uid3)
	if err != nil {
		t.Errorf("Expected successful retrieval, but %v received", err)
	}
	if req.Title != "title_3" {
		t.Errorf("Expected item with title_3, but %v received", *req)
	}

	lst, err := service.List(0, 200)
	if err != nil {
		t.Errorf("Expected successful list, but %v received", err)
	}
	assertItemsList(t, lst, []string{"title_2", "title_3"})
}

func TestRemoveAllItems(t *testing.T) {
	service := NewDummyItemService()
	//goland:noinspection SpellCheckingInspection
	items := make([]Item, 0)
	for i := 0; i < 20; i++ {
		uid, _ := service.Create(*NewItem(0, fmt.Sprintf("title_%v", i)))
		item, _ := service.Describe(uid)
		items = append(items, *item)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })
	for i, item := range items {
		err := service.Remove(item.Id)
		if err != nil {
			t.Errorf("Expected successful deletion, but %v received", err)
		}
		for j := i + 1; j < len(items); j++ {
			curItem := items[j]
			reqItem, err := service.Describe(curItem.Id)
			if err != nil {
				t.Errorf("Expected successful request, but %v received", err)
			}
			if *reqItem != curItem {
				t.Errorf("Expected item %v, but %v received", curItem, *reqItem)
			}
		}
	}
}
