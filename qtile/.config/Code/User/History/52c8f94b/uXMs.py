from flask import Blueprint, jsonify

from blog.user.service import UserService




user = Blueprint('user', __name__)
user_service = UserService()

@user.route('/user')
def create_user():
 
    return 'User created!'

@user.route('/user')
def get_all_users():
    users = user_service.get_all_users()
    return jsonify([user.serialize() for user in users])

