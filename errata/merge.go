package errata

import "slices"

func (e *Errata) Merge(other *Errata) {
	if other == nil {
		return
	}
	e.Disco.Merge(&other.Disco)
	e.Spec.Merge(&other.Spec)
	e.TestPlan.Merge(&other.TestPlan)
	e.SDK.Merge(&other.SDK)
}

func (d *Disco) Merge(other *Disco) {
	if other == nil {
		return
	}
	if d.Sections == nil {
		d.Sections = make(map[string]DiscoSection)
	}
	for k, v := range other.Sections {
		if s, ok := d.Sections[k]; ok {
			s.Merge(&v)
			d.Sections[k] = s
		} else {
			d.Sections[k] = v
		}
	}
}

func (ds *DiscoSection) Merge(other *DiscoSection) {
	if other == nil {
		return
	}
	if other.Skip != DiscoPurposeNone {
		ds.Skip = other.Skip
	}
}

func (s *Spec) Merge(other *Spec) {
	if other == nil {
		return
	}
	if other.UtilityInclude {
		s.UtilityInclude = true
	}
	if other.Domain != "" {
		s.Domain = other.Domain
	}
	if s.Sections == nil {
		s.Sections = make(map[string]SpecSection)
	}
	for k, v := range other.Sections {
		if ss, ok := s.Sections[k]; ok {
			ss.Merge(&v)
			s.Sections[k] = ss
		} else {
			s.Sections[k] = v
		}
	}
}

func (ss *SpecSection) Merge(other *SpecSection) {
	if other == nil {
		return
	}
	if other.Skip != SpecPurposeNone {
		ss.Skip = other.Skip
	}
}

func (tp *TestPlan) Merge(other *TestPlan) {
	if other == nil {
		return
	}
	if other.TestPlanPath != "" {
		tp.TestPlanPath = other.TestPlanPath
	}
	if tp.TestPlanPaths == nil {
		tp.TestPlanPaths = make(map[string]TestPlanPath)
	}
	for k, v := range other.TestPlanPaths {
		if existing, ok := tp.TestPlanPaths[k]; ok {
			if v.Path != "" {
				existing.Path = v.Path
			}
			tp.TestPlanPaths[k] = existing
		} else {
			tp.TestPlanPaths[k] = v
		}
	}
}

func (s *SDK) Merge(other *SDK) {
	if other == nil {
		return
	}
	if other.SkipFile {
		s.SkipFile = true
	}
	if other.SuppressAttributePermissions {
		s.SuppressAttributePermissions = true
	}
	if other.ClusterDefinePrefix != "" {
		s.ClusterDefinePrefix = other.ClusterDefinePrefix
	}
	if other.SuppressClusterDefinePrefix {
		s.SuppressClusterDefinePrefix = true
	}
	if s.DefineOverrides == nil {
		s.DefineOverrides = make(map[string]string)
	}
	for k, v := range other.DefineOverrides {
		s.DefineOverrides[k] = v
	}
	if other.ClusterName != "" {
		s.ClusterName = other.ClusterName
	}
	if s.ClusterAliases == nil {
		s.ClusterAliases = make(map[string][]string)
	}
	for k, v := range other.ClusterAliases {
		s.ClusterAliases[k] = v
	}
	if s.ClusterListKeys == nil {
		s.ClusterListKeys = make(map[string]string)
	}
	for k, v := range other.ClusterListKeys {
		s.ClusterListKeys[k] = v
	}
	if other.WritePrivilegeAsRole {
		s.WritePrivilegeAsRole = true
	}

	s.SeparateStructs.Merge(other.SeparateStructs)
	s.SeparateBitmaps.Merge(other.SeparateBitmaps)
	s.SeparateEnums.Merge(other.SeparateEnums)

	s.SharedBitmaps.Merge(other.SharedBitmaps)
	s.SharedEnums.Merge(other.SharedEnums)
	s.SharedStructs.Merge(other.SharedStructs)

	if other.TemplatePath != "" {
		s.TemplatePath = other.TemplatePath
	}
	if s.ClusterSplit == nil {
		s.ClusterSplit = make(map[string]string)
	}
	for k, v := range other.ClusterSplit {
		s.ClusterSplit[k] = v
	}
	for _, v := range other.ClusterSkip {
		if !slices.Contains(s.ClusterSkip, v) {
			s.ClusterSkip = append(s.ClusterSkip, v)
		}
	}
	if s.TypeNames == nil {
		s.TypeNames = make(map[string]string)
	}
	for k, v := range other.TypeNames {
		s.TypeNames[k] = v
	}

	if other.Types != nil {
		if s.Types == nil {
			s.Types = &SDKTypes{}
		}
		s.Types.Merge(other.Types)
	}
	if other.ExtraTypes != nil {
		if s.ExtraTypes == nil {
			s.ExtraTypes = &SDKTypes{}
		}
		s.ExtraTypes.Merge(other.ExtraTypes)
	}
}

func (u UniqueStringList) Merge(other UniqueStringList) {
	for k := range other {
		u[k] = struct{}{}
	}
}

func (st *SDKTypes) Merge(other *SDKTypes) {
	if other == nil {
		return
	}
	st.Attributes = mergeSDKTypeMap(st.Attributes, other.Attributes)
	st.Clusters = mergeSDKTypeMap(st.Clusters, other.Clusters)
	st.Enums = mergeSDKTypeMap(st.Enums, other.Enums)
	st.Bitmaps = mergeSDKTypeMap(st.Bitmaps, other.Bitmaps)
	st.Structs = mergeSDKTypeMap(st.Structs, other.Structs)
	st.Commands = mergeSDKTypeMap(st.Commands, other.Commands)
	st.Events = mergeSDKTypeMap(st.Events, other.Events)
	st.DeviceTypes = mergeSDKTypeMap(st.DeviceTypes, other.DeviceTypes)
}

func mergeSDKTypeMap(target, source map[string]*SDKType) map[string]*SDKType {
	if source == nil {
		return target
	}
	if target == nil {
		target = make(map[string]*SDKType)
	}
	for k, v := range source {
		if t, ok := target[k]; ok {
			t.Merge(v)
		} else {
			target[k] = v
		}
	}
	return target
}

func (t *SDKType) Merge(other *SDKType) {
	if other == nil {
		return
	}
	if other.Type != "" {
		t.Type = other.Type
	}
	if other.Name != "" {
		t.Name = other.Name
	}
	if other.OverrideName != "" {
		t.OverrideName = other.OverrideName
	}
	if other.OverrideType != "" {
		t.OverrideType = other.OverrideType
	}
	if other.List {
		t.List = true
	}
	t.Fields = mergeSDKTypeList(t.Fields, other.Fields)
	t.ExtraFields = mergeSDKTypeList(t.ExtraFields, other.ExtraFields)
	if other.Domain != "" {
		t.Domain = other.Domain
	}
	if other.Priority != "" {
		t.Priority = other.Priority
	}
	if other.Description != "" {
		t.Description = other.Description
	}
	if other.Bit != "" {
		t.Bit = other.Bit
	}
	if other.Code != "" {
		t.Code = other.Code
	}
	if other.Value != "" {
		t.Value = other.Value
	}
	if other.Constraint != "" {
		t.Constraint = other.Constraint
	}
	if other.Conformance != "" {
		t.Conformance = other.Conformance
	}
	if other.Fallback != "" {
		t.Fallback = other.Fallback
	}
	if other.Quality != "" {
		t.Quality = other.Quality
	}
	if other.Access != "" {
		t.Access = other.Access
	}
	if other.Direction != "" {
		t.Direction = other.Direction
	}
	if other.Response != "" {
		t.Response = other.Response
	}
	if other.FabricScoping != "" {
		t.FabricScoping = other.FabricScoping
	}
	if other.FabricSensitivity != "" {
		t.FabricSensitivity = other.FabricSensitivity
	}
	t.Attributes = mergeSDKTypeMap(t.Attributes, other.Attributes)
	t.Commands = mergeSDKTypeMap(t.Commands, other.Commands)
}

func mergeSDKTypeList(target, source []*SDKType) []*SDKType {
	if len(source) == 0 {
		return target
	}
	for _, s := range source {
		found := false
		for _, t := range target {
			if t.Name == s.Name {
				t.Merge(s)
				found = true
				break
			}
		}
		if !found {
			target = append(target, s)
		}
	}
	return target
}
