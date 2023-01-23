build:
	docker build -t main .

run:	build
		docker image ls
		docker run -d -p 8080:8080 main

del image:
	docker image rm main