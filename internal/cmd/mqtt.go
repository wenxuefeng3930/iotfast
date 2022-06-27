package cmd

import (
	"context"
	"crypto/tls"
	"fmt"
	"iotfast/library/libErr"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"iotfast/server/mqtt/config"
	"iotfast/server/mqtt/pkg/pidfile"
	"iotfast/server/mqtt/server"

	_ "iotfast/server/mqtt/persistence"
	_ "iotfast/server/mqtt/plugin/prometheus"
	_ "iotfast/server/mqtt/topicalias/fifo"
	// 在这里import所有的插件（为了调用对应的init方法）
	//  _ "iotfast/server/mqtt/plugin/admin"
	//  _ "iotfast/server/mqtt/plugin/auth"
	//  _ "iotfast/server/mqtt/plugin/prometheus"
	// _ "path/to/your/plugin"
)

var (
	ConfigFile string
	logger     *zap.Logger
)

func init() {
	configDir := filepath.Join(os.Getenv("programdata"), "gmqtt")
	ConfigFile = path.Join(configDir, "gmqttd.yml")
	//rootCmd.AddCommand(command.NewReloadCommand())
}

func installSignal(srv server.Server) {
	// reload
	reloadSignalCh := make(chan os.Signal, 1)
	signal.Notify(reloadSignalCh, syscall.SIGHUP)

	// stop
	stopSignalCh := make(chan os.Signal, 1)
	signal.Notify(stopSignalCh, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-reloadSignalCh:
			var c config.Config
			var err error
			c, err = config.ParseConfig(ConfigFile)
			if err != nil {
				logger.Error("reload error", zap.Error(err))
				return
			}
			srv.ApplyConfig(c)
			logger.Info("gmqtt reloaded")
		case <-stopSignalCh:
			err := srv.Stop(context.Background())
			if err != nil {
				fmt.Fprint(os.Stderr, err.Error())
			}
		}
	}

}

func GetListeners(c config.Config) (tcpListeners []net.Listener, websockets []*server.WsServer, err error) {
	for _, v := range c.Listeners {
		var ln net.Listener
		if v.Websocket != nil {
			ws := &server.WsServer{
				Server: &http.Server{Addr: v.Address},
				Path:   v.Websocket.Path,
			}
			if v.TLSOptions != nil {
				ws.KeyFile = v.Key
				ws.CertFile = v.Cert
			}
			websockets = append(websockets, ws)
			continue
		}
		if v.TLSOptions != nil {
			var cert tls.Certificate
			cert, err = tls.LoadX509KeyPair(v.Cert, v.Key)
			if err != nil {
				return
			}
			ln, err = tls.Listen("tcp", v.Address, &tls.Config{
				Certificates: []tls.Certificate{cert},
			})
		} else {
			ln, err = net.Listen("tcp", v.Address)
		}
		tcpListeners = append(tcpListeners, ln)
	}
	return
}

// NewStartCmd creates a *cobra.Command object for start command.
func NewStartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start gmqtt broker",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			libErr.ErrExit(err)
			c, err := config.ParseConfig(ConfigFile)
			if os.IsNotExist(err) {
				libErr.ErrExit(err)
			} else {
				libErr.ErrExit(err)
			}
			if c.PidFile != "" {
				pid, err := pidfile.New(c.PidFile)
				if err != nil {
					libErr.ErrExit(fmt.Errorf("open pid file failed: %s", err))
				}
				defer pid.Remove()
			}

			tcpListeners, websockets, err := GetListeners(c)
			libErr.ErrExit(err)
			l, err := c.GetLogger(c.Log)
			libErr.ErrExit(err)
			logger = l

			s := server.New(
				server.WithConfig(c),
				server.WithTCPListener(tcpListeners...),
				server.WithWebsocketServer(websockets...),
				server.WithLogger(l),
			)

			err = s.Init()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
				return
			}
			go installSignal(s)
			err = s.Run()
			if err != nil {
				fmt.Fprint(os.Stderr, err.Error())
				os.Exit(1)
				return
			}
		},
	}
	return cmd
}

func SetConfigPath(path string) {
	ConfigFile = path
	fmt.Printf("Set ConfigFile:%v \n", ConfigFile)
}

func NewSimpleCmd() (err error) {
	libErr.ErrExit(err)
	fmt.Printf("ConfigFile:%v \n", ConfigFile)
	c, err := config.ParseConfig(ConfigFile)
	if os.IsNotExist(err) {
		libErr.ErrExit(err)
	} else {
		libErr.ErrExit(err)
	}
	if c.PidFile != "" {
		pid, err := pidfile.New(c.PidFile)
		if err != nil {
			libErr.ErrExit(fmt.Errorf("open pid file failed: %s", err))
		}
		defer pid.Remove()
	}

	tcpListeners, websockets, err := GetListeners(c)
	libErr.ErrExit(err)
	l, err := c.GetLogger(c.Log)
	libErr.ErrExit(err)
	logger = l

	s := server.New(
		server.WithConfig(c),
		server.WithTCPListener(tcpListeners...),
		server.WithWebsocketServer(websockets...),
		server.WithLogger(l),
	)

	err = s.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	go installSignal(s)
	err = s.Run()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
		return
	}
	return
}
