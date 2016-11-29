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
	@/usr/lib/go/bin/go run src/basic/hello_world.go
