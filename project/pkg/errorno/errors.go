package errorno

import (
	"fmt"
	"time"
)

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (self ErrorMessage) Error() string {
	return fmt.Sprintf("Time: %s, Code: %d, Message: %s", time.Now().String(), self.Code, self.Message)
}

type ErrorsMessage []ErrorMessage
