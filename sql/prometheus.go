// Copyright (C) 2022 The PuzzleDB Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sql

import (
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/cybergarage/go-logger/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	DefaultPrometheusPort                            = 9181
	DefaultPrometheusConnectionTimeout time.Duration = time.Second * 60
)

// PrometheusExporter is a prometheus exporter service.
type PrometheusExporter struct {
	httpServer *http.Server
	Addr       string
	Port       int
}

// NewPrometheusExporter returns a new prometheus exporter service.
func NewPrometheusExporter() *PrometheusExporter {
	return &PrometheusExporter{
		httpServer: nil,
		Addr:       "",
		Port:       DefaultPrometheusPort,
	}
}

// SetAddress sets an address of the expoter.
func (expoter *PrometheusExporter) SetAddress(addr string) {
	expoter.Addr = addr
}

// SetPort sets a port number of the expoter.
func (expoter *PrometheusExporter) SetPort(port int) {
	expoter.Port = port
}

// Start starts the expoter.
func (expoter *PrometheusExporter) Start() error {
	err := expoter.Stop()
	if err != nil {
		return err
	}

	addr := net.JoinHostPort(expoter.Addr, strconv.Itoa(expoter.Port))
	expoter.httpServer = &http.Server{ // nolint:exhaustruct
		Addr:        addr,
		ReadTimeout: DefaultPrometheusConnectionTimeout,
		Handler:     promhttp.Handler(),
	}

	c := make(chan error)
	go func() {
		c <- expoter.httpServer.ListenAndServe()
	}()

	select {
	case err = <-c:
	case <-time.After(time.Millisecond * 500):
		err = nil
	}

	log.Infof("prometheus exporter (%s) started", addr)

	return err
}

// Stop stops the Grpc expoter.
func (expoter *PrometheusExporter) Stop() error {
	if expoter.httpServer == nil {
		return nil
	}

	err := expoter.httpServer.Close()
	if err != nil {
		return err
	}

	addr := net.JoinHostPort(expoter.Addr, strconv.Itoa(expoter.Port))
	log.Infof("prometheus exporter (%s) terminated", addr)

	return nil
}
