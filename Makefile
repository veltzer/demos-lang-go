.PHONY: all
all: tools.stamp
	@true

tools.stamp: config/deps.py
	$(info doing [$@])
	@pymakehelper touch_mkdir $@

.PHONY: run
run:
	$(info doing [$@])
	@go run src/basic/hello_world.go

.PHONY: build
build:
	$(info doing [$@])
	@go build -o out/hello_world src/basic/hello_world.go

.PHONY: clean_hard
clean_hard:
	@git clean -qffxd
.PHONY: clean
clean:
	@rm tools.stamp
