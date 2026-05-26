SHELL := /bin/bash

# Montage tunables: HEIGHT = matched row height, GUTTER = spacing/margin (px).
HEIGHT ?= 600
GUTTER ?= 20

.DEFAULT_GOAL := lgo-60-gopher-nix.png

ffraw.nix:
	termshot --raw-write ffraw.nix 'incus exec hello -- fastfetch -s none --pipe false'

ff.nix: ffraw.nix
	sed '/^[[:space:]]*$$/d' < ffraw.nix > ff.nix

ff.png: ff.nix
	termshot -f ff.png -- cat ff.nix

# Horizontal montage: gopher (transparent bg) + NixOS termshot window, each
# flattened onto black, scaled to the same height, with a black gutter between.
lgo-60-gopher-nix.png: leipzig-gopher.png ff.png
	montage \
		\( leipzig-gopher.png -background black -flatten \) \
		\( ff.png             -background black -flatten \) \
		-tile 2x1 -geometry x$(HEIGHT)+$(GUTTER)+0 -background black lgo-60-gopher-nix.png

.PHONY: clean
clean:
	rm -f ffraw.nix
	rm -f ff.nix
	rm -f ff.png
	rm -f lgo-60-gopher-nix.png

