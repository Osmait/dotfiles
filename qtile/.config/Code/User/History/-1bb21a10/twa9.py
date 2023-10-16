from blog.app.run import db



class User(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    username = db.Column(db.String(50), unique=True)
    email = db.Column(db.String(100), unique=True)

    def serialize(self):
        return {
            'id': self.id,
            'username': self.username,
            'email': self.email
        }