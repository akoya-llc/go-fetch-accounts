# Using a token to fetch a user's accounts

This golang code sample shows how use a token to fetch a user's list of accounts

## key concepts
* Using the token to authenticate the request to the Akoya API
* Making a GET request to the Akoya API /account-info to fetch a user's accounts
* Displaying the user's accounts in a rudimentary HTML page

## overview of project
The main.go module includes a single route "/" which drops the user on a basic HTML page rendered from the content in ./templates/index.tmml. This page has an input field and a button for the user to paste their token into the field and click the button to request a list of accounts.

The go.mod file includes the application's golang dependencies, and the run.sh file is a shell script that consolidates the dependency fetching and application compilation, and binary execution into a single step for convenience.


## running the sample
1. Initialize the ClientId and Secret variables in main.go with the client id and secret of your sandbox app
2. Open a terminal in the sandbox
3. Run the command "chmod +x run.sh"
4. Run the command "./run.sh"
5. A dilaog will show in code sandbox indicating "Port 8123 has been opened" with buttons to "Open Preview" and "Open Externally". Click "Open Externally" to open the application in the browser.
6. Follow the instructions in the browser once the application is opened.