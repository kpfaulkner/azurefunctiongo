This is purely an example of how to write an Azure Function in Go.
This is all based off the Azure Custom Handlers ( see https://docs.microsoft.com/en-us/azure/azure-functions/functions-custom-handlers )

Note, currently I can only get this working if using Go 1.13.0. If I try anything further I get runtime errors, a bug 
has been filed with the go github repo.

Instructions:

Build the main.go, and deploy the executable, host.json and testfunction/function.json to the Azure Function of your choice.

I use my webjobdeploy program ( see https://github.com/kpfaulkner/webjobdeploy ) although there are many other ways to do this, pick what suits you.


