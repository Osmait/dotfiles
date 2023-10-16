
from flask import Blueprint, make_response
from flask import request
# from domain.user import User
from domain.user_repository import UserRepository

user_repository = UserRepository()


userRoutes = Blueprint('user', __name__)


@userRoutes.route('/users', methods=['POST'])
def create_user():
    # email = request.json.get('email')
    # password = request.json.get('password')
    # user = User(email, password)
    # try:
    #     user_repository.create_user(user)
    # except Exception as e:
    #     return make_response(400)

    # return make_response(201)
    pass
