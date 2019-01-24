package zoop

type ErrImpl interface {
	Error() string
	GetStatus() string
	GetStatusCode() int
	GetType() string
	GetCategory() string
	GetMessage() string
}

type ErrResponse struct {
	ErrObject *ErrObject `json:"error"`
}

type ErrObject struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Type       string `json:"type"`
	Category   string `json:"category"`
	Message    string `json:"message"`
}

func (err ErrObject) Error() string {
	return err.Message
}

func (err ErrObject) GetStatus() string {
	return err.Status
}

func (err ErrObject) GetStatusCode() int {
	return err.StatusCode
}

func (err ErrObject) GetType() string {
	return err.Type
}

func (err ErrObject) GetCategory() string {
	return err.Category
}

func (err ErrObject) GetMessage() string {
	return err.Message
}
