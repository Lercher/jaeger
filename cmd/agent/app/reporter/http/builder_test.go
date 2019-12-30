// Copyright (c) 2019 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

var yamlConfig = `
collectorEndpoint: http://127.0.0.1:14268
`

func TestBuilderFromConfig(t *testing.T) {

	b := Builder{}
	err := yaml.Unmarshal([]byte(yamlConfig), &b)
	require.NoError(t, err)

	r, err := b.CreateReporter(zap.NewNop())
	require.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, "http://127.0.0.1:14268", r.url)

}

func TestBuilderEndpointFailure(t *testing.T) {

	b := Builder{}

	_, err := b.CreateReporter(zap.NewNop())
	require.Error(t, err)

}
