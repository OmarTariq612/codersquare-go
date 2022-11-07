dynamic:
	@go build

static:
	@go build -ldflags="-extldflags=-static" -tags sqlite_omit_load_extension
