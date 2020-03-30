## Search Endpoint

This route is used to search the Starship index for documents that have content
relevant to your search query.

If you would like to restrict who can search documents, the environment variable
`ACCESS_KEY` can be set with a password before starting Starship. Clients that do
not pass this password in the `X-ACCESS-KEY` header when making requests to this
endpoint will be denied access.

---

## Request

**GET - /v1/search**

This endpoint takes a GET request with the body passed as `json`.

```json
{
  "text": "do vacation and sick days carry over?"
}
```

---

## Responses

<h3 style="margin: 0;">SUCCESS</h3>
<p class="hug">
<code>200</code> - Ok
</p>

<h3 style="margin: 1rem 0 -.75rem 0;">RESPONSE BODY</h3>

`Results` - The root property of the `json` response.

`Distances` - An array of distances. Each element is the distance of the document at this index from your search query

`Documents` - An array of documents matching your search query.
  - - `id` - A uuid representing the id of the document.
  - - `sentenceID` - A uuid representing the id of the sentence that a match was found at.
  - - `annoyID` - The id in the index that the sentence embedding can be found at.
  - - `name` - The filename of the document.
  - - `path` - The full path to the document in the index.
  - - `downloadURL` - The full path to the document on object storage or on disk.
  - - `text` - The text that matches your search query.
  - - `rel` - The sentence score (relevance)

<br />

```json
{
    "results": {
        "distances": [
            0.7429964,
        ],
        "documents": [
            {
                "id": "xxx",
                "sentenceID": "xxx",
                "annoyID": 1,
                "name": "employee-handbook.pdf",
                "path": "_rootfolder_.employee-handbook.pdf",
                "downloadURL": "/2a8392ce-d20e-4194-bc6d-a1a2abbe750a/employee-handbook.pdf",
                "text": "Requests will be reviewed based on a number of factors, including, without limitations,  business needs and staffing requirements. If an employee’s earned vacation is not enough to cover the number  of days that they expect to be away from work, employees need to discuss and obtain approval for unpaid leave  of absence from the supervisor. There will be no carry-over of vacation days from one calendar year to the next.",
                "rel": 0.53906333
            }
        ]
    }
}
```
<br />

<h3 style="margin: 0;">FAILURE</h3>
<p class="hug">
<code>400</code> - Bad Request
</p>

```json
{
  "errors": [
    { "message": "validation error messages" }
  ]
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