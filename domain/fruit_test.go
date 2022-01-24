package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFruit(t *testing.T) {
	d := NewDomain()
	var idFruit uint64
	t.Run(`add fruit`, func(t *testing.T) {
		in := &InsertFruit_In{
			Name:        "apple",
			Description: "deskripsi",
		}
		out := d.InsertFruit(in)
		idFruit = out.Fruit.Id
		assert.Empty(t, out.Error)
		assert.NotEmpty(t, out.Fruit.Id)
		assert.Equal(t, out.Fruit.Name, "apple")
		assert.Equal(t, out.Fruit.Description, "deskripsi")
	})
	t.Run(`list fruit should contain inserted fruit`, func(t *testing.T) {
		in := &FruitList_In{
			Limit: 10,
		}
		out := d.FruitList(in)
		assert.Empty(t, out.Error)
		assert.Equal(t, out.Total, uint32(1))
		assert.Equal(t, out.Fruits[0].Name, "apple")
	})
	t.Run(`update fruit`, func(t *testing.T) {
		in := &UpdateFruit_In{
			Id:          idFruit,
			Name:        "pisang",
			Description: "deskripsi",
		}
		out := d.UpdateFruit(in)
		assert.Empty(t, out.Error)
		assert.Equal(t, out.Fruit.Id, idFruit)
		assert.Equal(t, out.Fruit.Name, "pisang")
		assert.Equal(t, out.Fruit.Description, "deskripsi")
	})
	t.Run(`list fruit should contain updated fruit`, func(t *testing.T) {
		in := &FruitList_In{
			Limit: 10,
		}
		out := d.FruitList(in)
		assert.Empty(t, out.Error)
		assert.Equal(t, out.Total, uint32(1))
		assert.Equal(t, out.Fruits[0].Name, "pisang")
	})
	t.Run(`delete fruit`, func(t *testing.T) {
		in := &DeleteFruit_In{
			Id: idFruit,
		}
		out := d.DeleteFruit(in)
		assert.Empty(t, out.Error)
		assert.Equal(t, out.Fruit.IsDeleted, true)
	})
	t.Run(`list fruit should contain isDeleted fruit true`, func(t *testing.T) {
		in := &FruitList_In{
			Limit: 10,
		}
		out := d.FruitList(in)
		assert.Empty(t, out.Error)
		assert.Equal(t, out.Fruits[0].IsDeleted, true)
	})
}
