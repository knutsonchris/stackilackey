package add

import (
	"fmt"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
)

type cart struct {
}

/*
Arguments

	[cart]

	The name of the cart to be created.
	Can also be a URL to a cart that we would like to download and add.


Parameters

	[authfile=string]

	Json formatted authentication file with Username and Password. This can also
	contain URLs for the cart. See examples below.

	[downloaddir=string]

	Directory to download to. Defaults /tmp.

	[downloadonly=boolean]

	If you just want to download them, set downloadonly=True.

	[file=string]

	Add a local cart from a compressed file.

	[password=string]

	If the remote cart download server requires authentication.

	[url=string]

	Add cart from a single url.

	[urlfile=string]

	Add multiple carts from a textfile with urls.
	The urlfile is a simple newline-separated list of URLS

	[username=string]

	If the remote cart download server requires authentication.
*/
func (cart *cart) Cart(cartName, authfile, downloaddir, file, password, url, urlfile, username string, downloadonly bool) ([]byte, error) {
	var downloadonlystr string
	if downloadonly == true {
		downloadonlystr = "true"
	} else {
		downloadonlystr = "false"
	}

	argKeys := []string{"authfile", "downloaddir", "downloadonly", "file", "password", "url", "urlfile", "username"}
	argValues := []interface{}{authfile, downloaddir, downloadonlystr, file, password, url, urlfile, username}
	baseCommand := fmt.Sprintf("add cart %s", cartName)

	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)

}
