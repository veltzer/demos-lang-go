.PHONY: all
all: tools.stamp
	@true

tools.stamp: config/deps.py
	$(info doing [$@])
	@pymakehelper touch_mkdir $@

.PHONY: run
run:
	$(info doing [$@])
	@/usr/bin/go run src/basic/hello_world.go

.PHONY: build
build:
	$(info doing [$@])
	@/usr/bin/go build -o out/hello_world src/basic/hello_world.go
