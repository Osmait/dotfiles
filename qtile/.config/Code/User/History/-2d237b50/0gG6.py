from domain.user import User


class CreateUser():

    def __init__(self, db) -> None:
        self.db = db

    def create_user(self, username, password):
        user = User(username=username, password=password)
        self.db.session.add(user)
        self.db.session.commit()
