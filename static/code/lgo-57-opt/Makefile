SHELL := /bin/bash

.PHONY: run
run:
	go run lgo-57-warehouse-opt.go

.PHONY: wrapup
wrapup:
	# go run lgo-57-warehouse-opt.go -prefix lgo-57-opt -W 1020 -H 320 -placement random -seed 100 -n 20 -delay 40
	go run lgo-57-warehouse-opt.go -prefix lgo-57-wrapup -W 1200 -H 320 -n 24 -delay 40 -seed 12


.PHONY: clean
clean:
	rm -f *.gif
	rm -f *.png
	rm -f *.svg
