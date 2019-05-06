// Package parser is an API blueprint parser and renderer
package parser

import (
	"bytes"
	"io"

	"github.com/bukalapak/snowboard/adapter/drafter"
	"github.com/bukalapak/snowboard/api"
	"github.com/bukalapak/snowboard/loader"
)

// Parse formats API blueprint as blueprint.API struct
func Parse(r io.Reader) (*api.API, error) {
	el, err := parseElement(r)
	if err != nil {
		return nil, err
	}

	return api.NewAPI(el)
}

// ParseAsJSON parse API blueprint as API Element JSON
func ParseAsJSON(r io.Reader) ([]byte, error) {
	return drafter.Parse(r)
}

// Validate validates API blueprint
func Validate(r io.Reader) (*api.API, error) {
	el, err := validateElement(r)
	if err == nil && el.Object() == nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return api.NewAPI(el)
}

// Load reads API blueprint from file as blueprint.API struct
func Load(name string) (*api.API, error) {
	b, err := loader.Load(name)
	if err != nil {
		return nil, err
	}

	return Parse(bytes.NewReader(b))
}

// LoadAsJSON reads API blueprint from file as API Element JSON
func LoadAsJSON(name string) ([]byte, error) {
	b, err := loader.Load(name)
	if err != nil {
		return nil, err
	}

	return ParseAsJSON(bytes.NewReader(b))
}

func parseElement(r io.Reader) (*api.Element, error) {
	b, err := ParseAsJSON(r)
	if err != nil {
		return nil, err
	}

	return api.ParseJSON(bytes.NewReader(b))
}

func validateElement(r io.Reader) (*api.Element, error) {
	b, err := drafter.Validate(r)
	if err != nil {
		return nil, err
	}

	if len(b) == 0 {
		return &api.Element{}, nil
	}

	return api.ParseJSON(bytes.NewReader(b))
}
