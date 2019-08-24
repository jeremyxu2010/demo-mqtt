cur_dir = $(shell pwd)

.PHONY: start-mqtt-broker
start-mqtt-broker:
	docker run --rm -p 1883:1883 -p 8080:8080 -v $(cur_dir)/conf/broker-config.yaml:/etc/volantmq/config.yaml -v $(cur_dir)/persistence/data:/persistence/data --env VOLANTMQ_CONFIG=/etc/volantmq/config.yaml volantmq/volantmq

.PHONY: start-mqtt-client
start-mqtt-client:
	go run $(cur_dir)/cmd/simple.go

.PHONY: dump-mqtt-packet
dump-mqtt-packet:
	tcpdump -n -w dump/demo-mqtt-qos-0.dmp -i lo0 'tcp and host 127.0.0.1 and port 1883'