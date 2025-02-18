# KYC Document Validation Application

Este proyecto es una aplicación de validación de documentos que utiliza una API externa para verificar la autenticidad de los documentos proporcionados por un usuario. La aplicación está dividida en dos partes: el backend (`kyc-api`) y el frontend (`kyc-app`).

## Tecnologías Principales

- **Frontend**: Vue 3, Pinia, Vue Router, Vuetify, Tailwind CSS
- **Backend**: Go, CHI, GORM
- **Base de Datos**: PostgreSQL

## Estructura del Proyecto

- **kyc-api**: Backend de la aplicación.
- **kyc-app**: Frontend de la aplicación.

## Configuración del Backend (kyc-api)

1. **Clonar el repositorio**:
   ```bash
   git clone <URL_DEL_REPOSITORIO>
   cd kyc-api
   ```

2. **Instalar dependencias**:
   ```bash
   go mod tidy
   ```

3. **Configurar variables de entorno**:
   Crea un archivo `.env` en la raíz del proyecto y define variables relacionadas con la base de datos donde se guardará la información, api keys, url api externa, puerto:


4. **Ejecutar la aplicación**:
   ```bash
   go run main.go
   ```

5. **Migrar la base de datos**:
   La aplicación migrará automáticamente las tablas necesarias al iniciar.

## Configuración del Frontend (kyc-app)

1. **Clonar el repositorio**:
   ```bash
   git clone <URL_DEL_REPOSITORIO>
   cd kyc-app
   ```

2. **Instalar dependencias**:
   ```bash
   npm install
   ```

3. **Ejecutar la aplicación**:
   ```bash
   npm run dev
   ```

4. **Abrir en el navegador**:
   Visita `http://localhost:5173` para acceder a la aplicación.

## Uso de la Aplicación

1. **Subir Documentos**: Los usuarios pueden subir imágenes de sus documentos (frente y reverso) en el paso 3 del proceso de onboarding.
2. **Validación**: Al enviar los documentos, la aplicación realiza una llamada a la API externa para validar los documentos.
3. **Resultados**: Los resultados de la validación se mostrarán en la pantalla después de un breve período.

## Rutas de la API

- **POST /validations/create**: Crea una nueva validación.
 - En el body espera un tipo de validación, tipo de documento, pais, timeout y autorización del usuario
- **GET /validations/result**: Obtiene el resultado de la validación.
- **GET /validations/get-config**: Obtiene la configuración de los países y tipos de documentos.

#### Modelo de Datos

- **ValidationData**: Modelo para la gestión de la información de validación.
- **Countries**: Modelo para la gestión de datos de países permitidos y donde se desea hacer la validación.
- **DocType**: Modelo para la gestión de tipo de documento de acuerdo con el país. 

## Contacto

*Para mas información puedes escribirme a mi correo grojas9807@gmail.com o contactarme por [LinkedIn](https://www.linkedin.com/in/gavrojas-dev/)*