terraform-provider-readthedocs: *.go */*.go go.mod
	go build .

install: terraform-provider-readthedocs
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/BarnabyShearer/readthedocs/0.1.0/linux_amd64
	cp $+ ~/.terraform.d/plugins/registry.terraform.io/BarnabyShearer/readthedocs/0.1.0/linux_amd64
	-rm .terraform.lock.hcl
	terraform init
