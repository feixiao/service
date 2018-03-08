package main

import (
	"flag"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr              = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
	uniformDomain     = flag.Float64("uniform.domain", 0.0002, "The domain for the uniform distribution.")
	normDomain        = flag.Float64("normal.domain", 0.0002, "The domain for the normal distribution.")
	normMean          = flag.Float64("normal.mean", 0.00001, "The mean for the normal distribution.")
	oscillationPeriod = flag.Duration("oscillation-period", 10*time.Minute, "The duration of the rate oscillation period.")
)

var (
	// Create a summary to track fictional interservice RPC latencies for three
	// distinct services with different latency distributions. These services are
	// differentiated via a "service" label.
	rpcDurations = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "rpc_durations_seconds",
			Help:       "RPC latency distributions.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"service"},
	)
	// The same as above, but now as a histogram, and only for the normal
	// distribution. The buckets are targeted to the parameters of the
	// normal distribution, with 20 buckets centered on the mean, each
	// half-sigma wide.
	rpcDurationsHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "rpc_durations_histogram_seconds",
		Help:    "RPC latency distributions.",
		Buckets: prometheus.LinearBuckets(*normMean-5**normDomain, .5**normDomain, 20),
	})
)

func init() {
	// Register the summary and the histogram with Prometheus's default registry.
	prometheus.MustRegister(rpcDurations)
	prometheus.MustRegister(rpcDurationsHistogram)
}

func main() {
	flag.Parse()

	start := time.Now()

	oscillationFactor := func() float64 {
		return 2 + math.Sin(math.Sin(2*math.Pi*float64(time.Since(start))/float64(*oscillationPeriod)))
	}

	// Periodically record some sample latencies for the three services.
	go func() {
		for {
			v := rand.Float64() * *uniformDomain
			rpcDurations.WithLabelValues("uniform").Observe(v)
			time.Sleep(time.Duration(100*oscillationFactor()) * time.Millisecond)
		}
	}()

	go func() {
		for {
			v := (rand.NormFloat64() * *normDomain) + *normMean
			rpcDurations.WithLabelValues("normal").Observe(v)
			rpcDurationsHistogram.Observe(v)
			time.Sleep(time.Duration(75*oscillationFactor()) * time.Millisecond)
		}
	}()

	go func() {
		for {
			v := rand.ExpFloat64() / 1e6
			rpcDurations.WithLabelValues("exponential").Observe(v)
			time.Sleep(time.Duration(50*oscillationFactor()) * time.Millisecond)
		}
	}()


	go func() {
		time.Sleep(30*time.Second)
		log.Println("fx_test")
		test := prometheus.NewHistogram(prometheus.HistogramOpts{
			Name:    "fx_test",
			Help:    "fx_test",
			Buckets: prometheus.LinearBuckets(*normMean-5**normDomain, .5**normDomain, 20),
		})
		prometheus.MustRegister(test)	// 验证中途添加新的采集字段也是可行的
		ctx,cancel := context.WithCancel(context.Background())
		go func() {
			for {
				select {
				case <- ctx.Done():
					return 
				case <-time.After(2 * time.Second): 
					v := rand.ExpFloat64() / 1e6
					test.Observe(v)
				}
			}
		}()
		time.Sleep(30*time.Second)
		cancel()
	}()


	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}