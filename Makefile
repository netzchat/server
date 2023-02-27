.PHONY: proto
proto:
	buf generate https://github.com/netzchat/proto.git \
		--path core/v1
