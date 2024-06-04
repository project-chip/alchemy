package generate

import (
	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
)

func renderConstraint(el *etree.Element, fs matter.FieldSet, f *matter.Field) {

	if f.Access.Write != matter.PrivilegeUnknown && f.Type != nil && f.Type.Size() > 2 {
		// ZAP can't handle min/max on types whose size is larger than two bytes, due to limitations in the template generation code
		// See https://github.com/project-chip/zap/blob/master/src-electron/generator/helper-endpointconfig.js#L465
		el.RemoveAttr("min")
		el.RemoveAttr("max")
		el.RemoveAttr("minLength")
		el.RemoveAttr("length")
		return
	}

	from, to := zap.GetMinMax(&matter.ConstraintContext{Field: f, Fields: fs}, f.Constraint)

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
			patchDataExtremeAttribute(el, "length", &to, f)
		}
		if from.Defined() {
			patchDataExtremeAttribute(el, "minLength", &from, f)
		}
		el.RemoveAttr("min")
		el.RemoveAttr("max")
	} else {
		if from.Defined() {
			patchDataExtremeAttribute(el, "min", &from, f)
		}
		if to.Defined() {
			patchDataExtremeAttribute(el, "max", &to, f)
		}
		el.RemoveAttr("minLength")
		el.RemoveAttr("length")
	}
}
