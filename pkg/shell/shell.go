package shell

// yaml "gopkg.in/yaml.v2"

type ErConfig struct {
	Nodes  []Node `yaml:"nodes"`
	Config Config `yaml:"config" mapstructure:"config"`
}

type Node struct {
	Host string `yaml:"host"`
	SID  string `yaml:"sid"`
}

type Config struct {
	SrcNode string `yaml:"mm_src_node" mapstructure:"mm_src_node"`
	Rules   []Rule `yaml:"rules" mapstructure:"rules"`
}

type Rule struct {
	DstNode      string   `yaml:"mm_dst_node" mapstructure:"mm_dst_node"`
	SrcAddr      string   `yaml:"mm_src_addr" mapstructure:"mm_src_addr"`
	VRF          int32    `yaml:"vrf" mapstructure:"vrf"`
	TransitNodes []string `yaml:"transit_nodes" mapstructure:"transit_nodes"`
}

type ErParam struct {
	Param Param `yaml:"params" mapstructure:"params"`
}

type Param struct {
	Method      string `yaml:"method"`
	PacketNum   int32  `yaml:"packet_num" mapstructure:"packet_num"`
	PacketSize  int32  `yaml:"packet_size" mapstructure:"packet_size"`
	RepeatNum   int32  `yaml:"repeat_num" mapstructure:"repeat_num"`
	MeasNum     int32  `yaml:"meas_num" mapstructure:"meas_num"`
	SmaInterval int32  `yaml:"sma_interval" mapstructure:"sma_interval"`
}
