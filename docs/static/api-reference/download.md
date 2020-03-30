## Download Endpoint

This route is used to download documents that have been stored on Starship. If
Starship has been configured to use object storage to store files, the document
will be downloaded from there. Otherwise, local storage is used and the document
is served from disk.

If you would like to restrict who can search documents, the environment variable
`ACCESS_KEY` can be set with a password before starting Starship. Clients that do
not pass this password in the `X-ACCESS-KEY` header when making requests to this
endpoint will be denied access.

---

## Request

**GET - /v1/download?file=[downloadURL]**

This endpoint responds to a GET request with the `downloadURL` of a file passed in
the file query parameter of the URL. The `downloadURL` of a file can be retrieved in
the results of a document in the response returned from the search endpoint.

---

## Responses

<h3 style="margin: 0;">SUCCESS</h3>
<p class="hug">
<code>200</code> - Ok
</p>

<h3 style="margin: 1rem 0 -.75rem 0;">RESPONSE BODY</h3>

The contents of the file is written to the http output stream. Headers for
`Content-Disposition` `Content-Type` `Content-Length` are set in the response.

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