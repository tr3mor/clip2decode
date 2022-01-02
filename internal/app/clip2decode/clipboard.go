//go:generate  mockgen -source=clipboard.go -destination="mocks\clipboard.go"
package clip2decode

import (
	"os/exec"
)

type ClipboardInterface interface {
	GetData() (string, error)
	WriteData(string) error
}

type Clipboard struct {
}

func NewClipboard() *Clipboard {
	return &Clipboard{}
}

func (c *Clipboard) GetData() (string, error) {
	output, err := exec.Command("pbpaste").Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func (c *Clipboard) WriteData(s string) error {
	cmd := exec.Command("pbcopy")
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	err = cmd.Start()
	if err != nil {
		return err
	}
	_, err = in.Write([]byte(s))
	err = in.Close()
	if err != nil {
		return err
	}
	return cmd.Wait()
}
