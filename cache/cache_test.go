package cache

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMockCache(t *testing.T) {
	mockCtrl, _ := gomock.WithContext(context.Background(), t)
	defer mockCtrl.Finish()

	mockComposer := NewMockComposer(mockCtrl)
	mockComposer.EXPECT().Initialize(1).AnyTimes()
	mockComposer.EXPECT().Put("a", "1").AnyTimes()
	mockComposer.EXPECT().Get("a").Return("1", nil).AnyTimes()
	mockComposer.EXPECT().Get("b").Return("", errors.New("empty")).AnyTimes()
	RegisterCache("test", mockComposer)

	c := NewCache(Config{Policy: "test", Size: 1})
	t.Run("value is not empty", func(t *testing.T) {
		c.Put("a", "1")
		v, err := c.Get("a")
		assert.NoError(t, err, "error should be nil")
		assert.Equal(t, "1", v)
	})

	t.Run("value is empty", func(t *testing.T) {
		v, err := c.Get("b")
		assert.Error(t, err, "error should be not nil")
		assert.Empty(t, v, "value should be empty")
	})
}
