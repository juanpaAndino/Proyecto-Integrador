from fastapi import FastAPI
from database import create_db_and_tables
from routers import auth_router

app = FastAPI(title="API de Autenticación con AES CBC")

@app.on_event("startup")
def on_startup():
    create_db_and_tables()

app.include_router(auth_router.router)