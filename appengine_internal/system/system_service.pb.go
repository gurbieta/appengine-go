// Code generated by protoc-gen-go from "system_service.proto"
// DO NOT EDIT!

package appengine

import proto "goprotobuf.googlecode.com/hg/proto"
import "math"
import "os"

// Reference proto, math & os imports to suppress error if they are not otherwise used.
var _ = proto.GetString
var _ = math.Inf
var _ os.Error

type SystemServiceError_ErrorCode int32

const (
	SystemServiceError_OK             SystemServiceError_ErrorCode = 0
	SystemServiceError_INTERNAL_ERROR SystemServiceError_ErrorCode = 1
)

var SystemServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "INTERNAL_ERROR",
}
var SystemServiceError_ErrorCode_value = map[string]int32{
	"OK":             0,
	"INTERNAL_ERROR": 1,
}

func NewSystemServiceError_ErrorCode(x SystemServiceError_ErrorCode) *SystemServiceError_ErrorCode {
	e := SystemServiceError_ErrorCode(x)
	return &e
}
func (x SystemServiceError_ErrorCode) String() string {
	return proto.EnumName(SystemServiceError_ErrorCode_name, int32(x))
}

type SystemServiceError struct {
	XXX_unrecognized []byte `json:",omitempty"`
}

func (this *SystemServiceError) Reset()         { *this = SystemServiceError{} }
func (this *SystemServiceError) String() string { return proto.CompactTextString(this) }

type SystemStat struct {
	Current          *float64 `protobuf:"fixed64,1,opt,name=current" json:"current,omitempty"`
	Average1M        *float64 `protobuf:"fixed64,3,opt,name=average1m" json:"average1m,omitempty"`
	Average10M       *float64 `protobuf:"fixed64,4,opt,name=average10m" json:"average10m,omitempty"`
	Total            *float64 `protobuf:"fixed64,2,opt,name=total" json:"total,omitempty"`
	Rate1M           *float64 `protobuf:"fixed64,5,opt,name=rate1m" json:"rate1m,omitempty"`
	Rate10M          *float64 `protobuf:"fixed64,6,opt,name=rate10m" json:"rate10m,omitempty"`
	XXX_unrecognized []byte   `json:",omitempty"`
}

func (this *SystemStat) Reset()         { *this = SystemStat{} }
func (this *SystemStat) String() string { return proto.CompactTextString(this) }

type GetSystemStatsRequest struct {
	XXX_unrecognized []byte `json:",omitempty"`
}

func (this *GetSystemStatsRequest) Reset()         { *this = GetSystemStatsRequest{} }
func (this *GetSystemStatsRequest) String() string { return proto.CompactTextString(this) }

type GetSystemStatsResponse struct {
	Cpu              *SystemStat `protobuf:"bytes,1,opt,name=cpu" json:"cpu,omitempty"`
	Memory           *SystemStat `protobuf:"bytes,2,opt,name=memory" json:"memory,omitempty"`
	XXX_unrecognized []byte      `json:",omitempty"`
}

func (this *GetSystemStatsResponse) Reset()         { *this = GetSystemStatsResponse{} }
func (this *GetSystemStatsResponse) String() string { return proto.CompactTextString(this) }

type StartBackgroundRequestRequest struct {
	XXX_unrecognized []byte `json:",omitempty"`
}

func (this *StartBackgroundRequestRequest) Reset()         { *this = StartBackgroundRequestRequest{} }
func (this *StartBackgroundRequestRequest) String() string { return proto.CompactTextString(this) }

type StartBackgroundRequestResponse struct {
	RequestId        *string `protobuf:"bytes,1,opt,name=request_id" json:"request_id,omitempty"`
	XXX_unrecognized []byte  `json:",omitempty"`
}

func (this *StartBackgroundRequestResponse) Reset()         { *this = StartBackgroundRequestResponse{} }
func (this *StartBackgroundRequestResponse) String() string { return proto.CompactTextString(this) }

func init() {
	proto.RegisterEnum("appengine.SystemServiceError_ErrorCode", SystemServiceError_ErrorCode_name, SystemServiceError_ErrorCode_value)
}
