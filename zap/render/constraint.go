package render

import (
	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

func (cr *configuratorRenderer) renderConstraint(el *etree.Element, fs matter.FieldSet, f *matter.Field) {

	if f.Access.Write != matter.PrivilegeUnknown && f.Type != nil && f.Type.Size() > 2 {
		// ZAP can't handle min/max on types whose size is larger than two bytes, due to limitations in the template generation code
		// See https://github.com/project-chip/zap/blob/master/src-electron/generator/helper-endpointconfig.js#L465
		el.RemoveAttr("min")
		el.RemoveAttr("max")
		el.RemoveAttr("minLength")
		el.RemoveAttr("length")
		return
	}

	from, to := zap.GetMinMax(matter.NewConstraintContext(f, fs), cr.configurator.Errata.OverrideConstraint(f))

	if !from.Defined() {
		el.RemoveAttr("min")
		el.RemoveAttr("minLength")
	}
	if !to.Defined() {
		el.RemoveAttr("max")
		el.RemoveAttr("length")
	}

	if f.Type != nil && (f.Type.HasLength() || f.Type.IsArray()) {
		if to.Defined() {
			patchDataExtremeAttribute(el, "length", to, f, types.DataExtremePurposeMaximum)
		}
		if from.Defined() {
			if from.IsZero() {
				el.RemoveAttr("minLength")
			} else {
				patchDataExtremeAttribute(el, "minLength", from, f, types.DataExtremePurposeMinimum)
			}
		}
		el.RemoveAttr("min")
		el.RemoveAttr("max")
	} else {
		if from.Defined() {
			patchDataExtremeAttribute(el, "min", from, f, types.DataExtremePurposeMinimum)
		}
		if to.Defined() {
			patchDataExtremeAttribute(el, "max", to, f, types.DataExtremePurposeMaximum)
		}
		el.RemoveAttr("minLength")
		el.RemoveAttr("length")
	}
}
