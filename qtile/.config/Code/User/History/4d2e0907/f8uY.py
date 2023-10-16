from pydantic import Field,BaseModel


class User(BaseModel):
    email:str = Field(min_length=3,max_length=255)
    password:str = Field(min_length=6, max_length=20)


    class Config:
        schema_extra = {
            "example":{
                "email":"example@example.com",
                "password":"example"
            }
        }