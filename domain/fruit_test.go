package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFruit(t *testing.T) {
	d := NewDomain()
	var idFruit uint64
	t.Run(`fruit name less than 3 character must error`, func(t *testing.T) {
		in := &InsertFruit_In{
			Name:        "ap",
			Description: "deskripsi",
		}
		out := d.InsertFruit(in)
		idFruit = out.Fruit.Id
		assert.Equal(t, out.Error, "Name must be at least 3 character")
	})
	t.Run(`list fruit should empty`, func(t *testing.T) {
		in := &FruitList_In{
			Limit: 10,
		}
		out := d.FruitList(in)
		assert.Empty(t, out.Error)
		assert.Empty(t, out.Fruits)
	})
	t.Run(`insert fail fruit description must be filled`, func(t *testing.T) {
		in := &InsertFruit_In{
			Name:        "apple",
			Description: "",
		}
		out := d.InsertFruit(in)
		idFruit = out.Fruit.Id
		assert.Equal(t, out.Error, "Description must be filled")
	})
	t.Run(`list fruit should empty`, func(t *testing.T) {
		in := &FruitList_In{
			Limit: 10,
		}
		out := d.FruitList(in)
		assert.Empty(t, out.Error)
		assert.Empty(t, out.Fruits)
	})
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
	t.Run(`update fruit id not found`, func(t *testing.T) {
		in := &UpdateFruit_In{
			Id:          idFruit + 1,
			Name:        "pisang",
			Description: "deskripsi",
		}
		out := d.UpdateFruit(in)
		assert.Equal(t, out.Error, "Fruit id not found on database")
	})
	t.Run(`list fruit should still same before update`, func(t *testing.T) {
		in := &FruitList_In{
			Limit: 10,
		}
		out := d.FruitList(in)
		assert.Empty(t, out.Error)
		assert.Equal(t, out.Total, uint32(1))
		assert.Equal(t, out.Fruits[0].Name, "apple")
	})
	t.Run(`update fruit failed name must be at least 3 character`, func(t *testing.T) {
		in := &UpdateFruit_In{
			Id:          idFruit,
			Name:        "pi",
			Description: "deskripsi",
		}
		out := d.UpdateFruit(in)
		assert.Equal(t, out.Error, "Name must be at least 3 character")
	})
	t.Run(`list fruit should still same before update`, func(t *testing.T) {
		in := &FruitList_In{
			Limit: 10,
		}
		out := d.FruitList(in)
		assert.Empty(t, out.Error)
		assert.Equal(t, out.Total, uint32(1))
		assert.Equal(t, out.Fruits[0].Name, "apple")
	})
	t.Run(`update fruit failed description must be filled`, func(t *testing.T) {
		in := &UpdateFruit_In{
			Id:          idFruit,
			Name:        "pisang",
			Description: "",
		}
		out := d.UpdateFruit(in)
		assert.Equal(t, out.Error, "Description must be filled")
	})
	t.Run(`list fruit should still same before update`, func(t *testing.T) {
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
	t.Run(`delete fruit id not found in database`, func(t *testing.T) {
		in := &DeleteFruit_In{
			Id: idFruit + 1,
		}
		out := d.DeleteFruit(in)
		assert.Equal(t, out.Error, "Fruit id not found on database")
	})
	t.Run(`list fruit should contain isDeleted fruit still false`, func(t *testing.T) {
		in := &FruitList_In{
			Limit: 10,
		}
		out := d.FruitList(in)
		assert.Empty(t, out.Error)
		assert.Equal(t, out.Fruits[0].IsDeleted, false)
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
	t.Run(`restore fruit id not found`, func(t *testing.T) {
		in := &RestoreFruit_In{
			Id: idFruit + 1,
		}
		out := d.RestoreFruit(in)
		assert.Equal(t, out.Error, "Fruit id not found on database")
	})
	t.Run(`list fruit should contain isDeleted fruit still true`, func(t *testing.T) {
		in := &FruitList_In{
			Limit: 10,
		}
		out := d.FruitList(in)
		assert.Empty(t, out.Error)
		assert.Equal(t, out.Fruits[0].IsDeleted, true)
	})
	t.Run(`restore fruit`, func(t *testing.T) {
		in := &RestoreFruit_In{
			Id: idFruit,
		}
		out := d.RestoreFruit(in)
		assert.Empty(t, out.Error)
		assert.Equal(t, out.Fruit.IsDeleted, false)
	})
	t.Run(`list fruit should contain isDeleted fruit false`, func(t *testing.T) {
		in := &FruitList_In{
			Limit: 10,
		}
		out := d.FruitList(in)
		assert.Empty(t, out.Error)
		assert.Equal(t, out.Fruits[0].IsDeleted, false)
	})
}
