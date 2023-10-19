package cloudcost

import (
	"time"

	"github.com/opencost/opencost/pkg/kubecost"
)

type Repository interface {
	Has(time.Time, string) (bool, error)
	Get(time.Time, string) (*kubecost.CloudCostSet, error)
	Keys() ([]string, error)
	Put(*kubecost.CloudCostSet) error
	Expire(time.Time) error
}
