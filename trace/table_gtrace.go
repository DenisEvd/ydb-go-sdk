// Code generated by gtrace. DO NOT EDIT.

package trace

import (
	"context"
)

// Compose returns a new Table which has functional fields composed
// both from t and x.
func (t Table) Compose(x Table) (ret Table) {
	switch {
	case t.OnSessionNew == nil:
		ret.OnSessionNew = x.OnSessionNew
	case x.OnSessionNew == nil:
		ret.OnSessionNew = t.OnSessionNew
	default:
		h1 := t.OnSessionNew
		h2 := x.OnSessionNew
		ret.OnSessionNew = func(s SessionNewStartInfo) func(SessionNewDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionNewDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnSessionDelete == nil:
		ret.OnSessionDelete = x.OnSessionDelete
	case x.OnSessionDelete == nil:
		ret.OnSessionDelete = t.OnSessionDelete
	default:
		h1 := t.OnSessionDelete
		h2 := x.OnSessionDelete
		ret.OnSessionDelete = func(s SessionDeleteStartInfo) func(SessionDeleteDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionDeleteDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnSessionKeepAlive == nil:
		ret.OnSessionKeepAlive = x.OnSessionKeepAlive
	case x.OnSessionKeepAlive == nil:
		ret.OnSessionKeepAlive = t.OnSessionKeepAlive
	default:
		h1 := t.OnSessionKeepAlive
		h2 := x.OnSessionKeepAlive
		ret.OnSessionKeepAlive = func(k KeepAliveStartInfo) func(KeepAliveDoneInfo) {
			r1 := h1(k)
			r2 := h2(k)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(k KeepAliveDoneInfo) {
					r1(k)
					r2(k)
				}
			}
		}
	}
	switch {
	case t.OnSessionQueryPrepare == nil:
		ret.OnSessionQueryPrepare = x.OnSessionQueryPrepare
	case x.OnSessionQueryPrepare == nil:
		ret.OnSessionQueryPrepare = t.OnSessionQueryPrepare
	default:
		h1 := t.OnSessionQueryPrepare
		h2 := x.OnSessionQueryPrepare
		ret.OnSessionQueryPrepare = func(s SessionQueryPrepareStartInfo) func(PrepareDataQueryDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PrepareDataQueryDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnSessionQueryExecute == nil:
		ret.OnSessionQueryExecute = x.OnSessionQueryExecute
	case x.OnSessionQueryExecute == nil:
		ret.OnSessionQueryExecute = t.OnSessionQueryExecute
	default:
		h1 := t.OnSessionQueryExecute
		h2 := x.OnSessionQueryExecute
		ret.OnSessionQueryExecute = func(e ExecuteDataQueryStartInfo) func(SessionQueryPrepareDoneInfo) {
			r1 := h1(e)
			r2 := h2(e)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionQueryPrepareDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnSessionQueryStreamExecute == nil:
		ret.OnSessionQueryStreamExecute = x.OnSessionQueryStreamExecute
	case x.OnSessionQueryStreamExecute == nil:
		ret.OnSessionQueryStreamExecute = t.OnSessionQueryStreamExecute
	default:
		h1 := t.OnSessionQueryStreamExecute
		h2 := x.OnSessionQueryStreamExecute
		ret.OnSessionQueryStreamExecute = func(s SessionQueryStreamExecuteStartInfo) func(SessionQueryStreamExecuteDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionQueryStreamExecuteDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnSessionQueryStreamRead == nil:
		ret.OnSessionQueryStreamRead = x.OnSessionQueryStreamRead
	case x.OnSessionQueryStreamRead == nil:
		ret.OnSessionQueryStreamRead = t.OnSessionQueryStreamRead
	default:
		h1 := t.OnSessionQueryStreamRead
		h2 := x.OnSessionQueryStreamRead
		ret.OnSessionQueryStreamRead = func(s SessionQueryStreamReadStartInfo) func(SessionQueryStreamReadDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionQueryStreamReadDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnSessionTransactionBegin == nil:
		ret.OnSessionTransactionBegin = x.OnSessionTransactionBegin
	case x.OnSessionTransactionBegin == nil:
		ret.OnSessionTransactionBegin = t.OnSessionTransactionBegin
	default:
		h1 := t.OnSessionTransactionBegin
		h2 := x.OnSessionTransactionBegin
		ret.OnSessionTransactionBegin = func(s SessionTransactionBeginStartInfo) func(SessionTransactionBeginDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionTransactionBeginDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnSessionTransactionCommit == nil:
		ret.OnSessionTransactionCommit = x.OnSessionTransactionCommit
	case x.OnSessionTransactionCommit == nil:
		ret.OnSessionTransactionCommit = t.OnSessionTransactionCommit
	default:
		h1 := t.OnSessionTransactionCommit
		h2 := x.OnSessionTransactionCommit
		ret.OnSessionTransactionCommit = func(s SessionTransactionCommitStartInfo) func(SessionTransactionCommitDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionTransactionCommitDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnSessionTransactionRollback == nil:
		ret.OnSessionTransactionRollback = x.OnSessionTransactionRollback
	case x.OnSessionTransactionRollback == nil:
		ret.OnSessionTransactionRollback = t.OnSessionTransactionRollback
	default:
		h1 := t.OnSessionTransactionRollback
		h2 := x.OnSessionTransactionRollback
		ret.OnSessionTransactionRollback = func(s SessionTransactionRollbackStartInfo) func(SessionTransactionRollbackDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionTransactionRollbackDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnPoolInit == nil:
		ret.OnPoolInit = x.OnPoolInit
	case x.OnPoolInit == nil:
		ret.OnPoolInit = t.OnPoolInit
	default:
		h1 := t.OnPoolInit
		h2 := x.OnPoolInit
		ret.OnPoolInit = func(p PoolInitStartInfo) func(PoolInitDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PoolInitDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnPoolClose == nil:
		ret.OnPoolClose = x.OnPoolClose
	case x.OnPoolClose == nil:
		ret.OnPoolClose = t.OnPoolClose
	default:
		h1 := t.OnPoolClose
		h2 := x.OnPoolClose
		ret.OnPoolClose = func(p PoolCloseStartInfo) func(PoolCloseDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PoolCloseDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnPoolRetry == nil:
		ret.OnPoolRetry = x.OnPoolRetry
	case x.OnPoolRetry == nil:
		ret.OnPoolRetry = t.OnPoolRetry
	default:
		h1 := t.OnPoolRetry
		h2 := x.OnPoolRetry
		ret.OnPoolRetry = func(p PoolRetryStartInfo) func(PoolRetryInternalInfo) func(PoolRetryDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(info PoolRetryInternalInfo) func(PoolRetryDoneInfo) {
					r11 := r1(info)
					r21 := r2(info)
					switch {
					case r11 == nil:
						return r21
					case r21 == nil:
						return r11
					default:
						return func(p PoolRetryDoneInfo) {
							r11(p)
							r21(p)
						}
					}
				}
			}
		}
	}
	switch {
	case t.OnPoolSessionNew == nil:
		ret.OnPoolSessionNew = x.OnPoolSessionNew
	case x.OnPoolSessionNew == nil:
		ret.OnPoolSessionNew = t.OnPoolSessionNew
	default:
		h1 := t.OnPoolSessionNew
		h2 := x.OnPoolSessionNew
		ret.OnPoolSessionNew = func(p PoolSessionNewStartInfo) func(PoolSessionNewDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PoolSessionNewDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnPoolSessionClose == nil:
		ret.OnPoolSessionClose = x.OnPoolSessionClose
	case x.OnPoolSessionClose == nil:
		ret.OnPoolSessionClose = t.OnPoolSessionClose
	default:
		h1 := t.OnPoolSessionClose
		h2 := x.OnPoolSessionClose
		ret.OnPoolSessionClose = func(p PoolSessionCloseStartInfo) func(PoolSessionCloseDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PoolSessionCloseDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnPoolPut == nil:
		ret.OnPoolPut = x.OnPoolPut
	case x.OnPoolPut == nil:
		ret.OnPoolPut = t.OnPoolPut
	default:
		h1 := t.OnPoolPut
		h2 := x.OnPoolPut
		ret.OnPoolPut = func(p PoolPutStartInfo) func(PoolPutDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PoolPutDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnPoolGet == nil:
		ret.OnPoolGet = x.OnPoolGet
	case x.OnPoolGet == nil:
		ret.OnPoolGet = t.OnPoolGet
	default:
		h1 := t.OnPoolGet
		h2 := x.OnPoolGet
		ret.OnPoolGet = func(p PoolGetStartInfo) func(PoolGetDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PoolGetDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnPoolWait == nil:
		ret.OnPoolWait = x.OnPoolWait
	case x.OnPoolWait == nil:
		ret.OnPoolWait = t.OnPoolWait
	default:
		h1 := t.OnPoolWait
		h2 := x.OnPoolWait
		ret.OnPoolWait = func(p PoolWaitStartInfo) func(PoolWaitDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PoolWaitDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnPoolTake == nil:
		ret.OnPoolTake = x.OnPoolTake
	case x.OnPoolTake == nil:
		ret.OnPoolTake = t.OnPoolTake
	default:
		h1 := t.OnPoolTake
		h2 := x.OnPoolTake
		ret.OnPoolTake = func(p PoolTakeStartInfo) func(PoolTakeWaitInfo) func(PoolTakeDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PoolTakeWaitInfo) func(PoolTakeDoneInfo) {
					r11 := r1(p)
					r21 := r2(p)
					switch {
					case r11 == nil:
						return r21
					case r21 == nil:
						return r11
					default:
						return func(p PoolTakeDoneInfo) {
							r11(p)
							r21(p)
						}
					}
				}
			}
		}
	}
	return ret
}
func (t Table) onSessionNew(s SessionNewStartInfo) func(SessionNewDoneInfo) {
	fn := t.OnSessionNew
	if fn == nil {
		return func(SessionNewDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionNewDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onSessionDelete(s SessionDeleteStartInfo) func(SessionDeleteDoneInfo) {
	fn := t.OnSessionDelete
	if fn == nil {
		return func(SessionDeleteDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionDeleteDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onSessionKeepAlive(k KeepAliveStartInfo) func(KeepAliveDoneInfo) {
	fn := t.OnSessionKeepAlive
	if fn == nil {
		return func(KeepAliveDoneInfo) {
			return
		}
	}
	res := fn(k)
	if res == nil {
		return func(KeepAliveDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onSessionQueryPrepare(s SessionQueryPrepareStartInfo) func(PrepareDataQueryDoneInfo) {
	fn := t.OnSessionQueryPrepare
	if fn == nil {
		return func(PrepareDataQueryDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(PrepareDataQueryDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onSessionQueryExecute(e ExecuteDataQueryStartInfo) func(SessionQueryPrepareDoneInfo) {
	fn := t.OnSessionQueryExecute
	if fn == nil {
		return func(SessionQueryPrepareDoneInfo) {
			return
		}
	}
	res := fn(e)
	if res == nil {
		return func(SessionQueryPrepareDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onSessionQueryStreamExecute(s SessionQueryStreamExecuteStartInfo) func(SessionQueryStreamExecuteDoneInfo) {
	fn := t.OnSessionQueryStreamExecute
	if fn == nil {
		return func(SessionQueryStreamExecuteDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionQueryStreamExecuteDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onSessionQueryStreamRead(s SessionQueryStreamReadStartInfo) func(SessionQueryStreamReadDoneInfo) {
	fn := t.OnSessionQueryStreamRead
	if fn == nil {
		return func(SessionQueryStreamReadDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionQueryStreamReadDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onSessionTransactionBegin(s SessionTransactionBeginStartInfo) func(SessionTransactionBeginDoneInfo) {
	fn := t.OnSessionTransactionBegin
	if fn == nil {
		return func(SessionTransactionBeginDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionTransactionBeginDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onSessionTransactionCommit(s SessionTransactionCommitStartInfo) func(SessionTransactionCommitDoneInfo) {
	fn := t.OnSessionTransactionCommit
	if fn == nil {
		return func(SessionTransactionCommitDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionTransactionCommitDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onSessionTransactionRollback(s SessionTransactionRollbackStartInfo) func(SessionTransactionRollbackDoneInfo) {
	fn := t.OnSessionTransactionRollback
	if fn == nil {
		return func(SessionTransactionRollbackDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionTransactionRollbackDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onPoolInit(p PoolInitStartInfo) func(PoolInitDoneInfo) {
	fn := t.OnPoolInit
	if fn == nil {
		return func(PoolInitDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PoolInitDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onPoolClose(p PoolCloseStartInfo) func(PoolCloseDoneInfo) {
	fn := t.OnPoolClose
	if fn == nil {
		return func(PoolCloseDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PoolCloseDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onPoolRetry(p PoolRetryStartInfo) func(info PoolRetryInternalInfo) func(PoolRetryDoneInfo) {
	fn := t.OnPoolRetry
	if fn == nil {
		return func(PoolRetryInternalInfo) func(PoolRetryDoneInfo) {
			return func(PoolRetryDoneInfo) {
				return
			}
		}
	}
	res := fn(p)
	if res == nil {
		return func(PoolRetryInternalInfo) func(PoolRetryDoneInfo) {
			return func(PoolRetryDoneInfo) {
				return
			}
		}
	}
	return func(info PoolRetryInternalInfo) func(PoolRetryDoneInfo) {
		res := res(info)
		if res == nil {
			return func(PoolRetryDoneInfo) {
				return
			}
		}
		return res
	}
}
func (t Table) onPoolSessionNew(p PoolSessionNewStartInfo) func(PoolSessionNewDoneInfo) {
	fn := t.OnPoolSessionNew
	if fn == nil {
		return func(PoolSessionNewDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PoolSessionNewDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onPoolSessionClose(p PoolSessionCloseStartInfo) func(PoolSessionCloseDoneInfo) {
	fn := t.OnPoolSessionClose
	if fn == nil {
		return func(PoolSessionCloseDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PoolSessionCloseDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onPoolPut(p PoolPutStartInfo) func(PoolPutDoneInfo) {
	fn := t.OnPoolPut
	if fn == nil {
		return func(PoolPutDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PoolPutDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onPoolGet(p PoolGetStartInfo) func(PoolGetDoneInfo) {
	fn := t.OnPoolGet
	if fn == nil {
		return func(PoolGetDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PoolGetDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onPoolWait(p PoolWaitStartInfo) func(PoolWaitDoneInfo) {
	fn := t.OnPoolWait
	if fn == nil {
		return func(PoolWaitDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PoolWaitDoneInfo) {
			return
		}
	}
	return res
}
func (t Table) onPoolTake(p PoolTakeStartInfo) func(PoolTakeWaitInfo) func(PoolTakeDoneInfo) {
	fn := t.OnPoolTake
	if fn == nil {
		return func(PoolTakeWaitInfo) func(PoolTakeDoneInfo) {
			return func(PoolTakeDoneInfo) {
				return
			}
		}
	}
	res := fn(p)
	if res == nil {
		return func(PoolTakeWaitInfo) func(PoolTakeDoneInfo) {
			return func(PoolTakeDoneInfo) {
				return
			}
		}
	}
	return func(p PoolTakeWaitInfo) func(PoolTakeDoneInfo) {
		res := res(p)
		if res == nil {
			return func(PoolTakeDoneInfo) {
				return
			}
		}
		return res
	}
}
func TableOnSessionNew(t Table, c context.Context) func(session sessionInfo, _ error) {
	var p SessionNewStartInfo
	p.Context = c
	res := t.onSessionNew(p)
	return func(session sessionInfo, e error) {
		var p SessionNewDoneInfo
		p.Session = session
		p.Error = e
		res(p)
	}
}
func TableOnSessionDelete(t Table, c context.Context, session sessionInfo) func(error) {
	var p SessionDeleteStartInfo
	p.Context = c
	p.Session = session
	res := t.onSessionDelete(p)
	return func(e error) {
		var p SessionDeleteDoneInfo
		p.Error = e
		res(p)
	}
}
func TableOnSessionKeepAlive(t Table, c context.Context, session sessionInfo) func(error) {
	var p KeepAliveStartInfo
	p.Context = c
	p.Session = session
	res := t.onSessionKeepAlive(p)
	return func(e error) {
		var p KeepAliveDoneInfo
		p.Error = e
		res(p)
	}
}
func TableOnSessionQueryPrepare(t Table, c context.Context, session sessionInfo, query string) func(query string, result dataQuery, cached bool, _ error) {
	var p SessionQueryPrepareStartInfo
	p.Context = c
	p.Session = session
	p.Query = query
	res := t.onSessionQueryPrepare(p)
	return func(query string, result dataQuery, cached bool, e error) {
		var p PrepareDataQueryDoneInfo
		p.Query = query
		p.Result = result
		p.Cached = cached
		p.Error = e
		res(p)
	}
}
func TableOnSessionQueryExecute(t Table, c context.Context, session sessionInfo, tx transactionInfo, query dataQuery, parameters queryParameters) func(prepared bool, result result, _ error) {
	var p ExecuteDataQueryStartInfo
	p.Context = c
	p.Session = session
	p.Tx = tx
	p.Query = query
	p.Parameters = parameters
	res := t.onSessionQueryExecute(p)
	return func(prepared bool, result result, e error) {
		var p SessionQueryPrepareDoneInfo
		p.Prepared = prepared
		p.Result = result
		p.Error = e
		res(p)
	}
}
func TableOnSessionQueryStreamExecute(t Table, c context.Context, session sessionInfo, query dataQuery, parameters queryParameters) func(result streamResult, _ error) {
	var p SessionQueryStreamExecuteStartInfo
	p.Context = c
	p.Session = session
	p.Query = query
	p.Parameters = parameters
	res := t.onSessionQueryStreamExecute(p)
	return func(result streamResult, e error) {
		var p SessionQueryStreamExecuteDoneInfo
		p.Result = result
		p.Error = e
		res(p)
	}
}
func TableOnSessionQueryStreamRead(t Table, c context.Context, session sessionInfo) func(result streamResult, _ error) {
	var p SessionQueryStreamReadStartInfo
	p.Context = c
	p.Session = session
	res := t.onSessionQueryStreamRead(p)
	return func(result streamResult, e error) {
		var p SessionQueryStreamReadDoneInfo
		p.Result = result
		p.Error = e
		res(p)
	}
}
func TableOnSessionTransactionBegin(t Table, c context.Context, session sessionInfo) func(tx transactionInfo, _ error) {
	var p SessionTransactionBeginStartInfo
	p.Context = c
	p.Session = session
	res := t.onSessionTransactionBegin(p)
	return func(tx transactionInfo, e error) {
		var p SessionTransactionBeginDoneInfo
		p.Tx = tx
		p.Error = e
		res(p)
	}
}
func TableOnSessionTransactionCommit(t Table, c context.Context, session sessionInfo, tx transactionInfo) func(error) {
	var p SessionTransactionCommitStartInfo
	p.Context = c
	p.Session = session
	p.Tx = tx
	res := t.onSessionTransactionCommit(p)
	return func(e error) {
		var p SessionTransactionCommitDoneInfo
		p.Error = e
		res(p)
	}
}
func TableOnSessionTransactionRollback(t Table, c context.Context, session sessionInfo, tx transactionInfo) func(error) {
	var p SessionTransactionRollbackStartInfo
	p.Context = c
	p.Session = session
	p.Tx = tx
	res := t.onSessionTransactionRollback(p)
	return func(e error) {
		var p SessionTransactionRollbackDoneInfo
		p.Error = e
		res(p)
	}
}
func TableOnPoolInit(t Table, c context.Context) func(limit int, keepAliveMinSize int) {
	var p PoolInitStartInfo
	p.Context = c
	res := t.onPoolInit(p)
	return func(limit int, keepAliveMinSize int) {
		var p PoolInitDoneInfo
		p.Limit = limit
		p.KeepAliveMinSize = keepAliveMinSize
		res(p)
	}
}
func TableOnPoolClose(t Table, c context.Context) func(error) {
	var p PoolCloseStartInfo
	p.Context = c
	res := t.onPoolClose(p)
	return func(e error) {
		var p PoolCloseDoneInfo
		p.Error = e
		res(p)
	}
}
func TableOnPoolRetry(t Table, c context.Context, idempotent bool) func(error) func(attempts int, _ error) {
	var p PoolRetryStartInfo
	p.Context = c
	p.Idempotent = idempotent
	res := t.onPoolRetry(p)
	return func(e error) func(int, error) {
		var p PoolRetryInternalInfo
		p.Error = e
		res := res(p)
		return func(attempts int, e error) {
			var p PoolRetryDoneInfo
			p.Attempts = attempts
			p.Error = e
			res(p)
		}
	}
}
func TableOnPoolSessionNew(t Table, c context.Context) func(session sessionInfo, _ error) {
	var p PoolSessionNewStartInfo
	p.Context = c
	res := t.onPoolSessionNew(p)
	return func(session sessionInfo, e error) {
		var p PoolSessionNewDoneInfo
		p.Session = session
		p.Error = e
		res(p)
	}
}
func TableOnPoolSessionClose(t Table, c context.Context, session sessionInfo) func() {
	var p PoolSessionCloseStartInfo
	p.Context = c
	p.Session = session
	res := t.onPoolSessionClose(p)
	return func() {
		var p PoolSessionCloseDoneInfo
		res(p)
	}
}
func TableOnPoolPut(t Table, c context.Context, session sessionInfo) func(error) {
	var p PoolPutStartInfo
	p.Context = c
	p.Session = session
	res := t.onPoolPut(p)
	return func(e error) {
		var p PoolPutDoneInfo
		p.Error = e
		res(p)
	}
}
func TableOnPoolGet(t Table, c context.Context) func(session sessionInfo, retryAttempts int, _ error) {
	var p PoolGetStartInfo
	p.Context = c
	res := t.onPoolGet(p)
	return func(session sessionInfo, retryAttempts int, e error) {
		var p PoolGetDoneInfo
		p.Session = session
		p.RetryAttempts = retryAttempts
		p.Error = e
		res(p)
	}
}
func TableOnPoolWait(t Table, c context.Context) func(session sessionInfo, _ error) {
	var p PoolWaitStartInfo
	p.Context = c
	res := t.onPoolWait(p)
	return func(session sessionInfo, e error) {
		var p PoolWaitDoneInfo
		p.Session = session
		p.Error = e
		res(p)
	}
}
func TableOnPoolTake(t Table, c context.Context, session sessionInfo) func() func(took bool, _ error) {
	var p PoolTakeStartInfo
	p.Context = c
	p.Session = session
	res := t.onPoolTake(p)
	return func() func(bool, error) {
		var p PoolTakeWaitInfo
		res := res(p)
		return func(took bool, e error) {
			var p PoolTakeDoneInfo
			p.Took = took
			p.Error = e
			res(p)
		}
	}
}
