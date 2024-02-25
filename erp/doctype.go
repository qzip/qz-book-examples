package erp

import "time"

// DocDef is a struct that is inspired by the DocDef in ERPNext
// It connects a JSON schema to a set of templates and rules
// The schema, templates and rules are looked up from the registry.
type DocDefCfg struct {
	Name       string `json:"name"`
	Schema     string `json:"schema"`
	Template   string `json:"template"`
	Rule       string `json:"rule"`
	DataAccess string `json:"dataAccess"`
}
type DocDef struct {
	Name     string
	rule     Rule
	schema   Schema
	template Template
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

type Role struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

type Permission struct {
	Name        string   `json:"name"`
	FieldGroups []string `json:"fieldGroups"`
}

// FieldGroup is a struct that is inspired ERPNext permission system
// empty fields means level applies to all fields
type FieldGroup struct {
	Name   string   `json:"name"`
	DocDef string   `json:"docDef"`
	Fields []string `json:"fields,omitempty"`
	Level  int      `json:"level"`
}

// Caser CAS Storage Provider, from imle/caser.go
type Caser interface {
	Store(data HashableData, ns string) (cas CasAddress, tmstamp time.Time, err error)
	Load(cas CasAddress) (data []byte, ns string, tmstamp time.Time, err error)
}

// HashableData must be implemnted for Caser storage
type HashableData interface {
	Hash() []byte
	Value() []byte
}
type CasAddress []byte

type DataAccessor interface {
	Caser
	Query(query string) (casAddresses []CasAddress, err error)
}
