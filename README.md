# Giphy API Integration Service  

## Features  
- **Giphy API Integration**: Fetch GIFs based on search terms.  
- **Response Formatting**: Support for JSON, XML, and Protobuf formats.
- **API Key Management**: Create and delete API keys using SQLite.  
- **Endpoints**:  
  - `/giphy`: Fetch GIFs using the Giphy API.  
  - `/api-key` (POST/DELETE): Manage API keys.  
- **Authentication**: API key-based authentication using middleware.
- **Caching**: Caching using `redis` for a faster response time when querying the same thing.
<img src="https://i.postimg.cc/K4XrK0Zq/241019-22h43m51s-screenshot.png" alt="caching response time">
_first request is before getting cached and second request is after getting cached and the significant improvement in response time_
<br>
- **Containerization**: Docker support with Docker Compose for easy setup.

---

## Endpoints  

1. **`GET /giphy`**  
   - **Query Params**:  
     - `q`: Search term for GIFs.  
     - `o`: (Optional) Response format - `json`, `xml`, or `protobuf`.  
   - **Response**: GIFs matching the search query in the requested format.  

2. **`POST /api-key`**  
   - **Description**: Create a new API key.  
   - **Response**: Returns the newly created API key.

3. **`DELETE /api-key`**  
   - **Description**: Delete an existing API key.  
   - **Parameters**:  
     ```
     /api-key/:id
     ```
   - **Response**: Confirmation of deletion.

---

## Setup  

1. **Clone the Repository**  
   ```bash
   git clone git@github.com:zokhcat/giphyscraper.git
   cd giphyscraper

2. **Configure Giphy API Key**

    Add your Giphy API key to .env file:

```bash
    GIPHY_API_KEY=<your-api-key>
```
    Run the Service Locally

```bash
    docker compose up
```
### Access the API
  Service will be available at http://localhost:8080.

### Middleware

The API uses a middleware to validate API keys for `/giphy` route

### Tech Stack

  - Golang
  - SQLite for API key storage
  - Redis for caching
  - Docker & Docker Compose for containerization


## Implementation Flow:
- [x] Giphy API integration
  - [x] Setup Giphy API client
  - [x] Function to fetch GIFs and parse Giphy API response
- [x] Response Formatting
  - [x] Implement JSON formatting
  - [x] Implement XML formatting
  - [x] Research and implement Protobuf formatting
- [x] API Key Management
  - [x] Implement SQLite db
  - [x] Create DB model
  - [x] API Key create, delete function
- [x] API endpoints(router)
  - [x] Implement `/giphy` endpoint
  - [x] Implement `/api-key` POST to create API key
  - [x] Implement `/api-key` DELETE to delete API key
  - [x] Write a middleware for API-key in `/giphy` route
- [x] Containerize the application using Docker
- [x] Docker Compose

## Extra Features I want to implement
- [x] Implement caching using redis