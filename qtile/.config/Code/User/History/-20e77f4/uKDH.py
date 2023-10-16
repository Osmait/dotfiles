
from flask import Blueprint, make_response
from flask import request

from service import CreateUser


userRoutes = Blueprint('user', __name__)


@userRoutes.route('/users', methods=['POST'])
def create_user():
    # email = request.json.get('email')
    # password = request.json.get('password')

    # try:

    # except Exception as e:
    #     return make_response(400)

    # return make_response(201)
    pass
