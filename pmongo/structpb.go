package pmongo

import structpb "google.golang.org/protobuf/types/known/structpb"

func InterfaceFromStructpb(s *structpb.Value) interface{} {
	switch v := s.Kind.(type) {
	case *structpb.Value_NullValue:
		return nil
	case *structpb.Value_NumberValue:
		return v.NumberValue
	case *structpb.Value_StringValue:
		return v.StringValue
	case *structpb.Value_BoolValue:
		return v.BoolValue
	case *structpb.Value_StructValue:
		return v.StructValue.AsMap()
	case *structpb.Value_ListValue:
		return v.ListValue.AsSlice()
	default:
		return nil
	}
}
