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
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/jaegertracing/jaeger/thrift-gen/jaeger"
	"github.com/jaegertracing/jaeger/thrift-gen/zipkincore"
	"go.uber.org/zap"
)

// Reporter forwards received spans to central collector tier over Http
type Reporter struct {
	url    string
	client *http.Client
	logger *zap.Logger
}

// NewReporter creates new Http-based Reporter
func NewReporter(url string, timeout time.Duration, logger *zap.Logger) *Reporter {

	r := &Reporter{
		url:    url,
		client: &http.Client{Timeout: timeout},
		logger: logger,
	}

	return r
}

// EmitBatch implements EmitBatch() of Reporter
func (r *Reporter) EmitBatch(batch *jaeger.Batch) error {

	body, err := serializeThrift(batch)
	if err != nil {
		r.logger.Error("Could not serialize jaeger batch", zap.Error(err))
		return err
	}

	req, err := http.NewRequest("POST", r.url, body)

	if err != nil {
		r.logger.Error("Error in creating new http request", zap.Error(err))
		return err
	}

	req.Header.Set("Content-Type", "application/x-thrift")
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		r.logger.Error("Error in sending spans over http", zap.Error(err))
		return err
	}

	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		r.logger.Error(fmt.Sprintf("Error from collector - Response Code :  %d", resp.StatusCode))
		return fmt.Errorf("error from collector: %d", resp.StatusCode)
	}

	r.logger.Debug("Span batch submitted by the agent", zap.Int64("span-count", int64(len(batch.Spans))))
	return nil

}

func serializeThrift(obj thrift.TStruct) (*bytes.Buffer, error) {
	t := thrift.NewTMemoryBuffer()
	p := thrift.NewTBinaryProtocolTransport(t)
	if err := obj.Write(p); err != nil {
		return nil, err
	}
	return t.Buffer, nil
}

// EmitZipkinBatch implements EmitZipkinBatch() of Reporter
func (r *Reporter) EmitZipkinBatch(spans []*zipkincore.Span) error {

	// discuss this whether or not to have backward support for zipkin spans for http reporter
	r.logger.Info("Zipkin spans currently not supported with http reporter")
	return nil

}
