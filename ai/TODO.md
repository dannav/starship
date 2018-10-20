Once we have the proper data we can pass it over to the universal-sentence-encoder to get a 512 dimensional vector.

512 is much to big for our use cases (since we want to offer a free tier) and it probably doesn't contain that much
variance to be significant once we get past ~100 components. Postgres will also only let us index up to 100 dimensions
of a vector with the `cube` extension. So we are limited by 100 if we want to do KNN through postgres and not GPU
accelerated in ram with something like `annoy`.

I would like to test performance and results of doing KNN with vectors in dimensions 7, 10, 15, and 30. We can get a
good idea of the number of dimensions that will give __good enough__ values by checking the variance of principal
components when applying PCA.

1. Take our 512d vector and apply the following dimension reducing algorithm as described in [this paper](https://arxiv.org/pdf/1708.03629v3.pdf).
    1. This paper describes an algorithm using PCA to reduce vector dimensions (pretty new)
2. Apply KNN through postgres as defined [here](https://dba.stackexchange.com/questions/163207/quick-nearest-neighbor-search-in-the-150-dimensional-space)

We can then perform tests to see how much ram something like this would take for each vector size and the quality
of search results.

This will tell us if this is even feasable a consumer first / b2b second, venture (think slack business model)

### References

* https://en.wikipedia.org/wiki/Principal_component_analysis#First_component
* https://github.com/gonum/gonum/blob/master/stat/pca_example_test.go
* https://arxiv.org/pdf/1708.03629v3.pdf
* https://dba.stackexchange.com/questions/163207/quick-nearest-neighbor-search-in-the-150-dimensional-space