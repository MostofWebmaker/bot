package box

var allEntities = []BoxInfo{}

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

func (s *BoxInfo) GetData() string {
	return "Заказ: №" + s.Order_id + "\n" +
		"Пункт выдачи: " + s.Point_address + "\n" +
		"Плановая дата доставки: " + s.Delivery_date + "\n" +
		"Статус доставки: " + s.getStatus() + "\n"
}
