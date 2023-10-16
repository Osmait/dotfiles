from fastapi import APIRouter
from typing import Union
from service import UserService


user_router = APIRouter()





@user_router.get("/")
def read_root():

    user=UserService("saulburgos6@","1234567")
    user.createUser()
    
    
    return 


@user_router.get("/items/{item_id}")
def read_item(item_id: int, q: Union[str, None] = None):
    return {"item_id": item_id, "q": q}