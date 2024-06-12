// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package liquidityincentive

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_Epoch_5_list)(nil)

type _Epoch_5_list struct {
	list *[]*Gauge
}

func (x *_Epoch_5_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Epoch_5_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Epoch_5_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Gauge)
	(*x.list)[i] = concreteValue
}

func (x *_Epoch_5_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Gauge)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Epoch_5_list) AppendMutable() protoreflect.Value {
	v := new(Gauge)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Epoch_5_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Epoch_5_list) NewElement() protoreflect.Value {
	v := new(Gauge)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Epoch_5_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Epoch              protoreflect.MessageDescriptor
	fd_Epoch_id           protoreflect.FieldDescriptor
	fd_Epoch_start_block  protoreflect.FieldDescriptor
	fd_Epoch_end_block    protoreflect.FieldDescriptor
	fd_Epoch_epoch_blocks protoreflect.FieldDescriptor
	fd_Epoch_gauges       protoreflect.FieldDescriptor
)

func init() {
	file_sunrise_liquidityincentive_epoch_proto_init()
	md_Epoch = File_sunrise_liquidityincentive_epoch_proto.Messages().ByName("Epoch")
	fd_Epoch_id = md_Epoch.Fields().ByName("id")
	fd_Epoch_start_block = md_Epoch.Fields().ByName("start_block")
	fd_Epoch_end_block = md_Epoch.Fields().ByName("end_block")
	fd_Epoch_epoch_blocks = md_Epoch.Fields().ByName("epoch_blocks")
	fd_Epoch_gauges = md_Epoch.Fields().ByName("gauges")
}

var _ protoreflect.Message = (*fastReflection_Epoch)(nil)

type fastReflection_Epoch Epoch

func (x *Epoch) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Epoch)(x)
}

func (x *Epoch) slowProtoReflect() protoreflect.Message {
	mi := &file_sunrise_liquidityincentive_epoch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Epoch_messageType fastReflection_Epoch_messageType
var _ protoreflect.MessageType = fastReflection_Epoch_messageType{}

type fastReflection_Epoch_messageType struct{}

func (x fastReflection_Epoch_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Epoch)(nil)
}
func (x fastReflection_Epoch_messageType) New() protoreflect.Message {
	return new(fastReflection_Epoch)
}
func (x fastReflection_Epoch_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Epoch
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Epoch) Descriptor() protoreflect.MessageDescriptor {
	return md_Epoch
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Epoch) Type() protoreflect.MessageType {
	return _fastReflection_Epoch_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Epoch) New() protoreflect.Message {
	return new(fastReflection_Epoch)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Epoch) Interface() protoreflect.ProtoMessage {
	return (*Epoch)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Epoch) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_Epoch_id, value) {
			return
		}
	}
	if x.StartBlock != int64(0) {
		value := protoreflect.ValueOfInt64(x.StartBlock)
		if !f(fd_Epoch_start_block, value) {
			return
		}
	}
	if x.EndBlock != int64(0) {
		value := protoreflect.ValueOfInt64(x.EndBlock)
		if !f(fd_Epoch_end_block, value) {
			return
		}
	}
	if x.EpochBlocks != int64(0) {
		value := protoreflect.ValueOfInt64(x.EpochBlocks)
		if !f(fd_Epoch_epoch_blocks, value) {
			return
		}
	}
	if len(x.Gauges) != 0 {
		value := protoreflect.ValueOfList(&_Epoch_5_list{list: &x.Gauges})
		if !f(fd_Epoch_gauges, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Epoch) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "sunrise.liquidityincentive.Epoch.id":
		return x.Id != uint64(0)
	case "sunrise.liquidityincentive.Epoch.start_block":
		return x.StartBlock != int64(0)
	case "sunrise.liquidityincentive.Epoch.end_block":
		return x.EndBlock != int64(0)
	case "sunrise.liquidityincentive.Epoch.epoch_blocks":
		return x.EpochBlocks != int64(0)
	case "sunrise.liquidityincentive.Epoch.gauges":
		return len(x.Gauges) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquidityincentive.Epoch"))
		}
		panic(fmt.Errorf("message sunrise.liquidityincentive.Epoch does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Epoch) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "sunrise.liquidityincentive.Epoch.id":
		x.Id = uint64(0)
	case "sunrise.liquidityincentive.Epoch.start_block":
		x.StartBlock = int64(0)
	case "sunrise.liquidityincentive.Epoch.end_block":
		x.EndBlock = int64(0)
	case "sunrise.liquidityincentive.Epoch.epoch_blocks":
		x.EpochBlocks = int64(0)
	case "sunrise.liquidityincentive.Epoch.gauges":
		x.Gauges = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquidityincentive.Epoch"))
		}
		panic(fmt.Errorf("message sunrise.liquidityincentive.Epoch does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Epoch) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "sunrise.liquidityincentive.Epoch.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "sunrise.liquidityincentive.Epoch.start_block":
		value := x.StartBlock
		return protoreflect.ValueOfInt64(value)
	case "sunrise.liquidityincentive.Epoch.end_block":
		value := x.EndBlock
		return protoreflect.ValueOfInt64(value)
	case "sunrise.liquidityincentive.Epoch.epoch_blocks":
		value := x.EpochBlocks
		return protoreflect.ValueOfInt64(value)
	case "sunrise.liquidityincentive.Epoch.gauges":
		if len(x.Gauges) == 0 {
			return protoreflect.ValueOfList(&_Epoch_5_list{})
		}
		listValue := &_Epoch_5_list{list: &x.Gauges}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquidityincentive.Epoch"))
		}
		panic(fmt.Errorf("message sunrise.liquidityincentive.Epoch does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Epoch) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "sunrise.liquidityincentive.Epoch.id":
		x.Id = value.Uint()
	case "sunrise.liquidityincentive.Epoch.start_block":
		x.StartBlock = value.Int()
	case "sunrise.liquidityincentive.Epoch.end_block":
		x.EndBlock = value.Int()
	case "sunrise.liquidityincentive.Epoch.epoch_blocks":
		x.EpochBlocks = value.Int()
	case "sunrise.liquidityincentive.Epoch.gauges":
		lv := value.List()
		clv := lv.(*_Epoch_5_list)
		x.Gauges = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquidityincentive.Epoch"))
		}
		panic(fmt.Errorf("message sunrise.liquidityincentive.Epoch does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Epoch) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "sunrise.liquidityincentive.Epoch.gauges":
		if x.Gauges == nil {
			x.Gauges = []*Gauge{}
		}
		value := &_Epoch_5_list{list: &x.Gauges}
		return protoreflect.ValueOfList(value)
	case "sunrise.liquidityincentive.Epoch.id":
		panic(fmt.Errorf("field id of message sunrise.liquidityincentive.Epoch is not mutable"))
	case "sunrise.liquidityincentive.Epoch.start_block":
		panic(fmt.Errorf("field start_block of message sunrise.liquidityincentive.Epoch is not mutable"))
	case "sunrise.liquidityincentive.Epoch.end_block":
		panic(fmt.Errorf("field end_block of message sunrise.liquidityincentive.Epoch is not mutable"))
	case "sunrise.liquidityincentive.Epoch.epoch_blocks":
		panic(fmt.Errorf("field epoch_blocks of message sunrise.liquidityincentive.Epoch is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquidityincentive.Epoch"))
		}
		panic(fmt.Errorf("message sunrise.liquidityincentive.Epoch does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Epoch) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "sunrise.liquidityincentive.Epoch.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "sunrise.liquidityincentive.Epoch.start_block":
		return protoreflect.ValueOfInt64(int64(0))
	case "sunrise.liquidityincentive.Epoch.end_block":
		return protoreflect.ValueOfInt64(int64(0))
	case "sunrise.liquidityincentive.Epoch.epoch_blocks":
		return protoreflect.ValueOfInt64(int64(0))
	case "sunrise.liquidityincentive.Epoch.gauges":
		list := []*Gauge{}
		return protoreflect.ValueOfList(&_Epoch_5_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquidityincentive.Epoch"))
		}
		panic(fmt.Errorf("message sunrise.liquidityincentive.Epoch does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Epoch) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in sunrise.liquidityincentive.Epoch", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Epoch) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Epoch) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Epoch) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Epoch) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Epoch)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		if x.StartBlock != 0 {
			n += 1 + runtime.Sov(uint64(x.StartBlock))
		}
		if x.EndBlock != 0 {
			n += 1 + runtime.Sov(uint64(x.EndBlock))
		}
		if x.EpochBlocks != 0 {
			n += 1 + runtime.Sov(uint64(x.EpochBlocks))
		}
		if len(x.Gauges) > 0 {
			for _, e := range x.Gauges {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Epoch)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Gauges) > 0 {
			for iNdEx := len(x.Gauges) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Gauges[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x2a
			}
		}
		if x.EpochBlocks != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.EpochBlocks))
			i--
			dAtA[i] = 0x20
		}
		if x.EndBlock != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.EndBlock))
			i--
			dAtA[i] = 0x18
		}
		if x.StartBlock != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.StartBlock))
			i--
			dAtA[i] = 0x10
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Epoch)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Epoch: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Epoch: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StartBlock", wireType)
				}
				x.StartBlock = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.StartBlock |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field EndBlock", wireType)
				}
				x.EndBlock = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.EndBlock |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field EpochBlocks", wireType)
				}
				x.EpochBlocks = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.EpochBlocks |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Gauges", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Gauges = append(x.Gauges, &Gauge{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Gauges[len(x.Gauges)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: sunrise/liquidityincentive/epoch.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Epoch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StartBlock  int64    `protobuf:"varint,2,opt,name=start_block,json=startBlock,proto3" json:"start_block,omitempty"`
	EndBlock    int64    `protobuf:"varint,3,opt,name=end_block,json=endBlock,proto3" json:"end_block,omitempty"`
	EpochBlocks int64    `protobuf:"varint,4,opt,name=epoch_blocks,json=epochBlocks,proto3" json:"epoch_blocks,omitempty"`
	Gauges      []*Gauge `protobuf:"bytes,5,rep,name=gauges,proto3" json:"gauges,omitempty"`
}

func (x *Epoch) Reset() {
	*x = Epoch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sunrise_liquidityincentive_epoch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Epoch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Epoch) ProtoMessage() {}

// Deprecated: Use Epoch.ProtoReflect.Descriptor instead.
func (*Epoch) Descriptor() ([]byte, []int) {
	return file_sunrise_liquidityincentive_epoch_proto_rawDescGZIP(), []int{0}
}

func (x *Epoch) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Epoch) GetStartBlock() int64 {
	if x != nil {
		return x.StartBlock
	}
	return 0
}

func (x *Epoch) GetEndBlock() int64 {
	if x != nil {
		return x.EndBlock
	}
	return 0
}

func (x *Epoch) GetEpochBlocks() int64 {
	if x != nil {
		return x.EpochBlocks
	}
	return 0
}

func (x *Epoch) GetGauges() []*Gauge {
	if x != nil {
		return x.Gauges
	}
	return nil
}

var File_sunrise_liquidityincentive_epoch_proto protoreflect.FileDescriptor

var file_sunrise_liquidityincentive_epoch_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x69, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x76, 0x65, 0x2f, 0x65, 0x70, 0x6f,
	0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73,
	0x65, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x69, 0x6e, 0x63, 0x65, 0x6e,
	0x74, 0x69, 0x76, 0x65, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6c,
	0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x69, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x76,
	0x65, 0x2f, 0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01,
	0x0a, 0x05, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x6e, 0x64,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x65, 0x70, 0x6f,
	0x63, 0x68, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x3f, 0x0a, 0x06, 0x67, 0x61, 0x75, 0x67,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x75, 0x6e, 0x72, 0x69,
	0x73, 0x65, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x69, 0x6e, 0x63, 0x65,
	0x6e, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x47, 0x61, 0x75, 0x67, 0x65, 0x42, 0x04, 0xc8, 0xde, 0x1f,
	0x00, 0x52, 0x06, 0x67, 0x61, 0x75, 0x67, 0x65, 0x73, 0x42, 0xe2, 0x01, 0x0a, 0x1e, 0x63, 0x6f,
	0x6d, 0x2e, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x69, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x76, 0x65, 0x42, 0x0a, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x6e,
	0x72, 0x69, 0x73, 0x65, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x69, 0x6e,
	0x63, 0x65, 0x6e, 0x74, 0x69, 0x76, 0x65, 0xa2, 0x02, 0x03, 0x53, 0x4c, 0x58, 0xaa, 0x02, 0x1a,
	0x53, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x69, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x76, 0x65, 0xca, 0x02, 0x1a, 0x53, 0x75, 0x6e,
	0x72, 0x69, 0x73, 0x65, 0x5c, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x69, 0x6e,
	0x63, 0x65, 0x6e, 0x74, 0x69, 0x76, 0x65, 0xe2, 0x02, 0x26, 0x53, 0x75, 0x6e, 0x72, 0x69, 0x73,
	0x65, 0x5c, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x69, 0x6e, 0x63, 0x65, 0x6e,
	0x74, 0x69, 0x76, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x1b, 0x53, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x3a, 0x3a, 0x4c, 0x69, 0x71, 0x75,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x69, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x76, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sunrise_liquidityincentive_epoch_proto_rawDescOnce sync.Once
	file_sunrise_liquidityincentive_epoch_proto_rawDescData = file_sunrise_liquidityincentive_epoch_proto_rawDesc
)

func file_sunrise_liquidityincentive_epoch_proto_rawDescGZIP() []byte {
	file_sunrise_liquidityincentive_epoch_proto_rawDescOnce.Do(func() {
		file_sunrise_liquidityincentive_epoch_proto_rawDescData = protoimpl.X.CompressGZIP(file_sunrise_liquidityincentive_epoch_proto_rawDescData)
	})
	return file_sunrise_liquidityincentive_epoch_proto_rawDescData
}

var file_sunrise_liquidityincentive_epoch_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sunrise_liquidityincentive_epoch_proto_goTypes = []interface{}{
	(*Epoch)(nil), // 0: sunrise.liquidityincentive.Epoch
	(*Gauge)(nil), // 1: sunrise.liquidityincentive.Gauge
}
var file_sunrise_liquidityincentive_epoch_proto_depIdxs = []int32{
	1, // 0: sunrise.liquidityincentive.Epoch.gauges:type_name -> sunrise.liquidityincentive.Gauge
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sunrise_liquidityincentive_epoch_proto_init() }
func file_sunrise_liquidityincentive_epoch_proto_init() {
	if File_sunrise_liquidityincentive_epoch_proto != nil {
		return
	}
	file_sunrise_liquidityincentive_gauge_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sunrise_liquidityincentive_epoch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Epoch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sunrise_liquidityincentive_epoch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sunrise_liquidityincentive_epoch_proto_goTypes,
		DependencyIndexes: file_sunrise_liquidityincentive_epoch_proto_depIdxs,
		MessageInfos:      file_sunrise_liquidityincentive_epoch_proto_msgTypes,
	}.Build()
	File_sunrise_liquidityincentive_epoch_proto = out.File
	file_sunrise_liquidityincentive_epoch_proto_rawDesc = nil
	file_sunrise_liquidityincentive_epoch_proto_goTypes = nil
	file_sunrise_liquidityincentive_epoch_proto_depIdxs = nil
}
