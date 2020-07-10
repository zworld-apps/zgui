package zgui

type Alignment int32

const (
	AlignStart Alignment = iota
	AlignCenter
	AlignEnd
)

type Direction int32

const (
	DirRow Direction = iota
	DirRowReverse
	DirColumn
	DirColumnReverse
)
