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
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/jaegertracing/jaeger/thrift-gen/jaeger"
)

func startHTTPServer() error {

	http.HandleFunc("/api/traces", func(w http.ResponseWriter, r *http.Request) {

		contentType := r.Header.Get("Content-Type")
		if contentType != "application/x-thrift" {
			// return error code
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		buffer := thrift.NewTMemoryBuffer()

		if _, err = buffer.Write(body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		transport := thrift.NewTBinaryProtocolTransport(buffer)
		batch := &jaeger.Batch{}

		if err = batch.Read(transport); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		// accepted error code is returned by default via WriteHeader

	})

	go func() { http.ListenAndServe(":9001", nil) }()

	return nil

}

var _ = startHTTPServer()

func TestJaegerHttpReporterSuccess(t *testing.T) {

	logger := zap.NewNop()
	r := NewReporter("http://127.0.0.1:9001/api/traces", 5*time.Second, logger)

	err := submitTestJaegerBatch(r)
	require.NoError(t, err)

}

func TestJaegerHttpReporterFailure(t *testing.T) {

	tests := []struct {
		endpoint string
		timeout  time.Duration
	}{
		{
			// invalid port case
			endpoint: "http://127.0.0.1:839939/api/traces",
			timeout:  2 * time.Second,
		},
		{
			// invalid endpoint
			endpoint: "tcp://127.0.0.1:80",
			timeout:  3 * time.Second,
		},
	}

	logger := zap.NewNop()

	for _, test := range tests {

		r := NewReporter(test.endpoint, test.timeout, logger)
		err := submitTestJaegerBatch(r)

		// require reporter to throw some error as these are all failure scenarios
		require.Error(t, err)
	}

}

func submitTestJaegerBatch(r *Reporter) error {

	batch := jaeger.NewBatch()
	batch.Process = jaeger.NewProcess()
	batch.Spans = []*jaeger.Span{{OperationName: "span1"}}

	return r.EmitBatch(batch)

}
