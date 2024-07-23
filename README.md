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
### from localhost
1. Initialize the ClientId and Secret variables in main.go with the client id and secret of your sandbox app
2. Open a terminal in the sandbox
3. Run the command "chmod +x run.sh"
4. Run the command "./run.sh"
5. The application will be available at http://localhost:8123
6. Follow the instructions in the browser once the application is opened.

### from code sandbox
1. Open the project in code sandbox
2. Click the "sign in" button with code sandbox
3. Choose "sign in" with github
4. Authenticate using your github credentials and grant codesandbox access to your repos
5. After signing in, click the dropdown next to the "sign in" button in codesandbox and choose "Fork project"
6. When prompted, search for this repository to fork from: https://github.com/fabrizio-akoya/go-fetch-accounts
7. Click the "fork" button to fork the project
8. The application will start to run in codesandbox, after which you'll be prompted to open preview via a "Preview: 8123" button
9. Click the "Preview: 8123" button to open the application in the browser
10. In the tool bar next to the preview URL bar, note the icons: Hover over the icons and look for the tooltip text "Open in a new tab" and click that
11. The application will start running in a new browser window where you can interact with it