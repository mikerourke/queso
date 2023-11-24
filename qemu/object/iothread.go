package object

import "github.com/mikerourke/queso"

// IOThread creates a dedicated event loop thread that devices can be assigned to.
// This is known as an IOThread. By default, device emulation happens in vCPU threads
// or the main event loop thread. This can become a scalability bottleneck.
// IOThreads allow device emulation and I/O to run on other host CPUs.
//
// The id parameter is a unique ID that will be used to reference this IOThread
// from a qemu.Device. Multiple devices can be assigned to an IOThread. Note that
// not all devices support an iothread parameter.
//
// IOThreads use an adaptive polling algorithm to reduce event loop latency.
// Instead of entering a blocking system call to monitor file descriptors and
// then pay the cost of being woken up when an event occurs, the polling algorithm
// spins waiting for events for a short time. The algorithm's default parameters
// are suitable for many cases but can be adjusted based on knowledge of the
// workload and/or host device latency.
//
// The pollMax parameter is the maximum number of nanoseconds to busy wait for
// events. Polling can be disabled by setting this value to 0.
//
// The pollGrow parameter is the multiplier used to increase the polling time
// when the algorithm detects it is missing events due to not polling long enough.
//
// The pollShrink parameter is the divisor used to decrease the polling time when
// the algorithm detects it is spending too long polling without encountering events.
//
// The aioMaxBatch parameter is the maximum number of requests in a batch for
// the AIO engine, 0 means that the engine will use its default.
func IOThread(
	id string,
	pollMax int,
	pollGrow int,
	pollShrink int,
	aioMaxBatch int,
) *queso.Option {
	props := []*queso.Property{
		queso.NewProperty("id", id),
		queso.NewProperty("poll-max-ns", pollMax),
		queso.NewProperty("poll-grow", pollGrow),
		queso.NewProperty("poll-shrink", pollShrink),
		queso.NewProperty("aio-max-batch", aioMaxBatch),
	}

	return queso.NewOption("object", "iothread", props...)
}
