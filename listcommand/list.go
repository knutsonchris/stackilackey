package listcommand

import (
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
	StoragePartition  Partition   `json:"partition"`
	HostSwitch        HostSwitch  `json:"switch"`
}

// emptyOutput can absorb an empty response from the stacki frontend
type emptyOutput struct {
	output interface{}
}

var errEmptyResponse = errors.New("empty")
