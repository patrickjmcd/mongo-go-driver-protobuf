package codecs

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pepper-iot/mongo-go-driver-protobuf/structpbbson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/pepper-iot/mongo-go-driver-protobuf/pmongo"
)

var (
	// Protobuf wrappers types
	boolValueType      = reflect.TypeOf(wrappers.BoolValue{})
	bytesValueType     = reflect.TypeOf(wrappers.BytesValue{})
	doubleValueType    = reflect.TypeOf(wrappers.DoubleValue{})
	floatValueType     = reflect.TypeOf(wrappers.FloatValue{})
	int32ValueType     = reflect.TypeOf(wrappers.Int32Value{})
	int64ValueType     = reflect.TypeOf(wrappers.Int64Value{})
	stringValueType    = reflect.TypeOf(wrappers.StringValue{})
	uint32ValueType    = reflect.TypeOf(wrappers.UInt32Value{})
	uint64ValueType    = reflect.TypeOf(wrappers.UInt64Value{})
	nullValueType      = reflect.TypeOf(bson.TypeNull)
	undefinedValueType = reflect.TypeOf(bson.TypeUndefined)

	// Protobuf Timestamp type
	timestampType = reflect.TypeOf(timestamp.Timestamp{})

	// Time type
	timeType = reflect.TypeOf(time.Time{})

	// ObjectId type
	objectIDType          = reflect.TypeOf(pmongo.ObjectId{})
	objectIDPrimitiveType = reflect.TypeOf(primitive.ObjectID{})
	objectIDPointerType   = reflect.TypeOf(&pmongo.ObjectId{})

	// Codecs
	wrapperValueCodecRef    = &wrapperValueCodec{}
	timestampCodecRef       = &timestampCodec{}
	objectIDCodecRef        = &objectIDCodec{}
	objectIDPointerCodecRef = &objectIDPointerCodec{}
)

// wrapperValueCodec is codec for Protobuf type wrappers
type wrapperValueCodec struct{}

// EncodeValue encodes Protobuf type wrapper value to BSON value
func (e *wrapperValueCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	val = val.FieldByName("Value")
	enc, err := ectx.LookupEncoder(val.Type())
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, val)
}

// DecodeValue decodes BSON value to Protobuf type wrapper value
func (e *wrapperValueCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	val = val.FieldByName("Value")
	enc, err := ectx.LookupDecoder(val.Type())
	if err != nil {
		return err
	}
	return enc.DecodeValue(ectx, vr, val)
}

// timestampCodec is codec for Protobuf Timestamp
type timestampCodec struct{}

// EncodeValue encodes Protobuf Timestamp value to BSON value
func (e *timestampCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	v := val.Interface().(timestamp.Timestamp)
	t, err := ptypes.Timestamp(&v)
	if err != nil {
		return err
	}
	enc, err := ectx.LookupEncoder(timeType)
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, reflect.ValueOf(t.In(time.UTC)))
}

// DecodeValue decodes BSON value to Timestamp value
func (e *timestampCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	enc, err := ectx.LookupDecoder(timeType)
	if err != nil {
		return err
	}
	var t time.Time
	if err = enc.DecodeValue(ectx, vr, reflect.ValueOf(&t).Elem()); err != nil {
		return err
	}
	ts, err := ptypes.TimestampProto(t.In(time.UTC))
	if err != nil {
		return err
	}
	val.Set(reflect.ValueOf(*ts))
	return nil
}

// objectIDCodec is codec for Protobuf ObjectId
type objectIDCodec struct{}

// EncodeValue encodes Protobuf ObjectId value to BSON value
func (e *objectIDCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	v := val.Interface().(pmongo.ObjectId)
	// Create primitive.ObjectId from string
	id, err := primitive.ObjectIDFromHex(v.Value)
	if err != nil {
		return err
	}
	enc, err := ectx.LookupEncoder(objectIDPrimitiveType)
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, reflect.ValueOf(id))
}

// DecodeValue decodes BSON value to ObjectId value
func (e *objectIDCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	enc, err := ectx.LookupDecoder(objectIDPrimitiveType)
	if err != nil {
		return err
	}
	var id primitive.ObjectID
	if err = enc.DecodeValue(ectx, vr, reflect.ValueOf(&id).Elem()); err != nil {
		return err
	}
	oid := *pmongo.NewObjectId(id)
	if err != nil {
		return err
	}
	val.Set(reflect.ValueOf(oid))
	return nil
}

// objectIDCodec is codec for Protobuf ObjectId
type objectIDPointerCodec struct{}

// EncodeValue encodes Protobuf ObjectId value to BSON value
func (e *objectIDPointerCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	v := val.Interface().(*pmongo.ObjectId)
	// Create primitive.ObjectId from string
	id, err := primitive.ObjectIDFromHex(v.Value)
	if err != nil {
		return err
	}
	enc, err := ectx.LookupEncoder(objectIDPrimitiveType)
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, reflect.ValueOf(id))
}

// DecodeValue decodes BSON value to ObjectId value
func (e *objectIDPointerCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	enc, err := ectx.LookupDecoder(objectIDPrimitiveType)
	if err != nil {
		return err
	}
	var id primitive.ObjectID
	if err = enc.DecodeValue(ectx, vr, reflect.ValueOf(&id).Elem()); err != nil {
		return err
	}
	oid := pmongo.NewObjectId(id)
	if err != nil {
		return err
	}
	val.Set(reflect.ValueOf(oid))
	return nil
}

func RegisterRegistry(rb *bsoncodec.Registry) *bsoncodec.Registry {
	// Types
	rb.RegisterTypeMapEntry(bson.TypeObjectID, objectIDType)

	// Decoders
	rb.RegisterTypeDecoder(boolValueType, wrapperValueCodecRef)
	rb.RegisterTypeDecoder(bytesValueType, wrapperValueCodecRef)
	rb.RegisterTypeDecoder(doubleValueType, wrapperValueCodecRef)
	rb.RegisterTypeDecoder(floatValueType, wrapperValueCodecRef)
	rb.RegisterTypeDecoder(int32ValueType, wrapperValueCodecRef)
	rb.RegisterTypeDecoder(int64ValueType, wrapperValueCodecRef)
	rb.RegisterTypeDecoder(stringValueType, wrapperValueCodecRef)
	rb.RegisterTypeDecoder(uint32ValueType, wrapperValueCodecRef)
	rb.RegisterTypeDecoder(uint64ValueType, wrapperValueCodecRef)
	rb.RegisterTypeDecoder(timestampType, timestampCodecRef)
	rb.RegisterTypeDecoder(objectIDType, objectIDCodecRef)
	rb.RegisterTypeDecoder(objectIDPointerType, objectIDPointerCodecRef)
	rb.RegisterTypeDecoder(structpbbson.ProtoStructType, structpbbson.StructCodec{})
	rb.RegisterTypeDecoder(structpbbson.ProtoValueType, structpbbson.ValueCodec{})
	rb.RegisterTypeDecoder(structpbbson.ProtoListValueType, structpbbson.ListCodec{})
	rb.RegisterTypeDecoder(structpbbson.ProtoValueNullType, structpbbson.ValueCodec{})
	rb.RegisterTypeDecoder(nullValueType, structpbbson.ValueCodec{})
	rb.RegisterTypeDecoder(undefinedValueType, structpbbson.ValueCodec{})

	// Encoders
	rb.RegisterTypeEncoder(boolValueType, wrapperValueCodecRef)
	rb.RegisterTypeEncoder(bytesValueType, wrapperValueCodecRef)
	rb.RegisterTypeEncoder(doubleValueType, wrapperValueCodecRef)
	rb.RegisterTypeEncoder(floatValueType, wrapperValueCodecRef)
	rb.RegisterTypeEncoder(int32ValueType, wrapperValueCodecRef)
	rb.RegisterTypeEncoder(int64ValueType, wrapperValueCodecRef)
	rb.RegisterTypeEncoder(stringValueType, wrapperValueCodecRef)
	rb.RegisterTypeEncoder(uint32ValueType, wrapperValueCodecRef)
	rb.RegisterTypeEncoder(uint64ValueType, wrapperValueCodecRef)
	rb.RegisterTypeEncoder(timestampType, timestampCodecRef)
	rb.RegisterTypeEncoder(objectIDType, objectIDCodecRef)
	rb.RegisterTypeEncoder(objectIDPointerType, objectIDPointerCodecRef)
	rb.RegisterTypeEncoder(structpbbson.ProtoStructType, structpbbson.StructCodec{})
	rb.RegisterTypeEncoder(structpbbson.ProtoValueType, structpbbson.ValueCodec{})
	rb.RegisterTypeEncoder(structpbbson.ProtoListValueType, structpbbson.ListCodec{})
	rb.RegisterTypeEncoder(structpbbson.ProtoValueNullType, structpbbson.ValueCodec{})
	rb.RegisterTypeEncoder(nullValueType, structpbbson.ValueCodec{})
	rb.RegisterTypeEncoder(undefinedValueType, structpbbson.ValueCodec{})

	return rb

}
