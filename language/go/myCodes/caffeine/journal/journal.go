package journal

import (
	"fmt"

	"github.com/docker/distribution/registry/storage/driver/filesystem"
)

// Module 3 - Caffeine journal showing statistics and graphs.
// x axis: day, y axis : amount of caffeine consumed.
// Showing a line graph of caffeine amount change

func DrawGraph() {
	fmt.Println("drawing a graph")
	filesystem.Driver()
}
