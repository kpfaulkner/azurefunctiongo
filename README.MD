This is purely an example of how to write an Azure Function in Go.
This is all based off the Azure Custom Handlers ( see https://docs.microsoft.com/en-us/azure/azure-functions/functions-custom-handlers )

Instructions:

Build the main.go, and deploy the executable, host.json and testfunction/function.json to the Azure Function of your choice.

I use my webjobdeploy program ( see https://github.com/kpfaulkner/webjobdeploy ) although there are many other ways to do this, pick what suits you.

First make a zipfile (need to automate this) that has azurefunctiongo.exe, host.json and testfunction/function.json in it.
Then execute the command:

  webjobdeploy.exe -deploy azurefunction -zipfilename .\testfunction.zip -appServiceName kentest

If the Azure App Service is called "kentest" and the Azure Function is called "testfunction", then the URL to hit would be http://kentest.azurewebsites.net/api/testfunction


