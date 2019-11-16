// Package std represents the interface codec implementation.
package std

import (
	"bytes"
	"encoding/json"
	"golang.org/x/xerrors"
)

// Codec predetermines the consistency of the interfaces codec implementation.
type Codec struct{}

// Marshal interface value marshalling into byte slice.
func (c *Codec) Marshal(v interface{}) ([]byte, error) {
	result, err := json.Marshal(v)
	if err != nil {
		return nil, xerrors.New(err.Error())
	}

	return result, nil
}

// MarshalIndent interface value marshalling with indent into byte slice.
func (c *Codec) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	result, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		return nil, xerrors.New(err.Error())
	}

	return result, nil
}

// Unmarshal unmarshalling byte slice value into interface value.
func (c *Codec) Unmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return xerrors.New(err.Error())
	}

	return nil
}

// UnmarshalWithDisallowUnknownFields unmarshalling byte slice value (with disallow unknown fields) into interface value.
func (c *Codec) UnmarshalWithDisallowUnknownFields(data []byte, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewBuffer(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(v)
	if err != nil {
		return xerrors.New(err.Error())
	}

	return nil
}

// New is a Codec constructor.
func New() *Codec {
	return &Codec{}
}
