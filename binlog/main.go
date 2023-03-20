package main

import (
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/qinxiaogit/k8sToGo/binlog/table"
	"github.com/qinxiaogit/k8sToGo/binlog/wechat"
	"github.com/siddontang/go-log/log"
)

type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	log.Infof("%s %v\n", e.Action, e.Rows)
	log.Info("-----\n",e.Table.Name,"------\n" )

	table := table.NewTable()
	table.Handle(e.Table.Name,e.Rows)

	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}


func main() {

	wechat:= wechat.Wechat{}
	wechat.SendSubscribeMsg()

	cfg := canal.NewDefaultConfig()
	cfg.Addr = "47.92.55.167:8306"
	cfg.User = "root"
	cfg.Password="zh123456"
	// We only care table canal_test in test db
	cfg.Dump.TableDB = "center_single"
	cfg.Dump.Tables = []string{"t_order_detail","t_order_info","t_center_user"}

	c, err := canal.NewCanal(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Register a handler to handle RowsEvent
	c.SetEventHandler(&MyEventHandler{})

	// Start canal
	_ = c.Run()
}
