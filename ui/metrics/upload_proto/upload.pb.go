// Copyright 2020 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.9.1
// source: upload.proto

package upload_proto

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

type Upload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The timestamp in milliseconds that the build was created.
	CreationTimestampMs *uint64 `protobuf:"varint,1,opt,name=creation_timestamp_ms,json=creationTimestampMs" json:"creation_timestamp_ms,omitempty"`
	// The timestamp in milliseconds when the build was completed.
	CompletionTimestampMs *uint64 `protobuf:"varint,2,opt,name=completion_timestamp_ms,json=completionTimestampMs" json:"completion_timestamp_ms,omitempty"`
	// The branch name.
	BranchName *string `protobuf:"bytes,3,opt,name=branch_name,json=branchName" json:"branch_name,omitempty"`
	// The target name.
	TargetName *string `protobuf:"bytes,4,opt,name=target_name,json=targetName" json:"target_name,omitempty"`
	// A list of metrics filepaths to upload.
	MetricsFiles []string `protobuf:"bytes,5,rep,name=metrics_files,json=metricsFiles" json:"metrics_files,omitempty"`
	// A list of directories to delete after the copy of metrics files
	// is completed for uploading.
	DirectoriesToDelete []string `protobuf:"bytes,6,rep,name=directories_to_delete,json=directoriesToDelete" json:"directories_to_delete,omitempty"`
}

func (x *Upload) Reset() {
	*x = Upload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Upload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Upload) ProtoMessage() {}

func (x *Upload) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Upload.ProtoReflect.Descriptor instead.
func (*Upload) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{0}
}

func (x *Upload) GetCreationTimestampMs() uint64 {
	if x != nil && x.CreationTimestampMs != nil {
		return *x.CreationTimestampMs
	}
	return 0
}

func (x *Upload) GetCompletionTimestampMs() uint64 {
	if x != nil && x.CompletionTimestampMs != nil {
		return *x.CompletionTimestampMs
	}
	return 0
}

func (x *Upload) GetBranchName() string {
	if x != nil && x.BranchName != nil {
		return *x.BranchName
	}
	return ""
}

func (x *Upload) GetTargetName() string {
	if x != nil && x.TargetName != nil {
		return *x.TargetName
	}
	return ""
}

func (x *Upload) GetMetricsFiles() []string {
	if x != nil {
		return x.MetricsFiles
	}
	return nil
}

func (x *Upload) GetDirectoriesToDelete() []string {
	if x != nil {
		return x.DirectoriesToDelete
	}
	return nil
}

var File_upload_proto protoreflect.FileDescriptor

var file_upload_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14,
	0x73, 0x6f, 0x6f, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x8f, 0x02, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x32, 0x0a, 0x15, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x4d, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x13, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x54, 0x6f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69,
	0x64, 0x2f, 0x73, 0x6f, 0x6f, 0x6e, 0x67, 0x2f, 0x75, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_upload_proto_rawDescOnce sync.Once
	file_upload_proto_rawDescData = file_upload_proto_rawDesc
)

func file_upload_proto_rawDescGZIP() []byte {
	file_upload_proto_rawDescOnce.Do(func() {
		file_upload_proto_rawDescData = protoimpl.X.CompressGZIP(file_upload_proto_rawDescData)
	})
	return file_upload_proto_rawDescData
}

var file_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_upload_proto_goTypes = []interface{}{
	(*Upload)(nil), // 0: soong_metrics_upload.Upload
}
var file_upload_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_upload_proto_init() }
func file_upload_proto_init() {
	if File_upload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_upload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Upload); i {
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
			RawDescriptor: file_upload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_upload_proto_goTypes,
		DependencyIndexes: file_upload_proto_depIdxs,
		MessageInfos:      file_upload_proto_msgTypes,
	}.Build()
	File_upload_proto = out.File
	file_upload_proto_rawDesc = nil
	file_upload_proto_goTypes = nil
	file_upload_proto_depIdxs = nil
}
