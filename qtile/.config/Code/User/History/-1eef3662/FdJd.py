from model import User


class UserRepository():
    def __init__(self, db_session):
        self.db_session = db_session

    def save(self, email, password) -> None:

        try:

            user = User(email=email, password=password)
            self.db_session.add(user)
            self.db_session.commit()
        except Exception as e:
            print(e)
            return e
        return

    # def update_user(self, user):
    #     self.db_session.add(user)
    #     self.db_session.commit()
    #     return user

    # def delete_user(self, user):
    #     self.db_session.delete(user)
    #     self.db_session.commit()
    #     return user
