package client

import (
	"context"

	"github.com/rajeshsharmax2c/tolling/types"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error
	GetInvoice(context.Context, int) (*types.Invoice, error)
}
