- global service
- manage permission
- do not share the root account
- the list privilege principle

Groups:
	only contains users

Policy:
	a json file to manage users' and groups' permissions
		- version number: 
			- always include "2012-10-17"
		- Id (optional)
			- an identifier for the policy
		- Statements:
			- Sid: 
				- an identifier for the statement(optional)
			- Effect:
				- Allow/Deny
			- 