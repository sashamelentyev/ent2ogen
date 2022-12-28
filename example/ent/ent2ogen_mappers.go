// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"

	openapi "github.com/ogen-go/ent2ogen/example/api"
)

// Following edges must be loaded:
//
//	switches
//	keycaps
func KeyboardSliceToOpenAPI(s []*Keyboard) (_ []openapi.Keyboard, err error) {
	result := make([]openapi.Keyboard, len(s))
	for i, v := range s {
		result[i], err = v.toOpenAPI()
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Following edges must be loaded:
//
//	switches
//	keycaps
func (e *Keyboard) ToOpenAPI() (*openapi.Keyboard, error) {
	t, err := e.toOpenAPI()
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (e *Keyboard) toOpenAPI() (t openapi.Keyboard, err error) {
	t.ID = e.ID
	t.Name = e.Name
	t.Price = e.Price
	if e.Discount != nil {
		t.Discount.SetTo(*e.Discount)
	} else {
		t.Discount.Null = true
	}
	// Edge 'switches'.
	if err := func() error {
		v, err := e.Edges.SwitchesOrErr()
		if err != nil {
			return fmt.Errorf("load: %w", err)
		}
		openapiType, err := v.toOpenAPI()
		if err != nil {
			return fmt.Errorf("convert to openapi: %w", err)
		}
		t.Switches = openapiType
		return nil
	}(); err != nil {
		return t, fmt.Errorf("edge 'switches': %w", err)
	}
	// Edge 'keycaps'.
	if err := func() error {
		v, err := e.Edges.KeycapsOrErr()
		if err != nil {
			return fmt.Errorf("load: %w", err)
		}
		openapiType, err := v.toOpenAPI()
		if err != nil {
			return fmt.Errorf("convert to openapi: %w", err)
		}
		t.Keycaps = openapiType
		return nil
	}(); err != nil {
		return t, fmt.Errorf("edge 'keycaps': %w", err)
	}
	return t, nil
}

func KeycapModelSliceToOpenAPI(s []*KeycapModel) (_ []openapi.Keycaps, err error) {
	result := make([]openapi.Keycaps, len(s))
	for i, v := range s {
		result[i], err = v.toOpenAPI()
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (e *KeycapModel) ToOpenAPI() (*openapi.Keycaps, error) {
	t, err := e.toOpenAPI()
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (e *KeycapModel) toOpenAPI() (t openapi.Keycaps, err error) {
	t.ID = e.ID
	t.Name = e.Name
	t.Profile = e.Profile
	t.Material = openapi.KeycapsMaterial(e.Material)
	return t, nil
}

func SwitchModelSliceToOpenAPI(s []*SwitchModel) (_ []openapi.Switches, err error) {
	result := make([]openapi.Switches, len(s))
	for i, v := range s {
		result[i], err = v.toOpenAPI()
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (e *SwitchModel) ToOpenAPI() (*openapi.Switches, error) {
	t, err := e.toOpenAPI()
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (e *SwitchModel) toOpenAPI() (t openapi.Switches, err error) {
	t.ID = e.ID
	t.Name = e.Name
	t.SwitchType = openapi.SwitchesSwitchType(e.SwitchType)
	return t, nil
}
