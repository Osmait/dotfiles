
from dto import User

class UserRepository:
    def __init__(self, db):
        self.db = db

    def get_user_by_id(self, user_id):
        return self.db.query(User).filter(User.id == user_id).first()

    def get_user_by_email(self, email):
        return self.db.query(User).filter(User.email == email).first()

    def create_user(self, user):
        self.db.add(user)
        self.db.commit()
        return user

    def update_user(self, user):
        self.db.add(user)
        self.db.commit()
        return user

    def delete_user(self, user):
        self.db.delete(user)
        self.db.commit()
        return user