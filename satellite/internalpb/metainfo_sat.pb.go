// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metainfo_sat.proto

package internalpb

import (
	fmt "fmt"
	math "math"
	time "time"

	proto "github.com/gogo/protobuf/proto"

	pb "storj.io/common/pb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type StreamID struct {
	Bucket               []byte               `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	EncryptedPath        []byte               `protobuf:"bytes,2,opt,name=encrypted_path,json=encryptedPath,proto3" json:"encrypted_path,omitempty"`
	Version              int32                `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Redundancy           *pb.RedundancyScheme `protobuf:"bytes,4,opt,name=redundancy,proto3" json:"redundancy,omitempty"`
	CreationDate         time.Time            `protobuf:"bytes,5,opt,name=creation_date,json=creationDate,proto3,stdtime" json:"creation_date"`
	ExpirationDate       time.Time            `protobuf:"bytes,6,opt,name=expiration_date,json=expirationDate,proto3,stdtime" json:"expiration_date"`
	SatelliteSignature   []byte               `protobuf:"bytes,9,opt,name=satellite_signature,json=satelliteSignature,proto3" json:"satellite_signature,omitempty"`
	StreamId             []byte               `protobuf:"bytes,10,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StreamID) Reset()         { *m = StreamID{} }
func (m *StreamID) String() string { return proto.CompactTextString(m) }
func (*StreamID) ProtoMessage()    {}
func (*StreamID) Descriptor() ([]byte, []int) {
	return fileDescriptor_47c60bd892d94aaf, []int{0}
}
func (m *StreamID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamID.Unmarshal(m, b)
}
func (m *StreamID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamID.Marshal(b, m, deterministic)
}
func (m *StreamID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamID.Merge(m, src)
}
func (m *StreamID) XXX_Size() int {
	return xxx_messageInfo_StreamID.Size(m)
}
func (m *StreamID) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamID.DiscardUnknown(m)
}

var xxx_messageInfo_StreamID proto.InternalMessageInfo

func (m *StreamID) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *StreamID) GetEncryptedPath() []byte {
	if m != nil {
		return m.EncryptedPath
	}
	return nil
}

func (m *StreamID) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *StreamID) GetRedundancy() *pb.RedundancyScheme {
	if m != nil {
		return m.Redundancy
	}
	return nil
}

func (m *StreamID) GetCreationDate() time.Time {
	if m != nil {
		return m.CreationDate
	}
	return time.Time{}
}

func (m *StreamID) GetExpirationDate() time.Time {
	if m != nil {
		return m.ExpirationDate
	}
	return time.Time{}
}

func (m *StreamID) GetSatelliteSignature() []byte {
	if m != nil {
		return m.SatelliteSignature
	}
	return nil
}

func (m *StreamID) GetStreamId() []byte {
	if m != nil {
		return m.StreamId
	}
	return nil
}

type SegmentID struct {
	StreamId             *StreamID                 `protobuf:"bytes,1,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	PartNumber           int32                     `protobuf:"varint,2,opt,name=part_number,json=partNumber,proto3" json:"part_number,omitempty"`
	Index                int32                     `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	RootPieceId          PieceID                   `protobuf:"bytes,5,opt,name=root_piece_id,json=rootPieceId,proto3,customtype=PieceID" json:"root_piece_id"`
	OriginalOrderLimits  []*pb.AddressedOrderLimit `protobuf:"bytes,6,rep,name=original_order_limits,json=originalOrderLimits,proto3" json:"original_order_limits,omitempty"`
	CreationDate         time.Time                 `protobuf:"bytes,7,opt,name=creation_date,json=creationDate,proto3,stdtime" json:"creation_date"`
	SatelliteSignature   []byte                    `protobuf:"bytes,8,opt,name=satellite_signature,json=satelliteSignature,proto3" json:"satellite_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *SegmentID) Reset()         { *m = SegmentID{} }
func (m *SegmentID) String() string { return proto.CompactTextString(m) }
func (*SegmentID) ProtoMessage()    {}
func (*SegmentID) Descriptor() ([]byte, []int) {
	return fileDescriptor_47c60bd892d94aaf, []int{1}
}
func (m *SegmentID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentID.Unmarshal(m, b)
}
func (m *SegmentID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentID.Marshal(b, m, deterministic)
}
func (m *SegmentID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentID.Merge(m, src)
}
func (m *SegmentID) XXX_Size() int {
	return xxx_messageInfo_SegmentID.Size(m)
}
func (m *SegmentID) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentID.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentID proto.InternalMessageInfo

func (m *SegmentID) GetStreamId() *StreamID {
	if m != nil {
		return m.StreamId
	}
	return nil
}

func (m *SegmentID) GetPartNumber() int32 {
	if m != nil {
		return m.PartNumber
	}
	return 0
}

func (m *SegmentID) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *SegmentID) GetOriginalOrderLimits() []*pb.AddressedOrderLimit {
	if m != nil {
		return m.OriginalOrderLimits
	}
	return nil
}

func (m *SegmentID) GetCreationDate() time.Time {
	if m != nil {
		return m.CreationDate
	}
	return time.Time{}
}

func (m *SegmentID) GetSatelliteSignature() []byte {
	if m != nil {
		return m.SatelliteSignature
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamID)(nil), "satellite.metainfo.StreamID")
	proto.RegisterType((*SegmentID)(nil), "satellite.metainfo.SegmentID")
}

func init() { proto.RegisterFile("metainfo_sat.proto", fileDescriptor_47c60bd892d94aaf) }

var fileDescriptor_47c60bd892d94aaf = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0xfc, 0x85, 0xfc, 0x4d, 0xfe, 0xa4, 0x29, 0x20, 0x2b, 0x05, 0x25, 0x2a, 0xaa, 0x94,
	0x95, 0x2d, 0xb5, 0x2b, 0xc4, 0x8a, 0x28, 0x1b, 0x4b, 0xfc, 0x14, 0x87, 0x15, 0x1b, 0x6b, 0xec,
	0xb9, 0x75, 0x06, 0xec, 0x19, 0x6b, 0xe6, 0x1a, 0xb5, 0x4b, 0x76, 0x2c, 0x79, 0x2c, 0x9e, 0x81,
	0x45, 0x79, 0x15, 0xe4, 0x71, 0x1c, 0x47, 0xaa, 0xba, 0x80, 0x9d, 0xcf, 0xb9, 0xe7, 0x1e, 0xdf,
	0x7b, 0xcf, 0x10, 0x9a, 0x03, 0x32, 0x21, 0xaf, 0x55, 0x64, 0x18, 0x7a, 0x85, 0x56, 0xa8, 0x28,
	0x35, 0x0c, 0x21, 0xcb, 0x04, 0x82, 0xd7, 0x54, 0xe7, 0x24, 0x55, 0xa9, 0xaa, 0xeb, 0xf3, 0x45,
	0xaa, 0x54, 0x9a, 0x81, 0x6f, 0x51, 0x5c, 0x5e, 0xfb, 0x28, 0x72, 0x30, 0xc8, 0xf2, 0x62, 0x2f,
	0x98, 0x15, 0x4a, 0x48, 0x04, 0xcd, 0xe3, 0x3d, 0x31, 0x6d, 0x7c, 0x6a, 0x7c, 0xf6, 0xbd, 0x43,
	0x06, 0x5b, 0xd4, 0xc0, 0xf2, 0x60, 0x43, 0x9f, 0x92, 0x5e, 0x5c, 0x26, 0x5f, 0x00, 0x5d, 0x67,
	0xe9, 0xac, 0xc6, 0xe1, 0x1e, 0xd1, 0x73, 0x32, 0x05, 0x99, 0xe8, 0xdb, 0x02, 0x81, 0x47, 0x05,
	0xc3, 0x9d, 0xfb, 0xbf, 0xad, 0x4f, 0x0e, 0xec, 0x15, 0xc3, 0x1d, 0x75, 0x49, 0xff, 0x2b, 0x68,
	0x23, 0x94, 0x74, 0x3b, 0x4b, 0x67, 0xd5, 0x0d, 0x1b, 0x48, 0x5f, 0x11, 0xa2, 0x81, 0x97, 0x92,
	0x33, 0x99, 0xdc, 0xba, 0x8f, 0x96, 0xce, 0x6a, 0x74, 0x71, 0xea, 0xb5, 0xb3, 0x85, 0x87, 0xe2,
	0x36, 0xd9, 0x41, 0x0e, 0xe1, 0x91, 0x9c, 0x06, 0x64, 0x92, 0x68, 0x60, 0x28, 0x94, 0x8c, 0x38,
	0x43, 0x70, 0xbb, 0xb6, 0x7f, 0xee, 0xd5, 0xcb, 0x7b, 0xcd, 0xf2, 0xde, 0xc7, 0x66, 0xf9, 0xf5,
	0xe0, 0xe7, 0xdd, 0xe2, 0xbf, 0x1f, 0xbf, 0x17, 0x4e, 0x38, 0x6e, 0x5a, 0x37, 0x0c, 0x81, 0xbe,
	0x25, 0x33, 0xb8, 0x29, 0x84, 0x3e, 0x32, 0xeb, 0xfd, 0x85, 0xd9, 0xb4, 0x6d, 0xb6, 0x76, 0x3e,
	0x39, 0x39, 0x04, 0x14, 0x19, 0x91, 0x4a, 0x86, 0xa5, 0x06, 0x77, 0x68, 0x8f, 0xd3, 0x66, 0xb7,
	0x6d, 0x2a, 0xf4, 0x94, 0x0c, 0x8d, 0x3d, 0x76, 0x24, 0xb8, 0x4b, 0xac, 0x6c, 0x50, 0x13, 0x01,
	0x3f, 0xfb, 0xd6, 0x21, 0xc3, 0x2d, 0xa4, 0x39, 0x48, 0x0c, 0x36, 0xf4, 0xe5, 0xb1, 0xd4, 0xb1,
	0x43, 0x3e, 0xf3, 0xee, 0x3f, 0x07, 0xaf, 0x09, 0xaf, 0x35, 0xa2, 0x0b, 0x32, 0x2a, 0x98, 0xc6,
	0x48, 0x96, 0x79, 0x0c, 0xda, 0x66, 0xd5, 0x0d, 0x49, 0x45, 0xbd, 0xb3, 0x0c, 0x7d, 0x4c, 0xba,
	0x42, 0x72, 0xb8, 0xd9, 0xc7, 0x54, 0x03, 0x7a, 0x49, 0x26, 0x5a, 0x29, 0x8c, 0x0a, 0x01, 0x09,
	0x54, 0x7f, 0xad, 0xee, 0x3c, 0x5e, 0xcf, 0xaa, 0xf5, 0x7f, 0xdd, 0x2d, 0xfa, 0x57, 0x15, 0x1f,
	0x6c, 0xc2, 0x51, 0xa5, 0xaa, 0x01, 0xa7, 0x1f, 0xc8, 0x13, 0xa5, 0x45, 0x2a, 0x24, 0xcb, 0x22,
	0xa5, 0x39, 0xe8, 0x28, 0x13, 0xb9, 0x40, 0xe3, 0xf6, 0x96, 0x9d, 0xd5, 0xe8, 0xe2, 0x79, 0x3b,
	0xe8, 0x6b, 0xce, 0x35, 0x18, 0x03, 0xfc, 0x7d, 0x25, 0x7b, 0x53, 0xa9, 0xc2, 0x93, 0xa6, 0xb7,
	0xe5, 0xcc, 0xfd, 0xbc, 0xfb, 0xff, 0x9c, 0xf7, 0x03, 0x01, 0x0d, 0x1e, 0x0a, 0x68, 0x7d, 0xfe,
	0xe9, 0x85, 0x41, 0xa5, 0x3f, 0x7b, 0x42, 0xf9, 0xf6, 0xc3, 0x3f, 0x88, 0x7c, 0xfb, 0x58, 0x25,
	0xcb, 0x8a, 0x38, 0xee, 0xd9, 0x19, 0x2e, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x46, 0x2d, 0x8f,
	0x40, 0xb4, 0x03, 0x00, 0x00,
}
