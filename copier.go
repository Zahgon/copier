package copier

import (
	"database/sql/driver"
	"reflect"
	"sync"
)

const (
	tagMust uint8 = 1 << iota

	tagNoPanic

	tagIgnore

	tagOverride

	hasCopied

	String  string  = ""
	Bool    bool    = false
	Int     int     = 0
	Float32 float32 = 0
	Float64 float64 = 0
)

type Option struct {
	IgnoreEmpty   bool
	CaseSensitive bool
	DeepCopy      bool
	Converters    []TypeConverter

	FieldNameMapping []FieldNameMapping

	Must bool

	NoPanic bool
}

func (opt Option) converters() map[converterPair]TypeConverter {
	_ = "STUB: not implemented"
	return nil
}

type TypeConverter struct {
	SrcType interface{}
	DstType interface{}
	Fn      func(src interface{}) (dst interface{}, err error)
}

type converterPair struct {
	SrcType reflect.Type
	DstType reflect.Type
}

func (opt Option) fieldNameMapping() map[converterPair]FieldNameMapping {
	_ = "STUB: not implemented"
	return nil
}

type FieldNameMapping struct {
	SrcType interface{}
	DstType interface{}
	Mapping map[string]string
}

type flags struct {
	BitFlags  map[string]uint8
	SrcNames  tagNameMapping
	DestNames tagNameMapping
}

type tagNameMapping struct {
	FieldNameToTag map[string]string
	TagToFieldName map[string]string
}

func Copy(toValue interface{}, fromValue interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func CopyWithOption(toValue interface{}, fromValue interface{}, opt Option) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func copier(toValue interface{}, fromValue interface{}, opt Option) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func getFieldNamesMapping(mappings map[converterPair]FieldNameMapping, fromType reflect.Type, toType reflect.Type) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func fieldByNameOrZeroValue(source reflect.Value, fieldName string) (value reflect.Value) {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func copyUnexportedStructFields(to, from reflect.Value) { _ = "STUB: not implemented"; return }

func shouldIgnore(v reflect.Value, bitFlags uint8, ignoreEmpty bool) bool {
	_ = "STUB: not implemented"
	return false
}

var deepFieldsLock sync.RWMutex
var deepFieldsMap = make(map[reflect.Type][]reflect.StructField)

func deepFields(reflectType reflect.Type) []reflect.StructField {
	_ = "STUB: not implemented"
	return nil
}

func indirect(reflectValue reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func indirectType(reflectType reflect.Type) (_ reflect.Type, isPtr bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Type), false
}

func set(to, from reflect.Value, deepCopy bool, converters map[converterPair]TypeConverter) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func lookupAndCopyWithConverter(to, from reflect.Value, converters map[converterPair]TypeConverter) (copied bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func parseTags(tag string) (flg uint8, name string, err error) {
	_ = "STUB: not implemented"
	return 0, "", nil
}

func getFlags(dest, src reflect.Value, toType, fromType reflect.Type, opt Option) (flags, error) {
	_ = "STUB: not implemented"
	return *new(flags), nil
}

func checkBitFlags(flagsList map[string]uint8) (err error) { _ = "STUB: not implemented"; return nil }

func getFieldName(fieldName string, flgs flags, fieldNameMapping map[string]string) (srcFieldName string, destFieldName string) {
	_ = "STUB: not implemented"
	return "", ""
}

func driverValuer(v reflect.Value) (i driver.Valuer, ok bool) {
	_ = "STUB: not implemented"
	return *new(driver.Valuer), false
}

func fieldByName(v reflect.Value, name string, caseSensitive bool) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}
