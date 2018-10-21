---
- simulator cli app to seed database with github readme info
    - use existing one in ai and move to this project
- chunk readmes on paragraph
- send to rest server to be indexed
---


# Old

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