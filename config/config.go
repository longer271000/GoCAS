package config
type APPConfig struct {
	AppName    string 'json:app_name'
	Port       string 'json:port'
	StaticPath string 'json:static_path'
	Mode       string 'json:mode'
}
