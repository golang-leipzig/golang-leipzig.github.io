crop.png: main.png
	convert -extract 400x200+105+80 main.png crop.png

main.pdf: main.tex
	pdflatex main.tex

main.png: main.pdf
	convert -define colorspace:auto-grayscale=false -background white -alpha remove -alpha off main.pdf main.png


.PHONY: clean
clean:
	rm -rf main.pdf
	rm -rf main.aux
	rm -rf main.log
	rm -rf main.png
	rm -rf crop.png

