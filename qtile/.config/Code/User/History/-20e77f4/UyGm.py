
from router import blueprint


@blueprint.route("/")
def hello_world():
    return "Hello World!"
