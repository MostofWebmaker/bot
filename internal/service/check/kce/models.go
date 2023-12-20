package kce

import (
	"fmt"
	"log"
)

var allEntities = []KCEInfo{
	{Data: "one"},
	{Data: "two"},
	{Data: "three"},
	{Data: "four"},
	{Data: "five"},
}

type KCEInfo struct {
	Found []interface{}
	Data  string
}

func (s *KCEInfo) getField(fieldName string) string {
	v := s.Found[0]
	data, ok := v.(map[string]interface{})
	if !ok {
		log.Panic("cannot convert")
	}

	value, ok := data[fieldName]
	if !ok {
		fmt.Println("trackNo is missing")
	}

	return value.(string)
}

func (s *KCEInfo) GetData() string {
	return "Заказ: №" + s.getField("Number") + "\n" +
		s.getField("Info") + "\n" +
		"Откуда: " + s.getField("FromGeo") + " " + s.getField("TakeDate") + "\n" +
		"Куда: " + s.getField("ToGeo") + " " + s.getField("DeliveryDate") + "\n" +
		"Статус доставки: " + s.getField("State") + "\n"
}
