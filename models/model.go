package models

type WatchConfigs struct {
	Service string   `yaml:"service" mapstructure:"service"`
	Command string   `yaml:"command" mapstructure:"command"`
	Args    []string `yaml:"args" mapstructure:"args"`
}

type DcmngrYaml struct {
	ShDefaultService       string       `yaml:"sh_default_service" mapstructure:"sh_default_service"`
	ShAlwaysAdmin          bool         `yaml:"sh_always_admin" mapstructure:"sh_always_admin"`
	BuildDefaultContainers []string     `yaml:"build_default_containers" mapstructure:"build_default_containers"`
	UpDefaultContainers    []string     `yaml:"up_default_containers" mapstructure:"up_default_containers"`
	WatchConfigs           WatchConfigs `yaml:"watch_configs" mapstructure:"watch_configs"`
}
