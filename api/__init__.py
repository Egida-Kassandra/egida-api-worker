from flask import Flask



def create_app():
    """Construct the core application."""
    app = Flask(__name__)

    app.app_context().push()
    # Imports Routes
    from .routes import security_socre_routes

    return app
