// Code generated by protoc-gen-go.
// DO NOT EDIT!

package appengine

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type ImagesServiceError_ErrorCode int32

const (
	ImagesServiceError_UNSPECIFIED_ERROR  ImagesServiceError_ErrorCode = 1
	ImagesServiceError_BAD_TRANSFORM_DATA ImagesServiceError_ErrorCode = 2
	ImagesServiceError_NOT_IMAGE          ImagesServiceError_ErrorCode = 3
	ImagesServiceError_BAD_IMAGE_DATA     ImagesServiceError_ErrorCode = 4
	ImagesServiceError_IMAGE_TOO_LARGE    ImagesServiceError_ErrorCode = 5
	ImagesServiceError_INVALID_BLOB_KEY   ImagesServiceError_ErrorCode = 6
	ImagesServiceError_ACCESS_DENIED      ImagesServiceError_ErrorCode = 7
	ImagesServiceError_OBJECT_NOT_FOUND   ImagesServiceError_ErrorCode = 8
)

var ImagesServiceError_ErrorCode_name = map[int32]string{
	1: "UNSPECIFIED_ERROR",
	2: "BAD_TRANSFORM_DATA",
	3: "NOT_IMAGE",
	4: "BAD_IMAGE_DATA",
	5: "IMAGE_TOO_LARGE",
	6: "INVALID_BLOB_KEY",
	7: "ACCESS_DENIED",
	8: "OBJECT_NOT_FOUND",
}
var ImagesServiceError_ErrorCode_value = map[string]int32{
	"UNSPECIFIED_ERROR":  1,
	"BAD_TRANSFORM_DATA": 2,
	"NOT_IMAGE":          3,
	"BAD_IMAGE_DATA":     4,
	"IMAGE_TOO_LARGE":    5,
	"INVALID_BLOB_KEY":   6,
	"ACCESS_DENIED":      7,
	"OBJECT_NOT_FOUND":   8,
}

func (x ImagesServiceError_ErrorCode) Enum() *ImagesServiceError_ErrorCode {
	p := new(ImagesServiceError_ErrorCode)
	*p = x
	return p
}
func (x ImagesServiceError_ErrorCode) String() string {
	return proto.EnumName(ImagesServiceError_ErrorCode_name, int32(x))
}
func (x ImagesServiceError_ErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *ImagesServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ImagesServiceError_ErrorCode_value, data, "ImagesServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = ImagesServiceError_ErrorCode(value)
	return nil
}

type ImagesServiceTransform_Type int32

const (
	ImagesServiceTransform_RESIZE           ImagesServiceTransform_Type = 1
	ImagesServiceTransform_ROTATE           ImagesServiceTransform_Type = 2
	ImagesServiceTransform_HORIZONTAL_FLIP  ImagesServiceTransform_Type = 3
	ImagesServiceTransform_VERTICAL_FLIP    ImagesServiceTransform_Type = 4
	ImagesServiceTransform_CROP             ImagesServiceTransform_Type = 5
	ImagesServiceTransform_IM_FEELING_LUCKY ImagesServiceTransform_Type = 6
)

var ImagesServiceTransform_Type_name = map[int32]string{
	1: "RESIZE",
	2: "ROTATE",
	3: "HORIZONTAL_FLIP",
	4: "VERTICAL_FLIP",
	5: "CROP",
	6: "IM_FEELING_LUCKY",
}
var ImagesServiceTransform_Type_value = map[string]int32{
	"RESIZE":           1,
	"ROTATE":           2,
	"HORIZONTAL_FLIP":  3,
	"VERTICAL_FLIP":    4,
	"CROP":             5,
	"IM_FEELING_LUCKY": 6,
}

func (x ImagesServiceTransform_Type) Enum() *ImagesServiceTransform_Type {
	p := new(ImagesServiceTransform_Type)
	*p = x
	return p
}
func (x ImagesServiceTransform_Type) String() string {
	return proto.EnumName(ImagesServiceTransform_Type_name, int32(x))
}
func (x ImagesServiceTransform_Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *ImagesServiceTransform_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ImagesServiceTransform_Type_value, data, "ImagesServiceTransform_Type")
	if err != nil {
		return err
	}
	*x = ImagesServiceTransform_Type(value)
	return nil
}

type InputSettings_ORIENTATION_CORRECTION_TYPE int32

const (
	InputSettings_UNCHANGED_ORIENTATION InputSettings_ORIENTATION_CORRECTION_TYPE = 0
	InputSettings_CORRECT_ORIENTATION   InputSettings_ORIENTATION_CORRECTION_TYPE = 1
)

var InputSettings_ORIENTATION_CORRECTION_TYPE_name = map[int32]string{
	0: "UNCHANGED_ORIENTATION",
	1: "CORRECT_ORIENTATION",
}
var InputSettings_ORIENTATION_CORRECTION_TYPE_value = map[string]int32{
	"UNCHANGED_ORIENTATION": 0,
	"CORRECT_ORIENTATION":   1,
}

func (x InputSettings_ORIENTATION_CORRECTION_TYPE) Enum() *InputSettings_ORIENTATION_CORRECTION_TYPE {
	p := new(InputSettings_ORIENTATION_CORRECTION_TYPE)
	*p = x
	return p
}
func (x InputSettings_ORIENTATION_CORRECTION_TYPE) String() string {
	return proto.EnumName(InputSettings_ORIENTATION_CORRECTION_TYPE_name, int32(x))
}
func (x InputSettings_ORIENTATION_CORRECTION_TYPE) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *InputSettings_ORIENTATION_CORRECTION_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(InputSettings_ORIENTATION_CORRECTION_TYPE_value, data, "InputSettings_ORIENTATION_CORRECTION_TYPE")
	if err != nil {
		return err
	}
	*x = InputSettings_ORIENTATION_CORRECTION_TYPE(value)
	return nil
}

type OutputSettings_MIME_TYPE int32

const (
	OutputSettings_PNG  OutputSettings_MIME_TYPE = 0
	OutputSettings_JPEG OutputSettings_MIME_TYPE = 1
	OutputSettings_WEBP OutputSettings_MIME_TYPE = 2
)

var OutputSettings_MIME_TYPE_name = map[int32]string{
	0: "PNG",
	1: "JPEG",
	2: "WEBP",
}
var OutputSettings_MIME_TYPE_value = map[string]int32{
	"PNG":  0,
	"JPEG": 1,
	"WEBP": 2,
}

func (x OutputSettings_MIME_TYPE) Enum() *OutputSettings_MIME_TYPE {
	p := new(OutputSettings_MIME_TYPE)
	*p = x
	return p
}
func (x OutputSettings_MIME_TYPE) String() string {
	return proto.EnumName(OutputSettings_MIME_TYPE_name, int32(x))
}
func (x OutputSettings_MIME_TYPE) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *OutputSettings_MIME_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OutputSettings_MIME_TYPE_value, data, "OutputSettings_MIME_TYPE")
	if err != nil {
		return err
	}
	*x = OutputSettings_MIME_TYPE(value)
	return nil
}

type CompositeImageOptions_ANCHOR int32

const (
	CompositeImageOptions_TOP_LEFT     CompositeImageOptions_ANCHOR = 0
	CompositeImageOptions_TOP          CompositeImageOptions_ANCHOR = 1
	CompositeImageOptions_TOP_RIGHT    CompositeImageOptions_ANCHOR = 2
	CompositeImageOptions_LEFT         CompositeImageOptions_ANCHOR = 3
	CompositeImageOptions_CENTER       CompositeImageOptions_ANCHOR = 4
	CompositeImageOptions_RIGHT        CompositeImageOptions_ANCHOR = 5
	CompositeImageOptions_BOTTOM_LEFT  CompositeImageOptions_ANCHOR = 6
	CompositeImageOptions_BOTTOM       CompositeImageOptions_ANCHOR = 7
	CompositeImageOptions_BOTTOM_RIGHT CompositeImageOptions_ANCHOR = 8
)

var CompositeImageOptions_ANCHOR_name = map[int32]string{
	0: "TOP_LEFT",
	1: "TOP",
	2: "TOP_RIGHT",
	3: "LEFT",
	4: "CENTER",
	5: "RIGHT",
	6: "BOTTOM_LEFT",
	7: "BOTTOM",
	8: "BOTTOM_RIGHT",
}
var CompositeImageOptions_ANCHOR_value = map[string]int32{
	"TOP_LEFT":     0,
	"TOP":          1,
	"TOP_RIGHT":    2,
	"LEFT":         3,
	"CENTER":       4,
	"RIGHT":        5,
	"BOTTOM_LEFT":  6,
	"BOTTOM":       7,
	"BOTTOM_RIGHT": 8,
}

func (x CompositeImageOptions_ANCHOR) Enum() *CompositeImageOptions_ANCHOR {
	p := new(CompositeImageOptions_ANCHOR)
	*p = x
	return p
}
func (x CompositeImageOptions_ANCHOR) String() string {
	return proto.EnumName(CompositeImageOptions_ANCHOR_name, int32(x))
}
func (x CompositeImageOptions_ANCHOR) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *CompositeImageOptions_ANCHOR) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CompositeImageOptions_ANCHOR_value, data, "CompositeImageOptions_ANCHOR")
	if err != nil {
		return err
	}
	*x = CompositeImageOptions_ANCHOR(value)
	return nil
}

type ImagesServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *ImagesServiceError) Reset()         { *this = ImagesServiceError{} }
func (this *ImagesServiceError) String() string { return proto.CompactTextString(this) }
func (*ImagesServiceError) ProtoMessage()       {}

type ImagesServiceTransform struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *ImagesServiceTransform) Reset()         { *this = ImagesServiceTransform{} }
func (this *ImagesServiceTransform) String() string { return proto.CompactTextString(this) }
func (*ImagesServiceTransform) ProtoMessage()       {}

type Transform struct {
	Width            *int32   `protobuf:"varint,1,opt,name=width" json:"width,omitempty"`
	Height           *int32   `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	CropToFit        *bool    `protobuf:"varint,11,opt,name=crop_to_fit,def=0" json:"crop_to_fit,omitempty"`
	CropOffsetX      *float32 `protobuf:"fixed32,12,opt,name=crop_offset_x,def=0.5" json:"crop_offset_x,omitempty"`
	CropOffsetY      *float32 `protobuf:"fixed32,13,opt,name=crop_offset_y,def=0.5" json:"crop_offset_y,omitempty"`
	Rotate           *int32   `protobuf:"varint,3,opt,name=rotate,def=0" json:"rotate,omitempty"`
	HorizontalFlip   *bool    `protobuf:"varint,4,opt,name=horizontal_flip,def=0" json:"horizontal_flip,omitempty"`
	VerticalFlip     *bool    `protobuf:"varint,5,opt,name=vertical_flip,def=0" json:"vertical_flip,omitempty"`
	CropLeftX        *float32 `protobuf:"fixed32,6,opt,name=crop_left_x,def=0" json:"crop_left_x,omitempty"`
	CropTopY         *float32 `protobuf:"fixed32,7,opt,name=crop_top_y,def=0" json:"crop_top_y,omitempty"`
	CropRightX       *float32 `protobuf:"fixed32,8,opt,name=crop_right_x,def=1" json:"crop_right_x,omitempty"`
	CropBottomY      *float32 `protobuf:"fixed32,9,opt,name=crop_bottom_y,def=1" json:"crop_bottom_y,omitempty"`
	Autolevels       *bool    `protobuf:"varint,10,opt,name=autolevels,def=0" json:"autolevels,omitempty"`
	AllowStretch     *bool    `protobuf:"varint,14,opt,name=allow_stretch,def=0" json:"allow_stretch,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *Transform) Reset()         { *this = Transform{} }
func (this *Transform) String() string { return proto.CompactTextString(this) }
func (*Transform) ProtoMessage()       {}

const Default_Transform_CropToFit bool = false
const Default_Transform_CropOffsetX float32 = 0.5
const Default_Transform_CropOffsetY float32 = 0.5
const Default_Transform_Rotate int32 = 0
const Default_Transform_HorizontalFlip bool = false
const Default_Transform_VerticalFlip bool = false
const Default_Transform_CropLeftX float32 = 0
const Default_Transform_CropTopY float32 = 0
const Default_Transform_CropRightX float32 = 1
const Default_Transform_CropBottomY float32 = 1
const Default_Transform_Autolevels bool = false
const Default_Transform_AllowStretch bool = false

func (this *Transform) GetWidth() int32 {
	if this != nil && this.Width != nil {
		return *this.Width
	}
	return 0
}

func (this *Transform) GetHeight() int32 {
	if this != nil && this.Height != nil {
		return *this.Height
	}
	return 0
}

func (this *Transform) GetCropToFit() bool {
	if this != nil && this.CropToFit != nil {
		return *this.CropToFit
	}
	return Default_Transform_CropToFit
}

func (this *Transform) GetCropOffsetX() float32 {
	if this != nil && this.CropOffsetX != nil {
		return *this.CropOffsetX
	}
	return Default_Transform_CropOffsetX
}

func (this *Transform) GetCropOffsetY() float32 {
	if this != nil && this.CropOffsetY != nil {
		return *this.CropOffsetY
	}
	return Default_Transform_CropOffsetY
}

func (this *Transform) GetRotate() int32 {
	if this != nil && this.Rotate != nil {
		return *this.Rotate
	}
	return Default_Transform_Rotate
}

func (this *Transform) GetHorizontalFlip() bool {
	if this != nil && this.HorizontalFlip != nil {
		return *this.HorizontalFlip
	}
	return Default_Transform_HorizontalFlip
}

func (this *Transform) GetVerticalFlip() bool {
	if this != nil && this.VerticalFlip != nil {
		return *this.VerticalFlip
	}
	return Default_Transform_VerticalFlip
}

func (this *Transform) GetCropLeftX() float32 {
	if this != nil && this.CropLeftX != nil {
		return *this.CropLeftX
	}
	return Default_Transform_CropLeftX
}

func (this *Transform) GetCropTopY() float32 {
	if this != nil && this.CropTopY != nil {
		return *this.CropTopY
	}
	return Default_Transform_CropTopY
}

func (this *Transform) GetCropRightX() float32 {
	if this != nil && this.CropRightX != nil {
		return *this.CropRightX
	}
	return Default_Transform_CropRightX
}

func (this *Transform) GetCropBottomY() float32 {
	if this != nil && this.CropBottomY != nil {
		return *this.CropBottomY
	}
	return Default_Transform_CropBottomY
}

func (this *Transform) GetAutolevels() bool {
	if this != nil && this.Autolevels != nil {
		return *this.Autolevels
	}
	return Default_Transform_Autolevels
}

func (this *Transform) GetAllowStretch() bool {
	if this != nil && this.AllowStretch != nil {
		return *this.AllowStretch
	}
	return Default_Transform_AllowStretch
}

type ImageData struct {
	Content          []byte  `protobuf:"bytes,1,req,name=content" json:"content,omitempty"`
	BlobKey          *string `protobuf:"bytes,2,opt,name=blob_key" json:"blob_key,omitempty"`
	Width            *int32  `protobuf:"varint,3,opt,name=width" json:"width,omitempty"`
	Height           *int32  `protobuf:"varint,4,opt,name=height" json:"height,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ImageData) Reset()         { *this = ImageData{} }
func (this *ImageData) String() string { return proto.CompactTextString(this) }
func (*ImageData) ProtoMessage()       {}

func (this *ImageData) GetContent() []byte {
	if this != nil {
		return this.Content
	}
	return nil
}

func (this *ImageData) GetBlobKey() string {
	if this != nil && this.BlobKey != nil {
		return *this.BlobKey
	}
	return ""
}

func (this *ImageData) GetWidth() int32 {
	if this != nil && this.Width != nil {
		return *this.Width
	}
	return 0
}

func (this *ImageData) GetHeight() int32 {
	if this != nil && this.Height != nil {
		return *this.Height
	}
	return 0
}

type InputSettings struct {
	CorrectExifOrientation     *InputSettings_ORIENTATION_CORRECTION_TYPE `protobuf:"varint,1,opt,name=correct_exif_orientation,enum=appengine.InputSettings_ORIENTATION_CORRECTION_TYPE,def=0" json:"correct_exif_orientation,omitempty"`
	ParseMetadata              *bool                                      `protobuf:"varint,2,opt,name=parse_metadata,def=0" json:"parse_metadata,omitempty"`
	TransparentSubstitutionRgb *int32                                     `protobuf:"varint,3,opt,name=transparent_substitution_rgb" json:"transparent_substitution_rgb,omitempty"`
	XXX_unrecognized           []byte                                     `json:"-"`
}

func (this *InputSettings) Reset()         { *this = InputSettings{} }
func (this *InputSettings) String() string { return proto.CompactTextString(this) }
func (*InputSettings) ProtoMessage()       {}

const Default_InputSettings_CorrectExifOrientation InputSettings_ORIENTATION_CORRECTION_TYPE = InputSettings_UNCHANGED_ORIENTATION
const Default_InputSettings_ParseMetadata bool = false

func (this *InputSettings) GetCorrectExifOrientation() InputSettings_ORIENTATION_CORRECTION_TYPE {
	if this != nil && this.CorrectExifOrientation != nil {
		return *this.CorrectExifOrientation
	}
	return Default_InputSettings_CorrectExifOrientation
}

func (this *InputSettings) GetParseMetadata() bool {
	if this != nil && this.ParseMetadata != nil {
		return *this.ParseMetadata
	}
	return Default_InputSettings_ParseMetadata
}

func (this *InputSettings) GetTransparentSubstitutionRgb() int32 {
	if this != nil && this.TransparentSubstitutionRgb != nil {
		return *this.TransparentSubstitutionRgb
	}
	return 0
}

type OutputSettings struct {
	MimeType         *OutputSettings_MIME_TYPE `protobuf:"varint,1,opt,name=mime_type,enum=appengine.OutputSettings_MIME_TYPE,def=0" json:"mime_type,omitempty"`
	Quality          *int32                    `protobuf:"varint,2,opt,name=quality" json:"quality,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (this *OutputSettings) Reset()         { *this = OutputSettings{} }
func (this *OutputSettings) String() string { return proto.CompactTextString(this) }
func (*OutputSettings) ProtoMessage()       {}

const Default_OutputSettings_MimeType OutputSettings_MIME_TYPE = OutputSettings_PNG

func (this *OutputSettings) GetMimeType() OutputSettings_MIME_TYPE {
	if this != nil && this.MimeType != nil {
		return *this.MimeType
	}
	return Default_OutputSettings_MimeType
}

func (this *OutputSettings) GetQuality() int32 {
	if this != nil && this.Quality != nil {
		return *this.Quality
	}
	return 0
}

type ImagesTransformRequest struct {
	Image            *ImageData      `protobuf:"bytes,1,req,name=image" json:"image,omitempty"`
	Transform        []*Transform    `protobuf:"bytes,2,rep,name=transform" json:"transform,omitempty"`
	Output           *OutputSettings `protobuf:"bytes,3,req,name=output" json:"output,omitempty"`
	Input            *InputSettings  `protobuf:"bytes,4,opt,name=input" json:"input,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (this *ImagesTransformRequest) Reset()         { *this = ImagesTransformRequest{} }
func (this *ImagesTransformRequest) String() string { return proto.CompactTextString(this) }
func (*ImagesTransformRequest) ProtoMessage()       {}

func (this *ImagesTransformRequest) GetImage() *ImageData {
	if this != nil {
		return this.Image
	}
	return nil
}

func (this *ImagesTransformRequest) GetOutput() *OutputSettings {
	if this != nil {
		return this.Output
	}
	return nil
}

func (this *ImagesTransformRequest) GetInput() *InputSettings {
	if this != nil {
		return this.Input
	}
	return nil
}

type ImagesTransformResponse struct {
	Image            *ImageData `protobuf:"bytes,1,req,name=image" json:"image,omitempty"`
	SourceMetadata   *string    `protobuf:"bytes,2,opt,name=source_metadata" json:"source_metadata,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (this *ImagesTransformResponse) Reset()         { *this = ImagesTransformResponse{} }
func (this *ImagesTransformResponse) String() string { return proto.CompactTextString(this) }
func (*ImagesTransformResponse) ProtoMessage()       {}

func (this *ImagesTransformResponse) GetImage() *ImageData {
	if this != nil {
		return this.Image
	}
	return nil
}

func (this *ImagesTransformResponse) GetSourceMetadata() string {
	if this != nil && this.SourceMetadata != nil {
		return *this.SourceMetadata
	}
	return ""
}

type CompositeImageOptions struct {
	SourceIndex      *int32                        `protobuf:"varint,1,req,name=source_index" json:"source_index,omitempty"`
	XOffset          *int32                        `protobuf:"varint,2,req,name=x_offset" json:"x_offset,omitempty"`
	YOffset          *int32                        `protobuf:"varint,3,req,name=y_offset" json:"y_offset,omitempty"`
	Opacity          *float32                      `protobuf:"fixed32,4,req,name=opacity" json:"opacity,omitempty"`
	Anchor           *CompositeImageOptions_ANCHOR `protobuf:"varint,5,req,name=anchor,enum=appengine.CompositeImageOptions_ANCHOR" json:"anchor,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (this *CompositeImageOptions) Reset()         { *this = CompositeImageOptions{} }
func (this *CompositeImageOptions) String() string { return proto.CompactTextString(this) }
func (*CompositeImageOptions) ProtoMessage()       {}

func (this *CompositeImageOptions) GetSourceIndex() int32 {
	if this != nil && this.SourceIndex != nil {
		return *this.SourceIndex
	}
	return 0
}

func (this *CompositeImageOptions) GetXOffset() int32 {
	if this != nil && this.XOffset != nil {
		return *this.XOffset
	}
	return 0
}

func (this *CompositeImageOptions) GetYOffset() int32 {
	if this != nil && this.YOffset != nil {
		return *this.YOffset
	}
	return 0
}

func (this *CompositeImageOptions) GetOpacity() float32 {
	if this != nil && this.Opacity != nil {
		return *this.Opacity
	}
	return 0
}

func (this *CompositeImageOptions) GetAnchor() CompositeImageOptions_ANCHOR {
	if this != nil && this.Anchor != nil {
		return *this.Anchor
	}
	return 0
}

type ImagesCanvas struct {
	Width            *int32          `protobuf:"varint,1,req,name=width" json:"width,omitempty"`
	Height           *int32          `protobuf:"varint,2,req,name=height" json:"height,omitempty"`
	Output           *OutputSettings `protobuf:"bytes,3,req,name=output" json:"output,omitempty"`
	Color            *int32          `protobuf:"varint,4,opt,name=color,def=-1" json:"color,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (this *ImagesCanvas) Reset()         { *this = ImagesCanvas{} }
func (this *ImagesCanvas) String() string { return proto.CompactTextString(this) }
func (*ImagesCanvas) ProtoMessage()       {}

const Default_ImagesCanvas_Color int32 = -1

func (this *ImagesCanvas) GetWidth() int32 {
	if this != nil && this.Width != nil {
		return *this.Width
	}
	return 0
}

func (this *ImagesCanvas) GetHeight() int32 {
	if this != nil && this.Height != nil {
		return *this.Height
	}
	return 0
}

func (this *ImagesCanvas) GetOutput() *OutputSettings {
	if this != nil {
		return this.Output
	}
	return nil
}

func (this *ImagesCanvas) GetColor() int32 {
	if this != nil && this.Color != nil {
		return *this.Color
	}
	return Default_ImagesCanvas_Color
}

type ImagesCompositeRequest struct {
	Image            []*ImageData             `protobuf:"bytes,1,rep,name=image" json:"image,omitempty"`
	Options          []*CompositeImageOptions `protobuf:"bytes,2,rep,name=options" json:"options,omitempty"`
	Canvas           *ImagesCanvas            `protobuf:"bytes,3,req,name=canvas" json:"canvas,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (this *ImagesCompositeRequest) Reset()         { *this = ImagesCompositeRequest{} }
func (this *ImagesCompositeRequest) String() string { return proto.CompactTextString(this) }
func (*ImagesCompositeRequest) ProtoMessage()       {}

func (this *ImagesCompositeRequest) GetCanvas() *ImagesCanvas {
	if this != nil {
		return this.Canvas
	}
	return nil
}

type ImagesCompositeResponse struct {
	Image            *ImageData `protobuf:"bytes,1,req,name=image" json:"image,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (this *ImagesCompositeResponse) Reset()         { *this = ImagesCompositeResponse{} }
func (this *ImagesCompositeResponse) String() string { return proto.CompactTextString(this) }
func (*ImagesCompositeResponse) ProtoMessage()       {}

func (this *ImagesCompositeResponse) GetImage() *ImageData {
	if this != nil {
		return this.Image
	}
	return nil
}

type ImagesHistogramRequest struct {
	Image            *ImageData `protobuf:"bytes,1,req,name=image" json:"image,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (this *ImagesHistogramRequest) Reset()         { *this = ImagesHistogramRequest{} }
func (this *ImagesHistogramRequest) String() string { return proto.CompactTextString(this) }
func (*ImagesHistogramRequest) ProtoMessage()       {}

func (this *ImagesHistogramRequest) GetImage() *ImageData {
	if this != nil {
		return this.Image
	}
	return nil
}

type ImagesHistogram struct {
	Red              []int32 `protobuf:"varint,1,rep,name=red" json:"red,omitempty"`
	Green            []int32 `protobuf:"varint,2,rep,name=green" json:"green,omitempty"`
	Blue             []int32 `protobuf:"varint,3,rep,name=blue" json:"blue,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ImagesHistogram) Reset()         { *this = ImagesHistogram{} }
func (this *ImagesHistogram) String() string { return proto.CompactTextString(this) }
func (*ImagesHistogram) ProtoMessage()       {}

type ImagesHistogramResponse struct {
	Histogram        *ImagesHistogram `protobuf:"bytes,1,req,name=histogram" json:"histogram,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (this *ImagesHistogramResponse) Reset()         { *this = ImagesHistogramResponse{} }
func (this *ImagesHistogramResponse) String() string { return proto.CompactTextString(this) }
func (*ImagesHistogramResponse) ProtoMessage()       {}

func (this *ImagesHistogramResponse) GetHistogram() *ImagesHistogram {
	if this != nil {
		return this.Histogram
	}
	return nil
}

type ImagesGetUrlBaseRequest struct {
	BlobKey          *string `protobuf:"bytes,1,req,name=blob_key" json:"blob_key,omitempty"`
	CreateSecureUrl  *bool   `protobuf:"varint,2,opt,name=create_secure_url,def=0" json:"create_secure_url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ImagesGetUrlBaseRequest) Reset()         { *this = ImagesGetUrlBaseRequest{} }
func (this *ImagesGetUrlBaseRequest) String() string { return proto.CompactTextString(this) }
func (*ImagesGetUrlBaseRequest) ProtoMessage()       {}

const Default_ImagesGetUrlBaseRequest_CreateSecureUrl bool = false

func (this *ImagesGetUrlBaseRequest) GetBlobKey() string {
	if this != nil && this.BlobKey != nil {
		return *this.BlobKey
	}
	return ""
}

func (this *ImagesGetUrlBaseRequest) GetCreateSecureUrl() bool {
	if this != nil && this.CreateSecureUrl != nil {
		return *this.CreateSecureUrl
	}
	return Default_ImagesGetUrlBaseRequest_CreateSecureUrl
}

type ImagesGetUrlBaseResponse struct {
	Url              *string `protobuf:"bytes,1,req,name=url" json:"url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ImagesGetUrlBaseResponse) Reset()         { *this = ImagesGetUrlBaseResponse{} }
func (this *ImagesGetUrlBaseResponse) String() string { return proto.CompactTextString(this) }
func (*ImagesGetUrlBaseResponse) ProtoMessage()       {}

func (this *ImagesGetUrlBaseResponse) GetUrl() string {
	if this != nil && this.Url != nil {
		return *this.Url
	}
	return ""
}

type ImagesDeleteUrlBaseRequest struct {
	BlobKey          *string `protobuf:"bytes,1,req,name=blob_key" json:"blob_key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *ImagesDeleteUrlBaseRequest) Reset()         { *this = ImagesDeleteUrlBaseRequest{} }
func (this *ImagesDeleteUrlBaseRequest) String() string { return proto.CompactTextString(this) }
func (*ImagesDeleteUrlBaseRequest) ProtoMessage()       {}

func (this *ImagesDeleteUrlBaseRequest) GetBlobKey() string {
	if this != nil && this.BlobKey != nil {
		return *this.BlobKey
	}
	return ""
}

type ImagesDeleteUrlBaseResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *ImagesDeleteUrlBaseResponse) Reset()         { *this = ImagesDeleteUrlBaseResponse{} }
func (this *ImagesDeleteUrlBaseResponse) String() string { return proto.CompactTextString(this) }
func (*ImagesDeleteUrlBaseResponse) ProtoMessage()       {}

func init() {
	proto.RegisterEnum("appengine.ImagesServiceError_ErrorCode", ImagesServiceError_ErrorCode_name, ImagesServiceError_ErrorCode_value)
	proto.RegisterEnum("appengine.ImagesServiceTransform_Type", ImagesServiceTransform_Type_name, ImagesServiceTransform_Type_value)
	proto.RegisterEnum("appengine.InputSettings_ORIENTATION_CORRECTION_TYPE", InputSettings_ORIENTATION_CORRECTION_TYPE_name, InputSettings_ORIENTATION_CORRECTION_TYPE_value)
	proto.RegisterEnum("appengine.OutputSettings_MIME_TYPE", OutputSettings_MIME_TYPE_name, OutputSettings_MIME_TYPE_value)
	proto.RegisterEnum("appengine.CompositeImageOptions_ANCHOR", CompositeImageOptions_ANCHOR_name, CompositeImageOptions_ANCHOR_value)
}