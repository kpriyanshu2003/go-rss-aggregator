# RSS Aggregator API Documentation

This documentation describes the available API endpoints as defined in the `postman/postman_collection.json` file for the Go RSS Aggregator project.  
Base URL: `http://localhost:3300/v1/`

---

## Users

### Create User

- **Endpoint:** `POST /users`
- **Body (JSON):**
  ```json
  {
    "name": "Inner Locus"
  }
  ```
- **Description:** Creates a new user.

### Get All Users

- **Endpoint:** `GET /users`
- **Description:** Returns a list of users.

---

## Feed

### Get All Feeds

- **Endpoint:** `GET /feed`
- **Description:** Returns all RSS feeds.

### Create Feed

- **Endpoint:** `POST /feed`
- **Headers:**
  - `Authorization: ApiKey <your_api_key>`
- **Body (JSON):**
  ```json
  {
    "name": "Inner Locus",
    "url": "https://innerlocus.com/feed"
  }
  ```
- **Description:** Creates a new RSS feed.

---

## Feed Follows

### Follow a Feed

- **Endpoint:** `POST /feed-follow`
- **Body (JSON):**
  ```json
  {
    "feed_id": "a55ca052-90c6-4cbc-ac42-7fc1cb996795"
  }
  ```
- **Description:** Follows a feed for the user.

### Get Followed Feeds

- **Endpoint:** `GET /feed-follow`
- **Description:** Returns the list of feeds the user is following.

### Unfollow a Feed

- **Endpoint:** `DELETE /feed-follow/:id`
- **Headers:**
  - `Authorization: ApiKey <your_api_key>`
- **Path Parameter:**
  - `id` - the follow relationship ID
- **Description:** Unfollows a feed.

---

## Posts

### Get All Posts by User

- **Endpoint:** `GET /posts`
- **Description:** Returns all posts for the authenticated user.

---

## Authentication

- Uses API Key authentication via the header:
  ```
  Authorization: ApiKey <your_api_key>
  ```
- Sample API Key (from collection):  
  `ApiKey 475ce2c168d07466bf4d362c2f3070b7043f2a5ea6750145cc712236ad811428`

---

## Notes

- Replace sample API keys and IDs with actual values from your application.
- All endpoints are prefixed with `/v1/`.
- All requests and responses use JSON.

---
