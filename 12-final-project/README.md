# Final Project

## Requirements

- Go 1.22.8
- PostgreSQL 15.4

## Setup

1. Copy `.env.example` to `.env` and fill the values

    ```bash
    cp .env.example .env
    ```

2. Run app 🚀

    ```bash
    make run
    ```

## API Documentation

### Auth

#### Register

```bash
curl -X POST \
  http://localhost:3000/register \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "reza",
    "email": "reza@example.com",
    "password": "123456",
    "age": 18
  }'
```

#### Login

```bash
curl -X POST \
  http://localhost:3000/login \
  -H 'Content-Type: application/json' \
  -d '{
    "email": "reza@example.com",
    "password": "123456"
  }'
```

### Photos

#### Get All

```bash
curl -X GET \
  http://localhost:3000/photos \
  -H 'Authorization: Bearer <token>'
```

#### Create

```bash
curl -X POST \
  http://localhost:3000/photos \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer <token>' \
  -d '{
    "title": "My Awesome Photo",
    "caption": "This is my awesome photo",
    "photo_url": "https://images.unsplash.com/photo-1683641931431-e1d8e0f4b0a4?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1050&q=80"
  }'
```

#### Get One

```bash
curl -X GET \
  http://localhost:3000/photos/1 \
  -H 'Authorization: Bearer <token>'
```

#### Delete

```bash
curl -X DELETE \
  http://localhost:3000/photos/1 \
  -H 'Authorization: Bearer <token>'
```

#### Update

```bash
curl -X PUT \
  http://localhost:3000/photos/1 \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer <token>' \
  -d '{
    "title": "My Awesome Photo",
    "caption": "This is my awesome photo",
    "photo_url": "https://images.unsplash.com/photo-1683641931431-e1d8e0f4b0a4?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1050&q=80"
  }'
```

### Comments

#### Get All

```bash
curl -X GET \
  http://localhost:3000/comments \
  -H 'Authorization: Bearer <token>'
```

#### Create

```bash
curl -X POST \
  http://localhost:3000/comments \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer <token>' \
  -d '{
    "message": "This is my awesome comment",
    "photo_id": 1
  }'
```

#### Get One

```bash
curl -X GET \
  http://localhost:3000/comments/1 \
  -H 'Authorization: Bearer <token>'
```

#### Delete

```bash
curl -X DELETE \
  http://localhost:3000/comments/1 \
  -H 'Authorization: Bearer <token>'
```

#### Update

```bash
curl -X PUT \
  http://localhost:3000/comments/1 \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer <token>' \
  -d '{
    "message": "This is my awesome comment",
    "photo_id": 1
  }'
```

### Social Media

#### Get All

```bash
curl -X GET \
  http://localhost:3000/social-media \
  -H 'Authorization: Bearer <token>'
```

#### Create

```bash
curl -X POST \
  http://localhost:3000/social-media \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer <token>' \
  -d '{
    "name": "My Awesome Social Media",
    "social_media_url": "https://www.example.com/reza"
  }'
```

#### Get One

```bash
curl -X GET \
  http://localhost:3000/social-media/1 \
  -H 'Authorization: Bearer <token>'
```

#### Delete

```bash
curl -X DELETE \
  http://localhost:3000/social-media/1 \
  -H 'Authorization: Bearer <token>'
```

#### Update

```bash
curl -X PUT \
  http://localhost:3000/social-media/1 \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer <token>' \
  -d '{
    "name": "My Awesome Social Media",
    "social_media_url": "https://www.example.com/reza"
  }'
```
