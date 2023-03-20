package table

type t interface {
	getTableName() string
}

type table struct {
}

func (t *table) Handle(tableName string,e [][]interface{}) {
	switch tableName {
	case centerUser{}.getTableName():
		centerUser{}.onInsert(e)
	}
}

func NewTable() *table {
	return &table{}
}
