package vault

type HealthResponse struct {
	ClusterID                  string `json:"cluster_id"`
	ClusterName                string `json:"cluster_name"`
	Initialized                bool   `json:"initialized"`
	PerformanceStandby         bool   `json:"performance_standby"`
	ReplicationDrMode          string `json:"replication_dr_mode"`
	ReplicationPerformanceMode string `json:"replication_performance_mode"`
	Sealed                     bool   `json:"sealed"`
	ServerTimeUtc              int    `json:"server_time_utc"`
	Standby                    bool   `json:"standby"`
	Version                    string `json:"version"`
}
