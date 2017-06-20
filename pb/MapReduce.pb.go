// Code generated by protoc-gen-go.
// source: MapReduce.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScanMetrics struct {
	Metrics          []*NameInt64Pair `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ScanMetrics) Reset()                    { *m = ScanMetrics{} }
func (m *ScanMetrics) String() string            { return proto.CompactTextString(m) }
func (*ScanMetrics) ProtoMessage()               {}
func (*ScanMetrics) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

func (m *ScanMetrics) GetMetrics() []*NameInt64Pair {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type TableSnapshotRegionSplit struct {
	Locations        []string     `protobuf:"bytes,2,rep,name=locations" json:"locations,omitempty"`
	Table            *TableSchema `protobuf:"bytes,3,opt,name=table" json:"table,omitempty"`
	Region           *RegionInfo  `protobuf:"bytes,4,opt,name=region" json:"region,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *TableSnapshotRegionSplit) Reset()                    { *m = TableSnapshotRegionSplit{} }
func (m *TableSnapshotRegionSplit) String() string            { return proto.CompactTextString(m) }
func (*TableSnapshotRegionSplit) ProtoMessage()               {}
func (*TableSnapshotRegionSplit) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{1} }

func (m *TableSnapshotRegionSplit) GetLocations() []string {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *TableSnapshotRegionSplit) GetTable() *TableSchema {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *TableSnapshotRegionSplit) GetRegion() *RegionInfo {
	if m != nil {
		return m.Region
	}
	return nil
}

func init() {
	proto.RegisterType((*ScanMetrics)(nil), "pb.ScanMetrics")
	proto.RegisterType((*TableSnapshotRegionSplit)(nil), "pb.TableSnapshotRegionSplit")
}

var fileDescriptor16 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x18, 0x27, 0x9b, 0x7f, 0x58, 0x0a, 0x0e, 0x7b, 0x0a, 0xe2, 0x61, 0x0c, 0x14, 0x41, 0x08, 0x28,
	0xe2, 0xc1, 0x63, 0x4f, 0xeb, 0x61, 0x32, 0x52, 0x5f, 0xe0, 0x4b, 0xfa, 0xad, 0x2d, 0xb4, 0x49,
	0x48, 0xb2, 0x67, 0xf0, 0x35, 0x7c, 0x54, 0x93, 0x65, 0x63, 0xb7, 0xe4, 0xf7, 0x97, 0xdf, 0x47,
	0x97, 0x5b, 0xb0, 0x02, 0xdb, 0x83, 0x42, 0x6e, 0x9d, 0x09, 0xa6, 0x9c, 0x59, 0xf9, 0x50, 0x6c,
	0x2a, 0xf0, 0x27, 0x60, 0xfd, 0x45, 0x8b, 0x46, 0x81, 0xde, 0x62, 0x70, 0x83, 0xf2, 0xe5, 0x2b,
	0xbd, 0x9d, 0xf2, 0x93, 0x91, 0xd5, 0xfc, 0xa5, 0x78, 0xbf, 0xe7, 0x56, 0xf2, 0x6f, 0x98, 0xb0,
	0xd6, 0xe1, 0xf3, 0x63, 0x07, 0x83, 0x13, 0x67, 0xc5, 0xfa, 0x97, 0x50, 0xf6, 0x03, 0x72, 0xc4,
	0x46, 0x83, 0xf5, 0xbd, 0x09, 0x02, 0xbb, 0xc1, 0xe8, 0xc6, 0x8e, 0x43, 0x28, 0x1f, 0xe9, 0x62,
	0x34, 0x0a, 0x42, 0x04, 0x3c, 0x9b, 0xc5, 0xac, 0x85, 0xb8, 0x00, 0xe5, 0x13, 0xbd, 0x0e, 0xc9,
	0xc9, 0xe6, 0x2b, 0x12, 0x5b, 0x96, 0xa9, 0x25, 0x47, 0xa9, 0x1e, 0x27, 0x10, 0x99, 0x2d, 0x9f,
	0xe9, 0x8d, 0x3b, 0x66, 0xb2, 0xab, 0xa3, 0xee, 0x2e, 0xe9, 0x72, 0x4b, 0xad, 0xf7, 0x46, 0x9c,
	0xd8, 0xaa, 0xa6, 0x6f, 0xc6, 0x75, 0x1c, 0x2c, 0x44, 0x3b, 0xef, 0xa1, 0x35, 0xc6, 0xf2, 0x5e,
	0xa6, 0x99, 0x3e, 0xfe, 0xb0, 0xcd, 0x6b, 0xe5, 0x61, 0xcf, 0x3b, 0xd4, 0xe8, 0x20, 0x60, 0x5b,
	0x5d, 0x8e, 0xb3, 0x4b, 0xa4, 0xdf, 0x90, 0x3f, 0x42, 0xfe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x41,
	0x9a, 0xb1, 0x4f, 0x33, 0x01, 0x00, 0x00,
}