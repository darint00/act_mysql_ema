// Copyright  OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package emareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mysqlreceiver"

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/scrapererror"
	"go.uber.org/zap"
)

const (
	picosecondsInNanoseconds int64 = 1000
)

type emaScraper struct {
	emaclient client
	logger    *zap.Logger
	config    *Config
	mb        *MetricsBuilder
}

func newEmaScraper(
	settings component.ReceiverCreateSettings,
	config *Config,
) *emaScraper {
	return &emaScraper{
		logger: settings.Logger,
		config: config,
		mb:     NewMetricsBuilder(config.Metrics, settings.BuildInfo),
	}
}

// start starts the scraper by initializing the db client connection.
func (m *emaScraper) start(_ context.Context, host component.Host) error {

	//newClient := newEmaClient(m.config)
	var emaclient emaClient

	//err := sqlclient.Connect()
	err := emaclient.Connect(m.config)
	//err := newClient.Connect()
	if err != nil {
		return err
	}

	c := emaclient.conn

	//m.emaclient := newClient

	fmt.Printf("Typeof c: %T\n", c)
	fmt.Fprintf(c, "cpuLoad\n")
	c.SetReadDeadline(time.Now().Add(1 * time.Second))
	reply := make([]byte, 1024)
	_, err = c.Read(reply)
	if err != nil {
		fmt.Println("Error : ", err.Error())

	} else {
		fmt.Println("Data: ", string(reply))
		fmt.Println("Connection: ", c)
		fmt.Printf("Connection Type: %T\n", c)
	}

	time.Sleep(2 * time.Second)
	//m.emaclient = client

	return nil
}

// shutdown closes the db connection
func (m *emaScraper) shutdown(context.Context) error {
	if m.emaclient == nil {
		return nil
	}
	return m.emaclient.Close()
}

// scrape scrapes the mysql db metric stats, transforms them and labels them into a metric slices.
func (m *emaScraper) scrape(context.Context) (pmetric.Metrics, error) {

	fmt.Println("Scraping Data")
	//time.Sleep((1 * time.Second))
	// if *m.conn == nil {
	// 	return pmetric.Metrics{}, errors.New("failed to connect to http client")
	// }

	fmt.Println("Scraping Data1")
	errs := &scrapererror.ScrapeErrors{}
	fmt.Println("Scraping Data2")

	// collect cpuLoad
	now := pcommon.NewTimestampFromTime(time.Now())
	fmt.Println("Scraping Data3")
	m.emaclient.getcpuLoad(now, errs)
	fmt.Println("Scraping Data4")

	m.mb.EmitForResource(WithMysqlInstanceEndpoint(m.config.Endpoint))

	return m.mb.Emit(), errs.Combine()
	//return nil, nil
}
