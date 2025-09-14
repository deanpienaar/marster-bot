package output

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
)

type Console struct {
	writer io.Writer
	colors colorsConfig
	reader bufio.Reader
	debug  bool
}

type colorsConfig struct {
	Header  *color.Color
	Success *color.Color
	Error   *color.Color
	Warning *color.Color
	Info    *color.Color
	Prompt  *color.Color
	Data    *color.Color
	Debug   *color.Color
}

func NewConsole(reader bufio.Reader, debug bool) *Console {
	return &Console{
		writer: os.Stdout,
		reader: reader,
		debug:  debug,
		colors: colorsConfig{
			Header:  color.New(color.FgCyan, color.Bold),
			Success: color.New(color.FgGreen),
			Error:   color.New(color.FgRed, color.Bold),
			Warning: color.New(color.FgYellow),
			Info:    color.New(color.FgWhite),
			Prompt:  color.New(color.FgMagenta),
			Data:    color.New(color.FgBlue, color.Bold),
			Debug:   color.New(color.FgHiBlack),
		},
	}
}

func (c *Console) Header(text string) {
	c.colors.Header.Fprintln(c.writer, text)
}

func (c *Console) HeaderWithBorder(text string) {
	border := strings.Repeat("=", len(text))
	c.colors.Header.Fprintln(c.writer, text)
	c.colors.Header.Fprintln(c.writer, border)
}

func (c *Console) Success(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	c.colors.Success.Fprintln(c.writer, "✓ "+message)
}

func (c *Console) Error(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	c.colors.Error.Fprintln(c.writer, "✗ "+message)
}

func (c *Console) Warning(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	c.colors.Warning.Fprintln(c.writer, "⚠ "+message)
}

func (c *Console) Info(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	c.colors.Info.Fprintln(c.writer, message)
}

func (c *Console) Prompt(text string) (string, error) {
	c.colors.Prompt.Fprint(c.writer, "▸ "+text)

	response, err := c.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	response = strings.TrimSpace(response)

	return response, nil

}

func (c *Console) Data(label string, value interface{}) {
	c.colors.Info.Fprint(c.writer, label+": ")
	c.colors.Data.Fprintln(c.writer, value)
}

func (c *Console) Blank() {
	fmt.Fprintln(c.writer)
}

func (c *Console) Divider() {
	c.colors.Info.Fprintln(c.writer, strings.Repeat("─", 40))
}

func (c *Console) Debug(format string, args ...interface{}) {
	if c.debug {
		message := fmt.Sprintf(format, args...)
		c.colors.Debug.Fprintln(c.writer, "[DEBUG] "+message)
	}
}
