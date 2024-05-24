build:
	sudo docker build -t forum .
run:
	sudo docker run --rm --name forum -p 8081:8081 forum

