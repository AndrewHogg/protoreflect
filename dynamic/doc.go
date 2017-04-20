// Package dynamic provides an implementation for a dynamic protobuf message.
//
// The dynamic message is essentially a message descriptor along with a map of
// tag numbers to values. It has a broad API for interacting with the message,
// including inspection and modification. Generally, most operations have two
// forms: a regular method that panics on bad input or error and a "Try" form
// of the method that will instead return an error.
//
// In addition to implementing the proto.Message interface, the included
// Message type also provides an XXX_MessageName() method, so it can work with
// proto.MessageName. And it provides a Descriptor() method that behaves just
// like the method of the same signature in messages generated by protoc. Because
// of this, it is actually compatible with proto.Message in many (though not all)
// contexts. In particular, it is compatible with proto.Marshal and proto.Unmarshal
// for serializing and de-serializing messages.
//
// This package also contains a several registries, for managing known types and
// descriptors. The MessageRegistry is used for interacting with Any messages
// where the actual embedded value may be a dynamic message. The KnownTypeRegistry
// allows de-serialization to use generated message types, instead of dynamic
// messages, for some kinds of nested message fields. And the ExtensionRegistry
// is for recognizing and parsing extensions fields (for proto2 messages).
//
// The dynamic message supports binary and text marshaling, using protobuf's
// well-defined binary format and the same text format that protoc-generated
// types use. It also supports JSON serialization/de-serialization by
// implementing the json.Marshaler and json.Unmarshaler interfaces.
package dynamic
