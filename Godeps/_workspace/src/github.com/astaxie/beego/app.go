// Copyright 2014 beego Author. All Rights Reserved.
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

package beego

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
	"time"

	"github.com/astaxie/beego/context"
)

// FilterFunc defines filter function type.
type FilterFunc func(*context.Context)

// App defines beego application with a new PatternServeMux.
type App struct {
	Handlers *ControllerRegistor
	Server   *http.Server
}

// NewApp returns a new beego application.
func NewApp() *App {
	cr := NewControllerRegister()
	app := &App{Handlers: cr, Server: &http.Server{}}
	return app
}

// Run beego application.
func (app *App) Run() {
	addr := HttpAddr

	if HttpPort != 0 {
		addr = fmt.Sprintf("%s:%d", HttpAddr, HttpPort)
	}

	var (
		err error
		l   net.Listener
	)
	endRunning := make(chan bool, 1)

	if UseFcgi {
		if HttpPort == 0 {
			l, err = net.Listen("unix", addr)
		} else {
			l, err = net.Listen("tcp", addr)
		}
		if err != nil {
			BeeLogger.Critical("Listen: ", err)
		}
		err = fcgi.Serve(l, app.Handlers)
	} else {
		app.Server.Addr = addr
		app.Server.Handler = app.Handlers
		app.Server.ReadTimeout = time.Duration(HttpServerTimeOut) * time.Second
		app.Server.WriteTimeout = time.Duration(HttpServerTimeOut) * time.Second

		if EnableHttpTLS {
			go func() {
				time.Sleep(20 * time.Microsecond)
				if HttpsPort != 0 {
					app.Server.Addr = fmt.Sprintf("%s:%d", HttpAddr, HttpsPort)
				}
				BeeLogger.Info("Running on %s", app.Server.Addr)
				err := app.Server.ListenAndServeTLS(HttpCertFile, HttpKeyFile)
				if err != nil {
					BeeLogger.Critical("ListenAndServeTLS: ", err)
					time.Sleep(100 * time.Microsecond)
					endRunning <- true
				}
			}()
		}

		if EnableHttpListen {
			go func() {
				app.Server.Addr = addr
				BeeLogger.Info("Running on %s", app.Server.Addr)
				err := app.Server.ListenAndServe()
				if err != nil {
					BeeLogger.Critical("ListenAndServe: ", err)
					time.Sleep(100 * time.Microsecond)
					endRunning <- true
				}
			}()
		}
	}

	<-endRunning
}
