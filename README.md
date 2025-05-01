# Healthcare Appointment Management API

## Project Overview
The **Healthcare Appointment Management API** is a Go-based backend application for managing patients and their appointments in a healthcare system. It provides endpoints to:

- Create and retrieve patients.
- Create and retrieve appointments linked to patients.
- Handle database interactions with GORM and SQLite.

The API includes Swagger documentation to provide interactive documentation for the available endpoints.

---

---

## Table of Contents

- [Technologies Used](#technologies-used)
- [Project Structure](#project-structure)
- [Installation & Setup](#installation-setup)
- [API Endpoints](#api-endpoints)
  - [Patients](#patients)
  - [Appointments](#appointments)
- [Database](#database)
- [Swagger Documentation](#swagger-documentation)

---

## Technologies Used
- **Go** - Backend language for the API.
- **Gin** - Web framework for routing and HTTP handling.
- **GORM** - ORM for database interaction.
- **SQLite** - Database for storing patient and appointment data.
- **Swagger** - API documentation tool.
- **CORS** - Cross-Origin Resource Sharing support via Gin middleware.

---

## Project Structure

```bash
healthcare-app/
├── config/
│   └── database.go         # Database connection and migration setup
├── controllers/
│   ├── appointment_controller.go  # Appointment-related API logic
│   └── patient_controller.go     # Patient-related API logic
├── middleware/
│   └── logger.go           # Custom logger middleware for request logging
├── models/
│   ├── appointment.go      # Appointment data model
│   └── patient.go          # Patient data model
├── routes/
│   └── routes.go           # Route setup for the API
├── docs/
│   └── swagger.json        # Swagger generated documentation
└── main.go                 # Main application entry point
```

---

## Installation & Setup

### Prerequisites
Ensure you have the following installed on your machine:
- **Go 1.15+**
- **Git**
- **SQLite** (SQLite is bundled with Go in this case, but you can install the command-line tools for managing the database)

### Steps to Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/bleronsy/healthcare-backend.git
   cd healthcare-app
   ```
2. **Install dependencies**
    ```bash
    go get ./...
    ```
3. **Run the application**
    ```bash
    go run main.go
    ```

By default, the application will run on `http://localhost:8000`.

## API Endpoints

### Patients

#### 1. `GET /patients`
Retrieve a paginated list of all patients.

**Query Parameters**:
- `page` (integer): Page number (default: `1`)
- `limit` (integer): Items per page (default: `10`)

**Responses**:
- `200 OK`: List of patients.
- `500 Internal Server Error`: Error retrieving patients.

---

#### 2. `GET /patients/{id}`
Retrieve a specific patient by their ID.

**Path Parameters**:
- `id` (integer): The ID of the patient.

**Responses**:
- `200 OK`: The requested patient.
- `400 Bad Request`: Invalid ID.
- `404 Not Found`: Patient not found.

---

#### 3. `POST /patients`
Create a new patient.

**Request Body**:
- `name` (string): Patient's name (required).
- `email` (string): Patient's email (required).

**Responses**:
- `201 Created`: Patient created.
- `400 Bad Request`: Invalid data.
- `500 Internal Server Error`: Error creating patient.

---

### Appointments

#### 1. `GET /appointments`
Retrieve a paginated list of all appointments.

**Query Parameters**:
- `page` (integer): Page number (default: `1`).
- `limit` (integer): Items per page (default: `10`).

**Responses**:
- `200 OK`: List of appointments.
- `500 Internal Server Error`: Error retrieving appointments.

---

#### 2. `POST /patients/{id}/appointments`
Create a new appointment for a specific patient.

**Path Parameters**:
- `id` (integer): The ID of the patient.

**Request Body**:
- `date` (string): Date of the appointment (required).
- `time` (string): Time of the appointment (required).
- `reason` (string): Reason for the appointment (optional).
- `notes` (string): Notes for the appointment (optional).

**Responses**:
- `201 Created`: Appointment created.
- `400 Bad Request`: Invalid data.
- `404 Not Found`: Patient not found.
- `500 Internal Server Error`: Error creating appointment.

---

#### 3. `GET /appointments/patient/{id}`
Retrieve all appointments for a specific patient by patient ID.

**Path Parameters**:
- `id` (integer): The ID of the patient.

**Responses**:
- `200 OK`: List of appointments for the patient.
- `404 Not Found`: Patient not found.
- `500 Internal Server Error`: Error retrieving appointments.

---

## Database

The project uses **SQLite** for storing patient and appointment data. The database file is named `healthcare.db` and is created automatically on the first run. The database schema is generated using GORM's `AutoMigrate` feature.

To initialize the database, the `InitDatabase` function in `config/database.go` will automatically set up the necessary tables for **patients** and **appointments**.

---

## Swagger Documentation

This project uses **Swagger** for automatic API documentation. The Swagger documentation is generated based on the Go code annotations and can be accessed via the `/swagger` route (once you set up Swagger UI).

To regenerate the Swagger documentation after modifying the code, run the following command:

```bash
swag init
```

This will regenerate the swagger.json and swagger.yaml files in the docs/ folder.