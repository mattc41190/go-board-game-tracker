# Road Map
## A general list of things this project needs to be considered usable

Use [emojis](https://gist.github.com/rxaviers/7360908) to commuincate status when possible

- See if you can run mysql and your app on the same lightsail instance. :white_check_mark:
- Begin storing suggested games in the database.
    - Create database :white_check_mark:
    - Create a user in the database :white_check_mark:
    - Create a game in the database :white_check_mark:
    - Add code that uses GoLang to look up a game :white_check_mark:
    - Add look up code to a route handler :white_check_mark:
    - Extract DB code into a module :white_check_mark:
    - Move from hardcoded storage to DB storage :white_check_mark:
    - Clean up hardcoded storage code :soon:
    - Write DB tests
    - Encrypt passwords
    - Look up user by password
- Create Travis Deployment to LightSail
- Get project on personal computer :white_check_mark:
- Send backups of DB to S3 daily 


## Resources

https://www.sohamkamani.com/blog/2017/10/18/golang-adding-database-to-web-application/
https://www.alexedwards.net/blog/organising-database-access