# figure out what GOROOT is supposed to be
GOROOT ?= $(shell printf 't:;@echo $$(GOROOT)\n' | gomake -f -)
include $(GOROOT)/src/Make.inc

VERSION=$(shell git describe --tags)

TARG=noeq
GOFILES=\
	main.go\

include $(GOROOT)/src/Make.cmd

goinstall:
	goinstall .

tar: clean goinstall
	tar -czf $(TARG)-$(VERSION)-$(GOOS)-$(GOARCH).tar.gz $(TARG) README.md
