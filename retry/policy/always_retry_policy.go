package policy

import "github.com/liuhailove/gretry/retry"

// AlwaysRetryPolicy 一种无穷重试策略
type AlwaysRetryPolicy struct {
	NeverRetryPolicy
}

func (a *AlwaysRetryPolicy) CanRetry(ctx retry.RtyContext) bool {
	return true
}
