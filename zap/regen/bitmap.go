package regen

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func clusterBitmapsHelper(spec *spec.Specification) func(clusterInfo ClusterInfo, options *raymond.Options) raymond.SafeString {
	return func(clusterInfo ClusterInfo, options *raymond.Options) raymond.SafeString {
		bitmaps := make(matter.BitmapSet, 0, len(clusterInfo.ReferencedBitmaps))
		for _, bm := range clusterInfo.ReferencedBitmaps {
			if len(filterBits(bm)) > 0 {
				bitmaps = append(bitmaps, bm)
			}
		}
		cluster := clusterInfo.Cluster
		if cluster.Features != nil {
			if len(filterBits(&cluster.Features.Bitmap)) > 0 {
				features := cluster.Features.Bitmap.Clone()
				// ZAP renames this for some reason
				features.Name = "Feature"
				bitmaps = append(bitmaps, features)
			}
		}
		slices.SortStableFunc(bitmaps, func(a *matter.Bitmap, b *matter.Bitmap) int {
			return strings.Compare(a.Name, b.Name)
		})

		return enumerateEntitiesHelper(bitmaps, spec, options)

	}
}

func bitmapTypeHelper(bitmap matter.Bitmap) raymond.SafeString {
	switch bitmap.Type.BaseType {
	case types.BaseDataTypeMap8:
		return raymond.SafeString("bitmap8")
	case types.BaseDataTypeMap16:
		return raymond.SafeString("bitmap16")
	case types.BaseDataTypeMap32:
		return raymond.SafeString("bitmap32")
	case types.BaseDataTypeMap64:
		return raymond.SafeString("bitmap24")
	default:
		return raymond.SafeString("unknown bitmap type")

	}
}

func bitNameHelper(bit any) raymond.SafeString {
	switch bit := bit.(type) {
	case matter.BitmapBit:
		return asUpperCamelCaseHelper(bit.Name())
	case matter.Feature:
		return asUpperCamelCaseHelper(bit.Name())
	default:
		return raymond.SafeString(fmt.Sprintf("unexpected bitName type: %T", bit))
	}
}

func bitMaskHelper(bit any) raymond.SafeString {
	switch bit := bit.(type) {
	case matter.BitmapBit:
		mask, err := bit.Mask()
		if err != nil {
			return raymond.SafeString(fmt.Sprintf("error converting bitmap mask: %v", err))
		}
		return raymond.SafeString(fmt.Sprintf("0x%X", mask))
	case matter.Feature:
		mask, err := bit.Mask()
		if err != nil {
			return raymond.SafeString(fmt.Sprintf("error converting feature mask: %v", err))
		}
		return raymond.SafeString(fmt.Sprintf("0x%X", mask))
	default:
		return raymond.SafeString(fmt.Sprintf("unexpected bitName type: %T", bit))
	}
}

func filterBits(bm *matter.Bitmap) (bits matter.BitSet) {
	for _, b := range bm.Bits {
		if conformance.IsZigbee(b.Conformance()) || zap.IsDisallowed(b, b.Conformance()) || conformance.IsDeprecated(b.Conformance()) {
			continue
		}
		bits = append(bits, b)
	}
	return
}

func bitmapBitsHelper(spec *spec.Specification) func(bm matter.Bitmap, options *raymond.Options) raymond.SafeString {
	return func(bm matter.Bitmap, options *raymond.Options) raymond.SafeString {
		return enumerateEntitiesHelper(filterBits(&bm), spec, options)
	}
}
