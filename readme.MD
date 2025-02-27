# Shorten URL API

## Overview
This API provides a URL shortening service. You can send a POST request with a long URL, and the API will return a shortened version of the URL.

## Endpoint
**Base URL:** `https://shorten-url-production.up.railway.app`

### **Shorten a URL**
- **Endpoint:** `/api/v1`
- **Method:** `POST`
- **Content-Type:** `application/json`

## Request Body
Send a JSON object containing the URL you want to shorten.

```json
{
  "url": "https://sjpl.co.id"
}
```

## Response
If successful, the API will return a shortened URL.

### **Success Response (200 OK)**
```json
{
  "short_url": "https://shorten-url-production.up.railway.app/abc123"
}
```

## Usage
You can test the API using `curl` or any API testing tool like Postman.

### **Using `curl` in Terminal**
```sh
curl -X POST "https://shorten-url-production.up.railway.app/api/v1" \
     -H "Content-Type: application/json" \
     -d '{"url": "https://sjpl.co.id"}'
```

### **Using an HTTP File**
If you're using VS Code with the REST Client extension, save the following as `shorten.http` and run it:

```http
POST https://shorten-url-production.up.railway.app/api/v1
Content-Type: application/json

{
  "url": "https://sjpl.co.id" | "sjpl.co.id"
}
```

## Notes
- Ensure the provided URL is valid and properly formatted.
- The shortened URL will be returned in the response if successful.
- The API have rate limits.

