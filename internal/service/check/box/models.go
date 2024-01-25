package box

import (
	"errors"
	"fmt"
)

var allEntities = []BoxInfo{}
var errProviderDataNotFound = errors.New("У провайдера нет данных по данном трек номеру! Попробуйте проверить другой.")

type BoxInfo struct {
	Track_id      string
	Order_id      string
	Point_address string
	Delivery_date string
	Statuses      []BoxInfoStatus
}

type BoxInfoStatus struct {
	Name string
}

func (s *BoxInfo) getStatus() string {
	return s.Statuses[len(s.Statuses)-1].Name
}

func (s *BoxInfo) GetData() (string, error) {
	if len(s.Statuses) == 0 {
		fmt.Printf("we don't have response")
		return "", errProviderDataNotFound
	}
	return "Заказ: №" + s.Order_id + "\n" +
		"Пункт выдачи: " + s.Point_address + "\n" +
		"Плановая дата доставки: " + s.Delivery_date + "\n" +
		"Статус доставки: " + s.getStatus() + "\n", nil
}
