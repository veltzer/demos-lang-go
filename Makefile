.PHONY: all
all:
	$(info doing [$@])

.PHONY: run
run:
	$(info doing [$@])
	@/usr/lib/go/bin/go run src/basic/hello_world.go
