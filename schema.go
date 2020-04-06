package openrpc_go_document

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/alecthomas/jsonschema"
	"github.com/go-openapi/spec"
)

func typeToSchema(opts *DocumentProviderParseOpts, ty reflect.Type) spec.Schema {
	if !jsonschemaPkgSupport(ty) {
		panic("FIXME")
	}

	rflctr := jsonschema.Reflector{
		AllowAdditionalProperties:  false,
		RequiredFromJSONSchemaTags: true,
		ExpandedStruct:             false,
		IgnoredTypes:               opts.SchemaIgnoredTypes,
		TypeMapper:                 opts.TypeMapper,
	}

	jsch := rflctr.ReflectFromType(ty)

	// Poor man's glue.
	// Need to get the type from the go struct -> json reflector package
	// to the swagger/go-openapi/jsonschema spec.
	// Do this with JSON marshaling.
	// Hacky? Maybe. Effective? Maybe.
	m, err := json.Marshal(jsch)
	if err != nil {
		log.Fatal(err)
	}
	sch := spec.Schema{}
	err = json.Unmarshal(m, &sch)
	if err != nil {
		log.Fatal(err)
	}

	// NOTE: Debug toggling.
	expand := true
	if expand {
		err = spec.ExpandSchema(&sch, &sch, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Again, this should be pluggable.
	handleDescriptionDefault := true
	if handleDescriptionDefault {
		if sch.Description == "" {
			sch.Description = fullTypeDescription(ty)
		}
	}

	return sch
}

func fullTypeDescription(ty reflect.Type) string {
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	pack := ty.PkgPath()
	if pack != "" {
		return fmt.Sprintf(`%s.%s`, ty.PkgPath(), ty.Name())
	}
	return ty.String()
}