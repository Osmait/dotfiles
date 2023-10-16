from repository import  UserRepository

userRepository = UserRepository()

class UserService:
    def __init__(self, user):
        self.user = user

    def createUser(self):
        userRepository.createUser(self.user)