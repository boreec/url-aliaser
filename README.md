# Usage

## payload

The payload you need to send needs to have the following structure:

```json
{
    "url" : "https://example.com",
    "length" : 10
}
```

- `url` represents the link to shorten (required)
- `length` represents the maximum length of the shortened link (optional)

The payload you receive has the following structure:

```json
{
    "url" : "http://localhost:8080/xxxx"
}
```

- `url` represents the shortened link

## request with Curl

In the following command, replace `https://example.com` with the url you want to shorten.

```bash
curl -X POST -H "Content-Type: application/json" -d '{"Url": "https://example.com"}' http://localhost:8080/shorten
```