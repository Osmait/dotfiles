from sqlalchemy.orm import Session


from sqlalchemy import Column, Integer, String
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class User(Base):
    __tablename__ = 'users'

    id = Column(Integer, primary_key=True)
    username = Column(String)
    password = Column(String)


class UserRepository:
    def __init__(self, session: Session):
        self.session = session

    def create_user(self, username: str, password: str):
        user = User(username=username, password=password)
        self.session.add(user)
        self.session.commit()
        return user

    def get_user_by_id(self, user_id: int):
        # return self.session.query(User).filter(User.id == user_id).first()
        pass
