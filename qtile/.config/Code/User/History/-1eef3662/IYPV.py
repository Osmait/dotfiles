

# from domain.user import User


class UserRepository:
    def __init__(self, db_session):
        self.db_session = db_session

    # def get_user_by_id(self, user_id):
    #     return self.db_session.query(User).filter(User.id == user_id).first()

    # def get_user_by_email(self, email):
    #     return self.db_session.query(User).filter(User.email == email).first()

    def create_user(self, user) -> None:

        try:

            self.db_session.add(user)
            self.db_session.commit()
        except Exception as e:
            print(e)
            return e
        return

    def update_user(self, user):
        self.db_session.add(user)
        self.db_session.commit()
        return user

    def delete_user(self, user):
        self.db_session.delete(user)
        self.db_session.commit()
        return user
