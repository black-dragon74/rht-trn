package interfaces

import "github.com/black-dragon74/rht-trn/types"

type Contract interface {
	ReadFromDisk() []types.Student
	WriteToDisk() error
}
