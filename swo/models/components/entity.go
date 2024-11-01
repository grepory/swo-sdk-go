// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Entity struct {
	ID   string  `json:"id"`
	Type string  `json:"type"`
	Name *string `json:"name,omitempty"`
}

func (o *Entity) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Entity) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *Entity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
