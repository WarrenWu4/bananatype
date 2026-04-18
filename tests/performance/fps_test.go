// Tests FPS performance of a typing test

package performance

import (
	"sync"
	"testing"
	"time"
)

type FPSCollector struct {
	frameChan chan bool
	count     int
	mu        sync.Mutex
}

func (c *FPSCollector) Name() string { return "Framerate" }
func (c *FPSCollector) Unit() string { return "FPS" }
func (c *FPSCollector) Collect(t time.Time) float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	// This is called by the ticker. Since ticker is 50ms, we scale up to 1s.
	// But actually, we want to measure frames since last collect.
	fps := float64(c.count) / 0.5 // ticker is 500ms
	c.count = 0
	return fps
}
func (c *FPSCollector) SetChannel(ch chan bool) {
	c.frameChan = ch
}
func (c *FPSCollector) Reset() {
	go func() {
		for range c.frameChan {
			c.mu.Lock()
			c.count++
			c.mu.Unlock()
		}
	}()
}

const minTargetFPS = 20.0
const targetAvgFPS = 30.0

func TestFPSTyping(t *testing.T) {
	data := runTypingTest(120, 100, &FPSCollector{})
	summary_stats := summarize(data)
	if *export {
		exportData(data, &FPSCollector{})
	}
	t.Logf("FPS - Min: %.2f FPS, Avg: %.2f FPS, Max: %.2f FPS", summary_stats.Min, summary_stats.Avg, summary_stats.Max)
}
