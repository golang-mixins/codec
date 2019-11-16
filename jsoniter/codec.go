// Package jsoniter represents the interface codec implementation.
package jsoniter

import (
	"bytes"
	"encoding/json"
	"golang.org/x/xerrors"

	jsoniter "github.com/json-iterator/go"
)

// Codec predetermines the consistency of the interfaces codec implementation.
type Codec struct {
	core jsoniter.API
}

// Marshal interface value marshalling into byte slice.
func (c *Codec) Marshal(v interface{}) ([]byte, error) {
	result, err := c.core.Marshal(v)
	if err != nil {
		return nil, xerrors.New(err.Error())
	}

	return result, nil
}

// MarshalIndent interface value marshalling with indent into byte slice.
// If the prefix is not empty, then the standard json is used, since in jsoniter prefix is not supported.
func (c *Codec) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	result, err := func() ([]byte, error) {
		if len(prefix) > 0 {
			return json.MarshalIndent(v, prefix, indent)
		}
		return c.core.MarshalIndent(v, prefix, indent)
	}()
	if err != nil {
		return nil, xerrors.New(err.Error())
	}

	return result, nil
}

// Unmarshal unmarshalling byte slice value into interface value.
func (c *Codec) Unmarshal(data []byte, v interface{}) error {
	err := c.core.Unmarshal(data, v)
	if err != nil {
		return xerrors.New(err.Error())
	}

	return nil
}

// UnmarshalWithDisallowUnknownFields unmarshalling byte slice value (with disallow unknown fields) into interface value.
func (c *Codec) UnmarshalWithDisallowUnknownFields(data []byte, v interface{}) error {
	decoder := c.core.NewDecoder(bytes.NewBuffer(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(v)
	if err != nil {
		return xerrors.New(err.Error())
	}

	return nil
}

// New is a Codec constructor.
func New() *Codec {
	return &Codec{jsoniter.ConfigCompatibleWithStandardLibrary}
}
