package blockchain

import (
	"log"
	"reflect"
)

// Unpack arguments conforming to model
func Unpack(s reflect.Value, unpackProfile reflect.Type) (out []interface{}) {
	//	unpackProfile := reflect.TypeOf(model)
	//	s := reflect.ValueOf(&inputParams).Elem()
	if unpackProfile.NumField() != (s.NumField() - 1) {
		panic("Model doesn't match number of argument of inputParams")
	}
	numField := unpackProfile.NumField()

	for i := 0; i < numField; i++ {
		f := s.Field(i + 1)
		log.Printf("%d: %s %s = %v\n", i,
			unpackProfile.Field(i).Type.Name(), f.Type(), f.Interface())
		switch unpackProfile.Field(i).Type.Name() {
		case "Address":
			out = append(out, f.Interface())
		case "String":
			out = append(out, f.Interface())
		case "Binary":
			out = append(out, f.Interface())
		case "Bool":
			out = append(out, f.Interface())
		case "BigInt":
			out = append(out, f.Interface())
		}
	}
	return nil
}
