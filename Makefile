start: build
	docker run \
		--name retina --detach \
		bodokaiser/retina

build:
	docker build -t bodokaiser/retina .
