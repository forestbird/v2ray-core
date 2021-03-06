package policy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Second struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Second) Reset()                    { *m = Second{} }
func (m *Second) String() string            { return proto.CompactTextString(m) }
func (*Second) ProtoMessage()               {}
func (*Second) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Second) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Policy struct {
	Timeout *Policy_Timeout `protobuf:"bytes,1,opt,name=timeout" json:"timeout,omitempty"`
	Stats   *Policy_Stats   `protobuf:"bytes,2,opt,name=stats" json:"stats,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Policy) GetTimeout() *Policy_Timeout {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *Policy) GetStats() *Policy_Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// Timeout is a message for timeout settings in various stages, in seconds.
type Policy_Timeout struct {
	Handshake      *Second `protobuf:"bytes,1,opt,name=handshake" json:"handshake,omitempty"`
	ConnectionIdle *Second `protobuf:"bytes,2,opt,name=connection_idle,json=connectionIdle" json:"connection_idle,omitempty"`
	UplinkOnly     *Second `protobuf:"bytes,3,opt,name=uplink_only,json=uplinkOnly" json:"uplink_only,omitempty"`
	DownlinkOnly   *Second `protobuf:"bytes,4,opt,name=downlink_only,json=downlinkOnly" json:"downlink_only,omitempty"`
}

func (m *Policy_Timeout) Reset()                    { *m = Policy_Timeout{} }
func (m *Policy_Timeout) String() string            { return proto.CompactTextString(m) }
func (*Policy_Timeout) ProtoMessage()               {}
func (*Policy_Timeout) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Policy_Timeout) GetHandshake() *Second {
	if m != nil {
		return m.Handshake
	}
	return nil
}

func (m *Policy_Timeout) GetConnectionIdle() *Second {
	if m != nil {
		return m.ConnectionIdle
	}
	return nil
}

func (m *Policy_Timeout) GetUplinkOnly() *Second {
	if m != nil {
		return m.UplinkOnly
	}
	return nil
}

func (m *Policy_Timeout) GetDownlinkOnly() *Second {
	if m != nil {
		return m.DownlinkOnly
	}
	return nil
}

type Policy_Stats struct {
	UserUplink   bool `protobuf:"varint,1,opt,name=user_uplink,json=userUplink" json:"user_uplink,omitempty"`
	UserDownlink bool `protobuf:"varint,2,opt,name=user_downlink,json=userDownlink" json:"user_downlink,omitempty"`
}

func (m *Policy_Stats) Reset()                    { *m = Policy_Stats{} }
func (m *Policy_Stats) String() string            { return proto.CompactTextString(m) }
func (*Policy_Stats) ProtoMessage()               {}
func (*Policy_Stats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

func (m *Policy_Stats) GetUserUplink() bool {
	if m != nil {
		return m.UserUplink
	}
	return false
}

func (m *Policy_Stats) GetUserDownlink() bool {
	if m != nil {
		return m.UserDownlink
	}
	return false
}

type SystemPolicy struct {
	Stats *SystemPolicy_Stats `protobuf:"bytes,1,opt,name=stats" json:"stats,omitempty"`
}

func (m *SystemPolicy) Reset()                    { *m = SystemPolicy{} }
func (m *SystemPolicy) String() string            { return proto.CompactTextString(m) }
func (*SystemPolicy) ProtoMessage()               {}
func (*SystemPolicy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SystemPolicy) GetStats() *SystemPolicy_Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type SystemPolicy_Stats struct {
	InboundUplink   bool `protobuf:"varint,1,opt,name=inbound_uplink,json=inboundUplink" json:"inbound_uplink,omitempty"`
	InboundDownlink bool `protobuf:"varint,2,opt,name=inbound_downlink,json=inboundDownlink" json:"inbound_downlink,omitempty"`
}

func (m *SystemPolicy_Stats) Reset()                    { *m = SystemPolicy_Stats{} }
func (m *SystemPolicy_Stats) String() string            { return proto.CompactTextString(m) }
func (*SystemPolicy_Stats) ProtoMessage()               {}
func (*SystemPolicy_Stats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *SystemPolicy_Stats) GetInboundUplink() bool {
	if m != nil {
		return m.InboundUplink
	}
	return false
}

func (m *SystemPolicy_Stats) GetInboundDownlink() bool {
	if m != nil {
		return m.InboundDownlink
	}
	return false
}

type Config struct {
	Level  map[uint32]*Policy `protobuf:"bytes,1,rep,name=level" json:"level,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	System *SystemPolicy      `protobuf:"bytes,2,opt,name=system" json:"system,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Config) GetLevel() map[uint32]*Policy {
	if m != nil {
		return m.Level
	}
	return nil
}

func (m *Config) GetSystem() *SystemPolicy {
	if m != nil {
		return m.System
	}
	return nil
}

func init() {
	proto.RegisterType((*Second)(nil), "v2ray.core.app.policy.Second")
	proto.RegisterType((*Policy)(nil), "v2ray.core.app.policy.Policy")
	proto.RegisterType((*Policy_Timeout)(nil), "v2ray.core.app.policy.Policy.Timeout")
	proto.RegisterType((*Policy_Stats)(nil), "v2ray.core.app.policy.Policy.Stats")
	proto.RegisterType((*SystemPolicy)(nil), "v2ray.core.app.policy.SystemPolicy")
	proto.RegisterType((*SystemPolicy_Stats)(nil), "v2ray.core.app.policy.SystemPolicy.Stats")
	proto.RegisterType((*Config)(nil), "v2ray.core.app.policy.Config")
}

func init() { proto.RegisterFile("v2ray.com/core/app/policy/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x18, 0xc5, 0x95, 0x96, 0x66, 0xe3, 0x6b, 0xbb, 0x4d, 0x16, 0x93, 0x4a, 0x25, 0x60, 0xea, 0x34,
	0xd4, 0xdd, 0xb8, 0x52, 0x76, 0x03, 0x4c, 0x0c, 0x31, 0xfe, 0x48, 0x48, 0x20, 0x26, 0x97, 0x3f,
	0x82, 0x9b, 0xca, 0x73, 0x0c, 0x8b, 0xea, 0xd8, 0x56, 0xe2, 0x14, 0xe5, 0x35, 0x78, 0x8c, 0x3d,
	0x14, 0xd7, 0x3c, 0x06, 0x8a, 0xed, 0x2c, 0x1b, 0x5a, 0xb7, 0xde, 0x25, 0x47, 0xbf, 0x73, 0x74,
	0x3e, 0xdb, 0x1f, 0x3c, 0x5e, 0x44, 0x19, 0x2d, 0x31, 0x53, 0xe9, 0x84, 0xa9, 0x8c, 0x4f, 0xa8,
	0xd6, 0x13, 0xad, 0x44, 0xc2, 0xca, 0x09, 0x53, 0xf2, 0x47, 0xf2, 0x13, 0xeb, 0x4c, 0x19, 0x85,
	0xb6, 0x6b, 0x2e, 0xe3, 0x98, 0x6a, 0x8d, 0x1d, 0x33, 0x7a, 0x08, 0xe1, 0x94, 0x33, 0x25, 0x63,
	0x74, 0x0f, 0x3a, 0x0b, 0x2a, 0x0a, 0x3e, 0x08, 0x76, 0x82, 0x71, 0x9f, 0xb8, 0x9f, 0xd1, 0xdf,
	0x36, 0x84, 0x27, 0x16, 0x45, 0x2f, 0x60, 0xcd, 0x24, 0x29, 0x57, 0x85, 0xb1, 0x48, 0x37, 0xda,
	0xc3, 0xd7, 0x66, 0x62, 0xc7, 0xe3, 0x4f, 0x0e, 0x26, 0xb5, 0x0b, 0x3d, 0x85, 0x4e, 0x6e, 0xa8,
	0xc9, 0x07, 0x2d, 0x6b, 0xdf, 0xbd, 0xd9, 0x3e, 0xad, 0x50, 0xe2, 0x1c, 0xc3, 0xdf, 0x2d, 0x58,
	0xf3, 0x79, 0xe8, 0x10, 0xee, 0x9e, 0x51, 0x19, 0xe7, 0x67, 0x74, 0xce, 0x7d, 0x93, 0x07, 0x4b,
	0xa2, 0xdc, 0x68, 0xa4, 0xe1, 0xd1, 0x5b, 0xd8, 0x64, 0x4a, 0x4a, 0xce, 0x4c, 0xa2, 0xe4, 0x2c,
	0x89, 0x05, 0xf7, 0x6d, 0x6e, 0x89, 0xd8, 0x68, 0x5c, 0xef, 0x62, 0xc1, 0xd1, 0x11, 0x74, 0x0b,
	0x2d, 0x12, 0x39, 0x9f, 0x29, 0x29, 0xca, 0x41, 0x7b, 0x95, 0x0c, 0x70, 0x8e, 0x8f, 0x52, 0x94,
	0xe8, 0x18, 0xfa, 0xb1, 0xfa, 0x25, 0x9b, 0x84, 0x3b, 0xab, 0x24, 0xf4, 0x6a, 0x4f, 0x95, 0x31,
	0xfc, 0x00, 0x1d, 0x7b, 0x48, 0xe8, 0x11, 0x74, 0x8b, 0x9c, 0x67, 0x33, 0x97, 0x6f, 0xcf, 0x64,
	0x9d, 0x40, 0x25, 0x7d, 0xb6, 0x0a, 0xda, 0x85, 0xbe, 0x05, 0x6a, 0xbb, 0x9d, 0x79, 0x9d, 0xf4,
	0x2a, 0xf1, 0xb5, 0xd7, 0x46, 0xe7, 0x01, 0xf4, 0xa6, 0x65, 0x6e, 0x78, 0x7a, 0x71, 0xe1, 0xfe,
	0xbe, 0xdc, 0x21, 0xef, 0x2f, 0xeb, 0x76, 0xc9, 0x73, 0xf5, 0xd6, 0xbe, 0xd5, 0x05, 0xf7, 0x60,
	0x23, 0x91, 0xa7, 0xaa, 0x90, 0xf1, 0xd5, 0x8e, 0x7d, 0xaf, 0xfa, 0x9a, 0xfb, 0xb0, 0x55, 0x63,
	0xff, 0x35, 0xdd, 0xf4, 0xfa, 0x45, 0xd9, 0x3f, 0x01, 0x84, 0xaf, 0xec, 0xfb, 0x46, 0x47, 0xd0,
	0x11, 0x7c, 0xc1, 0xc5, 0x20, 0xd8, 0x69, 0x8f, 0xbb, 0xd1, 0x78, 0x49, 0x4d, 0x47, 0xe3, 0xf7,
	0x15, 0xfa, 0x46, 0x9a, 0xac, 0x24, 0xce, 0x86, 0x0e, 0x21, 0xcc, 0xed, 0x08, 0xb7, 0xbc, 0xcb,
	0xcb, 0x73, 0x12, 0x6f, 0x19, 0x7e, 0x05, 0x68, 0x12, 0xd1, 0x16, 0xb4, 0xe7, 0xbc, 0xf4, 0x1b,
	0x54, 0x7d, 0xa2, 0x83, 0x7a, 0xab, 0x6e, 0x7e, 0x65, 0x3e, 0xd5, 0xb1, 0xcf, 0x5a, 0x4f, 0x82,
	0xe3, 0xe7, 0x70, 0x9f, 0xa9, 0xf4, 0x7a, 0xfc, 0x24, 0xf8, 0x1e, 0xba, 0xaf, 0xf3, 0xd6, 0xf6,
	0x97, 0x88, 0xd0, 0x6a, 0xba, 0x8c, 0xe3, 0x97, 0x5a, 0xfb, 0xa4, 0xd3, 0xd0, 0x6e, 0xfd, 0xc1,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x9f, 0x00, 0x31, 0x1f, 0x04, 0x00, 0x00,
}
