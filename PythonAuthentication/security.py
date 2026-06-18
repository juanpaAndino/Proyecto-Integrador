import os
import base64
import hashlib
from dotenv import load_dotenv
from Crypto.Cipher import AES
from Crypto.Util.Padding import pad, unpad
from Crypto.Random import get_random_bytes

load_dotenv()

SECRET_STR = os.getenv("PEPPER_SECRET", "ClavePorDefecto123")
AES_KEY = hashlib.sha256(SECRET_STR.encode('utf-8')).digest()

def encrypt_password(password: str) -> str:
    """Cifra la contraseña usando AES en modo CBC."""
    
    iv = get_random_bytes(16)
    
    cipher = AES.new(AES_KEY, AES.MODE_CBC, iv)
    
    padded_password = pad(password.encode('utf-8'), AES.block_size)
    
    encrypted_bytes = cipher.encrypt(padded_password)
    
    final_payload = iv + encrypted_bytes
    return base64.b64encode(final_payload).decode('utf-8')

def verify_password(plain_password: str, encrypted_password_b64: str) -> bool:
    """Desencripta el hash de la DB y verifica si coincide con la contraseña ingresada."""
    try:
        raw_data = base64.b64decode(encrypted_password_b64)
        
        iv = raw_data[:16]
        encrypted_bytes = raw_data[16:]
        
        cipher = AES.new(AES_KEY, AES.MODE_CBC, iv)
        
        decrypted_padded = cipher.decrypt(encrypted_bytes)
        decrypted_password = unpad(decrypted_padded, AES.block_size).decode('utf-8')
        
        return plain_password == decrypted_password
        
    except Exception:
        return False