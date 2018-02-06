.PHONY: all deps compile test cover clean

CURDIR=$(realpath .)
OUTDIR=$(CURDIR)/bin
SRVDIR=$(OUTDIR)/
CONSOLEIR=$(CURDIR)/admin

all:compile
	@echo "\nrun: cd $(OUTDIR) && ./console --root=./sources/ --perfix=/sources/"

compile:
	@make -C admin OUTDIR=$(SRVDIR)
	#@cp -f $(CURDIR)/scripts/* $(OUTDIR)/ 

deps:
	@make -C admin deps

test:

pack:
	#@tar zcf broker.tar.gz $(OUTDIR)

cover:

clean:
	@rm -fr $(OUTDIR)
	@echo "clean ok."
