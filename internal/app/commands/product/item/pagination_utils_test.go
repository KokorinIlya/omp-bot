package item

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/service/product/item"
	"github.com/stretchr/testify/assert"
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
	t.Parallel()
	mockedItemService := MockedItemService{
		Count: 0,
		Items: make([]item.Item, 0),
		Err:   nil,
	}

	text, keyboard, err := getPaginatedMessage(&mockedItemService, 0, 10)
	assert.Nil(t, err)
	assert.Equal(t, "Ни одного элемента", text)
	assert.Equal(t, 1, len(keyboard.InlineKeyboard))
	assert.Equal(t, 0, len(keyboard.InlineKeyboard[0]))
}

func TestPaginateError(t *testing.T) {
	t.Parallel()
	mockedItemService := MockedItemService{
		Count: 0,
		Items: make([]item.Item, 0),
		Err:   errors.New("bang"),
	}

	_, _, err := getPaginatedMessage(&mockedItemService, 0, 10)
	assert.NotNil(t, err)
}

func TestPaginateOnlyElement(t *testing.T) {
	t.Parallel()
	mockedItemService := MockedItemService{
		Count: 1,
		Items: []item.Item{
			*item.NewItem(10, 20, 30, "title"),
		},
		Err: nil,
	}

	text, keyboard, err := getPaginatedMessage(&mockedItemService, 0, 10)
	assert.Nil(t, err)
	expText := mockedItemService.Items[0].String() + ";\n"
	assert.Equal(t, expText, text)
	assert.Equal(t, 1, len(keyboard.InlineKeyboard))
	assert.Equal(t, 0, len(keyboard.InlineKeyboard[0]))
}

func TestPaginateOnlyPage(t *testing.T) {
	t.Parallel()
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
	assert.Nil(t, err)
	expText := fmt.Sprintf("%v;\n%v;\n%v;\n",
		mockedItemService.Items[0].String(),
		mockedItemService.Items[1].String(),
		mockedItemService.Items[2].String(),
	)
	assert.Equal(t, expText, text)
	assert.Equal(t, 1, len(keyboard.InlineKeyboard))
	assert.Equal(t, 0, len(keyboard.InlineKeyboard[0]))
}

func TestPaginateLeftmostPage(t *testing.T) {
	t.Parallel()
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
	assert.Nil(t, err)
	expText := fmt.Sprintf("%v;\n%v;\n%v;\n",
		mockedItemService.Items[0].String(),
		mockedItemService.Items[1].String(),
		mockedItemService.Items[2].String(),
	)
	assert.Equal(t, expText, text)
	assert.Equal(t, 1, len(keyboard.InlineKeyboard))
	assert.Equal(t, 1, len(keyboard.InlineKeyboard[0]))
	assert.Equal(t, "К следующей странице", keyboard.InlineKeyboard[0][0].Text)
	assert.Equal(t, "product__item__list__{\"offset\":3}", *keyboard.InlineKeyboard[0][0].CallbackData)
}

func TestPaginateRightmostPage(t *testing.T) {
	t.Parallel()
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
	assert.Nil(t, err)
	expText := fmt.Sprintf("%v;\n%v;\n%v;\n",
		mockedItemService.Items[0].String(),
		mockedItemService.Items[1].String(),
		mockedItemService.Items[2].String(),
	)
	assert.Equal(t, expText, text)
	assert.Equal(t, 1, len(keyboard.InlineKeyboard))
	assert.Equal(t, 1, len(keyboard.InlineKeyboard[0]))
	assert.Equal(t, "К предыдущей странице", keyboard.InlineKeyboard[0][0].Text)
	assert.Equal(t, "product__item__list__{\"offset\":87}", *keyboard.InlineKeyboard[0][0].CallbackData)
}

func TestPaginateMiddlePage(t *testing.T) {
	t.Parallel()
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
	assert.Nil(t, err)
	expText := fmt.Sprintf("%v;\n%v;\n%v;\n",
		mockedItemService.Items[0].String(),
		mockedItemService.Items[1].String(),
		mockedItemService.Items[2].String(),
	)
	assert.Equal(t, expText, text)
	assert.Equal(t, 1, len(keyboard.InlineKeyboard))
	assert.Equal(t, 2, len(keyboard.InlineKeyboard[0]))
	assert.Equal(t, "К предыдущей странице", keyboard.InlineKeyboard[0][0].Text)
	assert.Equal(t, "product__item__list__{\"offset\":47}", *keyboard.InlineKeyboard[0][0].CallbackData)
	assert.Equal(t, "К следующей странице", keyboard.InlineKeyboard[0][1].Text)
	assert.Equal(t, "product__item__list__{\"offset\":53}", *keyboard.InlineKeyboard[0][1].CallbackData)
}
