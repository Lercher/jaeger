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
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBindFlags(t *testing.T) {

	tests := []struct {
		flags   []string
		builder ConnBuilder
	}{
		{flags: []string{
			"--reporter.http.url=http://1.2.3.4:555",
		},
			// default timeout case
			builder: ConnBuilder{CollectorURLs: []string{"http://1.2.3.4:555"}, CollectorResponseTimeout: defaultCollectorResponseTimeout},
		},
		{
			flags: []string{
				"--reporter.http.url=http://1.2.3.4:666",
				"--reporter.http.response.timeout=15s",
			},
			builder: ConnBuilder{CollectorURLs: []string{"http://1.2.3.4:666"}, CollectorResponseTimeout: 15 * time.Second},
		},
	}

	for _, test := range tests {
		// Reset flags every iteration
		v := viper.New()
		command := cobra.Command{}

		flags := &flag.FlagSet{}
		AddFlags(flags)
		command.ResetFlags()
		command.PersistentFlags().AddGoFlagSet(flags)

		v.BindPFlags(command.PersistentFlags())

		err := command.ParseFlags(test.flags)
		require.NoError(t, err)
		b := ConnBuilder{}
		b.InitFromViper(v)
		assert.Equal(t, test.builder, b)
	}
}
