package main

import (
	"container/list"
	"fmt"
)

var (
	TestOrderType = [2]string{"3", "52a447d3-a64a-11e3-9aeb-50e549534c5e"}
	DrugOrderType = [2]string{"2", "131168f4-15f5-102d-96e4-000c29c2a5d7"}
)

const mapOrderTypeMetaData = "update amrs.order_type set uuid = '%s' where order_type_id = %s"

func MapOrderTypeMetadata() {
	orderList := mappedOrderTypeValues()
	for orderList.Len() > 0 {
		vals := orderList.Front().Value.([]string)
		kemrId := vals[0]
		uuid := vals[1]
		fmt.Println("-------Mapping Data Sets------")
		_, _ = amrs.Exec(fmt.Sprintf(mapOrderTypeMetaData, uuid, kemrId))
		orderList.Remove(orderList.Front())
	}
}

func mappedOrderTypeValues() *list.List {
	a := list.New()
	a.PushBack(TestOrderType)
	a.PushBack(DrugOrderType)
	return a
}
