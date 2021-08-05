// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.2
// source: block7.proto

package block7pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Column
type Column struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []int32 `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
}

func (x *Column) Reset() {
	*x = Column{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block7_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Column) ProtoMessage() {}

func (x *Column) ProtoReflect() protoreflect.Message {
	mi := &file_block7_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Column.ProtoReflect.Descriptor instead.
func (*Column) Descriptor() ([]byte, []int) {
	return file_block7_proto_rawDescGZIP(), []int{0}
}

func (x *Column) GetValues() []int32 {
	if x != nil {
		return x.Values
	}
	return nil
}

// SceneLayer - scene layer
type SceneLayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*Column `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *SceneLayer) Reset() {
	*x = SceneLayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block7_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneLayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneLayer) ProtoMessage() {}

func (x *SceneLayer) ProtoReflect() protoreflect.Message {
	mi := &file_block7_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneLayer.ProtoReflect.Descriptor instead.
func (*SceneLayer) Descriptor() ([]byte, []int) {
	return file_block7_proto_rawDescGZIP(), []int{1}
}

func (x *SceneLayer) GetValues() []*Column {
	if x != nil {
		return x.Values
	}
	return nil
}

// SpecialLayer - special layer
type SpecialLayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Special   int32   `protobuf:"varint,1,opt,name=special,proto3" json:"special,omitempty"`      // specialType
	Layer     int32   `protobuf:"varint,2,opt,name=layer,proto3" json:"layer,omitempty"`          // layer
	LayerType int32   `protobuf:"varint,3,opt,name=layerType,proto3" json:"layerType,omitempty"`  // layerType
	Values    []int32 `protobuf:"varint,4,rep,packed,name=values,proto3" json:"values,omitempty"` // pos, it like [x0,y0,z0,x1,y1,z1]
}

func (x *SpecialLayer) Reset() {
	*x = SpecialLayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block7_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecialLayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecialLayer) ProtoMessage() {}

func (x *SpecialLayer) ProtoReflect() protoreflect.Message {
	mi := &file_block7_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecialLayer.ProtoReflect.Descriptor instead.
func (*SpecialLayer) Descriptor() ([]byte, []int) {
	return file_block7_proto_rawDescGZIP(), []int{2}
}

func (x *SpecialLayer) GetSpecial() int32 {
	if x != nil {
		return x.Special
	}
	return 0
}

func (x *SpecialLayer) GetLayer() int32 {
	if x != nil {
		return x.Layer
	}
	return 0
}

func (x *SpecialLayer) GetLayerType() int32 {
	if x != nil {
		return x.LayerType
	}
	return 0
}

func (x *SpecialLayer) GetValues() []int32 {
	if x != nil {
		return x.Values
	}
	return nil
}

// Scene - scene
type Scene struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width  int32 `protobuf:"varint,1,opt,name=Width,proto3" json:"Width,omitempty"`   // width
	Height int32 `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"` // height
	Layers int32 `protobuf:"varint,3,opt,name=Layers,proto3" json:"Layers,omitempty"` // layers
	XOff   int32 `protobuf:"varint,4,opt,name=XOff,proto3" json:"XOff,omitempty"`     // x offset
	YOff   int32 `protobuf:"varint,5,opt,name=YOff,proto3" json:"YOff,omitempty"`     // y offset
	// Deprecated: Do not use.
	InitArr []*SceneLayer `protobuf:"bytes,6,rep,name=InitArr,proto3" json:"InitArr,omitempty"` // initial data
	// Deprecated: Do not use.
	History []*Column `protobuf:"bytes,7,rep,name=History,proto3" json:"History,omitempty"` // history
	// Deprecated: Do not use.
	ClickValues int32 `protobuf:"varint,8,opt,name=ClickValues,proto3" json:"ClickValues,omitempty"` // ClickValues
	// Deprecated: Do not use.
	FinishedPer float32 `protobuf:"fixed32,9,opt,name=FinishedPer,proto3" json:"FinishedPer,omitempty"` // FinishedPer
	Offset      string  `protobuf:"bytes,10,opt,name=Offset,proto3" json:"Offset,omitempty"`            // Offset
	// Deprecated: Do not use.
	IsOutputScene bool  `protobuf:"varint,11,opt,name=IsOutputScene,proto3" json:"IsOutputScene,omitempty"` // IsOutputScene
	StageID       int32 `protobuf:"varint,12,opt,name=StageID,proto3" json:"StageID,omitempty"`             // stage
	// Deprecated: Do not use.
	MapID        string          `protobuf:"bytes,13,opt,name=MapID,proto3" json:"MapID,omitempty"`               // map
	Version      int32           `protobuf:"varint,14,opt,name=Version,proto3" json:"Version,omitempty"`          // version
	SceneID      int64           `protobuf:"varint,15,opt,name=sceneID,proto3" json:"sceneID,omitempty"`          // id
	UserID       int64           `protobuf:"varint,16,opt,name=userID,proto3" json:"userID,omitempty"`            // userid
	HistoryID    int64           `protobuf:"varint,17,opt,name=historyID,proto3" json:"historyID,omitempty"`      // historyid
	MapID2       int32           `protobuf:"varint,18,opt,name=MapID2,proto3" json:"MapID2,omitempty"`            // map
	SpecialLayer []*SpecialLayer `protobuf:"bytes,19,rep,name=specialLayer,proto3" json:"specialLayer,omitempty"` // SpecialLayer
	InitArr2     []int32         `protobuf:"varint,20,rep,packed,name=InitArr2,proto3" json:"InitArr2,omitempty"` // initial data
	History2     []int32         `protobuf:"varint,21,rep,packed,name=History2,proto3" json:"History2,omitempty"` // history
	RngData      []int64         `protobuf:"varint,22,rep,packed,name=RngData,proto3" json:"RngData,omitempty"`   // rngdata
	GameState    int32           `protobuf:"varint,23,opt,name=GameState,proto3" json:"GameState,omitempty"`      // gamestate
}

func (x *Scene) Reset() {
	*x = Scene{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block7_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scene) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scene) ProtoMessage() {}

func (x *Scene) ProtoReflect() protoreflect.Message {
	mi := &file_block7_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scene.ProtoReflect.Descriptor instead.
func (*Scene) Descriptor() ([]byte, []int) {
	return file_block7_proto_rawDescGZIP(), []int{3}
}

func (x *Scene) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Scene) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Scene) GetLayers() int32 {
	if x != nil {
		return x.Layers
	}
	return 0
}

func (x *Scene) GetXOff() int32 {
	if x != nil {
		return x.XOff
	}
	return 0
}

func (x *Scene) GetYOff() int32 {
	if x != nil {
		return x.YOff
	}
	return 0
}

// Deprecated: Do not use.
func (x *Scene) GetInitArr() []*SceneLayer {
	if x != nil {
		return x.InitArr
	}
	return nil
}

// Deprecated: Do not use.
func (x *Scene) GetHistory() []*Column {
	if x != nil {
		return x.History
	}
	return nil
}

// Deprecated: Do not use.
func (x *Scene) GetClickValues() int32 {
	if x != nil {
		return x.ClickValues
	}
	return 0
}

// Deprecated: Do not use.
func (x *Scene) GetFinishedPer() float32 {
	if x != nil {
		return x.FinishedPer
	}
	return 0
}

func (x *Scene) GetOffset() string {
	if x != nil {
		return x.Offset
	}
	return ""
}

// Deprecated: Do not use.
func (x *Scene) GetIsOutputScene() bool {
	if x != nil {
		return x.IsOutputScene
	}
	return false
}

func (x *Scene) GetStageID() int32 {
	if x != nil {
		return x.StageID
	}
	return 0
}

// Deprecated: Do not use.
func (x *Scene) GetMapID() string {
	if x != nil {
		return x.MapID
	}
	return ""
}

func (x *Scene) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Scene) GetSceneID() int64 {
	if x != nil {
		return x.SceneID
	}
	return 0
}

func (x *Scene) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Scene) GetHistoryID() int64 {
	if x != nil {
		return x.HistoryID
	}
	return 0
}

func (x *Scene) GetMapID2() int32 {
	if x != nil {
		return x.MapID2
	}
	return 0
}

func (x *Scene) GetSpecialLayer() []*SpecialLayer {
	if x != nil {
		return x.SpecialLayer
	}
	return nil
}

func (x *Scene) GetInitArr2() []int32 {
	if x != nil {
		return x.InitArr2
	}
	return nil
}

func (x *Scene) GetHistory2() []int32 {
	if x != nil {
		return x.History2
	}
	return nil
}

func (x *Scene) GetRngData() []int64 {
	if x != nil {
		return x.RngData
	}
	return nil
}

func (x *Scene) GetGameState() int32 {
	if x != nil {
		return x.GameState
	}
	return 0
}

// MissionParams
type MissionParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MissionID int32 `protobuf:"varint,1,opt,name=MissionID,proto3" json:"MissionID,omitempty"` // mission id
}

func (x *MissionParams) Reset() {
	*x = MissionParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block7_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MissionParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MissionParams) ProtoMessage() {}

func (x *MissionParams) ProtoReflect() protoreflect.Message {
	mi := &file_block7_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MissionParams.ProtoReflect.Descriptor instead.
func (*MissionParams) Descriptor() ([]byte, []int) {
	return file_block7_proto_rawDescGZIP(), []int{4}
}

func (x *MissionParams) GetMissionID() int32 {
	if x != nil {
		return x.MissionID
	}
	return 0
}

// MissionResult
type MissionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MissionID   int32  `protobuf:"varint,1,opt,name=MissionID,proto3" json:"MissionID,omitempty"`    // mission id
	MissionHash string `protobuf:"bytes,2,opt,name=MissionHash,proto3" json:"MissionHash,omitempty"` // mission hash
	Scene       *Scene `protobuf:"bytes,3,opt,name=Scene,proto3" json:"Scene,omitempty"`             // scene
}

func (x *MissionResult) Reset() {
	*x = MissionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block7_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MissionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MissionResult) ProtoMessage() {}

func (x *MissionResult) ProtoReflect() protoreflect.Message {
	mi := &file_block7_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MissionResult.ProtoReflect.Descriptor instead.
func (*MissionResult) Descriptor() ([]byte, []int) {
	return file_block7_proto_rawDescGZIP(), []int{5}
}

func (x *MissionResult) GetMissionID() int32 {
	if x != nil {
		return x.MissionID
	}
	return 0
}

func (x *MissionResult) GetMissionHash() string {
	if x != nil {
		return x.MissionHash
	}
	return ""
}

func (x *MissionResult) GetScene() *Scene {
	if x != nil {
		return x.Scene
	}
	return nil
}

// UserDeviceInfo
type UserDeviceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserHash        string `protobuf:"bytes,1,opt,name=UserHash,proto3" json:"UserHash,omitempty"`               // user hash
	Game            string `protobuf:"bytes,2,opt,name=Game,proto3" json:"Game,omitempty"`                       // game code
	Platform        string `protobuf:"bytes,3,opt,name=Platform,proto3" json:"Platform,omitempty"`               // ios or android
	ADID            string `protobuf:"bytes,4,opt,name=ADID,proto3" json:"ADID,omitempty"`                       // ad id
	GUID            string `protobuf:"bytes,5,opt,name=GUID,proto3" json:"GUID,omitempty"`                       // guid
	PlatformInfo    string `protobuf:"bytes,6,opt,name=PlatformInfo,proto3" json:"PlatformInfo,omitempty"`       // platform infomation
	GameVersion     string `protobuf:"bytes,7,opt,name=GameVersion,proto3" json:"GameVersion,omitempty"`         // game version
	ResourceVersion string `protobuf:"bytes,8,opt,name=ResourceVersion,proto3" json:"ResourceVersion,omitempty"` // game resource version
	DeviceInfo      string `protobuf:"bytes,9,opt,name=DeviceInfo,proto3" json:"DeviceInfo,omitempty"`           // device infomation
	CreateTs        int64  `protobuf:"varint,10,opt,name=CreateTs,proto3" json:"CreateTs,omitempty"`             // create timestamp
	LoginTimes      int32  `protobuf:"varint,11,opt,name=LoginTimes,proto3" json:"LoginTimes,omitempty"`         // login times
	LastLoginTs     int64  `protobuf:"varint,12,opt,name=LastLoginTs,proto3" json:"LastLoginTs,omitempty"`       // last login timestamp
}

func (x *UserDeviceInfo) Reset() {
	*x = UserDeviceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block7_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDeviceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeviceInfo) ProtoMessage() {}

func (x *UserDeviceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_block7_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeviceInfo.ProtoReflect.Descriptor instead.
func (*UserDeviceInfo) Descriptor() ([]byte, []int) {
	return file_block7_proto_rawDescGZIP(), []int{6}
}

func (x *UserDeviceInfo) GetUserHash() string {
	if x != nil {
		return x.UserHash
	}
	return ""
}

func (x *UserDeviceInfo) GetGame() string {
	if x != nil {
		return x.Game
	}
	return ""
}

func (x *UserDeviceInfo) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *UserDeviceInfo) GetADID() string {
	if x != nil {
		return x.ADID
	}
	return ""
}

func (x *UserDeviceInfo) GetGUID() string {
	if x != nil {
		return x.GUID
	}
	return ""
}

func (x *UserDeviceInfo) GetPlatformInfo() string {
	if x != nil {
		return x.PlatformInfo
	}
	return ""
}

func (x *UserDeviceInfo) GetGameVersion() string {
	if x != nil {
		return x.GameVersion
	}
	return ""
}

func (x *UserDeviceInfo) GetResourceVersion() string {
	if x != nil {
		return x.ResourceVersion
	}
	return ""
}

func (x *UserDeviceInfo) GetDeviceInfo() string {
	if x != nil {
		return x.DeviceInfo
	}
	return ""
}

func (x *UserDeviceInfo) GetCreateTs() int64 {
	if x != nil {
		return x.CreateTs
	}
	return 0
}

func (x *UserDeviceInfo) GetLoginTimes() int32 {
	if x != nil {
		return x.LoginTimes
	}
	return 0
}

func (x *UserDeviceInfo) GetLastLoginTs() int64 {
	if x != nil {
		return x.LastLoginTs
	}
	return 0
}

// UserInfo
type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64             `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Data   []*UserDeviceInfo `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty"` // user device infomations
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block7_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_block7_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_block7_proto_rawDescGZIP(), []int{7}
}

func (x *UserInfo) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UserInfo) GetData() []*UserDeviceInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

// LoginParams
type LoginParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserHash        string `protobuf:"bytes,1,opt,name=UserHash,proto3" json:"UserHash,omitempty"`               // user hash
	Game            string `protobuf:"bytes,2,opt,name=Game,proto3" json:"Game,omitempty"`                       // game code
	Platform        string `protobuf:"bytes,3,opt,name=Platform,proto3" json:"Platform,omitempty"`               // ios or android
	ADID            string `protobuf:"bytes,4,opt,name=ADID,proto3" json:"ADID,omitempty"`                       // ad id
	GUID            string `protobuf:"bytes,5,opt,name=GUID,proto3" json:"GUID,omitempty"`                       // guid
	PlatformInfo    string `protobuf:"bytes,6,opt,name=PlatformInfo,proto3" json:"PlatformInfo,omitempty"`       // platform infomation
	GameVersion     string `protobuf:"bytes,7,opt,name=GameVersion,proto3" json:"GameVersion,omitempty"`         // game version
	ResourceVersion string `protobuf:"bytes,8,opt,name=ResourceVersion,proto3" json:"ResourceVersion,omitempty"` // game resource version
	DeviceInfo      string `protobuf:"bytes,9,opt,name=DeviceInfo,proto3" json:"DeviceInfo,omitempty"`           // device infomation
}

func (x *LoginParams) Reset() {
	*x = LoginParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block7_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginParams) ProtoMessage() {}

func (x *LoginParams) ProtoReflect() protoreflect.Message {
	mi := &file_block7_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginParams.ProtoReflect.Descriptor instead.
func (*LoginParams) Descriptor() ([]byte, []int) {
	return file_block7_proto_rawDescGZIP(), []int{8}
}

func (x *LoginParams) GetUserHash() string {
	if x != nil {
		return x.UserHash
	}
	return ""
}

func (x *LoginParams) GetGame() string {
	if x != nil {
		return x.Game
	}
	return ""
}

func (x *LoginParams) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *LoginParams) GetADID() string {
	if x != nil {
		return x.ADID
	}
	return ""
}

func (x *LoginParams) GetGUID() string {
	if x != nil {
		return x.GUID
	}
	return ""
}

func (x *LoginParams) GetPlatformInfo() string {
	if x != nil {
		return x.PlatformInfo
	}
	return ""
}

func (x *LoginParams) GetGameVersion() string {
	if x != nil {
		return x.GameVersion
	}
	return ""
}

func (x *LoginParams) GetResourceVersion() string {
	if x != nil {
		return x.ResourceVersion
	}
	return ""
}

func (x *LoginParams) GetDeviceInfo() string {
	if x != nil {
		return x.DeviceInfo
	}
	return ""
}

// LoginResult
type LoginResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserHash string `protobuf:"bytes,1,opt,name=UserHash,proto3" json:"UserHash,omitempty"` // user hash
	UserID   int64  `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`    // user id
}

func (x *LoginResult) Reset() {
	*x = LoginResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block7_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResult) ProtoMessage() {}

func (x *LoginResult) ProtoReflect() protoreflect.Message {
	mi := &file_block7_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResult.ProtoReflect.Descriptor instead.
func (*LoginResult) Descriptor() ([]byte, []int) {
	return file_block7_proto_rawDescGZIP(), []int{9}
}

func (x *LoginResult) GetUserHash() string {
	if x != nil {
		return x.UserHash
	}
	return ""
}

func (x *LoginResult) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

var File_block7_proto protoreflect.FileDescriptor

var file_block7_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x37, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x37, 0x70, 0x62, 0x22, 0x20, 0x0a, 0x06, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x0a, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x37, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x22, 0x74, 0x0a, 0x0c, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xc9, 0x05, 0x0a, 0x05, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x58, 0x4f, 0x66, 0x66,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x58, 0x4f, 0x66, 0x66, 0x12, 0x12, 0x0a, 0x04,
	0x59, 0x4f, 0x66, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x59, 0x4f, 0x66, 0x66,
	0x12, 0x32, 0x0a, 0x07, 0x49, 0x6e, 0x69, 0x74, 0x41, 0x72, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x37, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07, 0x49, 0x6e, 0x69,
	0x74, 0x41, 0x72, 0x72, 0x12, 0x2e, 0x0a, 0x07, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x37, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0b, 0x43,
	0x6c, 0x69, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0b, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x50, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x0b, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x50, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x0d, 0x49, 0x73, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x0d, 0x49, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x05,
	0x4d, 0x61, 0x70, 0x49, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x05, 0x4d, 0x61, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x4d, 0x61, 0x70, 0x49, 0x44, 0x32, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x4d, 0x61, 0x70, 0x49, 0x44, 0x32, 0x12, 0x3a, 0x0a, 0x0c, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x61, 0x6c, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x37, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61,
	0x6c, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x0c, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x4c,
	0x61, 0x79, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x69, 0x74, 0x41, 0x72, 0x72, 0x32,
	0x18, 0x14, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x49, 0x6e, 0x69, 0x74, 0x41, 0x72, 0x72, 0x32,
	0x12, 0x1a, 0x0a, 0x08, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x32, 0x18, 0x15, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x08, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x32, 0x12, 0x18, 0x0a, 0x07,
	0x52, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x16, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x52,
	0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x22, 0x2d, 0x0a, 0x0d, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x22, 0x76, 0x0a, 0x0d, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x25, 0x0a, 0x05, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x37, 0x70, 0x62, 0x2e, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x52, 0x05, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x22, 0xf2, 0x02, 0x0a, 0x0e,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x47, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x44,
	0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x44, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x47, 0x55, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x47, 0x55,
	0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x47, 0x61, 0x6d,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x73, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x73,
	0x22, 0x50, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x37, 0x70, 0x62, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x91, 0x02, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x47, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x12,
	0x0a, 0x04, 0x41, 0x44, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x44,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x47, 0x55, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x47, 0x55, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x61,
	0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x47, 0x61, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x41, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x73, 0x30, 0x30, 0x37, 0x2f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x37, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x37, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_block7_proto_rawDescOnce sync.Once
	file_block7_proto_rawDescData = file_block7_proto_rawDesc
)

func file_block7_proto_rawDescGZIP() []byte {
	file_block7_proto_rawDescOnce.Do(func() {
		file_block7_proto_rawDescData = protoimpl.X.CompressGZIP(file_block7_proto_rawDescData)
	})
	return file_block7_proto_rawDescData
}

var file_block7_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_block7_proto_goTypes = []interface{}{
	(*Column)(nil),         // 0: block7pb.Column
	(*SceneLayer)(nil),     // 1: block7pb.SceneLayer
	(*SpecialLayer)(nil),   // 2: block7pb.SpecialLayer
	(*Scene)(nil),          // 3: block7pb.Scene
	(*MissionParams)(nil),  // 4: block7pb.MissionParams
	(*MissionResult)(nil),  // 5: block7pb.MissionResult
	(*UserDeviceInfo)(nil), // 6: block7pb.UserDeviceInfo
	(*UserInfo)(nil),       // 7: block7pb.UserInfo
	(*LoginParams)(nil),    // 8: block7pb.LoginParams
	(*LoginResult)(nil),    // 9: block7pb.LoginResult
}
var file_block7_proto_depIdxs = []int32{
	0, // 0: block7pb.SceneLayer.values:type_name -> block7pb.Column
	1, // 1: block7pb.Scene.InitArr:type_name -> block7pb.SceneLayer
	0, // 2: block7pb.Scene.History:type_name -> block7pb.Column
	2, // 3: block7pb.Scene.specialLayer:type_name -> block7pb.SpecialLayer
	3, // 4: block7pb.MissionResult.Scene:type_name -> block7pb.Scene
	6, // 5: block7pb.UserInfo.Data:type_name -> block7pb.UserDeviceInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_block7_proto_init() }
func file_block7_proto_init() {
	if File_block7_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_block7_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Column); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_block7_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneLayer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_block7_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecialLayer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_block7_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scene); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_block7_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MissionParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_block7_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MissionResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_block7_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDeviceInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_block7_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_block7_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_block7_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_block7_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_block7_proto_goTypes,
		DependencyIndexes: file_block7_proto_depIdxs,
		MessageInfos:      file_block7_proto_msgTypes,
	}.Build()
	File_block7_proto = out.File
	file_block7_proto_rawDesc = nil
	file_block7_proto_goTypes = nil
	file_block7_proto_depIdxs = nil
}
