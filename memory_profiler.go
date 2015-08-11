package main


import (
	"net/http"
	"github.com/wblakecaldwell/profiler"
)
func main() {
	// add handlers to help us track memory usage - they don't track memory until they're told to
	profiler.AddMemoryProfilingHandlers()

	// Uncomment if you want to start profiling automatically
	profiler.StartProfiling()

	// listen on port 6060 (pick a port)
	http.ListenAndServe(":6060", nil)
}

/**
http://localhost:6060/profiler/stop : Stop recording memory statistics
http://localhost:6060/profiler/start : Start recording memory statistics
http://localhost:6060/profiler/info.html : Main page you should visit
http://localhost:6060/profiler/info : JSON data that feeds profiler/info.html
 */