# from flask import Flask

# from flask_sqlalchemy import SQLAlchemy
# from user.controller import userRoutes


# dbUser = "osmait"
# dbPassword = "admin123"
# dbName = "my_store"
# dbHost = "localhost"
# dbPort = 5432


# app = Flask(__name__)
# app.register_blueprint(userRoutes, url_prefix='/user')
# app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///project.db"
# db = SQLAlchemy(app)


# with app.app_context():
#     db.create_all()


from flask import Flask, request
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

from user.repository import UserRepository

app = Flask(__name__)

engine = create_engine('sqlite:///example.db')
Session = sessionmaker(bind=engine)


@app.route('/users', methods=['POST'])
def create_user():
    username = request.json.get('username')
    password = request.json.get('password')

    with Session() as session:
        user_repository = UserRepository(session)
        user = user_repository.create_user(username, password)

    return {'id': user.id, 'username': user.username}, 201


if __name__ == '__main__':
    app.run(debug=True)
