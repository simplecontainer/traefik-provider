package kinds

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
)

func New(definition []byte) (*Traefik, error) {
	var cr *v1.CustomDefinition

	if err := json.Unmarshal(definition, &cr); err != nil {
		return nil, err
	}

	var traefik *Traefik

	if err := mapstructure.Decode(cr.Spec, &traefik); err != nil {
		return nil, err
	}

	return traefik, nil
}
