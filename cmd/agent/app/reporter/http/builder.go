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
)

// ConnBuilder Struct to hold configurations
type ConnBuilder struct {
	// CollectorHostPorts is list of http://host:port Jaeger Collectors.
	CollectorHostPorts []string `yaml:"collectorHostPorts"`

	// Timeout for http response from collector
	CollectorResponseTimeout time.Duration `yaml:"collectorResponseTimeout"`
}

// NewConnBuilder creates a new grpc connection builder.
func NewConnBuilder() *ConnBuilder {
	return &ConnBuilder{}
}
