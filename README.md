Este proyecto es una tarea de practica de tareas con un backend hecho en Go y MongoDB, y un frontend en Nuxt 3 (Vue).

- Cómo ejecutar el backend:

1. Se debe tener Go y MongoDB instalados.
2. En el archivo `.env` en la carpeta del backend estan las variables de entorno con las rutas de conexion a la base de datos:

DB_URL=mongodb://localhost:27017
DB_NAME=Test_tasks
PORT=8080
DB_COLLECTION=tasks

3. En la carpeta del backend, se instalan dependencias con:
go mod tidy

4. Luego se corre el backend con:
go run main.go

El backend quedará corriendo en http://localhost:8080.

- Cómo ejecutar el frontend:

1. Se debe de tener Node.js y npm instalados.
2. En el archivo `.env` en la carpeta del frontend esta la ruta de conexion al backend:

BASE_URL_BACKEND=http://localhost:8080

3. En la carpeta del frontend, se instalan dependencias con:
npm install

4. Luego se corre el frontend con:
npm run dev

El software abrirá en http://localhost:3000.
