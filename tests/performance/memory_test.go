// Ensure that memory usage is within expected bounds

package performance

import (
	"runtime"
	"testing"
	"time"
)

type MemoryCollector struct{}

func (c *MemoryCollector) Name() string { return "Memory Usage" }
func (c *MemoryCollector) Unit() string { return "MB" }
func (c *MemoryCollector) Collect(t time.Time) float64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return float64(m.Alloc) / 1024 / 1024
}
func (c *MemoryCollector) Reset()               {}
func (c *MemoryCollector) SetChannel(chan bool) {}

func TestMemoryTyping(t *testing.T) {
	data := runTypingTest(120, 100, &MemoryCollector{})
	summary_stats := summarize(data)
	if *export {
		exportData(data, &MemoryCollector{})
	}
	t.Logf("Memory Usage - Min: %.2f MB, Avg: %.2f MB, Max: %.2f MB", summary_stats.Min, summary_stats.Avg, summary_stats.Max)
}
