##############
# parameters #
##############
# do you want to show the commands executed ?
DO_MKDBG?=0
# do you want dependency on the Makefile itself ?
DO_ALLDEP:=1
# do you want to build elf files?
DO_ELFS:=1

########
# code #
########
ALL:=

# silent stuff
ifeq ($(DO_MKDBG),1)
Q:=
# we are not silent in this branch
else # DO_MKDBG
Q:=@
#.SILENT:
endif # DO_MKDBG

SOURCES:=$(shell find src -name "*.go" -and -type f)
ELFS:=$(addsuffix .elf, $(basename $(SOURCES)))

ifeq ($(DO_ELFS),1)
ALL+=$(ELFS)
endif # DO_ELFS

#########
# rules #
#########
.PHONY: all
all: $(ALL)
	@true

.PHONY: debug
debug:
	$(info ALL is $(ALL))
	$(info SOURCES is $(SOURCES))
	$(info ELFS is $(ELFS))

.PHONY: run
run:
	$(info doing [$@])
	$(Q)go run src/basic/hello_world.go

.PHONY: clean_hard
clean_hard:
	$(Q)git clean -qffxd

.PHONY: clean
clean:
	$(Q)rm $(ALL)

############
# patterns #
############
$(ELFS): %.elf: %.go
	$(info doing [$@])
	$(Q)go build -o $@ $<
	$(Q)strip $@

##########
# alldep #
##########
ifeq ($(DO_ALLDEP),1)
.EXTRA_PREREQS+=$(foreach mk, ${MAKEFILE_LIST},$(abspath ${mk}))
endif # DO_ALLDEP
