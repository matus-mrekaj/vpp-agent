//  Copyright (c) 2019 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package vpp2001_324

import (
	govppapi "git.fd.io/govpp.git/api"
	"github.com/ligato/cn-infra/logging"
	vpe_vppcalls "github.com/ligato/vpp-agent/plugins/govppmux/vppcalls"
	vpe_vpp2001_324 "github.com/ligato/vpp-agent/plugins/govppmux/vppcalls/vpp2001_324"
	vpp_sr "github.com/ligato/vpp-agent/plugins/vpp/binapi/vpp2001_324/sr"
	vpp_vpe "github.com/ligato/vpp-agent/plugins/vpp/binapi/vpp2001_324/vpe"
	"github.com/ligato/vpp-agent/plugins/vpp/ifplugin/ifaceidx"
	"github.com/ligato/vpp-agent/plugins/vpp/srplugin/vppcalls"
)

func init() {
	var msgs []govppapi.Message
	msgs = append(msgs, vpp_sr.AllMessages()...)
	msgs = append(msgs, vpp_vpe.AllMessages()...) // using also vpe -> need to have correct vpp version also for vpe

	vppcalls.Versions["vpp2001_324"] = vppcalls.HandlerVersion{
		Msgs: msgs,
		New: func(ch govppapi.Channel, ifIndexes ifaceidx.IfaceMetadataIndex, log logging.Logger) vppcalls.SRv6VppAPI {
			return NewSRv6VppHandler(ch, ifIndexes, log)
		},
	}
}

// SRv6VppHandler is accessor for SRv6-related vppcalls methods
type SRv6VppHandler struct {
	vpe_vppcalls.VpeVppAPI

	log          logging.Logger
	ifIndexes    ifaceidx.IfaceMetadataIndex
	callsChannel govppapi.Channel
}

// NewSRv6VppHandler creates new instance of SRv6 vppcalls handler
func NewSRv6VppHandler(vppChan govppapi.Channel, ifIndexes ifaceidx.IfaceMetadataIndex, log logging.Logger) *SRv6VppHandler {
	return &SRv6VppHandler{
		callsChannel: vppChan,
		ifIndexes:    ifIndexes,
		log:          log,
		VpeVppAPI:    vpe_vpp2001_324.NewVpeHandler(vppChan),
	}
}
