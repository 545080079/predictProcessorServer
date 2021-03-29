target=predict_execute_server
$(shell GEN_PROTO_LIB ../proto/asw )
.PHONY: all clean proto

all: ${target}

${target}: clean
	go build -o ./bin/${target} ./server/main.go

clean:
	rm -f ./bin/${target}

proto:
	pwd
	GEN_PROTO_LIB ../proto/asw
	# 替换到pb中的omitempty,避免json序列化时的奇怪行为
	sed -i "" -e "s/,omitempty//g" ./pb3/asw/AswAuthSvr/*.go
	sed -i "" -e "s/,omitempty//g" ./pb3/asw/AswComponentSvr/*.go
	sed -i "" -e "s/,omitempty//g" ./pb3/asw/AswKafkaMsg/*.go
	sed -i "" -e "s/,omitempty//g" ./pb3/asw/AswProcessorSvr/*.go
	sed -i "" -e "s/,omitempty//g" ./pb3/asw/AswSchedulerSvr/*.go
	sed -i "" -e "s/,omitempty//g" ./pb3/asw/AswTemplateSvr/*.go