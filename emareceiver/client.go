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
	"fmt"
	"net"

	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/receiver/scrapererror"
	// registers the mysql driver
)

type client interface {
	Connect(m *Config) error
	//Connect() error
	getcpuLoad(now pcommon.Timestamp, errs *scrapererror.ScrapeErrors)
	Close() error
}

// type mySQLClient struct {
// 	connStr string
// 	client  *sql.DB
// }

type emaClient struct {
	conn              *net.TCPConn
	confignet.NetAddr `mapstructure:",squash"`
}

var _ client = (*emaClient)(nil)

func newEmaClient(conf *Config) client {

	return &emaClient{
		NetAddr: conf.NetAddr,
	}
}

// func (c *emaClient) Connect() error {

// 	var cl net.Conn

// 	cl, err := net.Dial("tcp", "localhost:43034")
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("client type: %T\n", cl)
// 	fmt.Printf("c.conn type: %T\n", c.conn)
// 	c.conn = cl.(*net.TCPConn)
// 	return nil
// }

func (c *emaClient) Connect(conf *Config) error {

	var cl net.Conn

	cl, err := net.Dial("tcp", "localhost:43034")
	if err != nil {
		return err
	}
	fmt.Printf("client type: %T\n", cl)
	fmt.Printf("c.conn type: %T\n", c.conn)
	c.conn = cl.(*net.TCPConn)
	return nil
}

// getcpuLoad
func (c *emaClient) getcpuLoad(now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {

	fmt.Printf("conn: %T\n", c.conn)
	fmt.Println("In getcpuLoad")
	//return nil
}

func (c *emaClient) Close() error {
	if c != nil {
		return c.Close()
	}
	return nil
}
