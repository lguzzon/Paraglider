package models

type TargetBackendHolder struct {
	LbName string
	BackendName string
	BackendIP string
	BackendPort int
	LbIp string
	LbPort int
}

type IptableConstantsHolder struct {
	Prerouting string
	Postrouting string
	Nat string
	Dnat string
	Snat string
	Tcp string
}