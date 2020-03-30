# AI powered semantic search

This project provides steps to build and serve the universal-sentence-encoder tensorflow model provided by google.
We use it to generate word embeddings on documents. We'll use these word embeddings to do semantic search on documents
by shoving all word embeddings into a database. We can then use the same encoder to generate embeddings for a query
against documents in the database. The KNN documents to that query are semantically similar documents. i.e. search
results.

Install Brew:
`make install-brew`

Install dependencies and build tensorflow model:
`make`