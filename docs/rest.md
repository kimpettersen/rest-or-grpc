
`POST:/library/books`

Request:
```
{
    "title": "Tuva the cat, a war story",
    "author": "Kim Pettersen",
    "description": "A thrilling story about an everlasting war agains a red dot"
}
```

Response:
```
{
    "data": {
        id": "87236"
    },
    "links":[{
            "href": "/library/books/87236",
            "type": "GET"
        },{
            "href": "/library/books/87236",
            "type": "PUT"
        },{
            "href": "/library/books/87236",
            "type": "DELETE"
        }]
}
```

`GET:/library/books/87236?t=html`
```
<!DOCTYPE html>
<html>
    <body>
        <h1>Tuva the cat, a war story</h1>
        <h2>Kim Pettersen</h2>
        <p>A thrilling story about an everlasting war agains a red dot on the wall</p>

        <a href="/library/books">All books</button>
</body>
</html>
```
