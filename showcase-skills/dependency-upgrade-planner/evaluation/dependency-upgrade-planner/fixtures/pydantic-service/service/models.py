from pydantic import BaseModel

class Profile(BaseModel):
    name: str

    class Config:
        orm_mode = True

def profile_from_row(row: object) -> Profile:
    return Profile.from_orm(row)

def parse_profile(payload: dict[str, object]) -> Profile:
    return Profile.parse_obj(payload)
