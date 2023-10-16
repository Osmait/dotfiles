from domain.user import User
from domain.user_repository import UserRepository
from share.data_base import Session


user_repository = UserRepository(Session)


class CreateUser():

    def __init__(self, db) -> None:
        self.db = db

    def create_user(self, username, password):
