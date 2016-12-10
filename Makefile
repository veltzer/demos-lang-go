.PHONY: all
all: tools.stamp
	@true

tools.stamp: templardefs/deps.py
	$(info doing [$@])
	@templar install_deps
	@make_helper touch-mkdir $@

.PHONY: run
run:
	$(info doing [$@])
	@/usr/bin/go run src/basic/hello_world.go

.PHONY: build
build:
	$(info doing [$@])
	@/usr/bin/go build -o out/hello_world src/basic/hello_world.go
