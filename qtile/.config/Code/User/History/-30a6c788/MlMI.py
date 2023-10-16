from flask import Blueprint


user = Blueprint('user', __name__)


@user.route('/', methods=['GET'])
def hello():
    return 'Hello World!'
