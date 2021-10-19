package item

import (
	"testing"
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

func TestUpdateChangeId(t *testing.T) {
	service := NewDummyItemService()
	item := NewItem(0, "title")
	uid, _ := service.Create(*item)
	newItem := NewItem(uid + 1, "new title")
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