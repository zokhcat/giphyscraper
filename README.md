## Implementation Flow:
- [x] Giphy API integration
  - [x] Setup Giphy API client
  - [x] Function to fetch GIFs and parse Giphy API response
- [x] Response Formatting
  - [x] Implement JSON formatting
  - [x] Implement XML formatting
  - [x] Research and implement Protobuf formatting
### - [x] Microservice
- [x] API Key Management
  - [x] Implement SQLite db
  - [x] Create DB model
  - [x] API Key create, delete function
- [x] API endpoints(router)
  - [x] Implement `/giphy` endpoint
  - [x] Implement `/api-key` POST to create API key
  - [x] Implement `/api-key` DELETE to delete API key
  - [x] Write a middleware for API-key in those routes
- [x] Containerize the application using Docker