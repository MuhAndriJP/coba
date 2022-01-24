package domain

import (
	"coba/model/mFruit/rqFruit"
	"coba/model/mFruit/wcFruit"

	"github.com/kokizzu/id64"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file template.go
//go:generate replacer 'Id" form' 'Id,string" form' type template.go
//go:generate replacer 'json:"id"' 'json:id,string" form' type template.go
//go:generate replacer 'By" form' 'By,string" form' type template.go
// go:generate msgp -tests=false -file template.go -o template__MSG.GEN.go

type (
	FruitList_In struct {
		RequestCommon
		Limit  uint32 `json:"limit" form:"limit" query:"limit" long:"limit" msg:"limit"`
		Offset uint32 `json:"offset" form:"offset" query:"offset" long:"offset" msg:"offset"`
	}
	FruitList_Out struct {
		ResponseCommon
		Limit  uint32 `json:"limit" form:"limit" query:"limit" long:"limit" msg:"limit"`
		Offset uint32 `json:"offset" form:"offset" query:"offset" long:"offset" msg:"offset"`
		Total  uint32 `json:"total" form:"total" query:"total" long:"total" msg:"total"`
		Fruits []*rqFruit.Fruits
	}
)

const FruitList_Url = `/Fruit`

func (d *Domain) FruitList(in *FruitList_In) (out FruitList_Out) {
	fruit := rqFruit.NewFruits(d.Taran)
	out.Limit = in.Limit
	out.Offset = in.Offset
	out.Total = uint32(fruit.Total())
	out.Fruits = fruit.FindOffsetLimit(in.Offset, in.Limit)
	return
}

type (
	DeleteFruit_In struct {
		RequestCommon
		Id uint64 `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	}
	DeleteFruit_Out struct {
		ResponseCommon
		Fruit rqFruit.Fruits
	}
)

const DeleteFruit_Url = `/DeleteFruit`

func (d *Domain) DeleteFruit(in *DeleteFruit_In) (out DeleteFruit_Out) {
	fruit := wcFruit.NewFruitsMutator(d.Taran)
	fruit.Id = in.Id
	if !fruit.FindById() {
		out.SetError(404, "Fruit id not found on database")
		return
	}

	fruit.SetDeletedAt(in.UnixNow())
	fruit.SetIsDeleted(true)
	fruit.SetUpdatedAt(in.UnixNow())
	if !fruit.DoUpdateById() {
		out.SetError(500, "Failed to delete fruit")
		return
	}
	out.Fruit = fruit.Fruits
	return
}

type (
	UpdateFruit_In struct {
		RequestCommon
		Id          uint64 `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
		Name        string `json:"name" form:"name" query:"name" long:"name" msg:"name"`
		Description string `json:"description" form:"description" query:"description" long:"description" msg:"description"`
	}
	UpdateFruit_Out struct {
		ResponseCommon
		Fruit rqFruit.Fruits
	}
)

const UpdateFruit_Url = `/UpdateFruit`

func (d *Domain) UpdateFruit(in *UpdateFruit_In) (out UpdateFruit_Out) {
	fruit := wcFruit.NewFruitsMutator(d.Taran)
	fruit.Id = in.Id
	if !fruit.FindById() {
		out.SetError(404, "Fruit id not found on database")
		return
	}
	fruit.SetName(in.Name)
	fruit.SetDescription(in.Description)

	if len(fruit.Name) < 3 {
		out.SetError(400, "Name must be at least 3 character")
		return
	}
	if fruit.Description == "" {
		out.SetError(400, "Description must be filled")
		return
	}
	fruit.SetUpdatedAt(in.UnixNow())
	if !fruit.DoUpdateById() {
		out.SetError(500, "Failed to update fruit")
		return
	}
	out.Fruit = fruit.Fruits
	return
}

type (
	InsertFruit_In struct {
		RequestCommon
		Name        string `json:"name" form:"name" query:"name" long:"name" msg:"name"`
		Description string `json:"description" form:"description" query:"description" long:"description" msg:"description"`
	}
	InsertFruit_Out struct {
		ResponseCommon
		Fruit rqFruit.Fruits
	}
)

const InsertFruit_Url = `/InsertFruit`

func (d *Domain) InsertFruit(in *InsertFruit_In) (out InsertFruit_Out) {
	fruit := wcFruit.NewFruitsMutator(d.Taran)
	fruit.Name = in.Name
	fruit.Description = in.Description
	if len(fruit.Name) < 3 {
		out.SetError(400, "Name must be at least 3 character")
		return
	}
	if fruit.Description == "" {
		out.SetError(400, "Description must be filled")
		return
	}
	fruit.Id = uint64(id64.ID())
	fruit.CreatedAt = in.UnixNow()
	fruit.UpdatedAt = in.UnixNow()
	if !fruit.DoInsert() {
		out.SetError(500, "Failed to insert fruit")
		return
	}
	out.Fruit = fruit.Fruits
	return
}

type (
	RestoreFruit_In struct {
		RequestCommon
		Id uint64
	}
	RestoreFruit_Out struct {
		ResponseCommon
		Fruit rqFruit.Fruits
	}
)

const RestoreFruit_Url = `/RestoreFruit`

func (d *Domain) RestoreFruit(in *RestoreFruit_In) (out RestoreFruit_Out) {
	fruit := wcFruit.NewFruitsMutator(d.Taran)
	fruit.Id = in.Id
	if !fruit.FindById() {
		out.SetError(404, "Fruit id not found on database")
		return
	}
	fruit.SetRestoredAt(in.UnixNow())
	fruit.SetIsDeleted(false)
	fruit.SetUpdatedAt(in.UnixNow())
	if !fruit.DoUpdateById() {
		out.SetError(500, "Failed to restore fruit")
		return
	}
	out.Fruit = fruit.Fruits
	return
}

// TODO: decrease stock after accepted by seller
