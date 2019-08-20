package add

import (
	"fmt"

	"github.td.teradata.com/ck250037/stackilackey/cmd"
)

type bootaction struct {
}

/*
Bootaction will add a bootaction specification to the system.
Arguments

	{action}

	Label name for the bootaction. You can see the bootaction label names by
	executing: 'stack list bootaction [host(s)]'.


Parameters

	[args=string]

	The second line for a pxelinux definition (e.g., ks ramdisk_size=150000
	lang= devfs=nomount pxe kssendmac selinux=0)

	[kernel=string]

	The name of the kernel that is associated with this boot action.

	[os=string]

	Operating System for the bootaction.
	The default os is Redhat.

	[ramdisk=string]

	The name of the ramdisk that is associated with this boot action.

	[type=string]

	Type of bootaction. Either 'os' or 'install'.
*/
func (bootaction *bootaction) Bootaction(action, args, kernel, os, ramdisk, typeName string) ([]byte, error) {

	argKeys := []string{"args", "kernel", "os", "ramdisk", "type"}
	argValues := []interface{}{args, kernel, os, ramdisk, typeName}
	baseCommand := fmt.Sprintf("add bootaction %s", action)

	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}
