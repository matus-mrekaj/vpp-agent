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

package vpp2001_379

import (
	ifs "github.com/ligato/vpp-agent/api/models/vpp/interfaces"
	vpp_afpacket "github.com/ligato/vpp-agent/plugins/vpp/binapi/vpp2001_379/af_packet"
)

// AddAfPacketInterface implements AfPacket handler.
func (h *InterfaceVppHandler) AddAfPacketInterface(ifName string, hwAddr string, afPacketIntf *ifs.AfpacketLink) (swIndex uint32, err error) {
	req := &vpp_afpacket.AfPacketCreate{
		HostIfName: afPacketIntf.HostIfName,
	}
	if hwAddr == "" {
		req.UseRandomHwAddr = true
	} else {
		mac, err := ParseMAC(hwAddr)
		if err != nil {
			return 0, err
		}
		req.HwAddr = mac
	}
	reply := &vpp_afpacket.AfPacketCreateReply{}

	if err = h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return 0, err
	}

	return uint32(reply.SwIfIndex), h.SetInterfaceTag(ifName, uint32(reply.SwIfIndex))
}

// DeleteAfPacketInterface implements AfPacket handler.
func (h *InterfaceVppHandler) DeleteAfPacketInterface(ifName string, idx uint32, afPacketIntf *ifs.AfpacketLink) error {
	req := &vpp_afpacket.AfPacketDelete{
		HostIfName: afPacketIntf.HostIfName,
	}
	reply := &vpp_afpacket.AfPacketDeleteReply{}

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}

	return h.RemoveInterfaceTag(ifName, idx)
}
