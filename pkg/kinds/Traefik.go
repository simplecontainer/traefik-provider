package kinds

import (
	"encoding/json"
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
)

func New(definition []byte) (*Traefik, error) {
	var cr *v1.CustomDefinition
	if err := json.Unmarshal(definition, &cr); err != nil {
		return nil, err
	}

	var traefik *Traefik

	bytes, err := json.Marshal(cr.Spec)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &traefik); err != nil {
		return nil, err
	}

	return traefik, nil
}
