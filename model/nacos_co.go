package model

type NacosCo struct {
	ClusterIps []string `json:"clusterIps" binding:"required,verify_ip"`
}
