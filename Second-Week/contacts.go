package main

import "fmt"

func createContacts(name string, phoneNumber string, contacts map[string]string) {
	if _, ok := contacts[name]; !ok {
		contacts[name] = phoneNumber
		fmt.Println("Contact created!")
	} else {
		fmt.Println("Unable create contact, you don't have this contact!")
	}
}
func readContacts(name string, contacts map[string]string) {
	if phone, ok := contacts[name]; ok {
		fmt.Printf("%s, %s\n", name, phone)
	} else {
		fmt.Println("Unable to read contact, you don't have this contact!")
	}

}
func updateContacts(name string, newPhone string, contacts map[string]string) {
	if _, ok := contacts[name]; ok {
		oldPhone := contacts[name]
		contacts[name] = newPhone
		fmt.Printf("Updated phone number %s to %s\n", oldPhone, contacts[name])
	} else {
		fmt.Println("Unable to update contact, you don't have this contact!")
	}
}
func deleteContacts(name string, contacts map[string]string) {
	if _, ok := contacts[name]; ok {
		delete(contacts, name)
		fmt.Println("Contact deleted!")
	} else {
		fmt.Println("Unable to delete contact, you don't have this contact!")
	}
}
