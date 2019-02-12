Starship is constructed as a collection of three main services and a database:

1. The Search API

   The main service for Starship. It handles indexing documents and provides the functionality
   to search against a collection of them.

2. Machine Learning Model API

   A service that returns the reponse of a machine learning model that does
   sentence embeddings. **This service is not bundled with Starship.**
   The API endpoints should be constructed in a way that matches a Tensorflow Serving API.
   If you want to get up and running quickly without creating your own API,
   visit the [Running Starship Section](/getting-started/running-starship) for information on how to use [Tensorflow Serving](https://www.tensorflow.org/serving/)
   with the [Universal Sentence Encoder](https://tfhub.dev/google/universal-sentence-encoder/2).

3. Document Parsing API

   A wrapper around an [Apache Tika](https://tika.apache.org/) instance.

4. A PostgreSQL Database Instance

   Starship uses [PostgreSQL](https://www.postgresql.org/) to store document content and index information.

## Architecture Diagram

Below is a graphic of how one could visualize the overall architecture of Starship.

![starship architecture diagram](/static/images/architecture.png "Starship Architecture Diagram")
