/*
Package for common io operation and const
*/
package utils

import (
	"encoding/binary"
	"io"
)

var (
	DEFAULT_ENCODING = binary.LittleEndian
)

// common writer method for binary
func ReadBinary(file io.Reader, obj interface{}) error {
	return binary.Read(file, DEFAULT_ENCODING, obj)
}

// common reader method for binary
func WriteBinary(file io.Writer, obj interface{}) error {
	return binary.Write(file, DEFAULT_ENCODING, obj)
}

func ReadInt64(file io.Reader, obj *int64) error {
	return ReadBinary(file, obj)
}

func WriteInt64(file io.Writer, obj int64) error {
	return WriteBinary(file, obj)
}

func ReadUInt64(file io.Reader, obj *uint64) error {
	return ReadBinary(file, obj)
}

func WriteUInt64(file io.Writer, obj uint64) error {
	return WriteBinary(file, obj)
}

func ReadInt32(file io.Reader, obj *int32) error {
	return ReadBinary(file, obj)
}

func WriteInt32(file io.Writer, obj int32) error {
	return WriteBinary(file, obj)
}

func ReadUInt32(file io.Reader, obj *uint32) error {
	return ReadBinary(file, obj)
}

func WriteUInt32(file io.Writer, obj uint32) error {
	return WriteBinary(file, obj)
}

func ReadInt16(file io.Reader, obj *int16) error {
	return ReadBinary(file, obj)
}

func WriteInt16(file io.Writer, obj int16) error {
	return WriteBinary(file, obj)
}

func ReadUInt16(file io.Reader, obj *uint16) error {
	return ReadBinary(file, obj)
}

func WriteUInt16(file io.Writer, obj uint16) error {
	return WriteBinary(file, obj)
}

func WriteByte(file io.Writer, obj byte) error {
	return WriteBinary(file, obj)
}

func ReadByte(file io.Reader, obj *byte) error {
	return ReadBinary(file, obj)
}

func WriteBytes(file io.Writer, obj []byte) error {
	return WriteBinary(file, obj)
}

func ReadBytes(file io.Reader, obj []byte) error {
	return ReadBinary(file, obj)
}
