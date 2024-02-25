package erp

// DocDef is a struct that is inspired by the DocDef in ERPNext
// It connects a JSON schema to a set of templates and rules
// The schema, templates and rules are looked up from the registry.
type DocDefCfg struct {
	Name     string `json:"name"`
	Schema   string `json:"schema"`
	Template string `json:"template"`
	Rule     string `json:"rule"`
}
type DocDef struct {
	rules     Rule
	schemas   Schema
	templates Template
}

func NewDocDef(cfg DocDefCfg) *DocDef {

}

type Rule interface {
	Validate(data interface{}) error
}

// Schema can be a JSON schema or ErpNext DocType schema or Ofbiz Entity schema
type Schema interface {
	Validate(data interface{}) error
	GetSchema() []byte
}

type Template interface{}
