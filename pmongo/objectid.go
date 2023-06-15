package pmongo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NewObjectId creates proto ObjectId from MongoDB ObjectID
func NewObjectId(id primitive.ObjectID) *ObjectId {
	return &ObjectId{Value: id.Hex()}
}

// GetObjectID returns MongoDB object ID
func (o *ObjectId) GetObjectID() (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(o.Value)
}

// NewObjectIDFromInterface creates proto ObjectId from MongoDB ObjectID as an interface
func NewObjectIDFromInterface(idIf interface{}) (*ObjectId, error) {
	var id ObjectId
	if idVal, ok := idIf.(ObjectId); ok {
		id = idVal
	} else {
		if idVal, ok := idIf.(primitive.ObjectID); ok {
			nid := NewObjectId(idVal)
			id = *nid
		} else {
			return nil, fmt.Errorf("error converting object id to ObjectId")
		}
	}
	return &id, nil
}
