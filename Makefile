# Author:  Tan Menglong <tanmenglong@gmail.com>
# Date:    Mon Dec  8 12:25:53 2014
#
# Make Target:
# ------------
# The Makefile provides the following targets to make:
#   $ make           compile and link
#   $ make clean     clean objects and the executable file
#
#===========================================================================

.PHONY : all dev deps output clean help main test

all : output

dev : deps
	go get github.com/goccmack/gocc

deps :
	go get github.com/colinmarc/hdfs
	go get github.com/awalterschulze/gographviz

output : main
	mkdir -p output/bin
	mv main/main output/bin/hpipe
	cp scripts/* output/bin/

main : deps
	cd main; go build

test : all
	cd dag/symbol; go test -v
	cd dag; go test -v

clean :
	rm -rf output

help :
	@echo 'Usage: make [TARGET]'
	@echo 'TARGETS:'
	@echo '  all       (=make) compile and link.'
	@echo '  dev       get dependencies for development.'
	@echo '  deps      get dependencies for compiling.'
	@echo '  clean     clean objects and the executable file.'
	@echo '  help      print this message.'
