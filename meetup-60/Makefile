TARGETS := phd-thesis-summary-by-glm-4.7-flash.md

.PHONY: all
all: $(TARGETS)

phd-thesis-summary-by-glm-4.7-flash.md: phd-thesis.pdf
	docli convert --image-export-mode placeholder phd-thesis.pdf | recap -m glm-4.7-flash:latest -t 30m -s paperplus > phd-thesis-summary-by-glm-4.7-flash.md

.PHONY: clean
clean:
	rm -f phd-thesis-summary-by-glm-4.7-flash.md

