## Implementation Flow:
- [x] Giphy API integration
  - [x] Setup Giphy API client
  - [x] Function to fetch GIFs and parse Giphy API response
- [ ] Response Formatting
  - [x] Implement JSON formatting
  - [x] Implement XML formatting
  - [x] Research and implement Protobuf formatting
### - [ ] Microservice
- [ ] API Key Management
  - [ ] Implement SQLite db
  - [ ] Create DB model
  - [ ] API Key create, delete function
  - [ ] API Key loading on startup
- [ ] API endpoints(router)
  - [ ] Implement `/giphy` endpoint
  - [ ] Implement `/api-key` POST to create API key
  - [ ] Implement `/api-key` DELETE to delete API key
  - [ ] Write a middleware for API-key in those routes
- [ ] Containerize the application using Docker