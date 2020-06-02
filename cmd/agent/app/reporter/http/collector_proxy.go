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
	"errors"
	"io"

	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/cmd/agent/app/configmanager"
	"github.com/jaegertracing/jaeger/cmd/agent/app/reporter"
)

// ProxyBuilder holds objects communicating with collector
type ProxyBuilder struct {
	reporter reporter.Reporter
	manager  configmanager.ClientConfigManager
}

// NewCollectorProxy creates ProxyBuilder
func NewCollectorProxy(builder *ConnBuilder, agentTags map[string]string, mFactory metrics.Factory, logger *zap.Logger) (*ProxyBuilder, error) {
	if len(builder.CollectorHostPorts) != 1 {
		return nil, errors.New("exactly one host:port for " + collectorHostPort + " is required")
	}
	r := NewReporter(builder.CollectorHostPorts[0], builder.CollectorResponseTimeout, agentTags, logger)

	httpMetrics := mFactory.Namespace(metrics.NSOptions{Name: "", Tags: map[string]string{"protocol": "http"}})

	return &ProxyBuilder{
		reporter: reporter.WrapWithMetrics(r, httpMetrics),
		manager:  nil,
	}, nil
}

// GetReporter returns Reporter
func (b ProxyBuilder) GetReporter() reporter.Reporter {
	return b.reporter
}

// GetManager returns manager
func (b ProxyBuilder) GetManager() configmanager.ClientConfigManager {
	return b.manager
}

// Close closes connections used by proxy.
func (b ProxyBuilder) Close() error {
	return b.reporter.(io.Closer).Close()
}
