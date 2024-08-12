package local

type ZtmLocalSpec interface {
	GetLocalUser() string
	GetLocalAgent() string
	GetLocalMesh() string
	GetLocalEndpointId() string
	GetLocalHostAddr() string
}
