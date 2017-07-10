PROJECT=tinyproject

RS=$(PROJECT).rc
TARGET=build/$(PROJECT)
SYSO=$(PROJECT).syso

allLangs = $(shell ls language/translate_*.go | sed 's/language\/translate_//' | sed 's/.go//')

all:clean $(TARGET)
	rm $(SYSO) -rf
	rm lang_*.go -rf


build/%:$(SYSO)
	for lang in $(allLangs);\
	do\
		echo $$lang && \
		cp language/translate_$$lang.go ./lang_map.go -rf && \
		go build -ldflags "-w -s -H windowsgui" -o $@_$$lang.exe &&\
		upx $@_$$lang.exe -9 && \
		rm ./lang_map.go ;\
	done

%.syso:$(RS)
	windres.exe -o $@ $(RS)

test:clean $(SYSO)
	cp language/translate_cn.go ./lang_map.go -rf
	go build -o $(TARGET).exe

debug:
	cp language/translate_cn.go ./lang_map.go -rf
	go test

clean:
	rm build/* -rf
	rm *.a -rf
	rm *.o -rf
	rm *.syso -rf
	rm *.txt -rf
	rm *.json -rf
	rm www -rf