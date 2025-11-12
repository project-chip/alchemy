package matter

type Revision struct {
	Number      *Number `json:"number,omitempty"`
	Description string  `json:"description,omitempty"`
}

type Revisions []*Revision

func (r Revisions) MostRecent() *Revision {
	var lastRevision *Revision
	var lastRevisionNumber uint64
	for _, rev := range r {
		if rev.Number.Valid() && rev.Number.Value() > lastRevisionNumber {
			lastRevision = rev
			lastRevisionNumber = rev.Number.Value()
		}
	}
	return lastRevision
}
