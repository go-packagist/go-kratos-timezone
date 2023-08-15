package timezone

import (
	"context"
	"time"
)

type options struct {
	timezone string
}

type Option func(o *options)

func WithTimezone(timezone string) Option {
	return func(o *options) {
		o.timezone = timezone
	}
}

func Timezone(opts ...Option) func(ctx context.Context) error {
	op := options{
		timezone: "UTC",
	}

	for _, opt := range opts {
		opt(&op)
	}

	return func(ctx context.Context) error {
		location, err := time.LoadLocation(op.timezone)
		if err != nil {
			return err
		}

		time.Local = location

		return nil
	}
}
