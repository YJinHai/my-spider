package log

import "context"

type Logger interface {
	Log(level Level, keyvals ...interface{}) error
}

type logger struct {
	logs      []Logger
	prefix    []interface{}
	hasValuer bool
	ctx       context.Context
}

func (c *logger) Log(level Level, keyvals ...interface{}) error {
	kvs := make([]interface{}, 0, len(c.prefix)+len(keyvals))
	kvs = append(kvs, c.prefix...)
	if c.hasValuer {
		bindValues(c.ctx, kvs)
	}
	kvs = append(kvs, keyvals...)
	for _, l := range c.logs {
		if err := l.Log(level, kvs...); err != nil {
			return err
		}
	}
	return nil
}

// WithContext returns a shallow copy of l with its context changed
// to ctx. The provided ctx must be non-nil.
func WithContext(ctx context.Context, l Logger) Logger {
	if c, ok := l.(*logger); ok {
		return &logger{
			logs:      c.logs,
			prefix:    c.prefix,
			hasValuer: c.hasValuer,
			ctx:       ctx,
		}
	}
	return &logger{logs: []Logger{l}, ctx: ctx}
}
