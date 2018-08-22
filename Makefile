all: mountpoint/ambex

vendor:
	glide install

ambex: mountpoint/ambex

mountpoint/ambex: vendor main.go
	sh -c "export GOOS=linux; go install ./..."
	cp $(GOBIN)/ambex mountpoint

clean:
	rm -rf vendor ambex mountpoint/ambex
