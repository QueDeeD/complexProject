# CyberStore App

## Front

Frontend is written in `Vue.js`

To run frontend, do

```bash
cd frontend/
npm install
npm run dev
```

## Backend

Backend is written in `Go` and is using `PostgreSQL` as a database

To run backend, do
```bash
cd backend/
go run .
```

## Example Data

Use `backend/client/client.go` to fill a database with custom data from `backend/data`

Create categories:
```bash
go run ./backend/client/client.go --mode cat --file ./backend/data/categories.json
```

Add products:
```bash
go run ./backend/client/client.go --mode prod --file ./backend/data/products.json
```
