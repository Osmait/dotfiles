from flask import Blueprint, request


user = Blueprint('user', __name__)


@user.route('/', methods=['POST'])
def hello():
    email = request.json.get('email')
    password = request.json.get('password')
    return f'Hello {email}!'
