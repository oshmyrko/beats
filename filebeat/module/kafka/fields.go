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

package kafka

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "kafka", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJysk09uszAQxfec4in7cAAW3+bbteouFxiFgVgY27KHtLl9Zdz8c51ApM4OD37vx5thi4FPDQbqBqoAUaK5weY9Pm8qoOWw98qJsqbBvwoA5h5G206aK6BTrNvQzK0tDI18lYslJ8cNem8n93NS0LyXuZXStr+clcQeCqZKsNr20MpwqG+aueOdKx9Z33XO3qQVhazjSA4zaV26N6reU2ITP3HRb+QQqOcXHcu3HvkVjfd2dNawkaL1wKdP69us9yTuWP/PkpADz9GrEH2U6dF5O9ZlEE0h/8oEIfyV0y0QvNGRkuBLCOJpXx5BvnArGIBdlIMyF4a4gXX2YmkJl2LBwnxW0f1KSRLuk5ywuLB4NrSVWMBH0oYjL7Ddle4xTjfp/Lf7G5bdgWfxczz5NL8DAAD//1wZSZY="
}
