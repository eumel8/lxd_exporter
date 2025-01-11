package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	lxd "github.com/lxc/incus/client"
	//lxd "github.com/lxc/lxd/client"
	"github.com/alecthomas/kingpin/v2"
	"github.com/eumel8/lxd_exporter/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	version = "staging-UNVERSIONED"

	port = kingpin.Flag(
		"port", "Provide the port to listen on").Default("9472").Int16()
)

func main() {
	logger := log.New(os.Stderr, "lxd_exporter: ", log.LstdFlags)

	kingpin.Version(version)
	kingpin.Parse()

	server, err := lxd.ConnectIncusUnix("/var/snap/lxd/common/lxd/unix.socket", nil)
	if err != nil {
		logger.Fatalf("Unable to contact LXD server: %s", err)
		return
	}

	prometheus.MustRegister(metrics.NewCollector(logger, server))
	http.Handle("/metrics", promhttp.Handler())

	serveAddress := fmt.Sprintf(":%d", *port)
	log.Print("Server listening on ", serveAddress)
	log.Fatal(http.ListenAndServe(serveAddress, nil))
}
