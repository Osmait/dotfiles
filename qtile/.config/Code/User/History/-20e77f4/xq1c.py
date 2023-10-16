
from flask import Blueprint, make_response
from flask import request

# from user_repository import UserRepository


userRoutes = Blueprint('user', __name__)

# user_repository = UserRepository()


@userRoutes.route('/users', methods=['POST'])
def create_user():
    # email = request.json.get('email')
    # password = request.json.get('password')

    # try:
    #     # user_repository.create_user(email, password)
    # except Exception as e:
    #     return make_response(400)

    # return make_response(201)
    pass
