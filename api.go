package gojsontmpl

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Config struct {
	SkipJSONValidation bool
	ValidateReaderFunc func(c *Config, buf *bytes.Buffer) (*bytes.Buffer, error)
}

func Default() *Config {
	return &Config{
		ValidateReaderFunc: validateJSONReader,
	}
}

func (c *Config) NewBuilder(tmpl string) *Builder {
	return &Builder{
		tmpl:   tmpl,
		config: c,
	}
}

func validateJSONReader(c *Config, buf *bytes.Buffer) (*bytes.Buffer, error) {
	if c.SkipJSONValidation {
		return buf, nil
	}

	var ob interface{}
	err := json.Unmarshal(buf.Bytes(), &ob)
	if err != nil {
		return nil, &ValueError{Err: err}
	}
	return bytes.NewBuffer(buf.Bytes()), nil
}

type ValueError struct {
	Err error
}

func (err *ValueError) Error() string {
	return fmt.Sprintf("value error: %s", err.Err.Error())
}

func (err *ValueError) Unwrap() error {
	return err.Err
}

type Builder struct {
	config *Config
	tmpl   string
}

func (b *Builder) ToReader(replacer func(string) string) (*bytes.Buffer, error) {
	tmpl := replacer(b.tmpl)
	buf := bytes.NewBufferString(tmpl)
	return b.config.ValidateReaderFunc(b.config, buf)
}
