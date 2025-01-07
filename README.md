# Oordy App - Backend Version 2 (Rideshare)

## Overview

Oordy is a rideshare platform built to connect riders with drivers. This repository contains the backend code for the Oordy app, developed using GoLang. Version 2 of the backend introduces several improvements and new features for scalability, performance, and security.

## Features

- **User Authentication**: Secure login and registration for both riders and drivers.
- **Ride Matching**: Algorithm to match riders with available drivers based on location.
- **Trip Management**: Track the status of trips, including start, in-progress, and completed.
- **Payment Integration**: Handle ride payments securely.
- **Admin Dashboard**: For managing users, rides, and payments.
- **Real-Time Notifications**: Via WebSocket to keep users updated on their ride status.

## Technologies Used

- **GoLang**: Main programming language for backend logic.
- **Fiber**: Web framework for building APIs.
- **GORM**: ORM for database interactions.
- **PostgreSQL**: Database to store user data, rides, and transactions.
- **Redis**: Used for caching and session management.
- **Docker**: Containerization for development and deployment.
- **JWT**: JSON Web Token for user authentication.
- **Stripe API**: For payment processing.
- **WebSocket**: For real-time notifications.

## Requirements

- Go 1.18 or higher
- PostgreSQL
- Redis
- Docker (optional)

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/oordy-backend-v2.git
   cd oordy-backend-v2
