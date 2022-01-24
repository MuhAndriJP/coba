package mFruit

import (
	"github.com/kokizzu/gotro/D/Tt"
)

// table fruits
const (
	TableFruits Tt.TableName = `fruits`
	Id                       = `id`
	CreatedBy                = `createdBy`
	CreatedAt                = `createdAt`
	UpdatedBy                = `updatedBy`
	UpdatedAt                = `updatedAt`
	DeletedBy                = `deletedBy`
	DeletedAt                = `deletedAt`
	IsDeleted                = `isDeleted`
	RestoredBy               = `restoredBy`
	RestoredAt               = `restoredAt`
	Name                     = `name`
	Description              = `description`
)

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	// can only adding fields on back, and must IsNullable: true
	// primary key must be first field and set to Unique: fieldName
	TableFruits: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{IsDeleted, Tt.Boolean},
			{RestoredAt, Tt.Integer},
			{RestoredBy, Tt.Unsigned},
			{Name, Tt.String},
			{Description, Tt.String},
		},
		AutoIncrementId: true,
	},
}

func GenerateORM() {
	Tt.GenerateOrm(TarantoolTables)
	//Ch.GenerateOrm(ClickhouseTables) // find d.InitClickhouseBuffer to create chBuffer on NewDomain
}
