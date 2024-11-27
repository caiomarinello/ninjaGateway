# Ninja API Gateway

The API Gateway provides a unified entry point to the Ninja Tech marketplace challenge system. It manages:

- Routing requests to appropriate backend services.
- Authentication and session management.

## Authentication & Authorization Endpoints

| **Method** | **Endpoint** | **Description**                   | **Access** | **Status**         |
| ---------- | ------------ | --------------------------------- | ---------- | ------------------ |
| POST       | `/register`  | Register a new user.              | User       | 🔴 Not Implemented |
| POST       | `/login`     | Authenticate a user in a session. | User       | 🔴 Not Implemented |
| POST       | `/logout`    | Log out the authenticated user.   | User       | 🔴 Not Implemented |

## Related Repositories

- **Backend API**: [Ninja Backend API](https://github.com/caiomarinello/ninja)
