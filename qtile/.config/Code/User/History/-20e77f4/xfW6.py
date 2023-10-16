
from flask import Blueprint
from flask import request
from application.user import User


user = Blueprint('user', __name__)


@user.route('/users', methods=['POST'])
def create_user():
    email = request.json.get('email')
    password = request.json.get('password')
    user = User()
    user.save()
    return user.json()
