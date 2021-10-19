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
