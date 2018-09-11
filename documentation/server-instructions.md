# Server Instructions
## A reasonably comprehensive guide to common server admin tasks.

- How do I get on my server to do junk there?
    - Get a terminal instance running.
    - Find your KeyPair `.pem` file
        - Mine is in `secrets/key.pem`
    - Call `ssh` and pass in the public IP for the server prefixed with a username and attaching the pem to the `-i` flag
    - `ssh ubuntu@18.214.39.35 -i key.pem`
    - Resources:
        - [SSH](https://www.howtogeek.com/311287/how-to-connect-to-an-ssh-server-from-windows-macos-or-linux/)

