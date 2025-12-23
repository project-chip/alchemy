package provisional

import (
	"fmt"
	"strings"
)

type Presence uint8

const (
	PresenceUnknown        Presence = iota
	PresenceBase                    = 1 << (iota - 1)
	PresenceBaseInProgress          = 1 << (iota - 1)
	PresenceHead                    = 1 << (iota - 1)
	PresenceHeadInProgress          = 1 << (iota - 1)
)

func (p Presence) Has(op Presence) bool {
	return p&op != PresenceUnknown
}

func (p Presence) String() string {
	if p == PresenceUnknown {
		return "unknown"
	}
	var ps strings.Builder
	for _, op := range []Presence{PresenceBase, PresenceBaseInProgress, PresenceHead, PresenceHeadInProgress} {
		if p.Has(op) {
			if ps.Len() > 0 {
				ps.WriteRune(',')
			}
			switch op {
			case PresenceBase:
				ps.WriteString("base")
			case PresenceBaseInProgress:
				ps.WriteString("baseInProgress")
			case PresenceHead:
				ps.WriteString("head")
			case PresenceHeadInProgress:
				ps.WriteString("headInProgress")
			default:
				return "invalid"
			}
		}
	}
	return ps.String()
}

func (p Presence) Novelty() (Novelty, error) {
	if p.Has(PresenceHeadInProgress) {
		if p.Has(PresenceHead) {
			if p.Has(PresenceBaseInProgress) {
				if p.Has(PresenceBase) {
					// This entity is in all versions of the spec
					return NoveltyNone, nil
				}
				// This PR adds an in-progress ifdef to an existing entity
				return NoveltyNone, nil
			}
			if p.Has(PresenceBase) {
				return NoveltyNone, fmt.Errorf("entity disappears when in-progress is set on base")
			}
			// This entity exists in both head and head-in-progress, but does not exist in base or base-in-progress,
			// so it's new but not ifdef'd
			return NoveltyNew, nil
		}
		// Does not exist in head
		if p.Has(PresenceBaseInProgress) {
			if p.Has(PresenceBase) {
				// This PR adds an in-progress ifdef
				return NoveltyNone, nil
			}
			// This entity is not new, and is properly if-def'd in both base and head
			return NoveltyNone, nil
		}
		// Does not exist in base-in-progress
		if p.Has(PresenceBase) {
			return NoveltyNone, fmt.Errorf("entity disappears when in-progress is set on base")
		}
		// This entity is new, and is properly if-def'd
		return NoveltyNew | NoveltyIfDefd, nil
	} else {
		if p.Has(PresenceHead) {
			return NoveltyNone, fmt.Errorf("entity disappears when in-progress is set on head")
		}
		if p.Has(PresenceBaseInProgress) {
			// This PR removes this entity
			return NoveltyNone, nil
		}
		if p.Has(PresenceBase) {
			return NoveltyNone, fmt.Errorf("entity disappears when in-progress is set on base")
		}
		return NoveltyNone, fmt.Errorf("entity is missing from base and head")
	}
}

type Novelty uint8

const (
	NoveltyNone   Novelty = iota
	NoveltyNew            = 1 << (iota - 1)
	NoveltyIfDefd         = 1 << (iota - 1)
)

func (n Novelty) Has(op Novelty) bool {
	return n&op != NoveltyNone
}

func (n Novelty) IsNew() bool {
	return n.Has(NoveltyNew)
}

func (n Novelty) IsIfDefd() bool {
	return n.Has(NoveltyIfDefd)
}
