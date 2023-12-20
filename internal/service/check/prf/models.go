package prf

import (
	"fmt"
	"log"
)

var allEntities = []PRFInfo{
	{Data: "one"},
	{Data: "two"},
	{Data: "three"},
	{Data: "four"},
	{Data: "five"},
}

type PRFInfo struct {
	DetailedTrackings []interface{}
	Data              string
}

func (s *PRFInfo) getField(fieldName string) interface{} {
	v := s.DetailedTrackings[0]
	data, ok := v.(map[string]interface{})
	if !ok {
		log.Panic("cannot convert")
	}

	item, ok := data["trackingItem"]
	if !ok {
		fmt.Println("trackNo is missing")
	}

	values, ok := item.(map[string]interface{})
	if !ok {
		log.Panic("cannot convert")
	}

	value, ok := values[fieldName]
	if !ok {
		fmt.Println("trackNo is missing")
	}

	return value
}

func (s *PRFInfo) getFieldString(fieldName string) string {
	value, ok := s.getField(fieldName).(string)
	if !ok {
		log.Panic("cannot convert")
	}

	return value
}

func (s *PRFInfo) getFieldStringFromFloat(fieldName string) string {
	value := s.getField(fieldName).(float64)

	return fmt.Sprintf("%.0f", value)
}

func (s *PRFInfo) getMoney() string {
	data := s.getField("insuranceMoney")

	if data == nil {
		return ""
	}

	values, ok := data.(map[string]interface{})
	if !ok {
		log.Panic("cannot convert")
	}

	sourceValue, ok := values["sourceValue"]
	if !ok {
		fmt.Println("trackNo is missing")
	}

	sourceValueItem, ok := sourceValue.(map[string]interface{})
	if !ok {
		log.Panic("cannot convert")
	}

	amount, ok := sourceValueItem["value"]
	if !ok {
		fmt.Println("trackNo is missing")
	}

	amountValue, ok := amount.(float64)
	if !ok {
		log.Panic("cannot convert")
	}

	currency, ok := sourceValueItem["currency"]
	if !ok {
		fmt.Println("trackNo is missing")
	}

	currencyString, ok := currency.(string)
	if !ok {
		log.Panic("cannot convert")
	}

	return fmt.Sprintf("%.0f", amountValue) + " " + currencyString
}

func (s *PRFInfo) GetData() string {
	return s.getFieldString("mailType") + ": №" + s.getFieldString("barcode") + " " + s.getFieldStringFromFloat("weight") + "г" + " " + s.getMoney() + "\n" +
		"Откуда: г. " + s.getFieldString("originCityName") + "\n" +
		"Куда: г. " + s.getFieldString("destinationCityName") + "\n" +
		"Отправитель: " + s.getFieldString("sender") + "\n" +
		"Получатель: " + s.getFieldString("recipient") + "\n" +
		"Статус доставки: " + s.getFieldString("commonStatus") + " " + s.getFieldString("lastOperationDateTime")
}
