package link

var allEntities = []Link{
	{Title: "one", Http: "https://yandex.ru", LinkType: "daily"},
	{Title: "two", Http: "https://yandex.ru", LinkType: "friday"},
	{Title: "three", Http: "https://yandex.ru", LinkType: "asia"},
	{Title: "four"},
	{Title: "five"},
}

type Link struct {
	Title    string
	Http     string
	LinkType string
}
