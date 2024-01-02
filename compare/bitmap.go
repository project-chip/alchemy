package compare

import (
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
)

func compareBits(specBit *matter.Bit, zapBit *matter.Bit) (diffs []any) {
	specMask, err := specBit.Mask()
	if err != nil {

	} else {
		zapMask, err := zapBit.Mask()
		if err != nil {

		} else if specMask != zapMask {
			diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyBit, Spec: specBit.Bit, ZAP: zapBit.Bit})
		}
	}
	if !namesEqual(specBit.Name, zapBit.Name) {
		diffs = append(diffs, &StringDiff{Type: DiffTypeMismatch, Property: DiffPropertyName, Spec: specBit.Name, ZAP: zapBit.Name})
	}
	return
}

func compareBitmapsByCode(specBitmap *matter.Bitmap, zapBitmap *matter.Bitmap) (diffs []any) {
	if specBitmap == nil {
		if zapBitmap == nil {
			return
		} else {
			diffs = append(diffs, newMissingDiff(zapBitmap.Name, types.EntityTypeBitmap, SourceSpec))
			return
		}
	} else if zapBitmap == nil {
		diffs = append(diffs, newMissingDiff(specBitmap.Name, types.EntityTypeBitmap, SourceZAP))
		return
	}
	specBitmapMap := make(map[string]*matter.Bit)
	for _, f := range specBitmap.Bits {
		specBitmapMap[strings.ToLower(f.Code)] = f
	}

	zapBitmapMap := make(map[string]*matter.Bit)
	for _, f := range zapBitmap.Bits {
		zapBitmapMap[strings.ToLower(f.Code)] = f
	}

	for code, zapBit := range zapBitmapMap {
		specBit, ok := specBitmapMap[code]
		if !ok {
			continue
		}
		delete(zapBitmapMap, code)
		delete(specBitmapMap, code)
		bitDiffs := compareBits(specBit, zapBit)
		if len(bitDiffs) > 0 {
			diffs = append(diffs, &IdentifiedDiff{Type: DiffTypeMismatch, Entity: types.EntityTypeFeature, Name: specBit.Name, Diffs: bitDiffs})
		}
	}
	for _, f := range specBitmapMap {
		diffs = append(diffs, &MissingDiff{Type: DiffTypeMissing, Entity: types.EntityTypeFeature, Code: f.Code, Source: SourceZAP})
	}
	for _, f := range zapBitmapMap {
		diffs = append(diffs, &MissingDiff{Type: DiffTypeMissing, Entity: types.EntityTypeFeature, Code: f.Code, Source: SourceSpec})
	}
	return
}

func compareBitmapsByMask(specBitmap *matter.Bitmap, zapBitmap *matter.Bitmap, entityType types.EntityType) (diffs []any) {
	if specBitmap == nil {
		if zapBitmap == nil {
			return
		} else {
			diffs = append(diffs, newMissingDiff(zapBitmap.Name, entityType, SourceSpec))
			return
		}
	} else if zapBitmap == nil {
		diffs = append(diffs, newMissingDiff(specBitmap.Name, entityType, SourceZAP))
		return
	}
	specBitmapMap := make(map[uint64]*matter.Bit)
	for _, f := range specBitmap.Bits {
		mask, err := f.Mask()
		if err == nil {
			specBitmapMap[mask] = f
		}
	}

	zapBitmapMap := make(map[uint64]*matter.Bit)
	for _, f := range zapBitmap.Bits {
		mask, err := f.Mask()
		if err == nil {
			zapBitmapMap[mask] = f
		}
	}

	for mask, zapBit := range zapBitmapMap {
		specBit, ok := specBitmapMap[mask]
		if !ok {
			continue
		}
		delete(zapBitmapMap, mask)
		delete(specBitmapMap, mask)
		bitDiffs := compareBits(specBit, zapBit)
		if len(bitDiffs) > 0 {
			diffs = append(diffs, &MaskDiff{Type: DiffTypeMismatch, Entity: entityType, Mask: mask, Diffs: bitDiffs})
		}
	}
	for _, f := range specBitmapMap {
		diffs = append(diffs, newMissingDiff(f.Name, DiffTypeMissing, entityType, f.Code, SourceZAP))
	}
	for _, f := range zapBitmapMap {
		diffs = append(diffs, newMissingDiff(f.Name, DiffTypeMissing, entityType, f.Code, SourceSpec))
	}
	return
}

func compareBitmaps(specBitmaps []*matter.Bitmap, zapBitmaps []*matter.Bitmap) (diffs []any) {
	specBitmapMap := make(map[string]*matter.Bitmap)
	for _, f := range specBitmaps {
		specBitmapMap[strings.ToLower(f.Name)] = f
	}

	zapBitmapMap := make(map[string]*matter.Bitmap)
	for _, f := range zapBitmaps {
		zapBitmapMap[strings.ToLower(f.Name)] = f
	}
	for name, zapBitmap := range zapBitmapMap {
		specName := name
		specBitmap, ok := specBitmapMap[specName]
		if !ok {
			specName = name + "bitmap"
			specBitmap, ok = specBitmapMap[specName]
			if !ok {
				continue
			}
		}
		delete(zapBitmapMap, name)
		delete(specBitmapMap, specName)
		bitmapDiffs := compareBitmapsByMask(specBitmap, zapBitmap, types.EntityTypeBitmap)
		if len(bitmapDiffs) > 0 {
			diffs = append(diffs, &IdentifiedDiff{Type: DiffTypeMismatch, Name: specBitmap.Name, Entity: types.EntityTypeBitmap, Diffs: bitmapDiffs})
		}
	}
	for _, f := range specBitmapMap {
		diffs = append(diffs, newMissingDiff(f.Name, DiffTypeMissing, types.EntityTypeBitmap, SourceZAP))
	}
	for _, f := range zapBitmapMap {
		diffs = append(diffs, newMissingDiff(f.Name, DiffTypeMissing, types.EntityTypeBitmap, SourceSpec))
	}
	return
}
