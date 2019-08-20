package add

import (
	"fmt"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
)

type storage struct {
}

/*
Controller will add a global storage controller configuration for all the hosts in the cluster.
Parameters

	{arrayid=string}

	The 'arrayid' is used to determine which disks are grouped as part
	of the same array. For example, all the disks with arrayid of '1' will
	be part of the same array.

	Arrayids must be integers starting at 1 or greater. If the arrayid is
	'global', then 'hotspare' must have at least one slot definition (this
	is how one specifies a global hotspare).

	In addition, the arrays will be created in arrayid order, that is,
	the array with arrayid equal to 1 will be created first, arrayid
	equal to 2 will be created second, etc.

	[adapter=integer]

	Adapter address.

	[enclosure=integer]

	Enclosure address.

	[hotspare=integer]

	Slot address(es) of the hotspares associated with this array id. This
	can be a comma-separated list (like the 'slot' parameter). If the
	'arrayid' is 'global', then the specified slots are global hotspares.

	[raidlevel=integer]

	RAID level. Raid 0, 1, 5, 6, 10, 50, 60 are currently supported.

	[slot=integer]

	Slot address(es). This can be a comma-separated list meaning all disks
	in the specified slots will be associated with the same array
*/
func (storage *storage) Controller(arrayID, adapter, enclosure, raidlevel int, slot string) ([]byte, error) {
	args := []interface{}{arrayID, adapter, enclosure, raidlevel, slot}
	c := fmt.Sprintf("add storage controller arrayid='%d' adapter='%d' enclosure='%d' raidlevel='%d' slot='%s'", args...)
	return cmd.RunCommand(c)
}

/*
Partition will add a partition configuration to the database.
Arguments

	{scope}

	Zero or one argument. The argument is the scope: a valid os (e.g.,
	'redhat'), a valid appliance (e.g., 'backend') or a valid host
	(e.g., 'backend-0-0). No argument means the scope is 'global'.


Parameters

	{device=string}

	Disk device on which we are creating partitions

	{options=string}

	Options that need to be supplied while adding partitions.

	{size=int}

	Size of the partition.

	[mountpoint=string]

	Mountpoint to create

	[partid=int]

	The relative partition id for this partition. Partitions will be
	created in ascending partition id order.

	[type=string]

	Type of partition E.g: ext4, ext3, xfs, raid, etc.
*/
func (storage *storage) Partition(scope, device, options, mountpoint, typeName string, size, partID int) ([]byte, error) {
	args := []interface{}{scope, device, options, size, mountpoint, partID, typeName}
	c := fmt.Sprintf("add storage partition %s device='%s' options='%s' size='%d' mountpoint='%s', partid='%d' type='%s'", args...)
	return cmd.RunCommand(c)
}
