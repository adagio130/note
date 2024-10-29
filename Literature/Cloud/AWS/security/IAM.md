#Security 
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
			- Principal:
				- account/user/role to which this policy applied to 
			- Action:
				- list of actions this policy allows or denies
			- Resource:
				- list of resources to which the actions applied to
			- Condition:(optional)
				- conditions for when this policy is in effect
		- 