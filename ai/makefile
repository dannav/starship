p = $(shell pwd)

setup: deps save-use-model test-model

install-brew:
	/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

deps:
	brew install python
	sudo easy_install pip
	pip install --user nose --ignore-installed tornado --ignore-installed tensorflow --ignore-installed tensorflow_hub --ignore-installed

save-use-model:
	python ./create_serving_use.py

test-model:
	export PATH=$(PATH):~/Library/Python/2.7/bin/; \
	saved_model_cli run --dir $p/tfserving/use/00000001 --tag_set serve --signature_def serving_default --input_exprs 'text=["test string"]'

# only needed for debugging. start serving from the docker-compose file in services project
start-serving:
	docker pull tensorflow/serving
	docker run -p 8501:8501 -v $p/tfserving/use:/models/universal_encoder -e MODEL_NAME=universal_encoder -t tensorflow/serving

build-and-run-serving: setup start-serving