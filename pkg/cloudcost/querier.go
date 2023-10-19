package cloudcost

import (
	"context"
	"time"

	filter "github.com/opencost/opencost/pkg/filter21"
	"github.com/opencost/opencost/pkg/kubecost"
)

// Querier allows for querying ranges of CloudCost data
type Querier interface {
	Query(QueryRequest, context.Context) (*kubecost.CloudCostSetRange, error)
}

type QueryRequest struct {
	Start       time.Time
	End         time.Time
	AggregateBy []string
	Accumulate  kubecost.AccumulateOption
	Filter      filter.Filter
}
