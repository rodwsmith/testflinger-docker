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
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "kafka", asset.ModuleFieldsPri, AssetKafka); err != nil {
		panic(err)
	}
}

// AssetKafka returns asset data.
// This is the base64 encoded zlib format compressed contents of module/kafka.
func AssetKafka() string {
	return "eJyskkFugzAQRfec4iv7cAAW3XTXqrtcYAQDsTAeZE/acvvK0ATi0oSo9Q5/9P7TePZoeSjQUt1SBqhRywV2r/F7lwEVh9KbXo24Ak8ZAIwZOqlOljOgNmyrUIzRHo46nnHx6NBzgcbLqf++WWFeY5YoK83lbg32K3A6k6yVBtY4DvkiTBuXraV0vTh2epWe+1sePsRXSXbDIp7nMxJ65NHIhNhjXIPaS5evi1gK4f8kXuidJuZDFnr0TGnVHzQOI2+EP+bhqeRVjXQlNkgAh4iDcReHuCN58uPamuDOC90bzya7H6+lk+6NOS29Og6BmnRas5nyZ7rdG7WAt4mNnrxC6tkuz74CAAD//xD+FUk="
}