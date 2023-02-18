package Strategy

import "testing"

func TestContext_Execute2(t *testing.T) {
	strategy := NewStrategyB()
	context := NewContext()
	context.SetStrategy(strategy)
	context.Execute()
}
