package structpbbson

import (
	"google.golang.org/protobuf/types/known/structpb"
	"reflect"
)

var (
	ProtoStructType       = reflect.TypeOf(structpb.Struct{})
	ProtoValueType        = reflect.TypeOf(structpb.Value{})
	ProtoListValueType    = reflect.TypeOf(structpb.ListValue{})
	ProtoValueStructType  = reflect.TypeOf(structpb.Value_StructValue{})
	ProtoValueNumberType  = reflect.TypeOf(structpb.Value_NumberValue{})
	ProtoValueStringType  = reflect.TypeOf(structpb.Value_StringValue{})
	ProtoValueBoolType    = reflect.TypeOf(structpb.Value_BoolValue{})
	ProtoValueNullType    = reflect.TypeOf(structpb.Value_NullValue{})
	ProtoValueListType    = reflect.TypeOf(structpb.Value_ListValue{})
	ProtoValueNullPtrType = reflect.TypeOf(&structpb.Value_NullValue{})
)
