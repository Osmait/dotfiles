from domain.user import User
from domain.user_repository import UserRepository
from app import db


user_repository = UserRepository(db)


class CreateUser:

    def __init__(self, db) -> None:
        self.db = db

    def create_user(self, username, password):
        user = User(username, password)
        user_repository.create_user(user)
