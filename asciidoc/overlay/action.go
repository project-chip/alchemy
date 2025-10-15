package overlay

type overlayAction uint8

const (
	overlayActionNone   overlayAction = iota
	overlayActionRemove               = (1 << iota)
	overlayActionReplace
	overlayActionOverrideParent
	overlayActionOverrideChildren
	overlayActionAppendElements
)

func (ppa overlayAction) Remove() bool {
	return ppa&overlayActionRemove != 0
}

func (ppa overlayAction) Replace() bool {
	return ppa&overlayActionReplace != 0
}

func (ppa overlayAction) OverrideChildren() bool {
	return ppa&overlayActionOverrideChildren != 0
}

func (ppa overlayAction) OverrideParent() bool {
	return ppa&overlayActionOverrideParent != 0
}

func (ppa overlayAction) Append() bool {
	return ppa&overlayActionAppendElements != 0
}
