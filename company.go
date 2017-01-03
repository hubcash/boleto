package goboleto

/*
Defines a company type,
holds the data of the emissor

@LegalName is important, you must set this correctly
its the 'Razão social' of your company

 */
type Company struct {
	Name 		string
	LegalName 	string
	Document 	string
	Address 	string
	Contact 	string
}