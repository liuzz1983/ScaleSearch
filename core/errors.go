package core

import (
	"errors"
)

var (
	ErrIsNotDir   = errors.New("this is not a directory")
	ErrIsReadOnly = errors.New("this directory is readonly")

	ErrFileNotExists = errors.New("this file is not exists")
	ErrFileExists    = errors.New("this file is exists")

	ErrFileFormat = errors.New("file format is wrong")
	ErrOutOfOrder = errors.New("out fo order")

	ErrNotImplement = errors.New("this method not implement ")
)
