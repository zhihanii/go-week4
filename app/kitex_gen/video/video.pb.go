// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.0
// source: cmd/video/idl/video.proto

package video

import (
	context "context"
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

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_idl_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_idl_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_cmd_video_idl_video_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResp) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BaseResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type FeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LatestTime int64 `protobuf:"varint,2,opt,name=latest_time,json=latestTime,proto3" json:"latest_time,omitempty"`
}

func (x *FeedRequest) Reset() {
	*x = FeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_idl_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedRequest) ProtoMessage() {}

func (x *FeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_idl_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedRequest.ProtoReflect.Descriptor instead.
func (*FeedRequest) Descriptor() ([]byte, []int) {
	return file_cmd_video_idl_video_proto_rawDescGZIP(), []int{1}
}

func (x *FeedRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FeedRequest) GetLatestTime() int64 {
	if x != nil {
		return x.LatestTime
	}
	return 0
}

type PubActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Title  string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *PubActionRequest) Reset() {
	*x = PubActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_idl_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubActionRequest) ProtoMessage() {}

func (x *PubActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_idl_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubActionRequest.ProtoReflect.Descriptor instead.
func (*PubActionRequest) Descriptor() ([]byte, []int) {
	return file_cmd_video_idl_video_proto_rawDescGZIP(), []int{2}
}

func (x *PubActionRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PubActionRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PubActionRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type PubListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AuthorId int64 `protobuf:"varint,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *PubListRequest) Reset() {
	*x = PubListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_idl_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubListRequest) ProtoMessage() {}

func (x *PubListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_idl_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubListRequest.ProtoReflect.Descriptor instead.
func (*PubListRequest) Descriptor() ([]byte, []int) {
	return file_cmd_video_idl_video_proto_rawDescGZIP(), []int{3}
}

func (x *PubListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PubListRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	FollowCount   int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3" json:"follow_count"`
	FollowerCount int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3" json:"follower_count"`
	IsFollow      bool   `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_idl_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_idl_video_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_cmd_video_idl_video_proto_rawDescGZIP(), []int{4}
}

func (x *Author) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *Author) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *Author) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	PlayUrl       string  `protobuf:"bytes,2,opt,name=play_url,json=playUrl,proto3" json:"play_url"`
	CoverUrl      string  `protobuf:"bytes,3,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url"`
	FavoriteCount int64   `protobuf:"varint,4,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count"`
	CommentCount  int64   `protobuf:"varint,5,opt,name=comment_count,json=commentCount,proto3" json:"comment_count"`
	IsFavorite    bool    `protobuf:"varint,6,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite"`
	Title         string  `protobuf:"bytes,7,opt,name=title,proto3" json:"title"`
	Author        *Author `protobuf:"bytes,8,opt,name=author,proto3" json:"author"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_idl_video_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_idl_video_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_cmd_video_idl_video_proto_rawDescGZIP(), []int{5}
}

func (x *Video) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Video) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type VideoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos []*Video `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos"`
}

func (x *VideoList) Reset() {
	*x = VideoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_idl_video_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoList) ProtoMessage() {}

func (x *VideoList) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_idl_video_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoList.ProtoReflect.Descriptor instead.
func (*VideoList) Descriptor() ([]byte, []int) {
	return file_cmd_video_idl_video_proto_rawDescGZIP(), []int{6}
}

func (x *VideoList) GetVideos() []*Video {
	if x != nil {
		return x.Videos
	}
	return nil
}

type FeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp  *BaseResp  `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp"`
	NextTime  int64      `protobuf:"varint,2,opt,name=next_time,json=nextTime,proto3" json:"next_time"`
	VideoList *VideoList `protobuf:"bytes,3,opt,name=video_list,json=videoList,proto3" json:"video_list"`
}

func (x *FeedResponse) Reset() {
	*x = FeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_idl_video_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedResponse) ProtoMessage() {}

func (x *FeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_idl_video_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedResponse.ProtoReflect.Descriptor instead.
func (*FeedResponse) Descriptor() ([]byte, []int) {
	return file_cmd_video_idl_video_proto_rawDescGZIP(), []int{7}
}

func (x *FeedResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *FeedResponse) GetNextTime() int64 {
	if x != nil {
		return x.NextTime
	}
	return 0
}

func (x *FeedResponse) GetVideoList() *VideoList {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type PubActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *PubActionResponse) Reset() {
	*x = PubActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_idl_video_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubActionResponse) ProtoMessage() {}

func (x *PubActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_idl_video_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubActionResponse.ProtoReflect.Descriptor instead.
func (*PubActionResponse) Descriptor() ([]byte, []int) {
	return file_cmd_video_idl_video_proto_rawDescGZIP(), []int{8}
}

func (x *PubActionResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type PubListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp  *BaseResp  `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	VideoList *VideoList `protobuf:"bytes,2,opt,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
}

func (x *PubListResponse) Reset() {
	*x = PubListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_video_idl_video_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubListResponse) ProtoMessage() {}

func (x *PubListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_video_idl_video_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubListResponse.ProtoReflect.Descriptor instead.
func (*PubListResponse) Descriptor() ([]byte, []int) {
	return file_cmd_video_idl_video_proto_rawDescGZIP(), []int{9}
}

func (x *PubListResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *PubListResponse) GetVideoList() *VideoList {
	if x != nil {
		return x.VideoList
	}
	return nil
}

var File_cmd_video_idl_video_proto protoreflect.FileDescriptor

var file_cmd_video_idl_video_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6d, 0x64, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x69, 0x64, 0x6c, 0x2f,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x22, 0x4a, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x47,
	0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x55, 0x0a, 0x10, 0x50, 0x75, 0x62, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x46,
	0x0a, 0x0e, 0x50, 0x75, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x22, 0xf9, 0x01, 0x0a,
	0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72,
	0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x25,
	0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73,
	0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x25, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x31, 0x0a, 0x09, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x0c,
	0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x09,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e,
	0x65, 0x78, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x09, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x11, 0x50, 0x75, 0x62, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a,
	0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x70, 0x0a, 0x0f, 0x50,
	0x75, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x0a,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xbf, 0x01,
	0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31,
	0x0a, 0x04, 0x46, 0x65, 0x65, 0x64, 0x12, 0x12, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x40, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e,
	0x50, 0x75, 0x62, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75,
	0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x1f, 0x5a, 0x1d, 0x67, 0x6f, 0x2d, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_video_idl_video_proto_rawDescOnce sync.Once
	file_cmd_video_idl_video_proto_rawDescData = file_cmd_video_idl_video_proto_rawDesc
)

func file_cmd_video_idl_video_proto_rawDescGZIP() []byte {
	file_cmd_video_idl_video_proto_rawDescOnce.Do(func() {
		file_cmd_video_idl_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_video_idl_video_proto_rawDescData)
	})
	return file_cmd_video_idl_video_proto_rawDescData
}

var file_cmd_video_idl_video_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_cmd_video_idl_video_proto_goTypes = []interface{}{
	(*BaseResp)(nil),          // 0: video.BaseResp
	(*FeedRequest)(nil),       // 1: video.FeedRequest
	(*PubActionRequest)(nil),  // 2: video.PubActionRequest
	(*PubListRequest)(nil),    // 3: video.PubListRequest
	(*Author)(nil),            // 4: video.Author
	(*Video)(nil),             // 5: video.Video
	(*VideoList)(nil),         // 6: video.VideoList
	(*FeedResponse)(nil),      // 7: video.FeedResponse
	(*PubActionResponse)(nil), // 8: video.PubActionResponse
	(*PubListResponse)(nil),   // 9: video.PubListResponse
}
var file_cmd_video_idl_video_proto_depIdxs = []int32{
	4,  // 0: video.Video.author:type_name -> video.Author
	5,  // 1: video.VideoList.videos:type_name -> video.Video
	0,  // 2: video.FeedResponse.base_resp:type_name -> video.BaseResp
	6,  // 3: video.FeedResponse.video_list:type_name -> video.VideoList
	0,  // 4: video.PubActionResponse.base_resp:type_name -> video.BaseResp
	0,  // 5: video.PubListResponse.base_resp:type_name -> video.BaseResp
	6,  // 6: video.PubListResponse.video_list:type_name -> video.VideoList
	1,  // 7: video.VideoService.Feed:input_type -> video.FeedRequest
	2,  // 8: video.VideoService.PubAction:input_type -> video.PubActionRequest
	3,  // 9: video.VideoService.PubList:input_type -> video.PubListRequest
	7,  // 10: video.VideoService.Feed:output_type -> video.FeedResponse
	8,  // 11: video.VideoService.PubAction:output_type -> video.PubActionResponse
	9,  // 12: video.VideoService.PubList:output_type -> video.PubListResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_cmd_video_idl_video_proto_init() }
func file_cmd_video_idl_video_proto_init() {
	if File_cmd_video_idl_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_video_idl_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
		file_cmd_video_idl_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedRequest); i {
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
		file_cmd_video_idl_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubActionRequest); i {
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
		file_cmd_video_idl_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubListRequest); i {
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
		file_cmd_video_idl_video_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_cmd_video_idl_video_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_cmd_video_idl_video_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoList); i {
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
		file_cmd_video_idl_video_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedResponse); i {
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
		file_cmd_video_idl_video_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubActionResponse); i {
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
		file_cmd_video_idl_video_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubListResponse); i {
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
			RawDescriptor: file_cmd_video_idl_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmd_video_idl_video_proto_goTypes,
		DependencyIndexes: file_cmd_video_idl_video_proto_depIdxs,
		MessageInfos:      file_cmd_video_idl_video_proto_msgTypes,
	}.Build()
	File_cmd_video_idl_video_proto = out.File
	file_cmd_video_idl_video_proto_rawDesc = nil
	file_cmd_video_idl_video_proto_goTypes = nil
	file_cmd_video_idl_video_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.3.1. DO NOT EDIT.

type VideoService interface {
	Feed(ctx context.Context, req *FeedRequest) (res *FeedResponse, err error)
	PubAction(ctx context.Context, req *PubActionRequest) (res *PubActionResponse, err error)
	PubList(ctx context.Context, req *PubListRequest) (res *PubListResponse, err error)
}
