SHELL = /bin/bash

2026EIOS.gif: everything-is-open-source.txt tileseq Makefile
	./tileseq -file everything-is-open-source.txt \
		-tile 7 \
		-scale 1 \
		-gap 2 \
		-linegap 3 \
		-dim 0.3 \
		-dimpad 2 \
		-textcolor white \
		-palette candy \
		-width 1200 \
		-height 560 \
		-freeze 500 \
		-scan 1 \
		-flux 48 \
		-reveal 72 \
		-delay 2 \
		-output 2026EIOS.gif
	gifsicle --colors 8 -O3 2026EIOS.gif > 2026EIOS.gif.wip && mv 2026EIOS.gif.wip 2026EIOS.gif

tileseq: tileseq.go
	go build tileseq.go

.PHONY: clean
clean:
	rm -f tileseq
