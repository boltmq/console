.PHONY: all deps compile test cover clean

OUTDIR=.
OUTBIN=console

all:compile

compile:
	@go build -o $(OUTDIR)/$(OUTBIN) main.go
	@echo "compile success [console] -> $(OUTDIR)/$(OUTBIN)."

deps:
	@go get -v ./... 

test:

cover:

clean:
	@rm -fr $(OUTDIR)/$(OUTBIN) 
	@echo "clean ok."
