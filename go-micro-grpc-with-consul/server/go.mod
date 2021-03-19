module server

go 1.15

require (
	github.com/Microsoft/go-winio v0.4.17-0.20210211115548-6eac466e5fa3 // indirect
	github.com/asim/go-micro/plugins/registry/consul/v3 v3.0.0-20210227064844-df90f2ca63ff
	github.com/asim/go-micro/v3 v3.5.0
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/go-git/go-git/v5 v5.2.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20201106050909-4977a11b4351 // indirect
	github.com/marten-seemann/qtls-go1-15 v0.1.4 // indirect
	github.com/miekg/dns v1.1.40 // indirect
	github.com/mitchellh/mapstructure v1.3.1 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	golang.org/x/crypto v0.0.0-20210314154223-e6e6c4f2bb5b // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20210314195730-07df6a141424 // indirect
	golang.org/x/text v0.3.5 // indirect
	grpc-lib v0.0.0
)

replace grpc-lib => ../grpc-lib
