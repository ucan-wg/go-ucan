package issue

import (
	"time"

	"github.com/ucan-wg/go-ucan/delegation"
	"github.com/ucan-wg/go-ucan/did"
)

//go:generate -command options go run github.com/launchdarkly/go-options
//go:generate options -type=config -prefix=With -output=issue_options.go -cmp=false -imports=time

type config struct {
	Meta         map[string]any
	NoExpiration bool
	NotBefore    time.Time
}

func ToDelegateOptions(sub did.DID, opts ...Option) ([]delegation.Option, error) {
	cfg, err := newConfig(opts...)
	if err != nil {
		return nil, err
	}

	return []delegation.Option{
		delegation.WithSubject(&sub),
		delegation.WithMeta(cfg.Meta),
		delegation.WithNoExpiration(cfg.NoExpiration),
		delegation.WithNotBefore(&cfg.NotBefore),
	}, nil
}
