.PHONY: all
all: tools.stamp
	$(info doing [$@])

tools.stamp: apt.yaml
	$(info doing [$@])
	@templar_cmd install_deps
	@make_helper touch-mkdir $@

.PHONY: run
run:
	$(info doing [$@])
	@/usr/lib/go/bin/go run src/basic/hello_world.go
