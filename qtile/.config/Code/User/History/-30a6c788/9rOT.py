from routes import blueprint


@blueprint.route('/user')
def hello():
    return 'Hello World!'
