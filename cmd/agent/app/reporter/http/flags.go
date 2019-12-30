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
	"flag"
	"time"

	"github.com/spf13/viper"
)

const (
	httpPrefix                      = "reporter.http."
	collectorEndPoint               = httpPrefix + "endpoint"
	collectorResponseTimeout        = httpPrefix + "response.timeout"
	defaultCollectorResponseTimeout = 1 * time.Second
)

// AddFlags adds flags for Builder.
func AddFlags(flags *flag.FlagSet) {

	flags.String(collectorEndPoint, "", "http endpoint of collector to connect to")
	flags.Duration(collectorResponseTimeout, defaultCollectorResponseTimeout, "sets the timeout for http response from collector")

}

// InitFromViper initializes Builder with properties retrieved from Viper.
func (b *Builder) InitFromViper(v *viper.Viper) *Builder {

	b.CollectorEndpoint = v.GetString(collectorEndPoint)
	b.CollectorResponseTimeout = v.GetDuration(collectorResponseTimeout)

	return b
}
