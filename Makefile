TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=jamakase.com
NAMESPACE=custom
NAME=airbyte
BINARY=terraform-provider-${NAME}
VERSION=0.0.1
OS_ARCH=darwin_amd64

default: install

build:
	go build -o ${BINARY}

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

test-tf: install
	rm -r examples/.terraform
	rm examples/.terraform.lock.hcl
	terraform -chdir=examples init
	terraform -chdir=examples plan
	TF_LOG=DEBUG terraform -chdir=examples apply --auto-approve

generate:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v5.2.1 generate \
		-i https://raw.githubusercontent.com/airbytehq/airbyte/master/airbyte-api/src/main/openapi/config.yaml \
		-g go \
		--package-name airbyte_sdk \
		--git-repo-id terraform-prodiver-airbyte \
		--git-user-id jamakase \
		--global-property models,supportingFiles,apis,modelDocs=false,apiDocs=false \
		-p enumClassPrefix=true \
		-o /local/airbyte_sdk
	rm "${PWD}/airbyte_sdk/go.mod" "${PWD}/airbyte_sdk/go.sum" "${PWD}/airbyte_sdk/.travis.yml" "${PWD}/airbyte_sdk/git_push.sh"
testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m