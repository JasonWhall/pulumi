package codegen

import "github.com/pulumi/pulumi/pkg/v3/codegen/schema"

func visitTypeClosure(t schema.Type, visitor func(t schema.Type), seen Set) {
	if seen.Has(t) {
		return
	}
	seen.Add(t)

	visitor(t)

	switch st := t.(type) {
	case *schema.ArrayType:
		visitTypeClosure(st.ElementType, visitor, seen)
	case *schema.MapType:
		visitTypeClosure(st.ElementType, visitor, seen)
	case *schema.ObjectType:
		for _, p := range st.Properties {
			visitTypeClosure(p.Type, visitor, seen)
		}
	case *schema.UnionType:
		for _, e := range st.ElementTypes {
			visitTypeClosure(e, visitor, seen)
		}
	case *schema.InputType:
		visitTypeClosure(st.ElementType, visitor, seen)
	case *schema.OptionalType:
		visitTypeClosure(st.ElementType, visitor, seen)
	}
}

func VisitTypeClosure(properties []*schema.Property, visitor func(t schema.Type)) {
	seen := Set{}
	for _, p := range properties {
		visitTypeClosure(p.Type, visitor, seen)
	}
}
