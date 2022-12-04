package adsb

import (
	"log"
	"reflect"
	"strconv"
)

func fieldPlacementInfo(st interface{}, fieldName string) (fieldLength int, offset int) {
	structField, ok := reflect.TypeOf(st).FieldByName(fieldName)
	if !ok {
		log.Panicf("field %s does not exist under %v", fieldName, st)
	}

	fieldLength, err := strconv.Atoi(structField.Tag.Get("adsbFieldLength"))
	if err != nil {
		panic(err)
	}

	fieldOffset, err := strconv.Atoi(structField.Tag.Get("adsbFieldOffset"))
	if err != nil {
		panic(err)
	}

	return fieldLength, fieldOffset
}

func applyFieldToPayload(payload *uint64, stp interface{}, fieldName string) {
	// dereference the supplied struct pointer
	st := reflect.ValueOf(stp).Elem().Interface()

	length, offset := fieldPlacementInfo(st, fieldName)
	fieldValue := reflect.ValueOf(st).FieldByName(fieldName).Interface()
	var fieldValueInt uint64 = 0

	switch v := fieldValue.(type) {
	case uint32:
		fieldValueInt = uint64(v)
	case uint16:
		fieldValueInt = uint64(v)
	case byte:
		fieldValueInt = uint64(v)
	case bool:
		if v {
			fieldValueInt = 1
		} else {
			fieldValueInt = 0
		}
	default:
		log.Panicf("unable to convert %T(%+v) to uint64", fieldValue, v)
	}

	*payload = *payload | fieldValueInt<<(64-length-offset)
}

func applyPayloadFieldToStruct(payload *uint64, stp interface{}, fieldName string) {

	// dereference the supplied struct pointer
	st := reflect.ValueOf(stp).Elem().Interface()

	// figure out where the data we want to extract exists
	length, offset := fieldPlacementInfo(st, fieldName)

	// pull it out into v
	v := (*payload) >> (64 - uint64(offset) - uint64(length))
	v = v & lengthToBitmask(uint(length))

	// grab the field of the struct we want to set
	structField := reflect.ValueOf(stp).Elem().FieldByName(fieldName)

	// get the type of the field we want to set
	fieldValue := reflect.ValueOf(st).FieldByName(fieldName).Interface()

	// set the value in the struct
	switch fieldValue.(type) {
	case uint32:
		structField.Set(reflect.ValueOf(uint32(v)))
	case uint16:
		structField.Set(reflect.ValueOf(uint16(v)))
	case bool:
		if v == 1 {
			structField.SetBool(true)
		} else if v == 0 {
			structField.SetBool(false)
		} else {
			panic("this should never happen")
		}
	case uint64:
		structField.SetUint(v)
	case byte:
		structField.Set(reflect.ValueOf(byte(v)))
	default:
		log.Panicf("unable to convert %T(%+v) to uint64", fieldValue, fieldValue)
	}

}

func lengthToBitmask(length uint) uint64 {
	var ret uint64 = 0
	for i := uint(0); i < length; i++ {
		ret = (ret << 1) | 0b1
	}
	return ret
}

func byteInSlice(i byte, s []byte) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false

}
