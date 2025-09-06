# brew install just

set shell := ["zsh", "-ic"]

default: which-shell

which-shell:
	echo ${SHELL}
	pwd
