"""App entry point."""
from api import create_app


def run(app):
    app.run(host='0.0.0.0', port=3000)

def main():
    app = create_app()
    run(app)

if __name__ == '__main__':
    main()
