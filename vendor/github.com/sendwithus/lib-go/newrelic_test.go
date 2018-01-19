package swu

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoopNewRelic(t *testing.T) {
	assert.NotPanics(t, func() {
		agent := NewNewRelicAgent(String("Testing"), nil)
		txn := agent.StartTransaction("Testing")
		defer txn.End()
		seg := agent.StartSegment("Test Segment", txn)
		defer seg.End()
	})
}
