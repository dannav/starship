## Running Starship

This guide covers setting up Starship to run in a local development environment or on a server for you and your team.
We'll cover how to setup a machine learning model, environment variables for configuring your Starship instance, and how
to run all of the required services using Docker Compose.

Before stepping through this guide please read [the architecture overview](/getting-started/architecture-overview.md) for information
on how Starship is architected.

---

Creating a fault tolerant environment for hosting Starship is out of scope of this document. If you need help setting up
such an environment, feel free to reach out to Danny Navarro at [navdgo@gmail.com](mailto:navdgo@gmail.com) for further information.

## System Requirements

System: Linux (Kernal 4.10 or later) or Mac OSX (High Sierra or later)
Memory: At least 4GB of ram
Storage: Storage required is dependent on how much data is going to be indexed by Starship and how big your machine learning model is.
It is recommended to have at least 20GB of storage on the server hosting Starship.

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

`PORT` - [Optional] Sets the port to listen on. Defaults to port `8080`.

### Required Environment Variables

`POSTGRES_DSN` - [Required] Contains the postgres connection string. Read the [postgres driver documentation](https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters) on how to set this up. Example: `user=postgres password=password host=db dbname=postgres`

`MODEL_URL` - [Required] Contains the URL to the machine learning model endpoint. It should point directly to the machine learning model and follow tensorflow serving conventions. i.e. `http://tfserving/v1/models/{model_name}`

`MODEL_VECTOR_DIMENSIONS` - [Required] Defines how many vector dimensions your machine learning model returns for a sentence embedding.

`TIKAD_URL` - [Required] Contains the URL to a running `tikad` service

### Storage Config

`STORAGE_PATH` - [Optional] Is the location to store the search index. If object storage is not configured documents are stored here has well.
Ensure that the user running this app has permissions to read / write files here. Defaults to `~/.starship`.

### Object Storage

Starship can be configured to store documents on any Amazon S3 compatible object storage provider. If these are not defined, Starship defaults to saving documents to disk.

`OBJECT_STORAGE_URL` - [Optional] Contains the URL to your object storage providers endpoint.

`OBJECT_STORAGE_BUCKET` - [Optional] Is the bucket to connect to with your object storage provider.

`OBJECT_STORAGE_KEY` - [Optional] Is the key to use when connecting to object storage.

`OBJECT_STORAGE_SECRET` - [Optional] Is the secret to use when connecting to object storage.

### Security

`INDEX_KEY` - [Optional] Contains an API key to limit who can index documents to clients which provide this key.
Clients should pass the header X-INDEX-KEY with value of this environment variable.

`ACCESS_KEY` - [Optional] Contains an API key to limit who can visit all Starship endpoints besides /index and /ready to
clients which provide this key. Clients should pass header X-ACCESS-KEY with value of this environment variable.

## Preparing A Machine Learning Model To Generate Sentence Embeddings

If you read [the architecture overview](/getting-started/architecture-overview.md) you'll notice that Starship requires an API that exposes a machine learning model to generate sentence embeddings. To do this we'll setup Tensorflow Serving with the Universal Sentence Encoder in this section.

Talk about getting the universal sentence encoder and what Tensorflow Serving is.

### Exposing Your Model Over An API With Tensorflow Serving

Now that we have the Universal Sentence Encoder setup, we need to expose it over an API so that the Starship search service can can communicate with it.

## Tieing It All Together With Docker Compose

Talk about setting up Docker.

Talk about installing Docker Compose.

### Creating A Docker Compose File

### Networking For The Containers

## Testing Everything Is Working Correctly

Install the Starship CLI and test an index and search operation

## Example Files
