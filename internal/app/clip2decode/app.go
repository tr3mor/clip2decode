package clip2decode

import (
	"encoding/base64"
)

type ApplicationInterface interface {
	DecodeData() (string, error)
}

type Application struct {
	clipboard ClipboardInterface
}

func NewApplication(c ClipboardInterface) *Application {
	return &Application{c}
}

func (a *Application) DecodeData() (string, error) {
	data, err := a.clipboard.GetData()
	if err != nil {
		return "Failed to get data from clipboard", err
	}
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "Failed to decode base64", err
	}
	err = a.clipboard.WriteData(string(decoded))
	if err != nil {
		return "Failed to write data to clipboard", err
	}
	return string(decoded), nil
}
