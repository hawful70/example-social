# Social App Example

This is a social platform project built as an example application using Go, Redis, and PostgreSQL. The project demonstrates how to implement key social media features such as user authentication, posting content, and interacting with others through likes and comments.

## Tech Stack

- **Go**: Backend development, handling the core logic, APIs, and business rules.
- **Redis**: Used for caching frequently accessed data like user sessions, and to improve performance for real-time features.
- **PostgreSQL**: Primary database for storing structured data such as users, posts, and interactions.

## Features

- **User Authentication**: Sign up, login, and session management.
- **Post Creation**: Users can create and share posts with others.
- **Likes & Comments**: Interaction features to engage with the content.
- **Feed Generation**: A simple timeline for users to view posts from others.

## Getting Started

### Prerequisites

- Go 1.20+
- Redis
- PostgreSQL 13+

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/hawful70/example-social
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Set up environment variables for Redis and PostgreSQL connections.

4. Run the application:
   ```bash
   go run cmd/api/main.go
   ```
