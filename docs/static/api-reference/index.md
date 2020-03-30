## Index Endpoint

This route is used to add a new document to your Starship instance. It handles
indexing the document so that it can be searched in the future. If an object
storage provider was configured to save documents to it will store the document
there. Otherwise, the document will be stored on the local machine so that it
can be downloaded from a client in the future.

If you would like to restrict who can index documents, the environment variable
`INDEX_KEY` can be set with a password before starting Starship. Clients that do
not pass this password in the `X-INDEX-KEY` header when making requests to this
endpoint will be denied access.

---

## Request

**POST - /v1/index**

A header `X-PATH` can be set in the request that tells Starship under what logical directory to store this
file. If `X-PATH` is not set, the file path will be set to the root of the index, i.e. `/`.

While indexing files, it's helpful to take on the mindset that you are placing files in a
directory structure on Starship. Other examples of valid paths are: `/company`, `/projects/my-project`,
`/devops/cookbooks`. The document to be indexed will be placed in the path provided, ending with it's
filename. For instance a file `employee-handbook.pdf` placed in `/company` will exist at `/company/employee-handbook.pdf`.

This endpoint takes a POST request with the body passed as `form-data`.

| Key      | Value       | Description                                                        |
|----------|-------------|--------------------------------------------------------------------|
| content  | [form-file] | Content contains the file that you would like to upload and index. |
| filename | foo.pdf     | Filename is the name of the file that you are uploading.           |

---

## Responses

<h3 style="margin: 0;">SUCCESS</h3>
<p class="hug">
<code>204</code> - No content
</p>
<br />

<h3 style="margin: 0;">FAILURE</h3>
<p class="hug">
<code>415</code> - Unsupported Media Type
</p>

<br />

<h3 style="margin: 0;">FAILURE</h3>
<p class="hug">
<code>500</code> - Internal Server Error
</p>

```json
{
  "errors": [
    { "message": "internal server error" }
  ]
}
```

<h3 style="margin: 0;">FAILURE</h3>
<p class="hug">
<code>503</code> - Service Unavailable
</p>

```json
{
  "errors": [
    { "message": "information on which service the search api could not connect to" }
  ]
}
```
<br />