# urlshort

Create short URLs for frequently access resouces.

## Steps to use
1. Edit 'pathToUrls' map in main.go line 13 to include your frequently accessed resources. The key in the map will be the short url you want to use to access it by
2. Update your browser proxy settings:
    1. Configure it to point to address: localhost port: 8080
    2. Add *.* as an exception - this will make only requests on the intranet use this proxy
3. Enter your short name into your browser as "shorturl/"

**Note: If the short URL is not recognized you will be sent to Google
