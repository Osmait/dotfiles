from flask import Flask
from infraestruture.controller.user import userRoutes
from flask_sqlalchemy import SQLAlchemy


dbUser = "osmait"
dbPassword = "admin123"
dbName = "my_store"
dbHost = "localhost"
dbPort = 5432

app = Flask(__name__)
app.register_blueprint(userRoutes, url_prefix='/api')
app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///project.db"
db = SQLAlchemy(app)


with app.app_context():
    db.create_all()

if __name__ == '__main__':
    app.run(debug=True)
