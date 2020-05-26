from flask import Flask



def create_app():
    """Construct the core application."""
    app = Flask(__name__)

    app.app_context().push()
    # Imports Routes
    from .routes import security_socre_routes
    from .routes import os_info_routes

    return app
