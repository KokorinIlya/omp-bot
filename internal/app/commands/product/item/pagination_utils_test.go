package item

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/service/product/item"
	"testing"
)

type MockedItemService struct {
	Count uint64
	Items []item.Item
	Err   error
}

func (itemService *MockedItemService) List(_ uint64, _ uint64) ([]item.Item, error) {
	return itemService.Items, itemService.Err
}

func (itemService *MockedItemService) ItemsCount() uint64 {
	return itemService.Count
}

func TestPaginateEmpty(t *testing.T) {
	mockedItemService := MockedItemService{
		Count: 0,
		Items: make([]item.Item, 0),
		Err:   nil,
	}

	text, keyboard, err := getPaginatedMessage(&mockedItemService, 0, 10)
	if err != nil {
		t.Errorf("Expected successful pagination, but received %v", err)
	}
	if text != "Ни одного элемента" {
		t.Errorf("Expected 'Ни одного элемента', but received %v", text)
	}
	if len(keyboard.InlineKeyboard) != 1 || len(keyboard.InlineKeyboard[0]) != 0 {
		t.Errorf("Expected no buttons, but received %v", keyboard.InlineKeyboard)
	}
}

func TestPaginateError(t *testing.T) {
	mockedItemService := MockedItemService{
		Count: 0,
		Items: make([]item.Item, 0),
		Err:   errors.New("bang"),
	}

	text, keyboard, err := getPaginatedMessage(&mockedItemService, 0, 10)
	if err == nil {
		t.Errorf("Expected error, but received %v, %v", text, *keyboard)
	}
}

func TestPaginateOnlyElement(t *testing.T) {
	mockedItemService := MockedItemService{
		Count: 1,
		Items: []item.Item{
			*item.NewItem(10, 20, 30, "title"),
		},
		Err: nil,
	}

	text, keyboard, err := getPaginatedMessage(&mockedItemService, 0, 10)
	if err != nil {
		t.Errorf("Expected successful pagination, but received %v", err)
	}
	expText := mockedItemService.Items[0].String() + ";\n"
	if text != expText {
		t.Errorf("Expected text %v, but received %v", expText, text)
	}
	if len(keyboard.InlineKeyboard) != 1 || len(keyboard.InlineKeyboard[0]) != 0 {
		t.Errorf("Expected no buttons, but received %v", keyboard.InlineKeyboard)
	}
}

func TestPaginateOnlyPage(t *testing.T) {
	mockedItemService := MockedItemService{
		Count: 1,
		Items: []item.Item{
			*item.NewItem(10, 20, 30, "title_1"),
			*item.NewItem(11, 20, 30, "title_2"),
			*item.NewItem(12, 20, 30, "title_3"),
		},
		Err: nil,
	}

	text, keyboard, err := getPaginatedMessage(&mockedItemService, 0, 10)
	if err != nil {
		t.Errorf("Expected successful pagination, but received %v", err)
	}
	expText := fmt.Sprintf("%v;\n%v;\n%v;\n",
		mockedItemService.Items[0].String(),
		mockedItemService.Items[1].String(),
		mockedItemService.Items[2].String(),
	)
	if text != expText {
		t.Errorf("Expected text %v, but received %v", expText, text)
	}
	if len(keyboard.InlineKeyboard) != 1 || len(keyboard.InlineKeyboard[0]) != 0 {
		t.Errorf("Expected no buttons, but received %v", keyboard.InlineKeyboard)
	}
}

func TestPaginateLeftmostPage(t *testing.T) {
	mockedItemService := MockedItemService{
		Count: 10,
		Items: []item.Item{
			*item.NewItem(10, 20, 30, "title_1"),
			*item.NewItem(11, 20, 30, "title_2"),
			*item.NewItem(12, 20, 30, "title_3"),
		},
		Err: nil,
	}

	text, keyboard, err := getPaginatedMessage(&mockedItemService, 0, 3)
	if err != nil {
		t.Errorf("Expected successful pagination, but received %v", err)
	}
	expText := fmt.Sprintf("%v;\n%v;\n%v;\n",
		mockedItemService.Items[0].String(),
		mockedItemService.Items[1].String(),
		mockedItemService.Items[2].String(),
	)
	if text != expText {
		t.Errorf("Expected text %v, but received %v", expText, text)
	}
	if len(keyboard.InlineKeyboard) != 1 || len(keyboard.InlineKeyboard[0]) != 1 ||
		keyboard.InlineKeyboard[0][0].Text != "К следующей странице" {
		t.Errorf("Expected only next page button, but received %v", keyboard.InlineKeyboard)
	}
}

func TestPaginateRightmostPage(t *testing.T) {
	mockedItemService := MockedItemService{
		Count: 100,
		Items: []item.Item{
			*item.NewItem(10, 20, 30, "title_1"),
			*item.NewItem(11, 20, 30, "title_2"),
			*item.NewItem(12, 20, 30, "title_3"),
		},
		Err: nil,
	}

	text, keyboard, err := getPaginatedMessage(&mockedItemService, 97, 10)
	if err != nil {
		t.Errorf("Expected successful pagination, but received %v", err)
	}
	expText := fmt.Sprintf("%v;\n%v;\n%v;\n",
		mockedItemService.Items[0].String(),
		mockedItemService.Items[1].String(),
		mockedItemService.Items[2].String(),
	)
	if text != expText {
		t.Errorf("Expected text %v, but received %v", expText, text)
	}
	if len(keyboard.InlineKeyboard) != 1 || len(keyboard.InlineKeyboard[0]) != 1 ||
		keyboard.InlineKeyboard[0][0].Text != "К предыдущей странице" {
		t.Errorf("Expected only next page button, but received %v", keyboard.InlineKeyboard)
	}
}

func TestPaginateMiddlePage(t *testing.T) {
	mockedItemService := MockedItemService{
		Count: 100,
		Items: []item.Item{
			*item.NewItem(10, 20, 30, "title_1"),
			*item.NewItem(11, 20, 30, "title_2"),
			*item.NewItem(12, 20, 30, "title_3"),
		},
		Err: nil,
	}

	text, keyboard, err := getPaginatedMessage(&mockedItemService, 50, 3)
	if err != nil {
		t.Errorf("Expected successful pagination, but received %v", err)
	}
	expText := fmt.Sprintf("%v;\n%v;\n%v;\n",
		mockedItemService.Items[0].String(),
		mockedItemService.Items[1].String(),
		mockedItemService.Items[2].String(),
	)
	if text != expText {
		t.Errorf("Expected text %v, but received %v", expText, text)
	}
	if len(keyboard.InlineKeyboard) != 1 || len(keyboard.InlineKeyboard[0]) != 2 ||
		keyboard.InlineKeyboard[0][0].Text != "К предыдущей странице" ||
		keyboard.InlineKeyboard[0][1].Text != "К следующей странице" {
		t.Errorf("Expected only next page button, but received %v", keyboard.InlineKeyboard)
	}
}
