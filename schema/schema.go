package schema

// Attribute represents an attribute of a ReferenceSchema.
type Attribute struct {
	Name            string       `json:"name"`
	Type            Type         `json:"type"`
	SubAttributes   []*Attribute `json:"subAttributes,omitempty"`
	MultiValued     bool         `json:"multiValued"`
	Description     string       `json:"description,omitempty"`
	Required        bool         `json:"required"`
	CanonicalValues []string     `json:"canonicalValues,omitempty"`
	CaseExact       bool         `json:"caseExact"`
	Mutability      Mutability   `json:"mutability"`
	Returned        Returned     `json:"returned"`
	Uniqueness      Uniqueness   `json:"uniqueness"`
	ReferenceTypes  []string     `json:"referenceTypes"`
}

// ForEachAttribute calls given function on itself all sub attributes recursively.
func (attribute *Attribute) ForEachAttribute(f func(attribute *Attribute)) {
	f(attribute)
	if attribute.Type == ComplexType {
		for _, subAttribute := range attribute.SubAttributes {
			subAttribute.ForEachAttribute(f)
		}
	}
}

type Mutability string

const (
	ReadOnly  Mutability = "readOnly"
	ReadWrite Mutability = "readWrite"
	Immutable Mutability = "immutable"
	WriteOnly Mutability = "writeOnly"
)

// ReferenceSchema represents a resource schema that is used to fuzz resources that are defined by this schema.
type ReferenceSchema struct {
	ID          string       `json:"id"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Attributes  []*Attribute `json:"attributes"`
}

// ForEachAttribute calls given function on all attributes recursively.
func (s ReferenceSchema) ForEachAttribute(f func(attribute *Attribute)) {
	for _, attribute := range s.Attributes {
		attribute.ForEachAttribute(f)
	}
}

type Returned string

const (
	Always  Returned = "always"
	Never   Returned = "never"
	Default Returned = "default"
	Request Returned = "request"
)

type Type string

const (
	StringType    Type = "string"
	BooleanType   Type = "boolean"
	BinaryType    Type = "binary"
	DecimalType   Type = "decimal"
	IntegerType   Type = "integer"
	DateTimeType  Type = "dateTime"
	ReferenceType Type = "reference"
	ComplexType   Type = "complex"
)

type Uniqueness string

const (
	None   Uniqueness = "none"
	Server Uniqueness = "server"
	Global Uniqueness = "global"
)