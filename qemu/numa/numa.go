// Package numa is used to define NUMA nodes for use with QEMU. It's sufficiently
// complex and different enough from the other options to merit its own package.
package numa

// Policy represents the NUMA policy to use for the host.
type Policy string

const (
	// PolicyDefault is the default host policy.
	PolicyDefault Policy = "default"

	// PolicyPreferred prefers the given host node list for allocation.
	PolicyPreferred Policy = "preferred"

	// PolicyBind restricts memory allocation to the given host node list.
	PolicyBind Policy = "bind"

	// PolicyInterleave interleaves memory allocations across the given host
	// node list.
	PolicyInterleave Policy = "interleave"
)
