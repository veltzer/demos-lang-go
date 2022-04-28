.PHONY: all
all: tools.stamp
	@true

tools.stamp: config/deps.py
	$(info doing [$@])
	$(Q)pymakehelper touch_mkdir $@

.PHONY: run
run:
	$(info doing [$@])
	$(Q)go run src/basic/hello_world.go

.PHONY: build
build:
	$(info doing [$@])
	$(Q)go build -o out/hello_world src/basic/hello_world.go

.PHONY: clean_hard
clean_hard:
	$(Q)git clean -qffxd

.PHONY: clean
clean:
	$(Q)rm tools.stamp
