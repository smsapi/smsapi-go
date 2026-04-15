# Changes

## 1.4.0
- Add NewAllClient factory accepting custom API URL with all APIs (incl. MMS/VMS)

## 1.3.0
- Add Callbacks, MFA, Opt-Out, SMS Templates and Shipment APIs
- Add Blacklist CSV import endpoint
- Add Subusers shares (sendernames, blacklist, templates) management
- Add contact custom fields options and available fields helpers
- Add contact trash management (clean/restore) and AssignContactToGroups
- Fix 415 error on contact, contact group and group permissions PUT/POST
  endpoints by sending application/x-www-form-urlencoded bodies
- Fix subuser update accidentally deactivating account; User.Active is
  now *bool so it can be omitted or set in either direction
- Fix MFA CreateMfaCode.Fast type to *bool to allow explicit false
- ListUsers now accepts an optional query filter

## 1.2.3
- Remove not supported API endpoint (sendername activation)

## 1.2.2
- Fix Sub-user API parameters naming

## 1.2.1
- FIX: Extract Profile and Sub-users API end-to-end tests

## 1.2.0
- Extract Profile and Sub-users API

## 1.1.4
- Fix update contacts custom field API method
- Fix activate sendername API method

## 1.1.3
- Fix MMS response points attribute type

## 1.1.2
- Fix sms,mms,vms GET endpoints

## 1.1.1
- Update http method for some endpoints

## 1.1.0
- Page iteration
- Add page iteration support for blacklist and contacts collection
- Blacklist fix data type for ExpireAt field

## 1.0.1
- Support blacklist API