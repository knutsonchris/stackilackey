package set

import (
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

type bootaction struct {
}

/*
Bootaction will set a bootaction specification.
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

	[ramdisk=string]

	The name of the ramdisk that is associated with this boot action.

	[type=string]

	Type of bootaction. Either 'os' or 'install'.
*/
func (bootaction *bootaction) Bootaction(action, args, kernel, os, ramdisk, typeName string) ([]byte, error) {
	a := []interface{}{action, args, kernel, os, ramdisk, typeName}
	c := fmt.Sprintf("set bootaction %s args='%s' kernel='%s' os-'%s' ramdisk='%s' type='%s'", a...)
	return cmd.RunCommand(c)
}

/*
Args will update the args for a bootaction.
Arguments

	{action}

	Name of the bootaction that needs to be updated.


Parameters

	{args=string}

	Updated args value.

	[os=string]

	os type of the bootaction.

	[type=string]

	type of the bootaction. Can be install or os.
*/
func (bootaction *bootaction) Args(action, args, os, typeName string) ([]byte, error) {
	a := []interface{}{action, args, os, typeName}
	c := fmt.Sprintf("set bootaction args %s args='%s' os='%s' type='%s'", a...)
	return cmd.RunCommand(c)
}

/*
Kernel will update the kernel for a bootaction.
Arguments

	{action}

	Name of the bootaction that needs to be updated.


Parameters

	{kernel=string}

	Updated kernel value.

	[os=string]

	os type of the bootaction.

	[type=string]

	type of the bootaction. Can be install or os.
*/
func (bootaction *bootaction) Kernel(action, kernel, os, typeName string) ([]byte, error) {
	a := []interface{}{action, kernel, os, typeName}
	c := fmt.Sprintf("set bootaction kernel %s kernel='%s' os='%s' type='%s'", a...)
	return cmd.RunCommand(c)
}

/*
Ramdisk will update the ramdisk for a bootaction.
Arguments

	{action}

	Name of the bootaction that needs to be updated.


Parameters

	{ramdisk=string}

	Updated ramdisk value.

	[os=string]

	os type of the bootaction.

	[type=string]

	type of the bootaction. Can be install or os.
*/
func (bootaction *bootaction) Ramdisk(action, ramdisk, os, typeName string) ([]byte, error) {
	a := []interface{}{action, ramdisk, os, typeName}
	c := fmt.Sprintf("set bootaction ramdisk %s ramdisk='%s' os='%s' type='%s'", a...)
	return cmd.RunCommand(c)
}
