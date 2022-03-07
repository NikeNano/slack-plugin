requirement: 
	poetry export -f requirements.txt -o requirements.txt --without-hashes

build: 
	docker build -f ../../python.Dockerfile --tag nikenano/pluginserverpython .

run:
	docker run -p 4355:4355 -e SLACK_BOT_TOKEN="" nikenano/pluginserverpython:202203061753