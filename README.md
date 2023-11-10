# Weather-App Readme

## Overview

Welcome to the Weather-App! This application utilizes the Gin framework to provide weather information based on location. Users can retrieve weather data by sending an HTTP request to the specified endpoint.

## Setup

1. **Install Dependencies:** Make sure you have Go and the Gin framework installed on your system. You can install Gin using the following command:

    ```bash
    go get -u github.com/gin-gonic/gin
    ```

2. **Clone the Repository:**

    ```bash
    git clone [repository_url]
    ```

3. **Run the Application:**

    Navigate to the project directory and run the following command to start the application:

    ```bash
    go run main.go
    ```

    The application will start on `http://localhost:8080`.

## Usage

To retrieve weather information, use the following `curl` command:

```bash
curl "http://localhost:8080/weather?location=Purwokerto"
```


