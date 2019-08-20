package listcommand

import "testing"

func TestHost_Host(t *testing.T) {
	// run a list command for all hosts and make sure that one of the hosts returned is our frontend (there should only be one)
	h := Host{}
	hosts, err := h.Host("")
	if err != nil {
		t.Fatalf("list host failed with error %s", err)
	}
	for _, host := range hosts {
		if host.Appliance == "frontend" {
			return
		}
	}
	t.Error("list host failed. unable to find expectd host with appliance frontend")
}

func TestHost_Attr(t *testing.T) {
	// run a list command for all hosts, grab the name of one of them, then do a search for host attrs and make sure is can find appliance
	h := Host{}
	hosts, err := h.Host("")
	if err != nil {
		t.Fatalf("list host attr set up failed. list host failed with error %s", err)
	}
	var hostName string
	for _, host := range hosts {
		if host.Appliance == "frontend" {
			hostName = host.Name
		}
	}
	if hostName == "" {
		t.Fatal("list host attr set up failed. unable to find the frontend in the list of hosts")
	}
	hostAttrs, err := h.Attr(hostName)
	if err != nil {
		t.Fatalf("list host attr failed with error %s", err)
	}
	for _, attr := range hostAttrs {
		if attr.AttrName == "appliance" {
			return
		}
	}
	t.Error("list host attr failed. unable to find expected appliance attr in list")
}

func TestHost_Firewall(t *testing.T) {
	// run a list command fot all hosts, grab the name of one of them, then search for host firewall rules and make sure there is one for SSH
	h := Host{}
	hosts, err := h.Host("")
	if err != nil {
		t.Fatalf("list host firewall set up failed. list host failed with error %s", err)
	}
	var hostName string
	for _, host := range hosts {
		if host.Appliance == "frontend" {
			hostName = host.Name
		}
	}
	if hostName == "" {
		t.Fatal("list host firewall set up failed. unable to find the frontend in the list of hosts")
	}
	firewalls, err := h.Firewall(hostName)
	if err != nil {
		t.Fatalf("list host firewall failed with error %s", err)
	}
	for _, firewall := range firewalls {
		if firewall.Name == "SSH" {
			return
		}
	}
	t.Error("list host firewall failed. unable to find expected rule SSH")
}

func TestHost_StorageController(t *testing.T) {
	// run a list host storage controller command and make sure there is at least one
	h := Host{}
	controllers, err := h.StorageController("")
	if err != nil {
		t.Fatalf("list host storage controller failed with error %s", err)
	}
	if len(controllers) == 0 {
		t.Error("list host storage controller failed. unable to find any storage controllers")
	}
}

func TestHost_Route(t *testing.T) {
	// run a list host network command and make sure there is at least one
	h := Host{}
	routes, err := h.Route("")
	if err != nil {
		t.Fatalf("list host route failed with error %s", err)
	}
	if len(routes) == 0 {
		t.Error("list host route failed. unable to find any host routes on the frontend")
	}
}

func TestHost_Interface(t *testing.T) {
	// run a list host interface command and make sure there is at least one
	h := Host{}
	ifaces, err := h.Interface("")
	if err != nil {
		t.Fatalf("list host interface failed with error %s", err)
	}
	if len(ifaces) == 0 {
		t.Error("list host interface failed. unable to find any host interfaces on the frontend")
	}
}
