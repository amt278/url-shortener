# URL Shortener API

A simple URL shortener API built with **Go** and **Firebase Firestore** as the database backend. This API allows users to shorten long URLs and retrieve the original URL by accessing the shortened URL.

## Table of Contents

1. [Introduction](#introduction)
2. [Getting Started](#getting-started)
3. [Folder Structure](#folder-structure)
4. [Environment Variables](#environment-variables)
5. [API Endpoints](#api-endpoints)
6. [Run the Application](#run-the-application)
7. [Dependencies](#dependencies)
8. [License](#license)

## Introduction

This project is a URL shortener API built in Go. It uses **Firebase Firestore** to store shortened URLs and their original counterparts. Users can shorten long URLs and later retrieve them by visiting the shortened URL.

## Getting Started

To run this project locally, you will need **Go** installed on your system and a **Firebase** account to generate a service account key for accessing Firestore.

### Prerequisites

- **Go**: Download and install Go from the official site [Go Downloads](https://go.dev/dl/).
- **Firebase**: Set up a Firebase project and create a service account key for Firestore. You can download the key JSON file from [Firebase Console](https://console.firebase.google.com/).

### Install Dependencies

1. Clone the repository:
    ```bash
    git clone https://github.com/amt27/url-shortener.git
    cd url-shortener
    ```

2. Initialize Go modules and install dependencies:
    ```bash
    go mod init url_shortener
    go get firebase.google.com/go
    ```

3. Set the environment variable `GOOGLE_APPLICATION_CREDENTIALS` to point to your Firebase service account JSON file:
    ```bash
    export GOOGLE_APPLICATION_CREDENTIALS="path/to/your/firebase-adminsdk.json"
    ```

## Folder Structure

The project structure is organized as follows:


- **`cmd/shorturl/main.go`**: The entry point for the application, where the server is started and routes are defined.
- **`db/db.go`**: Contains the logic for initializing the Firebase database connection using the Firestore client.
- **`handlers/urlhandlers.go`**: Contains the HTTP handler functions for creating shortened URLs and redirecting to the original URL.
- **`utils/urlshortener.go`**: Contains utility functions, such as generating random shortened URLs.

## Environment Variables

Before running the project, ensure you have set the following environment variables:

- **`GOOGLE_APPLICATION_CREDENTIALS`**: Path to your Firebase service account JSON file. For example:
    ```bash
    export GOOGLE_APPLICATION_CREDENTIALS="path/to/firebase-adminsdk.json"
    ```

## API Endpoints

### POST `/shorten`
- **Description**: Shortens a given long URL.
- **Request Body**:
    ```json
    {
        "url": "https://www.example.com"
    }
    ```
- **Response**:
    ```json
    {
        "shortened URL": "shortened-url"
    }
    ```

### GET `/original/{shortened}`
- **Description**: Redirects to the original URL associated with the shortened URL.
- **Response**: A 302 redirect to the original URL.

## Run the Application

1. Start the Go server:
    ```bash
    go run cmd/shorturl/main.go
    ```

2. The application will be accessible at `http://localhost:8080`.

## Dependencies

- **Firebase Admin SDK**: Used to interact with Firebase Firestore.
- **Go**: Programming language used for building the server.
- **Go modules**: To manage dependencies.

### Install Go dependencies:
```bash
go get firebase.google.com/go
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
