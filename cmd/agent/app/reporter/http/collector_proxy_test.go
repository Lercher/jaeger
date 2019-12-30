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
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/cmd/agent/app/reporter"
)

func TestCollectorProxy(t *testing.T) {

	cfg := &Builder{CollectorEndpoint: "http://localhost:14268/api/traces"}
	mFactory := metrics.NullFactory
	logger := zap.NewNop()

	proxy, err := NewCollectorProxy(cfg, mFactory, logger)
	require.NoError(t, err)
	assert.NotNil(t, proxy)
	assert.NotNil(t, proxy.GetReporter())
	// assert.NotNil(t, proxy.GetManager())

	r, _ := cfg.CreateReporter(logger)
	assert.Equal(t, reporter.WrapWithMetrics(r, mFactory), proxy.GetReporter())

}
