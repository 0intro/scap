package oval_def

func (defs *OvalDefinitions) DefinitionClasses() []DefinitionClass {
	res := []DefinitionClass{}
	seen := map[string]struct{}{}

	for _, def := range defs.Definitions.Definition {
		if _, found := seen[def.Class]; !found {
			res = append(res, DefinitionClass(def.Class))
			seen[def.Class] = struct{}{}
		}
	}
	return res
}
