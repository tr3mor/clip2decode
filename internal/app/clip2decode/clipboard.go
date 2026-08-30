//go:generate  mockgen -source=clipboard.go -destination="mocks\clipboard.go"
package clip2decode

import (
	"context"
	"errors"
	"log"

	"golang.design/x/clipboard"
)

type ClipboardInterface interface {
	GetData() (string, error)
	WriteData(string) error
}

type Clipboard struct {
}

func NewClipboard() *Clipboard {
	err := clipboard.Init()
	if err != nil {
		log.Fatalf("Clipboard init failed: %s", err)
	}
	return &Clipboard{}
}

func (c *Clipboard) GetData() (string, error) {
	output, err := clipboard.Read(context.Background(), clipboard.FmtText)
	if err != nil {
		return "", err
	}
	if output == nil {
		return "", errors.New("failed to get data from clipboard via systemcall")
	}
	return string(output), nil
}

func (c *Clipboard) WriteData(s string) error {
	status, err := clipboard.Write(context.Background(), clipboard.FmtText, []byte(s))
	if err != nil {
		return err
	}
	if status == nil {
		return errors.New("failed to write to clipboard via systemcall")
	}
	return nil
}
