from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from user.routes import user_bp


app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///database.db'
db = SQLAlchemy(app)


app.register_blueprint(user_bp)



if __name__ == "__main__":
    app.run(debug=True)
