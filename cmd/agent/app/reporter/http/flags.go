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
	"strings"
	"time"

	"github.com/spf13/viper"
)

const (
	httpPrefix                      = "reporter.http"
	collectorURL                    = httpPrefix + ".url"
	collectorResponseTimeout        = httpPrefix + ".response.timeout"
	defaultCollectorResponseTimeout = 3 * time.Second
)

// AddFlags adds flags for Builder.
func AddFlags(flags *flag.FlagSet) {
	flags.String(collectorURL, "", "url of a static collector to connect to directly (N.B.: standard is http://server:14268/api/traces), if path is missing /api/traces is appended")
	flags.Duration(collectorResponseTimeout, defaultCollectorResponseTimeout, "sets the timeout for http response from collector")
}

// InitFromViper initializes Builder with properties retrieved from Viper.
func (b *ConnBuilder) InitFromViper(v *viper.Viper) *ConnBuilder {
	urls := v.GetString(collectorURL)
	if urls != "" {
		b.CollectorURLs = strings.Split(urls, ",")
	}
	b.CollectorResponseTimeout = v.GetDuration(collectorResponseTimeout)

	return b
}
