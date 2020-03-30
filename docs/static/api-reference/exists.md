## Exists Endpoint

This route is used to check if a file exists at the given path in the document index.

If you would like to restrict who can search documents, the environment variable
`ACCESS_KEY` can be set with a password before starting Starship. Clients that do
not pass this password in the `X-ACCESS-KEY` header when making requests to this
endpoint will be denied access.

---

## Request

**GET - /v1/exists?path=/foo/bar.pdf**

This endpoint responds to a GET request with the `path` of a file passed in
the path query parameter of the URL. The value of `path` should be the full
logical path to a file on the index, including it's filename.

---

## Responses

<h3 style="margin: 0;">SUCCESS</h3>
<p class="hug">
<code>200</code> - Ok
</p>

<h3 style="margin: 1rem 0 -.75rem 0;">RESPONSE BODY</h3>

```json
{
  "results": true / false
}
```

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