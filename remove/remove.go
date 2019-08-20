package remove

// Remove contains stack remove *
type Remove struct {
	Appliance         appliance   `json:"appliance"`
	Attr              attr        `json:"attr"`
	Bootaction        bootaction  `json:"bootaction"`
	Box               box         `json:"box"`
	Cart              cart        `json:"cart"`
	Environment       environment `json:"environment"`
	Firewall          firewall    `json:"firewall"`
	Group             group       `json:"group"`
	Host              host        `json:"host"`
	Network           network     `json:"network"`
	OS                os          `json:"os"`
	Pallet            pallet      `json:"pallet"`
	Route             route       `json:"route"`
	StorageController controller  `json:"controller"`
}
