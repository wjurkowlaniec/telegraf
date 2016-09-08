package min

import (
	"fmt"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/aggregators"
)

type Min struct {
	cache map[uint64]map[string]interface{}
}

var sampleConfig = `
`

func (m *Min) SampleConfig() string {
	return sampleConfig
}

func (m *Min) Description() string {
	return "Aggregate the minimum value of each numerical field."
}

func (m *Min) Apply(in telegraf.Metric) {
	fmt.Println("MIN AGGREGATOR: ", in.String())
}

func (m *Min) Start(acc telegraf.Accumulator) {
	m.cache = make(map[uint64]map[string]interface{})
}

func (m *Min) Stop() {
}

func init() {
	aggregators.Add("min", func() telegraf.Aggregator {
		return &Min{}
	})
}
