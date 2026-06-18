from fastapi import APIRouter, HTTPException, Depends
from sqlmodel import Session, select
from database import get_session
from models import User, UserCredentials
from security import encrypt_password, verify_password

router = APIRouter(tags=["Autenticación"])

@router.post("/register")
def register(user_data: UserCredentials, session: Session = Depends(get_session)):
    statement = select(User).where(User.username == user_data.username)
    existing_user = session.exec(statement).first()
    
    if existing_user:
        raise HTTPException(status_code=400, detail="El nombre de usuario ya está en uso")
        
    hashed_pwd = encrypt_password(user_data.password)
    
    new_user = User(username=user_data.username, password_hash=hashed_pwd)
    session.add(new_user)
    session.commit()
    session.refresh(new_user)
    
    return {"message": "Usuario registrado exitosamente", "user_id": new_user.id}

@router.post("/login")
def login(user_data: UserCredentials, session: Session = Depends(get_session)):
    statement = select(User).where(User.username == user_data.username)
    db_user = session.exec(statement).first()
    
    error_msg = "Credenciales inválidas"
    
    if not db_user:
        raise HTTPException(status_code=401, detail=error_msg)
        
    if not verify_password(user_data.password, db_user.password_hash):
        raise HTTPException(status_code=401, detail=error_msg)
        
    return {"message": "Inicio de sesión exitoso", "username": db_user.username}