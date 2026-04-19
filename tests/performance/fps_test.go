// Tests FPS performance of a typing test

package performance

import (
	"sync"
	"testing"
	"time"
)

type FPSCollector struct {
	frameChan   chan bool
	count       int
	mu          sync.Mutex
	lastCollect time.Time
}

func (c *FPSCollector) Name() string { return "Framerate" }
func (c *FPSCollector) Unit() string { return "FPS" }

func (c *FPSCollector) Collect(t time.Time) float64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	elapsed := t.Sub(c.lastCollect).Seconds()
	var fps float64
	if elapsed > 0 {
		fps = float64(c.count) / elapsed
	}

	c.count = 0
	c.lastCollect = t
	return fps
}

func (c *FPSCollector) SetChannel(ch chan bool) {
	c.frameChan = ch
}

func (c *FPSCollector) Reset() {
	c.mu.Lock()
	c.count = 0
	c.lastCollect = time.Now()
	c.mu.Unlock()

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
