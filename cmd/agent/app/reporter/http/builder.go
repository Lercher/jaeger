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
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Builder struct to hold configurations
type Builder struct {

	// Endpoint for connecting to Jaeger Collector
	CollectorEndpoint string `yaml:"collectorEndpoint"`

	// Timeout for http response from collector
	CollectorResponseTimeout time.Duration `yaml:"collectorResponseTimeout"`
}

// NewBuilder creates a new reporter builder
func NewBuilder() *Builder {
	return &Builder{}
}

// CreateReporter creates a http-based reporter
func (b *Builder) CreateReporter(logger *zap.Logger) (*Reporter, error) {

	if b.CollectorEndpoint == "" {
		return nil, errors.New("Collector Endpoint can not be empty")
	}

	reporter := NewReporter(b.CollectorEndpoint, b.CollectorResponseTimeout, logger)

	return reporter, nil

}
