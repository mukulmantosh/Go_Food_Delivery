package announcements

type FlashEvents struct {
	Headline string `json:"headline"`
	Message  string `json:"message"`
}

type Announcement interface {
	FlashEvents() (*[]FlashEvents, error)
}
