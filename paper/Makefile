# NI-DKG paper. Requires LuaLaTeX (fontspec) + biber.
MAIN := main

.PHONY: all clean distclean deps
all: $(MAIN).pdf

$(MAIN).pdf: $(MAIN).tex $(wildcard appendix*.tex) $(wildcard fig-*.tex) references.bib
	latexmk -lualatex -interaction=nonstopmode -file-line-error $(MAIN).tex

clean:
	latexmk -c $(MAIN).tex
	rm -f $(MAIN).bbl

distclean:
	latexmk -C $(MAIN).tex

deps:
	sudo apt-get install -y --no-install-recommends \
	  texlive-latex-recommended texlive-latex-extra texlive-fonts-recommended \
	  texlive-science texlive-luatex latexmk biber
