package wcFruit

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"coba/model/mFruit/rqFruit"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcFruit__ORM.GEN.go
//go:generate replacer 'Id" form' 'Id,string" form' type wcFruit__ORM.GEN.go
//go:generate replacer 'json:"id"' 'json:"id,string"' type wcFruit__ORM.GEN.go
//go:generate replacer 'By" form' 'By,string" form' type wcFruit__ORM.GEN.go
// go:generate msgp -tests=false -file wcFruit__ORM.GEN.go -o wcFruit__MSG.GEN.go

type FruitsMutator struct {
	rqFruit.Fruits
	mutations []A.X
}

func NewFruitsMutator(adapter *Tt.Adapter) *FruitsMutator {
	return &FruitsMutator{Fruits: rqFruit.Fruits{Adapter: adapter}}
}

func (f *FruitsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(f.mutations) > 0
}

// Overwrite all columns, error if not exists
func (f *FruitsMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := f.Adapter.Update(f.SpaceName(), f.UniqueIndexId(), A.X{f.Id}, f.ToUpdateArray())
	return !L.IsError(err, `Fruits.DoOverwriteById failed: `+f.SpaceName())
}

// Update only mutated, error if not exists, use Find* and Set* methods instead of direct assignment
func (f *FruitsMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !f.HaveMutation() {
		return true
	}
	_, err := f.Adapter.Update(f.SpaceName(), f.UniqueIndexId(), A.X{f.Id}, f.mutations)
	return !L.IsError(err, `Fruits.DoUpdateById failed: `+f.SpaceName())
}

func (f *FruitsMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := f.Adapter.Delete(f.SpaceName(), f.UniqueIndexId(), A.X{f.Id})
	return !L.IsError(err, `Fruits.DoDeletePermanentById failed: `+f.SpaceName())
}

// func (f *FruitsMutator) DoUpsert() bool { //nolint:dupl false positive
//	_, err := f.Adapter.Upsert(f.SpaceName(), f.ToArray(), A.X{
//		A.X{`=`, 0, f.Id},
//		A.X{`=`, 1, f.CreatedAt},
//		A.X{`=`, 2, f.CreatedBy},
//		A.X{`=`, 3, f.UpdatedAt},
//		A.X{`=`, 4, f.UpdatedBy},
//		A.X{`=`, 5, f.DeletedAt},
//		A.X{`=`, 6, f.DeletedBy},
//		A.X{`=`, 7, f.IsDeleted},
//		A.X{`=`, 8, f.RestoredAt},
//		A.X{`=`, 9, f.RestoredBy},
//		A.X{`=`, 10, f.Name},
//		A.X{`=`, 11, f.Description},
//	})
//	return !L.IsError(err, `Fruits.DoUpsert failed: `+f.SpaceName())
// }

// insert, error if exists
func (f *FruitsMutator) DoInsert() bool { //nolint:dupl false positive
	row, err := f.Adapter.Insert(f.SpaceName(), f.ToArray())
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			f.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Fruits.DoInsert failed: `+f.SpaceName())
}

// replace = upsert, only error when there's unique secondary key
func (f *FruitsMutator) DoReplace() bool { //nolint:dupl false positive
	_, err := f.Adapter.Replace(f.SpaceName(), f.ToArray())
	return !L.IsError(err, `Fruits.DoReplace failed: `+f.SpaceName())
}

func (f *FruitsMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != f.Id {
		f.mutations = append(f.mutations, A.X{`=`, 0, val})
		f.Id = val
		return true
	}
	return false
}

func (f *FruitsMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != f.CreatedAt {
		f.mutations = append(f.mutations, A.X{`=`, 1, val})
		f.CreatedAt = val
		return true
	}
	return false
}

func (f *FruitsMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != f.CreatedBy {
		f.mutations = append(f.mutations, A.X{`=`, 2, val})
		f.CreatedBy = val
		return true
	}
	return false
}

func (f *FruitsMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != f.UpdatedAt {
		f.mutations = append(f.mutations, A.X{`=`, 3, val})
		f.UpdatedAt = val
		return true
	}
	return false
}

func (f *FruitsMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != f.UpdatedBy {
		f.mutations = append(f.mutations, A.X{`=`, 4, val})
		f.UpdatedBy = val
		return true
	}
	return false
}

func (f *FruitsMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != f.DeletedAt {
		f.mutations = append(f.mutations, A.X{`=`, 5, val})
		f.DeletedAt = val
		return true
	}
	return false
}

func (f *FruitsMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != f.DeletedBy {
		f.mutations = append(f.mutations, A.X{`=`, 6, val})
		f.DeletedBy = val
		return true
	}
	return false
}

func (f *FruitsMutator) SetIsDeleted(val bool) bool { //nolint:dupl false positive
	if val != f.IsDeleted {
		f.mutations = append(f.mutations, A.X{`=`, 7, val})
		f.IsDeleted = val
		return true
	}
	return false
}

func (f *FruitsMutator) SetRestoredAt(val int64) bool { //nolint:dupl false positive
	if val != f.RestoredAt {
		f.mutations = append(f.mutations, A.X{`=`, 8, val})
		f.RestoredAt = val
		return true
	}
	return false
}

func (f *FruitsMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != f.RestoredBy {
		f.mutations = append(f.mutations, A.X{`=`, 9, val})
		f.RestoredBy = val
		return true
	}
	return false
}

func (f *FruitsMutator) SetName(val string) bool { //nolint:dupl false positive
	if val != f.Name {
		f.mutations = append(f.mutations, A.X{`=`, 10, val})
		f.Name = val
		return true
	}
	return false
}

func (f *FruitsMutator) SetDescription(val string) bool { //nolint:dupl false positive
	if val != f.Description {
		f.mutations = append(f.mutations, A.X{`=`, 11, val})
		f.Description = val
		return true
	}
	return false
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
