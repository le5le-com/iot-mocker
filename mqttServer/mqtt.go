package mqttServer

import (
	"context"
	"crypto/tls"
	"le5le/iot-mocker/config"
	"le5le/iot-mocker/utils"
	"log/slog"
	"os"
	"time"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/rs/zerolog/log"
)

var MqttBroker *mqtt.Server

var ConnectedCount int64 // 连接数量
var interval int64
var LastPing int64

func Init() {
	level := new(slog.LevelVar)
	level.Set(slog.LevelError)
	mqttLog := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))

	mqttOpts := &mqtt.Options{
		Capabilities:             mqtt.NewDefaultServerCapabilities(),
		ClientNetWriteBufferSize: 4096,
		ClientNetReadBufferSize:  4096,
		SysTopicResendInterval:   1,
		InlineClient:             true,
		Logger:                   mqttLog,
	}
	if config.App.MqttBroker.MaximumMessageExpiryInterval < 0 {
		mqttOpts.Capabilities.MaximumMessageExpiryInterval = 0
	}

	MqttBroker = mqtt.New(mqttOpts)

	auths := make(auth.AuthRules, 0)
	if len(config.App.MqttBroker.Auths) > 0 {
		for _, item := range config.App.MqttBroker.Auths {
			auths = append(auths, auth.AuthRule{Username: auth.RString(item.Username), Password: auth.RString(item.Password), Allow: true})
		}
	} else if config.App.MqttBroker.Username != "" {
		auths = append(auths, auth.AuthRule{Username: auth.RString(config.App.MqttBroker.Username), Password: auth.RString(config.App.MqttBroker.Password), Allow: true})
	} else {
		auths = append(auths, auth.AuthRule{Remote: "*", Allow: true})
	}

	authRules := &auth.Ledger{
		Auth: auths,
		ACL: auth.ACLRules{
			{
				// 允许指定主题订阅与发布
				Filters: auth.Filters{
					"le5le-iot/properties/test": auth.ReadOnly,
					"le5le-iot/subscribe/ping":  auth.WriteOnly,
				},
			},
			{
				// 限制其他主题订阅与发布
				Filters: auth.Filters{
					"#": auth.Deny,
				},
			},
		},
	}

	err := MqttBroker.AddHook(new(auth.Hook), &auth.Options{
		Ledger: authRules,
	})
	if err != nil {
		log.Error().Err(err).Msgf("MQTT访问控制失败：%+v", authRules)
		return
	}

	tcp := listeners.NewTCP(listeners.Config{ID: "tcp", Address: ":" + config.App.MqttBroker.TcpPort})
	err = MqttBroker.AddListener(tcp)
	if err != nil {
		log.Error().Err(err).Msgf("MQTT监听TCP端口%s失败", config.App.MqttBroker.TcpPort)
		return
	}

	ws := listeners.NewWebsocket(listeners.Config{
		ID:      "ws",
		Address: ":" + config.App.MqttBroker.WsPort,
	})
	err = MqttBroker.AddListener(ws)
	if err != nil {
		log.Fatal().Err(err).Msgf("MQTT监听Websocket端口%s失败", config.App.MqttBroker.WsPort)
		return
	}

	// 加载TLS
	if config.App.Certfile != "" {
		cert, err := tls.LoadX509KeyPair(config.App.Certfile, config.App.Keyfile)
		if err == nil {
			tlsConfig := &tls.Config{
				Certificates: []tls.Certificate{cert},
			}

			tcp := listeners.NewTCP(listeners.Config{
				ID:        "tls",
				Address:   ":" + config.App.MqttBroker.TlsPort,
				TLSConfig: tlsConfig,
			})
			err = MqttBroker.AddListener(tcp)
			if err != nil {
				log.Fatal().Err(err).Msgf("MQTT监听TLS端口%s失败", config.App.MqttBroker.TlsPort)
			}

			ws := listeners.NewWebsocket(listeners.Config{
				ID:        "wss",
				Address:   ":" + config.App.MqttBroker.WssPort,
				TLSConfig: tlsConfig,
			})
			err = MqttBroker.AddListener(ws)
			if err != nil {
				log.Fatal().Err(err).Msgf("MQTT监听WSS端口%s失败", config.App.MqttBroker.WssPort)
			}
		}
	}

	MqttBroker.Subscribe("le5le-iot/subscribe/ping", 1, callbackFn)

	go func() {
		err := MqttBroker.Serve()
		if err != nil {
			log.Error().Err(err).Msgf("启动MQTT失败")
		}
	}()

	go Start(utils.Ctx)
}

func Start(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	var err error
	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("Mqtt虚拟设备服务退出")
			return
		case <-ticker.C:
			if ConnectedCount < 1 {
				continue
			}

			interval++

			// 10分钟没有ping，视为无订阅者
			if time.Now().Unix()-LastPing > 10*60 {
				ConnectedCount = 0
				// log.Info().Msgf("没有订阅者")
			}

			if interval >= config.App.Interval {
				interval = 0
				err = MqttBroker.Publish("le5le-iot/properties/test", []byte(GetDataJson()), false, 0)
				if err != nil {
					log.Err(err).Msgf(`Mqtt publish error`)
				}
			}

		}
	}
}

func callbackFn(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
	LastPing = time.Now().Unix()
}
