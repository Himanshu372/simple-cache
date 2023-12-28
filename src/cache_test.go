package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalCache(t *testing.T) {
	m := map[string]string{
		"KB5031009": "Security update",
		"KB5031010": "Security update",
		"KB5029917": "Monthly update",
		"KB5029916": "Monthly update",
		"KB5029915": "Monthly update",
	}
	lc, err := NewLocalCache(5)
	if err != nil {
		t.Errorf("following error occurred when creating LocalCache: %s", err.Error())
	}
	for k, v := range m {
		lc.Set(k, v)
	}
	lc.Get("KB5029917")
	lc.Get("KB5029916")
	lc.Get("KB5029915")
	lc.Set("KB5029918", "Monthly update")
	expectedLC := map[string]string{
		"KB5029918": "Monthly update",
		"KB5029917": "Monthly update",
		"KB5029916": "Monthly update",
		"KB5029915": "Monthly update",
		"KB5031010": "Security update",
	}
	expectedRU := map[string]int{
		"KB5029917": 1,
		"KB5029916": 1,
		"KB5029915": 1,
		"KB5031010": 0,
		"KB5029918": 0,
	}
	assert.Equal(t, expectedLC, lc.c, "cache items should match")
	assert.Equal(t, expectedRU, lc.ru, "recently used count map should match")
}
