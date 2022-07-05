package types

type NodeType string

const (
	GPUShareNode     NodeType = "GPUShare"
	QGPUNode         NodeType = "QGPU"
	GPUExclusiveNode NodeType = "GPUExclusive"
	GPUTopologyNode  NodeType = "GPUTopology"
	NormalNode       NodeType = "Normal"
	UnknownNode      NodeType = "unknown"
	AllKnownNode     NodeType = ""
)

type NodeTypeInfo struct {
	Name      NodeType
	Alias     string
	Shorthand string
}

var NodeTypeSlice = []NodeTypeInfo{
	{
		Name:      NormalNode,
		Alias:     "none",
		Shorthand: "n",
	},
	{
		Name:      GPUExclusiveNode,
		Alias:     "exclusive",
		Shorthand: "e",
	},
	{
		Name:      GPUTopologyNode,
		Alias:     "topology",
		Shorthand: "t",
	},
	{
		Name:      GPUShareNode,
		Alias:     "share",
		Shorthand: "s",
	},
	{
		Name:      QGPUNode,
		Alias:     "qgpu",
		Shorthand: "q",
	},
}

type CommonNodeInfo struct {
	Name        string   `json:"name" yaml:"name"`
	Description string   `json:"description" yaml:"description"`
	IP          string   `json:"ip" yaml:"ip"`
	Status      string   `json:"status" yaml:"status"`
	Role        string   `json:"role" yaml:"role"`
	Type        NodeType `json:"type" yaml:"type"`
}

/*
type CommonGPUNodeInfo struct {
	TotalGPUs           int             `json:"totalGPUs" yaml:"totalGPUs"`
	AllocatedGPUs       int             `json:"allocatedGPUs" yaml:"allocatedGPUs"`
	UnhealthyGPUs       int             `json:"unhealthyGPUs" yaml:"unhealthyGPUs"`
	TotalGPUMemory      float64         `json:"totalGPUMemory" yaml:"totalGPUMemory"`
	AllocatedGPUMemory  float64         `json:"allocatedGPUMemory" yaml:"allocatedGPUMemory"`
	UsedGPUMemory       float64         `json:"usedGPUMemory" yaml:"usedGPUMemory"`
	Devices             []GPUDeviceInfo `json:"devices" yaml:"devices"`
	DutyCycle           float64         `json:"dutyCycle" yaml:"dutyCycle"`
	GPUMetrics          NodeGpuMetric   `json:"gpuMetrics" yaml:"gpuMetrics"`
	GPUMetricsIsEnabled bool            `json:"gpuMetricsIsEnabled" yaml:"gpuMetricsIsEnabled"`
}
*/

type CommonGPUNodeInfo struct {
	TotalGPUs     float64              `json:"totalGPUs" yaml:"totalGPUs"`
	AllocatedGPUs float64              `json:"allocatedGPUs" yaml:"allocatedGPUs"`
	UnhealthyGPUs float64              `json:"unhealthyGPUs" yaml:"unhealthyGPUs"`
	GPUMetrics    []*AdvancedGpuMetric `json:"gpuMetrics" yaml:"gpuMetrics"`
}

type GPUDeviceInfo struct {
	ID                 string  `json:"id" yaml:"id"`
	TotalGPUMemory     float64 `json:"totalGPUMemory" yaml:"totalGPUMemory"`
	AllocatedGPUMemory float64 `json:"allocatedGPUMemory" yaml:"allocatedGPUMemory"`
	UsedGPUMemory      float64 `json:"usedGPUMemory" yaml:"usedGPUMemory"`
	DutyCycle          float64 `json:"dutyCycle" yaml:"dutyCycle"`
}

type GPUShareNodeInfo struct {
	PodInfos           []GPUSharePodInfo    `json:"instances" yaml:"instances"`
	TotalGPUMemory     float64              `json:"totalGPUMemory" yaml:"totalGPUMemory"`
	AllocatedGPUMemory float64              `json:"allocatedGPUMemory" yaml:"allocatedGPUMemory"`
	Devices            []GPUShareNodeDevice `json:"devices" yaml:"devices"`
	CommonGPUNodeInfo  `yaml:",inline" json:",inline"`
	CommonNodeInfo     `yaml:",inline" json:",inline"`
}

type GPUSharePodInfo struct {
	Name          string         `json:"name" yaml:"name"`
	Namespace     string         `json:"namespace" yaml:"namespace"`
	Status        string         `json:"status" yaml:"status"`
	RequestMemory int            `json:"requestGPUMemory" yaml:"requestGPUMemory"`
	Allocation    map[string]int `json:"allocation" yaml:"allocation"`
}

type GPUShareNodeDevice struct {
	Id                 string  `json:"id" yaml:"id"`
	TotalGPUMemory     float64 `json:"totalGPUMemory" yaml:"totalGPUMemory"`
	AllocatedGPUMemory float64 `json:"allocatedGPUMemory" yaml:"allocatedGPUMemory"`
}

type GPUExclusiveNodeInfo struct {
	PodInfos          []GPUExclusivePodInfo `json:"instances" yaml:"instances"`
	CommonNodeInfo    `yaml:",inline" json:",inline"`
	CommonGPUNodeInfo `yaml:",inline" json:",inline"`
}

type GPUExclusivePodInfo struct {
	Name       string `json:"name" yaml:"name"`
	Namespace  string `json:"namespace" yaml:"namespace"`
	Status     string `json:"status" yaml:"status"`
	RequestGPU int    `json:"requestGPUs" yaml:"requestGPUs"`
}

type NormalNodeInfo struct {
	CommonNodeInfo `yaml:",inline" json:",inline"`
}

type AllNodeInfo map[string][]interface{}

type GPUTopologyNodeInfo struct {
	PodInfos          []GPUTopologyPodInfo `json:"instances" yaml:"instances"`
	GPUTopology       GPUTopology          `json:"gpuTopology" yaml:"gpuTopology"`
	CommonGPUNodeInfo `yaml:",inline" json:",inline"`
	CommonNodeInfo    `yaml:",inline" json:",inline"`
	Devices           []GPUTopologyNodeDevice `yaml:"devices" yaml:"devices"`
}

type GPUTopology struct {
	LinkMatrix      [][]string  `json:"linkMatrix" yaml:"linkMatrix"`
	BandwidthMatrix [][]float32 `json:"bandwidthMatrix" yaml:"bandwidthMatrix"`
}

type GPUTopologyPodInfo struct {
	Name        string   `json:"name" yaml:"name"`
	Namespace   string   `json:"namespace" yaml:"namespace"`
	Status      string   `json:"status" yaml:"status"`
	RequestGPU  int      `json:"requestGPUs" yaml:"requestGPUs"`
	Allocation  []string `json:"allocation" yaml:"allocation"`
	VisibleGPUs []string `json:"visibleGPUs" yaml:"visibleGPUs"`
}

type GPUTopologyNodeDevice struct {
	Id      string `json:"id" yaml:"id"`
	Healthy bool   `json:"healthy" yaml:"healthy"`
	Status  string `json:"status" yaml:"status"`
}

type QGPUNodeInfo struct {
	PodInfos           []QGPUPodInfo    `json:"instances" yaml:"instances"`
	TotalGPUMemory     float64          `json:"totalGPUMemory" yaml:"totalGPUMemory"`
	AllocatedGPUMemory float64          `json:"allocatedGPUMemory" yaml:"allocatedGPUMemory"`
	Devices            []QGPUNodeDevice `json:"devices" yaml:"devices"`
	CommonGPUNodeInfo  `yaml:",inline" json:",inline"`
	CommonNodeInfo     `yaml:",inline" json:",inline"`
}

type QGPUPodInfo struct {
	Name          string             `json:"name" yaml:"name"`
	Namespace     string             `json:"namespace" yaml:"namespace"`
	Status        string             `json:"status" yaml:"status"`
	RequestMemory float64            `json:"requestGPUMemory" yaml:"requestGPUMemory"`
	Allocation    map[string]float64 `json:"allocation" yaml:"allocation"`
}

type QGPUNodeDevice struct {
	Id                 string  `json:"id" yaml:"id"`
	TotalGPUMemory     float64 `json:"totalGPUMemory" yaml:"totalGPUMemory"`
	AllocatedGPUMemory float64 `json:"allocatedGPUMemory" yaml:"allocatedGPUMemory"`
}
