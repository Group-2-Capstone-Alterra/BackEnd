# **PetPal API**

![PetPal Logo](/docs/petpal.png)

## **About the Project**

PetPal App is an e-commerce platform specifically designed for pet owners. The application provides various features that make it easy for pet owners to buy products and consult with veterinarians, either online or offline. With a user-friendly UI, PetPal App allows users to find high-quality pet products and receive health advice from professional veterinarians.

## **Getting Started**

### Installation

```bash
$ git clone https://github.com/Group-2-Capstone-Alterra/BackEnd.git
```

### Running the Server

```bash
$ go run main.go
```

### Base URL

`http://localhost:5000/api`

## **API Endpoints**

### GET /rolls

- Retrieve a list of all sushi rolls
- Response: JSON array of Roll objects

### GET /rolls/:id

- Retrieve a single sushi roll by ID
- Response: JSON object of Roll

### POST /rolls

- Create a new sushi roll
- Request: JSON object of Roll
- Response: JSON object of created Roll

## **Prototype**

[View Prototype on Figma](https://www.figma.com/design/hVqvSWqgOSIv9V0oWxO9NL/Untitled?node-id=0-1)

## **Documentation (SwaggerHub)**

[View API Documentation on SwaggerHub](https://app.swaggerhub.com/apis-docs/WFHADIT/PETPAL/1.0.0)

## **ERD**

![ERD Diagram](/docs/erd.png)

## **Contributors**

- [Aditya Ramadhan](https://www.linkedin.com/in/adit6/)
- [Bagas Alfaristo Putra](https://www.linkedin.com/in/bagas-alfaristo-putra/)

## **License**

MIT License
