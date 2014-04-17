// Code generated by protoc-gen-go.
// source: search.proto
// DO NOT EDIT!

package main

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// 检索命令
type SearchRequest struct {
	Query            *string `protobuf:"bytes,1,req,name=query" json:"query,omitempty"`
	Pn               *int32  `protobuf:"varint,2,opt,name=pn,def=0" json:"pn,omitempty"`
	Rn               *int32  `protobuf:"varint,3,opt,name=rn,def=10" json:"rn,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}

const Default_SearchRequest_Pn int32 = 0
const Default_SearchRequest_Rn int32 = 10

func (m *SearchRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *SearchRequest) GetPn() int32 {
	if m != nil && m.Pn != nil {
		return *m.Pn
	}
	return Default_SearchRequest_Pn
}

func (m *SearchRequest) GetRn() int32 {
	if m != nil && m.Rn != nil {
		return *m.Rn
	}
	return Default_SearchRequest_Rn
}

// 检索结果
type SearchResponse struct {
	RetNum           *int32                  `protobuf:"varint,1,req" json:"RetNum,omitempty"`
	DispNum          *int32                  `protobuf:"varint,2,req" json:"DispNum,omitempty"`
	Result           []*SearchResponseOneRes `protobuf:"bytes,3,rep" json:"Result,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}

func (m *SearchResponse) GetRetNum() int32 {
	if m != nil && m.RetNum != nil {
		return *m.RetNum
	}
	return 0
}

func (m *SearchResponse) GetDispNum() int32 {
	if m != nil && m.DispNum != nil {
		return *m.DispNum
	}
	return 0
}

func (m *SearchResponse) GetResult() []*SearchResponseOneRes {
	if m != nil {
		return m.Result
	}
	return nil
}

type SearchResponseOneRes struct {
	Data             []byte `protobuf:"bytes,1,req,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SearchResponseOneRes) Reset()         { *m = SearchResponseOneRes{} }
func (m *SearchResponseOneRes) String() string { return proto.CompactTextString(m) }
func (*SearchResponseOneRes) ProtoMessage()    {}

func (m *SearchResponseOneRes) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
}
