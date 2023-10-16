
from app import db
from model import User
from repository import UserRepository


user_repository = UserRepository(db)


class CreateUser():
    def __init__(self, username, password):
        self.username = username
        self.password = password

    def create_user(self,):
        user = User(self.username, self.password)
        user_repository.save(user)
