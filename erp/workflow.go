package erp

import "time"

/**


 */

type URI string

type TuppleNameURI map[string]URI

type VM string

const (
	JS VM = "js"  // https://github.com/dop251/goja
	CEL VM = "cel"  //https://github.com/google/cel-go
	JOKER VM = "joker"  // https://github.com/candid82/joker
	GNO  VM = "gno"  // https://github.com/gnolang/gno/tree/master/gnovm
	SKYLARK VM = "starlark" // https://github.com/google/starlark-go/blob/master/doc/spec.md	
 )


// Metadata common for all the schema compound objects
type Metadata struct {
	Name          string
	Title         string
	Description   string
	LastUpdatedOn time.Time
	LastUpdatedBy URI
	ResourceURI   URI
	Roles         []string
	Tags          []string
	Context       map[string]interface{}
}

// Applet top level entity, a collection of Workflows
type Applet struct {
	Meta      Metadata
	Workflows []TuppleNameURI
}

// Workflow is a collection of Forms and Scriptlets
type Workflow struct {
	Meta      Metadata
	// map[step, action]
	Forms     map[string]Form
	Scripts   map[string]Scriptlet
	// can be form or scriptlet
	FirstStep string       
}


// Scriptlet part of the workflow can be executed on Applet Container (client side)  
// Microservices are called from the scriptlet
// variable assumed: nextStep "error" is a halt state. "ok" is completion state
type Scriptlet struct {
	VM  VM
	code []byte
}

//Form is a JSON form https://jsonforms.io/
type Form struct {
	Meta Metadata
    SchemaData URI
    SchemaUI   URI
	// map[EventName, ScriptletStepName]
	OnActions   map[string]string 
}