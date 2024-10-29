SHELL := /bin/bash

anim.webm: anim.gif
	ffmpeg -y -i anim.gif -c vp9 -b:v 0 -crf 40 anim.webm

anim.gif: main.py
	source .venv/bin/activate && python main.py && deactivate

.PHONY: venv
venv:
	python -m venv .venv && source .venv/bin/activate && pip install -U uv && uv pip install -r requirements.txt

.PHONY: clean
clean:
	# manually run:
	# rm -f anim.gif
	# rm -f anim.webm
