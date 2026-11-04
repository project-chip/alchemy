package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
)

func validateBitmaps(spec *Specification, bitmaps []*matter.Bitmap) {
	for _, bm := range bitmaps {
		validateBitmap(spec, bm)
	}
}

func validateBitmap(spec *Specification, en *matter.Bitmap) {
	bits := make(map[uint64]matter.Bit)
	nu := make(nameUniqueness[matter.Bit])
	cv := make(conformanceValidation)
	for _, b := range en.Bits {
		nu.check(spec, b)
		cv.add(b, b.Conformance())
		from, to, err := b.Bits()
		if err != nil {
			slog.Warn("Unable to determine range of bitmap values", log.Path("source", b), matter.LogEntity("parent", en), slog.String("name", b.Name()), slog.Any("error", err))
			continue
		}
		for i := from; i <= to; i++ {
			existing, ok := bits[i]
			if ok {
				slog.Error("Overlapping bitmap bit range", log.Path("source", b), matter.LogEntity("parent", en), slog.Uint64("bitFrom", from), slog.Uint64("bitTo", to), slog.Uint64("bitOverlap", i), slog.String("bitName", b.Name()), slog.String("previousBitName", existing.Name()))
				spec.addError(&DuplicateEntityIDError{Entity: b, Previous: existing})
			}
			bits[i] = b
		}
	}
	cv.check(spec)
}
