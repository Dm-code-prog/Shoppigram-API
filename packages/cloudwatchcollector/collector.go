package cloudwatchcollector

import (
	"context"
	"log"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

type (
	metricsCollector struct {
		client     *cloudwatch.Client
		namespace  string
		mu         sync.Mutex
		counters   map[string]float64
		dimensions map[string][]types.Dimension
		ticker     *time.Ticker
		done       chan bool
	}

	// Dimensions are used to add labels to metrics.
	Dimensions map[string]string
)

var (
	collector *metricsCollector
	once      sync.Once
)

// Init initializes the metricsCollector with the given namespace.
// It must be called before using the Increment function.
func Init(namespace string) {
	once.Do(func() {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatalf("Failed to load AWS config: %v", err)
		}
		client := cloudwatch.NewFromConfig(cfg)
		collector = &metricsCollector{
			client:     client,
			namespace:  namespace,
			counters:   make(map[string]float64),
			dimensions: make(map[string][]types.Dimension),
			ticker:     time.NewTicker(1 * time.Minute),
			done:       make(chan bool),
		}
		go collector.run()
	})
}

// Increment increases the counter for the given metric name and dimensions.
// It panics if Init has not been called.
func Increment(metricName string, dims map[string]string) {
	if collector == nil {
		panic("cloudwatch package not initialized. Please call cloudwatch.Init(namespace) before using Increment.")
	}
	key, dimensions := metricKey(metricName, dims)
	collector.mu.Lock()
	collector.counters[key] += 1
	collector.dimensions[key] = dimensions
	collector.mu.Unlock()
}

// metricKey generates a unique key for the metric based on its name and sorted dimensions.
func metricKey(name string, dims map[string]string) (string, []types.Dimension) {
	var dimensionKeys []string
	for k := range dims {
		dimensionKeys = append(dimensionKeys, k)
	}
	sort.Strings(dimensionKeys)

	var dimensionParts []string
	var dimensions []types.Dimension
	for _, k := range dimensionKeys {
		v := dims[k]
		dimensionParts = append(dimensionParts, k+"="+v)
		dimensions = append(dimensions, types.Dimension{
			Name:  aws.String(k),
			Value: aws.String(v),
		})
	}

	key := name + "|" + strings.Join(dimensionParts, ",")
	return key, dimensions
}

// run starts the periodic publishing of metrics to CloudWatch.
func (mc *metricsCollector) run() {
	for {
		select {
		case <-mc.ticker.C:
			mc.publishMetrics()
		case <-mc.done:
			mc.ticker.Stop()
			return
		}
	}
}

// publishMetrics sends the accumulated metrics to AWS CloudWatch.
func (mc *metricsCollector) publishMetrics() {
	mc.mu.Lock()
	if len(mc.counters) == 0 {
		mc.mu.Unlock()
		return
	}

	var metricData []types.MetricDatum
	for key, value := range mc.counters {
		metricName := extractMetricName(key)
		dimensions := mc.dimensions[key]

		metricDatum := types.MetricDatum{
			MetricName: aws.String(metricName),
			Dimensions: dimensions,
			Timestamp:  aws.Time(time.Now()),
			Value:      aws.Float64(value),
			Unit:       types.StandardUnitCount,
		}
		metricData = append(metricData, metricDatum)
	}

	// Clear the counters and dimensions after copying
	mc.counters = make(map[string]float64)
	mc.dimensions = make(map[string][]types.Dimension)
	mc.mu.Unlock()

	input := &cloudwatch.PutMetricDataInput{
		Namespace:  aws.String(mc.namespace),
		MetricData: metricData,
	}

	_, err := mc.client.PutMetricData(context.TODO(), input)
	if err != nil {
		log.Printf("Failed to put metric data: %v", err)
	}
}

// extractMetricName retrieves the metric name from the key.
func extractMetricName(key string) string {
	parts := strings.SplitN(key, "|", 2)
	return parts[0]
}

// Shutdown gracefully stops the metricsCollector.
func Shutdown() {
	if collector != nil {
		collector.done <- true
	}
}
