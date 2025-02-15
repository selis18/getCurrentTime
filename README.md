# GetCurrentTime API

## Overview
The GetCurrentTime API is a simple HTTP server that provides the current time based on the specified timezone. It returns the time in JSON format, allowing users to easily access and display the data in their applications.

## Features
- Retrieve the current time in various timezones.
- Response in JSON format with easy-to-read structure.

## Usage
Start the server:

```bash
go run main.go
```

Access the time endpoint:

```http
GET http://localhost:8080/time?timezone=Asia/Kolkata
```

## Example Response
```json
{
  "time": "14-08-2023 15:30:45",
  "timezone": "Asia/Yekaterinburg"
}
```
