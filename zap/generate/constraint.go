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
			el.CreateAttr("length", to.ZapString(f.Type))
		}
		if from.Defined() {
			el.CreateAttr("minLength", from.ZapString(f.Type))
		}
		el.RemoveAttr("min")
		el.RemoveAttr("max")
	} else {
		if from.Defined() {
			el.CreateAttr("min", from.ZapString(f.Type))
		}
		if to.Defined() {
			el.CreateAttr("max", to.ZapString(f.Type))
		}
		el.RemoveAttr("minLength")
		el.RemoveAttr("length")
	}
}
