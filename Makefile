
CONFIG ?= config.toml,config/links.yml,config/ads.yml
ARGS ?= --config $(CONFIG) 

run:
	hugo serve --buildFuture --buildDrafts

build:
	hugo  --gc --minify --cleanDestinationDir
	touch public/.nojekyll

clean:
	rm -rf public resources

compress: compress.png

compress.prepare:
	sudo apt-get update && sudo apt-get install -y pngquant

compress.png:
	find ./ -type f -name "*.png" -not -name "*-fs8.png" |xargs pngquant -f

clean.png-fs8:
	find ./ -type f -name "*-fs8.png.tmp" |xargs rm -f 
	find ./ -type f -name "*-fs8.png" |xargs rm -f 