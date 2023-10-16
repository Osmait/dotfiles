from flask import Blueprint, request


user = Blueprint('user', __name__)


@user.route('/', methods=['GET'])
def hello():
    username = request.json.get('email')
    password = request.json.get('password')
    return f'Hello {username}!'
