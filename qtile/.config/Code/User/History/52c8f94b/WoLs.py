from flask import Blueprint, jsonify, request

from services.service import UserService



user_bp = Blueprint('user', __name__)
user_service = UserService()

@user_bp.route('/users', methods=['GET'])
def get_all_users():
    users = user_service.get_all_users()
    return jsonify([user.serialize() for user in users])

@user_bp.route('/users/<user_id>', methods=['GET'])
def get_user(user_id):
    user = user_service.get_user_by_id(user_id)
    if user is None:
        return jsonify({'error': 'User not found'}), 404
    return jsonify(user.serialize())

@user_bp.route('/users', methods=['POST'])
def create_user():
    user_data = request.get_json()
    user = user_service.create_user(user_data)
    return jsonify(user.serialize()), 201

@user_bp.route('/users/<user_id>', methods=['PUT'])
def update_user(user_id):
    user_data = request.get_json()
    user = user_service.update_user(user_id, user_data)
    if user is None:
        return jsonify({'error': 'User not found'}), 404
    return jsonify(user.serialize())

@user_bp.route('/users/<user_id>', methods=['DELETE'])
def delete_user(user_id):
    user = user_service.delete_user(user_id)
    if user is None:
        return jsonify({'error': 'User not found'}), 404
    return jsonify(user.serialize())