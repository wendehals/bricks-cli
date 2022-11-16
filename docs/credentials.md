# Credentials

Credentials can be provdided by passing an option providing a credentials file. The content of the
credentials file looks like:

    {
        "username": "<rebrickable username>",
        "password": "<rebrickable password>",
        "api_key": "<rebrickable api key>"
    }

You don't need to pass the credentials file each time you call *bricks*. The easiest way to not
pass it manually is to place the credentials in a file `.bricks-credentials.json` in your user's
home directory. It will then be automatically loaded.

You need to obtain the Rebrickable API key from your profile's settings page on rebrickable.com.
The login credentials will be send in a secured way.

