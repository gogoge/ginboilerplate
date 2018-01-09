.PHONY:

build-builder-image:
	docker build -f docker/Dockerfile.multistage -t ginboilerplate:builder --target build-env .
build:
	docker build -f docker/Dockerfile.multistage -t ginboilerplate:multistage .
start:
	docker run --rm -it -p 10443:10443 \
		-v `pwd`/logs:/logs \
		ginboilerplate:multistage 

		# -e "TZ=Asia/Taipei" \
		# -v /etc/localtime:/etc/localtime:ro \
		# -v /etc/timezone:/etc/timezone:ro \