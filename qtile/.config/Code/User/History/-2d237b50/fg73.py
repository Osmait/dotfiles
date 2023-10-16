
from app import db
from model import User
from repository import UserRepository


user_repository = UserRepository(db)


class CreateUser:

    def __init__(self, db) -> None:
        self.db = db

    def save(self, username, password):
        user = User(username, password)
        user_repository.save(user)
