package model

import (
	pb "github.com/raystack/optimus/protos/raystack/optimus/core/v1beta1"
)

type ResourceSpec struct {
	Version int                    `yaml:"version"`
	Name    string                 `yaml:"name"`
	Type    string                 `yaml:"type"`
	Labels  map[string]string      `yaml:"labels"`
	Spec    map[string]interface{} `yaml:"spec"`
	Path    string                 `yaml:"-"`
}

func (r ResourceSpec) ToProto() (*pb.ResourceSpecification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: check if we really need assets
