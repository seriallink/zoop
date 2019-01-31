package zoop

type ErrImpl interface {
	Error() string
}

type ErrResponse struct {
	ErrObject *ErrObject `json:"error"`
}

type ErrObject struct {
	Status     string   `json:"status"`
	StatusCode int      `json:"status_code"`
	Type       string   `json:"type"`
	Category   string   `json:"category"`
	Message    string   `json:"message"`
	Reasons    []string `json:"reasons"`
}

func (err ErrObject) Error() string {
	return err.Message
}
