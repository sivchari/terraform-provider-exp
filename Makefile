default: testacc

.PHONY: testacc
testacc:
	TF_ACC=1 go test -v ./...
