os	:=$(shell expr substr $(shell uname -s) 1 5)


ifeq ($(os),MINGW)
GOBIN :=$(subst \,/,$(GOBIN))
GOPATH :=$(subst \,/,$(GOPATH))
$(info path alter $$GOBIN=$(GOBIN))
$(info path alter $$GOPATH=$(GOPATH))

GOBIN :=/$(shell echo "$(GOBIN)" | cut -d';' -f1 | sed 's/://g')
GOPATH :=/$(shell echo "$(GOPATH)" | cut -d';' -f1 | sed 's/://g')
$(info cut first $$GOBIN=$(GOBIN))
$(info cut first $$GOPATH=$(GOPATH))

endif

BIN	:=""

# check GOBIN
ifneq ($(GOBIN),)
	BIN=$(GOBIN)
else
# check GOPATH
	ifneq ($(GOPATH),)
		BIN=$(GOPATH)/bin
	endif
endif

$(info $$BIN=$(BIN))

.PHONY: info build

info:
	@echo $(GOBIN)
	@echo $(GOPATH)

build:
	@cd cmd/ && go build && cd - &> /dev/null
	@mkdir -p '$(BIN)'
	@cp ./cmd/cmd.exe '$(BIN)'