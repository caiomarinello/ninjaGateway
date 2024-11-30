# Ninja API Gateway

The API Gateway provides a unified entry point to the Ninja Tech marketplace challenge system. It manages:

- Routing requests to appropriate backend services.
- Authentication and session management.

## Table of Contents

1. [Project Setup (Docker)](#project-setup-docker)
2. [Alternative Project Setup (Without Docker)](#alternative-project-setup-without-docker)
3. [Access Endpoints](#access-endpoints)
4. [Related Repositories](#related-repositories)

---

## Project Setup (docker)

This project setup requires you to have docker installed. If you prefer to run `ninja` and `ninjaGateway` services manually, refer to the Alternative Project Setup section.

### 1. Clone the Required Repositories
Clone both the `ninja` and `ninjaGateway` repositories into a directory:

```bash
git clone https://github.com/caiomarinello/ninja.git
git clone https://github.com/caiomarinello/ninjaGateway.git
```

### 2. Add Setup Files
Add the following setup files to the root directory:

- [docker-compose.yml](https://github.com/caiomarinello/ninjaSetupFiles/blob/main/docker-compose.yml)
- [init.sql](https://github.com/caiomarinello/ninjaSetupFiles/blob/main/init.sql)
- `.env`

Ensure that these files are placed in the root directory of your project, alongside the cloned `ninja` and `ninjaGateway` repositories. 

The `.env` should contain the following variables:
```env
MYSQL_ROOT_PASSWORD="root_password_here"
MYSQL_DATABASE="database_name"
MYSQL_USER="database_username"
MYSQL_PASSWORD="database_password"
```

### 3. Setup .env files in ninja and ninjaGateway services.
A `.env.example` file is provided in both projects. Ensure that the Database configuration variables match with the`.env` variables in the previous step.


### 4. Build and Start the Services
In the root directory where all repositories are cloned, run the following command to build and start the services:

```bash
docker-compose up --build
```

This will set up the environment, build the Docker containers, and start the services as defined in the `docker-compose.yml` file.

---

## Alternative Project Setup (Without Docker)

If you prefer to run the project without Docker, you can follow the steps below to set up and run both the `ninja` and `ninjaGateway` services manually. This setup requires you to have **Go** and **MySQL** installed on your local machine.

### 1. Clone the Required Repositories
Clone both the `ninja` and `ninjaGateway` repositories to your local machine:

```bash
git clone https://github.com/caiomarinello/ninja.git
git clone https://github.com/caiomarinello/ninjaGateway.git
```
### 2. Setup .env files
Refer to steps 2 and 3 of Project setup with docker.

### 3. Database setup
Run `init.sql` [file](https://github.com/caiomarinello/ninjaSetupFiles/blob/main/init.sql) to Set Up the Database.  You can execute this file using the MySQL command line or a database management tool like MySQL Workbench.

Example command to run init.sql via MySQL command line:

```bash
mysql -u user -p < path/to/init.sql
```
### 4. Run both services
Navigate to the `ninja` directory and run it using:
```bash
go run main.go
```
Navigate to the `ninjaGateway` directory and run it using:
```bash
go run main.go
```

---

## Access Endpoints

| **Method** | **Endpoint** | **Description**                   | **Access** | **Status**     |
| ---------- | ------------ | --------------------------------- | ---------- | -------------- |
| POST       | `/register`  | Register a new user.              | User       | 🟢 Implemented |
| POST       | `/login`     | Authenticate a user in a session. | User       | 🟢 Implemented |
| POST       | `/logout`    | Log out the authenticated user.   | User       | 🟢 Implemented |

More endpoints can be accessed via the ninja service, refer to the following README:

- [Ninja README](https://github.com/caiomarinello/ninja/blob/main/README.md)

## Related Repositories

- **Backend API**: [Ninja Backend API](https://github.com/caiomarinello/ninja)
