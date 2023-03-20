package table

type centerUser struct {
}

func (c centerUser) getTableName() string {
	return "t_center_user"
}

func onUpdateRow() {

}

func (c *centerUser) onInsert(e [][]interface{}) {
	for _,row := range e{
		
	}
}
