import secrets

# Generate a secure random 32-byte secret key
jwt_secret_key = secrets.token_urlsafe(32)
print("Your secure JWT secret key:", jwt_secret_key)
