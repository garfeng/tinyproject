PROJECT=tinyproject

RS=$(PROJECT).rc
TARGET=build/$(PROJECT)
SYSO=$(PROJECT).syso

allLangs = $(shell ls language/translate_*.go | sed 's/language\/translate_//' | sed 's/.go//')

all:clean build_dir $(TARGET)
	rm $(SYSO) -rf
	rm lang_*.go -rf

build_dir:
	if [ ! -d build ] ; then mkdir build ; fi

# -H windowsgui
build/%:$(SYSO)
	for lang in $(allLangs);\
	do\
		echo $$lang && \
		cp language/translate_$$lang.go ./lang_map.go -rf && \
		go build -ldflags "-w -s" -o $@_$$lang.exe &&\
		upx $@_$$lang.exe -9 && \
		rm ./lang_map.go ;\
	done
	cp pngquant.exe build/
	cp config.json build/

%.syso:$(RS)
	windres.exe -o $@ $(RS)

test:clean build_dir $(SYSO)
	cp language/translate_cn.go ./lang_map.go -rf
	go build -o $(TARGET).exe

debug:
	cp language/translate_cn.go ./lang_map.go -rf
	go test

clean:
	rm build -rf
