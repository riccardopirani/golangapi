package mongo

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

// BuildDocument is a wrapper for constructing a mongo document
// e.g. {"$sort": {"name": 1}}
//       <-key->  <- value ->
func BuildDocument(key string, value interface{}) bson.D {
	return bson.D{{Key: key, Value: value}}
}

// Build Element is a wrapper for constructing part of an object
// e.g. {"$sort": {"name":      1    }}
//                 <-key->  <-value->
func BuildElement(key string, value interface{}) bson.E {
	return bson.E{Key: key, Value: value}
}

func BuildArray(items interface{}) (arr bson.A, exist bool) {
	val := reflect.ValueOf(items)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		return
	}

	arrLength := val.Len()
	if arrLength > 0 {
		exist = true
	}

	for i := 0; i < arrLength; i++ {
		arr = append(arr, val.Index(i).Interface())
	}
	return
}

func BuildLookupByIDStage(field, fromCol string) bson.D {
	return bson.D{
		{Key: StageLookup, Value: bson.D{
			{Key: MetaFrom, Value: fromCol},
			{Key: MetaLocalField, Value: field},
			{Key: MetaForeignField, Value: "_id"},
			{Key: MetaAs, Value: field},
		},
		},
	}
}

func BuildUnwindStage(field string) bson.D {
	return bson.D{{Key: StageUnwind, Value: bson.D{
		{Key: MetaPath, Value: "$" + field},
		{Key: MetaPreserveNullAndEmptyArrays, Value: true},
	}}}
}
