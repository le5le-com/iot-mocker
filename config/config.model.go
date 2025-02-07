package config

// Config App config model.
type Config struct {
	Name    string `json:"name" yaml:"name"`
	Version string `json:"version" yaml:"version"`
	Host    string `json:"host" yaml:"host"`
	Port    string `json:"port" yaml:"port"`
	CPU     int    `json:"cpu" yaml:"cpu"`

	// SSL证书路径
	Certfile string `json:"certfile" yaml:"certfile"`
	Keyfile  string `json:"keyfile" yaml:"keyfile"`

	Log struct {
		Level      int8   `json:"level" yaml:"level"`
		Filename   string `json:"filename" yaml:"filename"`
		MaxSize    int    `json:"maxSize" yaml:"maxSize"`
		MaxBackups int    `json:"maxBackups" yaml:"maxBackups"`
		MaxAge     int    `json:"maxAge" yaml:"maxAge"`
	} `json:"log" yaml:"log"`

	MqttBroker struct {
		Host          string `json:"host" yaml:"host"`
		TcpPort       string `json:"tcpPort" yaml:"tcpPort"`
		TlsPort       string `json:"tlsPort" yaml:"tlsPort"`
		WsPort        string `json:"wsPort" yaml:"wsPort"`
		WssPort       string `json:"wssPort" yaml:"wssPort"`
		TlsServerName string `json:"tlsServerName" yaml:"tlsServerName"`
		Username      string `json:"username" yaml:"username"`
		Password      string `json:"password" yaml:"password"`

		MaximumMessageExpiryInterval int64 `json:"maximumMessageExpiryInterval" yaml:"maximumMessageExpiryInterval"`
		Auths                        []struct {
			Username string `json:"username" yaml:"username"`
			Password string `json:"password" yaml:"password"`
		} `json:"auths" yaml:"auths"`
	} `json:"mqttBroker" yaml:"mqttBroker"`

	Interval int64 `json:"interval" yaml:"interval"`
}
