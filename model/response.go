package model

import (
	"encoding/json"
	"io"
)

type CountResponse struct {
	Msg   string `json:"message"`
}

func (res *CountResponse) ToJSON(w io.Writer) error {
	err := json.NewEncoder(w).Encode(res)
	return err
}

