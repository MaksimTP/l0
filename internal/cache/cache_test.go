package cache

import (
	"main/internal/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := New()
	assert.NotNil(t, c)
}

func TestSaveData(t *testing.T) {
	c := New()
	data := types.Order{OrderUid: "qwe"}
	c.SaveData(data)
	assert.Equal(t, c.cachedData["qwe"], data)
}

func TestOrderById(t *testing.T) {
	c := New()
	data := types.Order{OrderUid: "qwe"}
	c.SaveData(data)
	res, err := c.GetOrderById("qwe")
	assert.Nil(t, err)
	assert.Equal(t, res, data)
}
