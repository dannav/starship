p = $(shell pwd)

setup: deps save-use-model test-model

deps:
	pip install tensorflow tensorflow_hub

save-use-model:
	python ./create_serving_use.py

test-model:
	saved_model_cli run --dir ./tfserving/use/00000001 --tag_set serve --signature_def serving_default --input_exprs 'text=["test string"]'

start-serving:
	docker pull tensorflow/serving
	docker run -p 8501:8501 -v $p/tfserving/use:/models/universal_encoder -e MODEL_NAME=universal_encoder -t tensorflow/serving

build-and-run-serving: setup start-serving