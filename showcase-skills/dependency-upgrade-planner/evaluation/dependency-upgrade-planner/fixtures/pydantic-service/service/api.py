from fastapi import FastAPI
from .models import Profile

app = FastAPI()

@app.post("/profiles", response_model=Profile)
def create_profile(profile: Profile) -> Profile:
    return profile
