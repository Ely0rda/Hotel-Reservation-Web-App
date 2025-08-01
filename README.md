# Hotel Reservation Web App

![Go](https://img.shields.io/badge/Go-1.19+-blue.svg)
![Bootstrap](https://img.shields.io/badge/Bootstrap-4.3.1-purple.svg)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-12+-green.svg)

A full-featured hotel reservation system built with Go, allowing guests to browse rooms, check availability, and make reservations with an elegant user interface.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Technology Stack](#technology-stack)
- [Architecture](#architecture)
- [Screenshots](#screenshots)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Database](#database)
- [Testing](#testing)
- [Skills Demonstrated](#skills-demonstrated)
- [License](#license)

## Overview

This Hotel Reservation Web App is a complete booking system for a bed and breakfast establishment. It allows customers to view room details, check availability, and make reservations. The application follows clean architecture principles with separation of concerns and provides a responsive user interface built with Bootstrap.

## Features

- **Room Browsing:** View detailed information about different room types (General's Quarters, Major's Suite)
- **Availability Search:** Check room availability by date using a date range picker
- **Reservation System:** Make and manage room reservations with form validation
- **User Authentication:** User registration and login functionality (in progress)
- **Admin Dashboard:** For managing rooms, reservations, and restrictions (in progress)
- **Responsive Design:** Works on mobile, tablet, and desktop browsers

## Technology Stack

- **Backend:** Go (Golang 1.19+)
- **Frontend:** HTML, CSS, JavaScript, Bootstrap 4
- **Database:** PostgreSQL
- **Session Management:** SCS session manager
- **Form Validation:** Custom form validation with Go
- **Security:** CSRF protection with NoSurf
- **Testing:** Go's built-in testing package

## Architecture

The application follows a clean architecture approach with clear separation of concerns:

- **Handlers:** Process HTTP requests and responses
- **Models:** Define data structures
- **Render:** Manage template rendering
- **Forms:** Handle form validation
- **Config:** Store application configuration
- **Helpers:** Provide utility functions
- **Migrations:** Manage database schema changes

The app uses the repository pattern to abstract database operations and middleware for cross-cutting concerns.

## Screenshots

(Screenshots would be placed here)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/Hotel-Reservation-Web-App.git
cd Hotel-Reservation-Web-App
```

2. Install dependencies:

```bash
go mod download
```

3. Set up the database:

```bash
# Create database.yml from the example file
cp database.yml.example database.yml
# Edit database.yml with your PostgreSQL credentials
```

4. Run database migrations:

```bash
# Install soda CLI (Buffalo's migration tool)
go install github.com/gobuffalo/pop/v6/soda@latest

# Run migrations
soda migrate up
```

5. Run the application:

```bash
./run.sh
# Or directly with Go
go run cmd/main/*.go
```

## Usage

Once the application is running, access it at http://localhost:8080

- Browse room types from the navigation menu
- Check room availability by clicking on "Book Now" or "Check Availability" buttons
- Make a reservation by filling out the reservation form

## Project Structure

```
├── cmd/            # Application entrypoint
│   └── main/       # Main application package
├── internal/       # Private application code
│   ├── config/     # Application configuration
│   ├── forms/      # Form validation
│   ├── handlers/   # HTTP request handlers
│   ├── helpers/    # Helper functions
│   ├── models/     # Data models
│   └── render/     # Template rendering
├── migrations/     # Database migrations
├── static/         # Static files (CSS, JS, images)
├── templates/      # HTML templates
├── html-source/    # Original HTML source files
└── bookings/       # Compiled application
```

## Database

The application uses PostgreSQL with the following main tables:

- `users`: Store user information
- `rooms`: Available room types
- `reservations`: Customer booking information
- `restrictions`: Booking restrictions (e.g., already booked, owner block)
- `room_restrictions`: Join table for room availability

## Testing

The application includes comprehensive tests for handlers, forms, and rendering:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests for a specific package
go test -v ./internal/handlers
```

## Skills Demonstrated

- **Go Programming:** Server-side development with Go, including routing, middleware, and handlers
- **Web Development:** HTML, CSS, JavaScript, Bootstrap for responsive design
- **Database Management:** PostgreSQL schema design and migrations
- **Form Handling:** Client and server-side validation
- **Session Management:** User sessions and authentication
- **Software Architecture:** Clean architecture with separation of concerns
- **Testing:** Unit testing with Go's testing package
- **Security:** CSRF protection, form validation, secure sessions

## License

This project is open source and available under the [MIT License](LICENSE).
