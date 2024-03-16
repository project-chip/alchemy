package generate

import (
	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/zap"
)

func renderConstraint(el *etree.Element, fs matter.FieldSet, f *matter.Field) {

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
