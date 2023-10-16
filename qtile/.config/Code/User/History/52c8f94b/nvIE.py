from flask import Blueprint


from user.model import User


user = Blueprint('user', __name__)


@user.route('/user')
def create_user():
 
    return 'User created!'
