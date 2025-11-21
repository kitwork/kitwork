package action

type Kind string

const (
	KindFull      Kind = "full"
	KindShort     Kind = "short"
	KindValue     Kind = "value"
	KindList      Kind = "list"
	KindPrimitive Kind = "primitive"
	KindUnknown   Kind = "unknown"
)
