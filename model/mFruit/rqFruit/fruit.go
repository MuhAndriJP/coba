package rqFruit

import (
	"coba/conf"

	"github.com/kokizzu/gotro/A"

	"github.com/kokizzu/gotro/L"

	"github.com/kokizzu/gotro/X"
)

func (p *Fruits) FindOffsetLimit(offset, limit uint32) (res []*Fruits) {
	query := `
SELECT ` + p.sqlSelectAllFields() + `
FROM ` + p.sqlTableName() + `
ORDER BY ` + p.sqlId() + `
LIMIT ` + X.ToS(limit) + `
OFFSET ` + X.ToS(offset)
	if conf.DEBUG_MODE {
		L.Print(query)
	}
	p.Adapter.QuerySql(query, func(row []interface{}) {
		obj := &Fruits{}
		obj.FromArray(row)
		res = append(res, obj)
	})
	return
}

func (p *Fruits) FindByIds(ids ...uint64) (res []*Fruits) {
	query := `
SELECT ` + p.sqlSelectAllFields() + `
FROM ` + p.sqlTableName() + `
WHERE ` + p.sqlId() + ` IN (` + A.UIntJoin(ids, `,`) + `)
ORDER BY ` + p.sqlId()
	if conf.DEBUG_MODE {
		L.Print(query)
	}
	p.Adapter.QuerySql(query, func(row []interface{}) {
		obj := &Fruits{}
		obj.FromArray(row)
		res = append(res, obj)
	})
	return
}

func (p *Fruits) DeleteByIds(ids ...uint64) (res []*Fruits) {
	query := `
DELETE ` + p.sqlSelectAllFields() + `
FROM ` + p.sqlTableName() + `
WHERE ` + p.sqlId() + ` IN (` + A.UIntJoin(ids, `,`) + `)
ORDER BY ` + p.sqlId()
	if conf.DEBUG_MODE {
		L.Print(query)
	}
	p.Adapter.QuerySql(query, func(row []interface{}) {
		obj := &Fruits{}
		obj.FromArray(row)
		res = append(res, obj)
	})
	return
}

func (p *Fruits) UpdateByIds(ids ...uint64) (res []*Fruits) {
	query := `
UPDATE ` + p.sqlSelectAllFields() + `
FROM ` + p.sqlTableName() + `
WHERE ` + p.sqlId() + ` IN (` + A.UIntJoin(ids, `,`) + `)
ORDER BY ` + p.sqlId()
	if conf.DEBUG_MODE {
		L.Print(query)
	}
	p.Adapter.QuerySql(query, func(row []interface{}) {
		obj := &Fruits{}
		obj.FromArray(row)
		res = append(res, obj)
	})
	return
}
