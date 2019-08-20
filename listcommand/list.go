package listcommand

import (
	"encoding/json"
	"errors"
)

// List contains stack list *
type List struct {
	API               API         `json:"api"`
	Appliance         Appliance   `json:"appliance"`
	Attr              Attr        `json:"attr"`
	Bootaction        Bootaction  `json:"bootaction"`
	Box               Box         `json:"box"`
	Cart              Cart        `json:"cart"`
	Environment       Environment `json:"environment"`
	Firewall          Firewall    `json:"firwall"`
	Group             Group       `json:"group"`
	Host              Host        `json:"host"`
	Network           Network     `json:"network"`
	OS                OS          `json:"os"`
	Pallet            Pallet      `json:"pallet"`
	Route             Route       `json:"route"`
	StorageController Controller  `json:"controller"`
	HostSwitch        HostSwitch  `json:"switch"`
}

// emptyOutput can absorb an empty response from the stacki frontend
type emptyOutput struct {
	output interface{}
}

var errEmptyResponse = errors.New("empty")

// responseHandler accepts a []byte output from a stacki frontend and returns a struct representing a Stacki resource
// this will handle enpty output from the frontend as well
func responseHandler(b []byte) (interface{}, error) {

	appliances := []Appliance{}
	attrs := []Attr{}
	firewalls := []Firewall{}
	routes := []Route{}
	controllers := []Controller{}
	arrays := []Array{}
	emptyOutput := emptyOutput{}

	switch {
	case json.Unmarshal(b, &appliances) == nil:
		return appliances, nil
	case json.Unmarshal(b, &attrs) == nil:
		return attrs, nil
	case json.Unmarshal(b, &firewalls) == nil:
		return firewalls, nil
	case json.Unmarshal(b, &routes) == nil:
		return routes, nil
	case json.Unmarshal(b, &controllers) == nil:
		return controllers, nil
	case json.Unmarshal(b, &arrays) == nil:
		return arrays, nil
	case json.Unmarshal(b, &emptyOutput) == nil:
		return nil, errEmptyResponse
	default:
		return nil, errors.New("cannot unmarshal " + string(b) + " into a known struct")
	}
}
