// Code generated by Kitex v0.11.3. DO NOT EDIT.

package base

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/cloudwego/gopkg/protocol/thrift"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.STOP
)

func (p *TrafficEnv) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.BOOL {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_TrafficEnv[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *TrafficEnv) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field bool
	if v, l, err := thrift.Binary.ReadBool(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Open = _field
	return offset, nil
}

func (p *TrafficEnv) FastReadField2(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Env = _field
	return offset, nil
}

// for compatibility
func (p *TrafficEnv) FastWrite(buf []byte) int {
	return 0
}

func (p *TrafficEnv) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
		offset += p.fastWriteField2(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *TrafficEnv) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *TrafficEnv) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.BOOL, 1)
	offset += thrift.Binary.WriteBool(buf[offset:], p.Open)
	return offset
}

func (p *TrafficEnv) fastWriteField2(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 2)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.Env)
	return offset
}

func (p *TrafficEnv) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.BoolLength()
	return l
}

func (p *TrafficEnv) field2Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.Env)
	return l
}

func (p *Base) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.MAP {
				l, err = p.FastReadField6(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Base[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *Base) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.LogID = _field
	return offset, nil
}

func (p *Base) FastReadField2(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Caller = _field
	return offset, nil
}

func (p *Base) FastReadField3(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Addr = _field
	return offset, nil
}

func (p *Base) FastReadField4(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.Client = _field
	return offset, nil
}

func (p *Base) FastReadField5(buf []byte) (int, error) {
	offset := 0
	_field := NewTrafficEnv()
	if l, err := _field.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.TrafficEnv = _field
	return offset, nil
}

func (p *Base) FastReadField6(buf []byte) (int, error) {
	offset := 0

	_, _, size, l, err := thrift.Binary.ReadMapBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	_field := make(map[string]string, size)
	for i := 0; i < size; i++ {
		var _key string
		if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
			_key = v
		}

		var _val string
		if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
			_val = v
		}

		_field[_key] = _val
	}
	p.Extra = _field
	return offset, nil
}

// for compatibility
func (p *Base) FastWrite(buf []byte) int {
	return 0
}

func (p *Base) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
		offset += p.fastWriteField2(buf[offset:], w)
		offset += p.fastWriteField3(buf[offset:], w)
		offset += p.fastWriteField4(buf[offset:], w)
		offset += p.fastWriteField5(buf[offset:], w)
		offset += p.fastWriteField6(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *Base) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *Base) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 1)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.LogID)
	return offset
}

func (p *Base) fastWriteField2(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 2)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.Caller)
	return offset
}

func (p *Base) fastWriteField3(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 3)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.Addr)
	return offset
}

func (p *Base) fastWriteField4(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 4)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.Client)
	return offset
}

func (p *Base) fastWriteField5(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetTrafficEnv() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRUCT, 5)
		offset += p.TrafficEnv.FastWriteNocopy(buf[offset:], w)
	}
	return offset
}

func (p *Base) fastWriteField6(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetExtra() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.MAP, 6)
		mapBeginOffset := offset
		offset += thrift.Binary.MapBeginLength()
		var length int
		for k, v := range p.Extra {
			length++
			offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, k)
			offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, v)
		}
		thrift.Binary.WriteMapBegin(buf[mapBeginOffset:], thrift.STRING, thrift.STRING, length)
	}
	return offset
}

func (p *Base) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.LogID)
	return l
}

func (p *Base) field2Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.Caller)
	return l
}

func (p *Base) field3Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.Addr)
	return l
}

func (p *Base) field4Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.Client)
	return l
}

func (p *Base) field5Length() int {
	l := 0
	if p.IsSetTrafficEnv() {
		l += thrift.Binary.FieldBeginLength()
		l += p.TrafficEnv.BLength()
	}
	return l
}

func (p *Base) field6Length() int {
	l := 0
	if p.IsSetExtra() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.MapBeginLength()
		for k, v := range p.Extra {
			_, _ = k, v

			l += thrift.Binary.StringLengthNocopy(k)
			l += thrift.Binary.StringLengthNocopy(v)
		}
	}
	return l
}

func (p *BaseResp) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.MAP {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_BaseResp[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *BaseResp) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.StatusMessage = _field
	return offset, nil
}

func (p *BaseResp) FastReadField2(buf []byte) (int, error) {
	offset := 0

	var _field int32
	if v, l, err := thrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = v
	}
	p.StatusCode = _field
	return offset, nil
}

func (p *BaseResp) FastReadField3(buf []byte) (int, error) {
	offset := 0

	_, _, size, l, err := thrift.Binary.ReadMapBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	_field := make(map[string]string, size)
	for i := 0; i < size; i++ {
		var _key string
		if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
			_key = v
		}

		var _val string
		if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
			_val = v
		}

		_field[_key] = _val
	}
	p.Extra = _field
	return offset, nil
}

// for compatibility
func (p *BaseResp) FastWrite(buf []byte) int {
	return 0
}

func (p *BaseResp) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField2(buf[offset:], w)
		offset += p.fastWriteField1(buf[offset:], w)
		offset += p.fastWriteField3(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *BaseResp) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *BaseResp) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 1)
	offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, p.StatusMessage)
	return offset
}

func (p *BaseResp) fastWriteField2(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I32, 2)
	offset += thrift.Binary.WriteI32(buf[offset:], p.StatusCode)
	return offset
}

func (p *BaseResp) fastWriteField3(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetExtra() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.MAP, 3)
		mapBeginOffset := offset
		offset += thrift.Binary.MapBeginLength()
		var length int
		for k, v := range p.Extra {
			length++
			offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, k)
			offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, v)
		}
		thrift.Binary.WriteMapBegin(buf[mapBeginOffset:], thrift.STRING, thrift.STRING, length)
	}
	return offset
}

func (p *BaseResp) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.StringLengthNocopy(p.StatusMessage)
	return l
}

func (p *BaseResp) field2Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += thrift.Binary.I32Length()
	return l
}

func (p *BaseResp) field3Length() int {
	l := 0
	if p.IsSetExtra() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.MapBeginLength()
		for k, v := range p.Extra {
			_, _ = k, v

			l += thrift.Binary.StringLengthNocopy(k)
			l += thrift.Binary.StringLengthNocopy(v)
		}
	}
	return l
}

func (p *BaseElem) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_BaseElem[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *BaseElem) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field *int32
	if v, l, err := thrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.A = _field
	return offset, nil
}

// for compatibility
func (p *BaseElem) FastWrite(buf []byte) int {
	return 0
}

func (p *BaseElem) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *BaseElem) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *BaseElem) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetA() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I32, 1)
		offset += thrift.Binary.WriteI32(buf[offset:], *p.A)
	}
	return offset
}

func (p *BaseElem) field1Length() int {
	l := 0
	if p.IsSetA() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.I32Length()
	}
	return l
}
