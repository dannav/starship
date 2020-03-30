## Running Starship

This guide covers setting up Starship to run in a local development environment or on a server for you and your team.
We'll cover how to setup a machine learning model, environment variables for configuring your Starship instance, and how
to run all of the required services using Docker Compose.

Before stepping through this guide please read [the architecture overview](/getting-started/architecture-overview) for information
on how Starship is architected.

---

Creating a fault tolerant environment for hosting Starship is out of scope of this document. If you need help setting up
such an environment, feel free to reach out to Danny Navarro at [navdgo@gmail.com](mailto:navdgo@gmail.com) for further information.

---

To get up and running immediately, check out the [TLDR section](#tldr) below.

## System Requirements

**System**: Linux (Kernal 4.10 or later) or Mac OSX (High Sierra or later)

**Memory**: At least 4GB of free RAM

**Storage**: Storage **Required** is dependent on how much data is going to be indexed by Starship and how big your machine learning model is.
It is recommended to have at least 20GB of storage on the server hosting Starship.

For development purposes go version `1.10+` is required.

## Docker

To make building and running Starship straight forward, Docker is used for build and server environments. Follow the [Docker documentation](https://docs.docker.com/install/)
on how to install Docker CE.

## Setup Go

Starship is built in go. Visit the [download page of the go website](https://golang.org/dl/) and follow the instructions to setup go on your machine.

## Building The Project

Starship uses `dep` to vendor project dependencies. Visit the [dep repository](https://github.com/golang/dep) for instructions on how to install `dep`.

Since some dependencies for Starship require CGO to be enabled, Docker containers are distributed with the source to create a reproducible build environment.
The Docker containers for the `searchd` service use Ubuntu to build the project. If you are going to distribute and run the resulting binary on another
operating system, you'll want to update the container to reflect that.

1. Get the Starship source code by running `go get -u github.com/starship-fyi/starship`. The project source code should be placed in your GOPATH.
2. A `Makefile` is included in the root of the repository. Browse to the starship source code directory and run `make build`.
    - When the build process is done, `tikad` and `searchd` binaries will be located in the `bin` directory.
    - `tikad` is the binary for the Document Parsing API, while `searchd` is the Search API.
3. Read the environment variables section below on how to configure both of the compiled services.

## Environment Variables

Environment variables are used to configure Starship services at runtime. The following is each environment variable in detail for each service:

## Tikad

`PORT` - Sets the port to listen on. Defaults to port `8080`.

## Searchd

`PORT` - [**Optional**] Sets the port to listen on. Defaults to port `8080`.

### Required Environment Variables

`POSTGRES_DSN` - [**Required**] Contains the postgres connection string. Read the [postgres driver documentation](https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters) on how to set this up. Example: `user=postgres password=password host=db dbname=postgres`

`MODEL_URL` - [**Required**] Contains the URL to the machine learning model endpoint. It should point directly to the machine learning model and follow tensorflow serving conventions. i.e. `http://tfserving/v1/models/{model_name}`

`MODEL_VECTOR_DIMENSIONS` - [**Required**] Defines how many vector dimensions your machine learning model returns for a sentence embedding.

`TIKAD_URL` - [**Required**] Contains the URL to a running `tikad` service

### Storage Config

`STORAGE_PATH` - [**Optional**] Is the location to store the search index. If object storage is not configured documents are stored here has well.
Ensure that the user running this app has permissions to read / write files here. Defaults to `~/.starship`.

### Object Storage

Starship can be configured to store documents on any Amazon S3 compatible object storage provider. If these are not defined, Starship defaults to saving documents to disk.

`OBJECT_STORAGE_URL` - [**Optional**] Contains the URL to your object storage providers endpoint.

`OBJECT_STORAGE_BUCKET` - [**Optional**] Is the bucket to connect to with your object storage provider.

`OBJECT_STORAGE_KEY` - [**Optional**] Is the key to use when connecting to object storage.

`OBJECT_STORAGE_SECRET` - [**Optional**] Is the secret to use when connecting to object storage.

### Security

`INDEX_KEY` - [**Optional**] Contains an API key to limit who can index documents to clients which provide this key.
Clients should pass the header X-INDEX-KEY with value of this environment variable.

`ACCESS_KEY` - [**Optional**] Contains an API key to limit who can visit all Starship endpoints besides /index and /ready to
clients which provide this key. Clients should pass header X-ACCESS-KEY with value of this environment variable.

## Preparing A Machine Learning Model To Generate Sentence Embeddings

If you read [the architecture overview](/getting-started/architecture-overview.md) you'll notice that Starship requires an API that exposes a machine learning model to generate sentence embeddings. To do this we'll setup Tensorflow Serving with the Universal Sentence Encoder in this section.

The Universal Sentence Encoder is a machine learning model which takes text and converts it to high dimensional vectors that represents the content processed.
While Tensorflow Serving is an application that takes a model and serves it over a GRPC and REST API.

For more information on the Universal Sentence Encoder and how it works, you can read the information page on [Tensorflow Hub](https://tfhub.dev/google/universal-sentence-encoder/2).

Before we can prepare our machine learning model, you will need to install Tensorflow and Python by visiting the [install guide](https://www.tensorflow.org/install).

Create a directory to hold the model information and run this python script in the same directory:

```
import tensorflow as tf
import tensorflow_hub as hub
from tensorflow.saved_model import simple_save

export_dir = "./use/00000001"
with tf.Session(graph=tf.Graph()) as sess:
    module = hub.Module("https://tfhub.dev/google/universal-sentence-encoder/2")
    input_params = module.get_input_info_dict()
    text_input = tf.placeholder(name='text', dtype=input_params['text'].dtype, shape=input_params['text'].get_shape())
    sess.run([tf.global_variables_initializer(), tf.tables_initializer()])
    embeddings = module(text_input)

    simple_save(sess,
        export_dir,
        inputs={'text': text_input},
        outputs={'embeddings': embeddings},
        legacy_init_op=tf.tables_initializer())
```

The above script downloads the Universal Sentence Encoder from Tensorflow Hub and saves it to a directory `./use/00000001`. The `00000001` in this directory represents the version of
the model and will be used by Tensorflow Serving in the next step of this guide. This script downloads a large machine learning model and will take a couple of minutes to run.

## Exposing Your Model Over An API With Tensorflow Serving

Now that we have the Universal Sentence Encoder setup, we need to expose it over an API so that the Starship search service can can communicate with it. To do this we will load
the machine learning model into [Tensorflow Serving](https://www.tensorflow.org/serving) and expose it over a REST API.

The easiest way to setup TensorFlow Serving is with the official Docker container `tensorflow/serving`. When running this container you will want to mount the `./use` directory
created above into the directory `/models/universal_encoder` of the `tensorflow/serving` Docker container.

You can run the container with the following commands:

1. Get `tensorflow/serving` with `docker pull tensorflow/serving`
2. `docker run -t --rm -p 8501:8501 -v "$PWD/use:/models/universal_encoder" -e MODEL_NAME=universal_encoder tensorflow/serving`

When Tensorflow Serving has finished loading you should be able to make a curl request to `http://localhost:8501/v1/models/universal_encoder`, and receive information regarding the
status of the model.

Test generating a sentence embedding by making a POST request to `http://localhost:8501/v1/models/universal_encoder:predict` with the following request body.

```
{
	"inputs": [
		"hello world"
	]
}
```

## Tieing It All Together With Docker Compose

We will use [Docker Compose](https://docs.docker.com/compose/) to run all of our services and network them together.

Take a look at the [running-starship example](https://google.com) in the Starship repo for how to setup Docker Compose and all required services for running Starship.

<h2 id="tldr">TLDR - Just Show Me How To Setup Starship Now</h2>

A `Makefile` is located in the [running-starship example](https://google.com) of the Starship repo. You can download the Universal Sentence Encoder model by running
`make download_model` in that directory.

When the model is finished downloading, run `make setup` to build the application binaries and Docker containers

Finally, you can run Starship with `make up`. Feel free to make modifications to the `docker-compose.yaml` file in the directory, such as adjusting environment variables.

---

**Note** - this Docker Compose file contains a section that sets up a local PostgreSQL instance to be used by Starship.

---

The Starship search API will be ready on `http://localhost:8080`

To stop all Docker containers related to Starship and delete application data, run `make down` in this directory.

## Testing Everything Is Working Correctly

[Use the Starship CLI](/getting-started/using-the-starship-cli) to test indexing and search operations on your new local instance of Starship.

