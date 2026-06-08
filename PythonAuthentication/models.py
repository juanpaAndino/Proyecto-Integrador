from sqlmodel import SQLModel, Field
from pydantic import BaseModel
from typing import Optional

class User(SQLModel, table=True):
    id: Optional[int] = Field(default=None, primary_key=True)
    username: str = Field(index=True, unique=True)
    password_hash: str

class UserCredentials(BaseModel):
    username: str
    password: str