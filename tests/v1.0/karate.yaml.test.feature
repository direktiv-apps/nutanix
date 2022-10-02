
Feature: Basic

# The secrects can be used in the payload with the following syntax #(mysecretname)
Background:
* def nutanixpwd = karate.properties['nutanixpwd']


Scenario: get request

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
		"auth": {
			"host": "https://SERVER:9440",
			"password": #(nutanixpwd),
			"username": "USER",
			"skipVerify": true
		},
		"api": {
			"path": "/vms/list",
			"method": "POST",
			"body": {
				"kind": "vm"
			}
		}
	}
	"""
	When method POST
	Then status 200
	