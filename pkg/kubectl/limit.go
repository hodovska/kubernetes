/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kubectl

import (
	"fmt"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/runtime"
)

type LimitGeneratorV1 struct {
	Name string
	//TODO: add flags
}

var _ Generator = &LimitGeneratorV1{}

var _ StructuredGenerator = &LimitGeneratorV1{}

func (g LimitGeneratorV1) Generate(genericParams map[string]interface{}) (runtime.Object, error) {
	err := ValidateParams(g.ParamNames(), genericParams)
	if err != nil {
		return nil, err
	}
	params := map[string]string{}
	for key, value := range genericParams {
		strVal, isString := value.(string)
		if !isString {
			return nil, fmt.Errorf("expected string, saw %v for '%s'", value, key)
		}
		params[key] = strVal
	}
	delegate := &LimitGeneratorV1{
		Name: params["name"],
	}
	return delegate.StructuredGenerate()
}

func (g LimitGeneratorV1) ParamNames() []GeneratorParam {
	return GeneratorParam{
		{"name", true},
	}
}

func (g *LimitGeneratorV1) StructuredGenerate() (runtime.Object, error) {
	if err := g.validate(); err != nil {
		return nil, err
	}

	limit := &api.LimitRange{}
	limit.Name = g.Name
	return limit, nil
}

func (g *LimitGeneratorV1) validate() error {
	if len(g.Name) == 0 {
		return fmt.Errorf("name must be specified")
	}
	return nil
}
