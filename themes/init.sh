#!/bin/bash

cd $(dirname $0)

HUGO_THEME="git@github.com:tangx/maupassant-hugo.git"

git clone $HUGO_THEME maupassant-hugo
