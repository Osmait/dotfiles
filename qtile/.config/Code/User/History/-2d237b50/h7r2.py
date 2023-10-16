
from app import db
from model import User
from repository import UserRepository


user_repository = UserRepository(db)


class CreateUser:

    def create_user(self, username, password):
        user = User(username, password)
        user_repository.save(user)
