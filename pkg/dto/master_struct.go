package dto

import (
	"net/http"
	"time"

	"github.com/storeapp/pkg/constants"
)

type MasterStruct struct {
	Message   Message     `json:"message,omitempty"`
	Payload   interface{} `json:"payload,omitempty"`
	Timestamp time.Time   `json:"timestamp,omitempty"`
}
type Message struct {
	Code              int    `json:"code,omitempty"`
	Description       string `json:"description,omitempty"`
	DetailDescription string `json:"detailDescription,omitempty"`
}

func CreateMasterStruct(payloadDetails interface{}) *MasterStruct {

	return &MasterStruct{
		Message: Message{
			Code:              http.StatusOK,
			Description:       constants.SuccessText,
			DetailDescription: constants.SuccessDetailDescription,
		},
		Payload:   payloadDetails,
		Timestamp: time.Now(),
	}
}
