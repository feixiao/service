// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"net/http"
	_ "net/http/pprof"

	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/node_exporter/collector"
	"github.com/prometheus/common/version"
)

func init() {
	
}


var addr = flag.String("listen-address", ":9100", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()
	//http.Handle("/metrics", promhttp.Handler())
	prometheus.MustRegister(version.NewCollector("node_exporter"))
	nc, err := collector.NewNodeCollector()
	if err != nil {
		log.Fatalf("Couldn't create collector: %s", err)
	}
	log.Infof("Enabled collectors:")
	for n := range nc.Collectors {
		log.Infof(" - %s", n)
	}

	if err := prometheus.Register(nc); err != nil {
		log.Fatalf("Couldn't register collector: %s", err)
	}
	handler := promhttp.HandlerFor(prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			ErrorLog:      log.NewErrorLogger(),
			ErrorHandling: promhttp.ContinueOnError,
		})

	// TODO(ts): Remove deprecated and problematic InstrumentHandler usage.
	http.Handle("/metrics", prometheus.InstrumentHandler("prometheus", handler))

	log.Fatal(http.ListenAndServe(*addr, nil))
}


