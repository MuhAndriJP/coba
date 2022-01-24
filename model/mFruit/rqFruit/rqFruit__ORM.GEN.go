package rqFruit

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"coba/model/mFruit"

	"github.com/tarantool/go-tarantool"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqFruit__ORM.GEN.go
//go:generate replacer 'Id" form' 'Id,string" form' type rqFruit__ORM.GEN.go
//go:generate replacer 'json:"id"' 'json:"id,string"' type rqFruit__ORM.GEN.go
//go:generate replacer 'By" form' 'By,string" form' type rqFruit__ORM.GEN.go
// go:generate msgp -tests=false -file rqFruit__ORM.GEN.go -o rqFruit__MSG.GEN.go

type Fruits struct {
	Adapter     *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id          uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	CreatedAt   int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy   uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt   int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy   uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt   int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy   uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	IsDeleted   bool        `json:"isDeleted" form:"isDeleted" query:"isDeleted" long:"isDeleted" msg:"isDeleted"`
	RestoredAt  int64       `json:"restoredAt" form:"restoredAt" query:"restoredAt" long:"restoredAt" msg:"restoredAt"`
	RestoredBy  uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
	Name        string      `json:"name" form:"name" query:"name" long:"name" msg:"name"`
	Description string      `json:"description" form:"description" query:"description" long:"description" msg:"description"`
}

func NewFruits(adapter *Tt.Adapter) *Fruits {
	return &Fruits{Adapter: adapter}
}

func (f *Fruits) SpaceName() string { //nolint:dupl false positive
	return string(mFruit.TableFruits)
}

func (f *Fruits) sqlTableName() string { //nolint:dupl false positive
	return `"fruits"`
}

func (f *Fruits) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

func (f *Fruits) FindById() bool { //nolint:dupl false positive
	res, err := f.Adapter.Select(f.SpaceName(), f.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{f.Id})
	if L.IsError(err, `Fruits.FindById failed: `+f.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		f.FromArray(rows[0])
		return true
	}
	return false
}

func (f *Fruits) sqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "isDeleted"
	, "restoredAt"
	, "restoredBy"
	, "name"
	, "description"
	`
}

func (f *Fruits) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, f.Id},
		A.X{`=`, 1, f.CreatedAt},
		A.X{`=`, 2, f.CreatedBy},
		A.X{`=`, 3, f.UpdatedAt},
		A.X{`=`, 4, f.UpdatedBy},
		A.X{`=`, 5, f.DeletedAt},
		A.X{`=`, 6, f.DeletedBy},
		A.X{`=`, 7, f.IsDeleted},
		A.X{`=`, 8, f.RestoredAt},
		A.X{`=`, 9, f.RestoredBy},
		A.X{`=`, 10, f.Name},
		A.X{`=`, 11, f.Description},
	}
}

func (f *Fruits) IdxId() int { //nolint:dupl false positive
	return 0
}

func (f *Fruits) sqlId() string { //nolint:dupl false positive
	return `"id"`
}

func (f *Fruits) IdxCreatedAt() int { //nolint:dupl false positive
	return 1
}

func (f *Fruits) sqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

func (f *Fruits) IdxCreatedBy() int { //nolint:dupl false positive
	return 2
}

func (f *Fruits) sqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

func (f *Fruits) IdxUpdatedAt() int { //nolint:dupl false positive
	return 3
}

func (f *Fruits) sqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

func (f *Fruits) IdxUpdatedBy() int { //nolint:dupl false positive
	return 4
}

func (f *Fruits) sqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

func (f *Fruits) IdxDeletedAt() int { //nolint:dupl false positive
	return 5
}

func (f *Fruits) sqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

func (f *Fruits) IdxDeletedBy() int { //nolint:dupl false positive
	return 6
}

func (f *Fruits) sqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

func (f *Fruits) IdxIsDeleted() int { //nolint:dupl false positive
	return 7
}

func (f *Fruits) sqlIsDeleted() string { //nolint:dupl false positive
	return `"isDeleted"`
}

func (f *Fruits) IdxRestoredAt() int { //nolint:dupl false positive
	return 8
}

func (f *Fruits) sqlRestoredAt() string { //nolint:dupl false positive
	return `"restoredAt"`
}

func (f *Fruits) IdxRestoredBy() int { //nolint:dupl false positive
	return 9
}

func (f *Fruits) sqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

func (f *Fruits) IdxName() int { //nolint:dupl false positive
	return 10
}

func (f *Fruits) sqlName() string { //nolint:dupl false positive
	return `"name"`
}

func (f *Fruits) IdxDescription() int { //nolint:dupl false positive
	return 11
}

func (f *Fruits) sqlDescription() string { //nolint:dupl false positive
	return `"description"`
}

func (f *Fruits) ToArray() A.X { //nolint:dupl false positive
	var id interface{} = nil
	if f.Id != 0 {
		id = f.Id
	}
	return A.X{
		id,
		f.CreatedAt,   // 1
		f.CreatedBy,   // 2
		f.UpdatedAt,   // 3
		f.UpdatedBy,   // 4
		f.DeletedAt,   // 5
		f.DeletedBy,   // 6
		f.IsDeleted,   // 7
		f.RestoredAt,  // 8
		f.RestoredBy,  // 9
		f.Name,        // 10
		f.Description, // 11
	}
}

func (f *Fruits) FromArray(a A.X) *Fruits { //nolint:dupl false positive
	f.Id = X.ToU(a[0])
	f.CreatedAt = X.ToI(a[1])
	f.CreatedBy = X.ToU(a[2])
	f.UpdatedAt = X.ToI(a[3])
	f.UpdatedBy = X.ToU(a[4])
	f.DeletedAt = X.ToI(a[5])
	f.DeletedBy = X.ToU(a[6])
	f.IsDeleted = X.ToBool(a[7])
	f.RestoredAt = X.ToI(a[8])
	f.RestoredBy = X.ToU(a[9])
	f.Name = X.ToS(a[10])
	f.Description = X.ToS(a[11])
	return f
}

func (f *Fruits) Total() int64 { //nolint:dupl false positive
	rows := f.Adapter.CallBoxSpace(f.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
