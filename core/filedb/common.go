/*
Package for common io operation and const
*/
package filedb

import (
	"encoding/binary"
	"io"
)

var (
	DEFAULT_ENCODING = binary.LittleEndian
	DEFAULT_MAGIC    = []byte("HSH3")
)

func readBinary(file io.Reader, obj interface{}) error {
	return binary.Read(file, DEFAULT_ENCODING, obj)
}

func writeBinary(file io.Writer, obj interface{}) error {
	return binary.Write(file, DEFAULT_ENCODING, obj)
}
