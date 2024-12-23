package model_event_logger

type Event struct {
	UnixTS      int64  `json:"unix_ts"`
	UserID      int    `json:"user_id"`
	EventName   string `json:"event_name"`
	Id          int    `binding:"-"`
	EventSource string `binding:"-"`
}
