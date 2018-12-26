// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package dhcpv4

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "dhcpv4", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
