// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by x-pack/dev-tools/cmd/buildspec/buildspec.go - DO NOT EDIT.

package program

import (
	"strings"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/packer"
)

var Supported []Spec
var SupportedMap map[string]Spec

func init() {
	// Packed Files
	// spec/apm-server.yml
	// spec/endpoint.yml
	// spec/filebeat.yml
	// spec/fleet-server.yml
	// spec/heartbeat.yml
	// spec/metricbeat.yml
	// spec/osquerybeat.yml
	// spec/packetbeat.yml
	unpacked := packer.MustUnpack("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")
	SupportedMap = make(map[string]Spec)

	for f, v := range unpacked {
		s, err := NewSpecFromBytes(v)
		if err != nil {
			panic("Cannot read spec from " + f)
		}
		Supported = append(Supported, s)
		SupportedMap[strings.ToLower(s.Cmd)] = s
	}
}
