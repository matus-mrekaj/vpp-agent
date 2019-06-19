// Code generated by GoVPP binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/core/session.api.json

/*
Package session is a generated from VPP binary API module 'session'.

 The session module consists of:
	 38 messages
	 19 services
*/
package session

import api "git.fd.io/govpp.git/api"
import bytes "bytes"
import context "context"
import strconv "strconv"
import struc "github.com/lunixbochs/struc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = strconv.Itoa
var _ = struc.Pack

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

const (
	// ModuleName is the name of this module.
	ModuleName = "session"
	// APIVersion is the API version of this module.
	APIVersion = "1.6.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x299f3131
)

/* Messages */

// AppCutThroughRegistrationAdd represents VPP binary API message 'app_cut_through_registration_add':
type AppCutThroughRegistrationAdd struct {
	EvtQAddress     uint64
	PeerEvtQAddress uint64
	WrkIndex        uint32
	NFds            uint8
	FdFlags         uint8
}

func (*AppCutThroughRegistrationAdd) GetMessageName() string {
	return "app_cut_through_registration_add"
}
func (*AppCutThroughRegistrationAdd) GetCrcString() string {
	return "6d73b1b9"
}
func (*AppCutThroughRegistrationAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AppCutThroughRegistrationAddReply represents VPP binary API message 'app_cut_through_registration_add_reply':
type AppCutThroughRegistrationAddReply struct {
	Retval int32
}

func (*AppCutThroughRegistrationAddReply) GetMessageName() string {
	return "app_cut_through_registration_add_reply"
}
func (*AppCutThroughRegistrationAddReply) GetCrcString() string {
	return "e8d4e804"
}
func (*AppCutThroughRegistrationAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// AppNamespaceAddDel represents VPP binary API message 'app_namespace_add_del':
type AppNamespaceAddDel struct {
	Secret         uint64
	SwIfIndex      uint32
	IP4FibID       uint32
	IP6FibID       uint32
	NamespaceIDLen uint8
	NamespaceID    []byte `struc:"[64]byte"`
}

func (*AppNamespaceAddDel) GetMessageName() string {
	return "app_namespace_add_del"
}
func (*AppNamespaceAddDel) GetCrcString() string {
	return "dd074c65"
}
func (*AppNamespaceAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AppNamespaceAddDelReply represents VPP binary API message 'app_namespace_add_del_reply':
type AppNamespaceAddDelReply struct {
	Retval     int32
	AppnsIndex uint32
}

func (*AppNamespaceAddDelReply) GetMessageName() string {
	return "app_namespace_add_del_reply"
}
func (*AppNamespaceAddDelReply) GetCrcString() string {
	return "85137120"
}
func (*AppNamespaceAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// AppWorkerAddDel represents VPP binary API message 'app_worker_add_del':
type AppWorkerAddDel struct {
	AppIndex uint32
	WrkIndex uint32
	IsAdd    uint8
}

func (*AppWorkerAddDel) GetMessageName() string {
	return "app_worker_add_del"
}
func (*AppWorkerAddDel) GetCrcString() string {
	return "6d2b2279"
}
func (*AppWorkerAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// AppWorkerAddDelReply represents VPP binary API message 'app_worker_add_del_reply':
type AppWorkerAddDelReply struct {
	Retval               int32
	WrkIndex             uint32
	AppEventQueueAddress uint64
	NFds                 uint8
	FdFlags              uint8
	SegmentNameLength    uint8
	SegmentName          []byte `struc:"[128]byte"`
	SegmentHandle        uint64
	IsAdd                uint8
}

func (*AppWorkerAddDelReply) GetMessageName() string {
	return "app_worker_add_del_reply"
}
func (*AppWorkerAddDelReply) GetCrcString() string {
	return "56b21abc"
}
func (*AppWorkerAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// ApplicationAttach represents VPP binary API message 'application_attach':
type ApplicationAttach struct {
	InitialSegmentSize uint32
	Options            []uint64 `struc:"[16]uint64"`
	NamespaceIDLen     uint8
	NamespaceID        []byte `struc:"[64]byte"`
}

func (*ApplicationAttach) GetMessageName() string {
	return "application_attach"
}
func (*ApplicationAttach) GetCrcString() string {
	return "81d4f974"
}
func (*ApplicationAttach) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// ApplicationAttachReply represents VPP binary API message 'application_attach_reply':
type ApplicationAttachReply struct {
	Retval               int32
	AppEventQueueAddress uint64
	NFds                 uint8
	FdFlags              uint8
	SegmentSize          uint32
	SegmentNameLength    uint8
	SegmentName          []byte `struc:"[128]byte"`
	AppIndex             uint32
	SegmentHandle        uint64
}

func (*ApplicationAttachReply) GetMessageName() string {
	return "application_attach_reply"
}
func (*ApplicationAttachReply) GetCrcString() string {
	return "581866e8"
}
func (*ApplicationAttachReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// ApplicationDetach represents VPP binary API message 'application_detach':
type ApplicationDetach struct{}

func (*ApplicationDetach) GetMessageName() string {
	return "application_detach"
}
func (*ApplicationDetach) GetCrcString() string {
	return "51077d14"
}
func (*ApplicationDetach) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// ApplicationDetachReply represents VPP binary API message 'application_detach_reply':
type ApplicationDetachReply struct {
	Retval int32
}

func (*ApplicationDetachReply) GetMessageName() string {
	return "application_detach_reply"
}
func (*ApplicationDetachReply) GetCrcString() string {
	return "e8d4e804"
}
func (*ApplicationDetachReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// ApplicationTLSCertAdd represents VPP binary API message 'application_tls_cert_add':
type ApplicationTLSCertAdd struct {
	AppIndex uint32
	CertLen  uint16 `struc:"sizeof=Cert"`
	Cert     []byte
}

func (*ApplicationTLSCertAdd) GetMessageName() string {
	return "application_tls_cert_add"
}
func (*ApplicationTLSCertAdd) GetCrcString() string {
	return "3f5cfe45"
}
func (*ApplicationTLSCertAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// ApplicationTLSCertAddReply represents VPP binary API message 'application_tls_cert_add_reply':
type ApplicationTLSCertAddReply struct {
	Retval int32
}

func (*ApplicationTLSCertAddReply) GetMessageName() string {
	return "application_tls_cert_add_reply"
}
func (*ApplicationTLSCertAddReply) GetCrcString() string {
	return "e8d4e804"
}
func (*ApplicationTLSCertAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// ApplicationTLSKeyAdd represents VPP binary API message 'application_tls_key_add':
type ApplicationTLSKeyAdd struct {
	AppIndex uint32
	KeyLen   uint16 `struc:"sizeof=Key"`
	Key      []byte
}

func (*ApplicationTLSKeyAdd) GetMessageName() string {
	return "application_tls_key_add"
}
func (*ApplicationTLSKeyAdd) GetCrcString() string {
	return "5eaf70cd"
}
func (*ApplicationTLSKeyAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// ApplicationTLSKeyAddReply represents VPP binary API message 'application_tls_key_add_reply':
type ApplicationTLSKeyAddReply struct {
	Retval int32
}

func (*ApplicationTLSKeyAddReply) GetMessageName() string {
	return "application_tls_key_add_reply"
}
func (*ApplicationTLSKeyAddReply) GetCrcString() string {
	return "e8d4e804"
}
func (*ApplicationTLSKeyAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BindSock represents VPP binary API message 'bind_sock':
type BindSock struct {
	WrkIndex uint32
	Vrf      uint32
	IsIP4    uint8
	IP       []byte `struc:"[16]byte"`
	Port     uint16
	Proto    uint8
	Options  []uint64 `struc:"[16]uint64"`
}

func (*BindSock) GetMessageName() string {
	return "bind_sock"
}
func (*BindSock) GetCrcString() string {
	return "0394633f"
}
func (*BindSock) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BindSockReply represents VPP binary API message 'bind_sock_reply':
type BindSockReply struct {
	Retval int32
}

func (*BindSockReply) GetMessageName() string {
	return "bind_sock_reply"
}
func (*BindSockReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BindSockReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BindURI represents VPP binary API message 'bind_uri':
type BindURI struct {
	AcceptCookie uint32
	URI          []byte `struc:"[128]byte"`
}

func (*BindURI) GetMessageName() string {
	return "bind_uri"
}
func (*BindURI) GetCrcString() string {
	return "fae140cb"
}
func (*BindURI) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BindURIReply represents VPP binary API message 'bind_uri_reply':
type BindURIReply struct {
	Retval int32
}

func (*BindURIReply) GetMessageName() string {
	return "bind_uri_reply"
}
func (*BindURIReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BindURIReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// ConnectSock represents VPP binary API message 'connect_sock':
type ConnectSock struct {
	WrkIndex           uint32
	ClientQueueAddress uint64
	Options            []uint64 `struc:"[16]uint64"`
	Vrf                uint32
	IsIP4              uint8
	IP                 []byte `struc:"[16]byte"`
	Port               uint16
	Proto              uint8
	HostnameLen        uint8 `struc:"sizeof=Hostname"`
	Hostname           []byte
}

func (*ConnectSock) GetMessageName() string {
	return "connect_sock"
}
func (*ConnectSock) GetCrcString() string {
	return "a916aa77"
}
func (*ConnectSock) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// ConnectSockReply represents VPP binary API message 'connect_sock_reply':
type ConnectSockReply struct {
	Retval int32
}

func (*ConnectSockReply) GetMessageName() string {
	return "connect_sock_reply"
}
func (*ConnectSockReply) GetCrcString() string {
	return "e8d4e804"
}
func (*ConnectSockReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// ConnectURI represents VPP binary API message 'connect_uri':
type ConnectURI struct {
	ClientQueueAddress uint64
	Options            []uint64 `struc:"[16]uint64"`
	URI                []byte   `struc:"[128]byte"`
}

func (*ConnectURI) GetMessageName() string {
	return "connect_uri"
}
func (*ConnectURI) GetCrcString() string {
	return "a36143d6"
}
func (*ConnectURI) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// ConnectURIReply represents VPP binary API message 'connect_uri_reply':
type ConnectURIReply struct {
	Retval int32
}

func (*ConnectURIReply) GetMessageName() string {
	return "connect_uri_reply"
}
func (*ConnectURIReply) GetCrcString() string {
	return "e8d4e804"
}
func (*ConnectURIReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DisconnectSession represents VPP binary API message 'disconnect_session':
type DisconnectSession struct {
	Handle uint64
}

func (*DisconnectSession) GetMessageName() string {
	return "disconnect_session"
}
func (*DisconnectSession) GetCrcString() string {
	return "7279205b"
}
func (*DisconnectSession) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DisconnectSessionReply represents VPP binary API message 'disconnect_session_reply':
type DisconnectSessionReply struct {
	Retval int32
	Handle uint64
}

func (*DisconnectSessionReply) GetMessageName() string {
	return "disconnect_session_reply"
}
func (*DisconnectSessionReply) GetCrcString() string {
	return "d6960a03"
}
func (*DisconnectSessionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MapAnotherSegment represents VPP binary API message 'map_another_segment':
type MapAnotherSegment struct {
	FdFlags       uint8
	SegmentSize   uint32
	SegmentName   []byte `struc:"[128]byte"`
	SegmentHandle uint64
}

func (*MapAnotherSegment) GetMessageName() string {
	return "map_another_segment"
}
func (*MapAnotherSegment) GetCrcString() string {
	return "dc2d630b"
}
func (*MapAnotherSegment) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MapAnotherSegmentReply represents VPP binary API message 'map_another_segment_reply':
type MapAnotherSegmentReply struct {
	Retval int32
}

func (*MapAnotherSegmentReply) GetMessageName() string {
	return "map_another_segment_reply"
}
func (*MapAnotherSegmentReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MapAnotherSegmentReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SessionEnableDisable represents VPP binary API message 'session_enable_disable':
type SessionEnableDisable struct {
	IsEnable uint8
}

func (*SessionEnableDisable) GetMessageName() string {
	return "session_enable_disable"
}
func (*SessionEnableDisable) GetCrcString() string {
	return "30ac9be7"
}
func (*SessionEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SessionEnableDisableReply represents VPP binary API message 'session_enable_disable_reply':
type SessionEnableDisableReply struct {
	Retval int32
}

func (*SessionEnableDisableReply) GetMessageName() string {
	return "session_enable_disable_reply"
}
func (*SessionEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SessionEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SessionRuleAddDel represents VPP binary API message 'session_rule_add_del':
type SessionRuleAddDel struct {
	TransportProto uint8
	IsIP4          uint8
	LclIP          []byte `struc:"[16]byte"`
	LclPlen        uint8
	RmtIP          []byte `struc:"[16]byte"`
	RmtPlen        uint8
	LclPort        uint16
	RmtPort        uint16
	ActionIndex    uint32
	IsAdd          uint8
	AppnsIndex     uint32
	Scope          uint8
	Tag            []byte `struc:"[64]byte"`
}

func (*SessionRuleAddDel) GetMessageName() string {
	return "session_rule_add_del"
}
func (*SessionRuleAddDel) GetCrcString() string {
	return "4ab2eb06"
}
func (*SessionRuleAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SessionRuleAddDelReply represents VPP binary API message 'session_rule_add_del_reply':
type SessionRuleAddDelReply struct {
	Retval int32
}

func (*SessionRuleAddDelReply) GetMessageName() string {
	return "session_rule_add_del_reply"
}
func (*SessionRuleAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SessionRuleAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SessionRulesDetails represents VPP binary API message 'session_rules_details':
type SessionRulesDetails struct {
	TransportProto uint8
	IsIP4          uint8
	LclIP          []byte `struc:"[16]byte"`
	LclPlen        uint8
	RmtIP          []byte `struc:"[16]byte"`
	RmtPlen        uint8
	LclPort        uint16
	RmtPort        uint16
	ActionIndex    uint32
	AppnsIndex     uint32
	Scope          uint8
	Tag            []byte `struc:"[64]byte"`
}

func (*SessionRulesDetails) GetMessageName() string {
	return "session_rules_details"
}
func (*SessionRulesDetails) GetCrcString() string {
	return "a52b0e96"
}
func (*SessionRulesDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SessionRulesDump represents VPP binary API message 'session_rules_dump':
type SessionRulesDump struct{}

func (*SessionRulesDump) GetMessageName() string {
	return "session_rules_dump"
}
func (*SessionRulesDump) GetCrcString() string {
	return "51077d14"
}
func (*SessionRulesDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// UnbindSock represents VPP binary API message 'unbind_sock':
type UnbindSock struct {
	WrkIndex uint32
	Handle   uint64
}

func (*UnbindSock) GetMessageName() string {
	return "unbind_sock"
}
func (*UnbindSock) GetCrcString() string {
	return "08880908"
}
func (*UnbindSock) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// UnbindSockReply represents VPP binary API message 'unbind_sock_reply':
type UnbindSockReply struct {
	Retval int32
}

func (*UnbindSockReply) GetMessageName() string {
	return "unbind_sock_reply"
}
func (*UnbindSockReply) GetCrcString() string {
	return "e8d4e804"
}
func (*UnbindSockReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// UnbindURI represents VPP binary API message 'unbind_uri':
type UnbindURI struct {
	URI []byte `struc:"[128]byte"`
}

func (*UnbindURI) GetMessageName() string {
	return "unbind_uri"
}
func (*UnbindURI) GetCrcString() string {
	return "294cf07d"
}
func (*UnbindURI) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// UnbindURIReply represents VPP binary API message 'unbind_uri_reply':
type UnbindURIReply struct {
	Retval int32
}

func (*UnbindURIReply) GetMessageName() string {
	return "unbind_uri_reply"
}
func (*UnbindURIReply) GetCrcString() string {
	return "e8d4e804"
}
func (*UnbindURIReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// UnmapSegment represents VPP binary API message 'unmap_segment':
type UnmapSegment struct {
	SegmentHandle uint64
}

func (*UnmapSegment) GetMessageName() string {
	return "unmap_segment"
}
func (*UnmapSegment) GetCrcString() string {
	return "f77096f6"
}
func (*UnmapSegment) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// UnmapSegmentReply represents VPP binary API message 'unmap_segment_reply':
type UnmapSegmentReply struct {
	Retval int32
}

func (*UnmapSegmentReply) GetMessageName() string {
	return "unmap_segment_reply"
}
func (*UnmapSegmentReply) GetCrcString() string {
	return "e8d4e804"
}
func (*UnmapSegmentReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*AppCutThroughRegistrationAdd)(nil), "session.AppCutThroughRegistrationAdd")
	api.RegisterMessage((*AppCutThroughRegistrationAddReply)(nil), "session.AppCutThroughRegistrationAddReply")
	api.RegisterMessage((*AppNamespaceAddDel)(nil), "session.AppNamespaceAddDel")
	api.RegisterMessage((*AppNamespaceAddDelReply)(nil), "session.AppNamespaceAddDelReply")
	api.RegisterMessage((*AppWorkerAddDel)(nil), "session.AppWorkerAddDel")
	api.RegisterMessage((*AppWorkerAddDelReply)(nil), "session.AppWorkerAddDelReply")
	api.RegisterMessage((*ApplicationAttach)(nil), "session.ApplicationAttach")
	api.RegisterMessage((*ApplicationAttachReply)(nil), "session.ApplicationAttachReply")
	api.RegisterMessage((*ApplicationDetach)(nil), "session.ApplicationDetach")
	api.RegisterMessage((*ApplicationDetachReply)(nil), "session.ApplicationDetachReply")
	api.RegisterMessage((*ApplicationTLSCertAdd)(nil), "session.ApplicationTLSCertAdd")
	api.RegisterMessage((*ApplicationTLSCertAddReply)(nil), "session.ApplicationTLSCertAddReply")
	api.RegisterMessage((*ApplicationTLSKeyAdd)(nil), "session.ApplicationTLSKeyAdd")
	api.RegisterMessage((*ApplicationTLSKeyAddReply)(nil), "session.ApplicationTLSKeyAddReply")
	api.RegisterMessage((*BindSock)(nil), "session.BindSock")
	api.RegisterMessage((*BindSockReply)(nil), "session.BindSockReply")
	api.RegisterMessage((*BindURI)(nil), "session.BindURI")
	api.RegisterMessage((*BindURIReply)(nil), "session.BindURIReply")
	api.RegisterMessage((*ConnectSock)(nil), "session.ConnectSock")
	api.RegisterMessage((*ConnectSockReply)(nil), "session.ConnectSockReply")
	api.RegisterMessage((*ConnectURI)(nil), "session.ConnectURI")
	api.RegisterMessage((*ConnectURIReply)(nil), "session.ConnectURIReply")
	api.RegisterMessage((*DisconnectSession)(nil), "session.DisconnectSession")
	api.RegisterMessage((*DisconnectSessionReply)(nil), "session.DisconnectSessionReply")
	api.RegisterMessage((*MapAnotherSegment)(nil), "session.MapAnotherSegment")
	api.RegisterMessage((*MapAnotherSegmentReply)(nil), "session.MapAnotherSegmentReply")
	api.RegisterMessage((*SessionEnableDisable)(nil), "session.SessionEnableDisable")
	api.RegisterMessage((*SessionEnableDisableReply)(nil), "session.SessionEnableDisableReply")
	api.RegisterMessage((*SessionRuleAddDel)(nil), "session.SessionRuleAddDel")
	api.RegisterMessage((*SessionRuleAddDelReply)(nil), "session.SessionRuleAddDelReply")
	api.RegisterMessage((*SessionRulesDetails)(nil), "session.SessionRulesDetails")
	api.RegisterMessage((*SessionRulesDump)(nil), "session.SessionRulesDump")
	api.RegisterMessage((*UnbindSock)(nil), "session.UnbindSock")
	api.RegisterMessage((*UnbindSockReply)(nil), "session.UnbindSockReply")
	api.RegisterMessage((*UnbindURI)(nil), "session.UnbindURI")
	api.RegisterMessage((*UnbindURIReply)(nil), "session.UnbindURIReply")
	api.RegisterMessage((*UnmapSegment)(nil), "session.UnmapSegment")
	api.RegisterMessage((*UnmapSegmentReply)(nil), "session.UnmapSegmentReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*AppCutThroughRegistrationAdd)(nil),
		(*AppCutThroughRegistrationAddReply)(nil),
		(*AppNamespaceAddDel)(nil),
		(*AppNamespaceAddDelReply)(nil),
		(*AppWorkerAddDel)(nil),
		(*AppWorkerAddDelReply)(nil),
		(*ApplicationAttach)(nil),
		(*ApplicationAttachReply)(nil),
		(*ApplicationDetach)(nil),
		(*ApplicationDetachReply)(nil),
		(*ApplicationTLSCertAdd)(nil),
		(*ApplicationTLSCertAddReply)(nil),
		(*ApplicationTLSKeyAdd)(nil),
		(*ApplicationTLSKeyAddReply)(nil),
		(*BindSock)(nil),
		(*BindSockReply)(nil),
		(*BindURI)(nil),
		(*BindURIReply)(nil),
		(*ConnectSock)(nil),
		(*ConnectSockReply)(nil),
		(*ConnectURI)(nil),
		(*ConnectURIReply)(nil),
		(*DisconnectSession)(nil),
		(*DisconnectSessionReply)(nil),
		(*MapAnotherSegment)(nil),
		(*MapAnotherSegmentReply)(nil),
		(*SessionEnableDisable)(nil),
		(*SessionEnableDisableReply)(nil),
		(*SessionRuleAddDel)(nil),
		(*SessionRuleAddDelReply)(nil),
		(*SessionRulesDetails)(nil),
		(*SessionRulesDump)(nil),
		(*UnbindSock)(nil),
		(*UnbindSockReply)(nil),
		(*UnbindURI)(nil),
		(*UnbindURIReply)(nil),
		(*UnmapSegment)(nil),
		(*UnmapSegmentReply)(nil),
	}
}
