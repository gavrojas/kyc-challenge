# KYC Document Validation Application

This project is a document validation application that uses an external API to verify the authenticity of documents provided by a user. The application is divided into two parts: the backend (`kyc-api`) and the frontend (`kyc-app`).

## Main Technologies

- **Frontend**: Vue 3, Pinia, Vue Router, Vuetify, Tailwind CSS
- **Backend**: Go, CHI, GORM
- **Database**: PostgreSQL

## Project Structure

- **kyc-api**: Backend of the application.
- **kyc-app**: Frontend of the application.

## Backend Setup (kyc-api)

1. **Clone the repository**:
   ```bash
   git clone <REPOSITORY_URL>
   cd kyc-api
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Configure environment variables**:
   Create a `.env` file in the root of the project and define variables related to the database where information will be stored, API keys, external API URL, and port.

4. **Run the application**:
   ```bash
   go run main.go
   ```

5. **Migrate the database**:
   The application will automatically migrate the necessary tables upon startup.

## Frontend Setup (kyc-app)

1. **Clone the repository**:
   ```bash
   git clone <REPOSITORY_URL>
   cd kyc-app
   ```

2. **Install dependencies**:
   ```bash
   npm install
   ```

3. **Run the application**:
   ```bash
   npm run dev
   ```

4. **Open in the browser**:
   Visit `http://localhost:5173` to access the application.

## Application Usage

1. **Upload Documents**: Users can upload images of their documents (front and back) in step 3 of the onboarding process.
2. **Validation**: Upon submitting the documents, the application makes a call to the external API to validate the documents.
3. **Results**: The results of the validation will be displayed on the screen after a brief period.

## API Routes

- **POST /validations/create**: Creates a new validation.
  - The body expects a validation type, document type, country, timeout, and user authorization.
- **GET /validations/result**: Retrieves the result of the validation.
- **GET /validations/get-config**: Retrieves the configuration of allowed countries and document types.


#### Data Models

- **ValidationData**: Model for managing validation information.
- **Countries**: Model for managing data of allowed countries where validation is desired.
- **DocType**: Model for managing document types according to the country.

## Contacto

*For more information, you can email me at grojas9807@gmail.com or contact me on [LinkedIn](https://www.linkedin.com/in/gavrojas-dev/)*