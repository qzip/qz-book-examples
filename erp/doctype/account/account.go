package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	message string
}

func (e *ValidationError) Error() string {
	return e.message
}

type Account struct {
	ParentAccount string
	AccountType   string
	IsGroup       bool
	Company       string
}

func (a *Account) ValidateParentChildAccountType() error {
	if a.ParentAccount != "" {
		if a.AccountType == "Direct Income" ||
			a.AccountType == "Indirect Income" ||
			a.AccountType == "Current Asset" ||
			a.AccountType == "Current Liability" ||
			a.AccountType == "Direct Expense" ||
			a.AccountType == "Indirect Expense" {
			parentAccountType := GetAccountType(a.ParentAccount)
			if parentAccountType == a.AccountType {
				return errors.New(fmt.Sprintf("Only Parent can be of type %s", a.AccountType))
			}
		}
	}
	return nil
}

func (a *Account) ValidateParent() error {
	if a.ParentAccount != "" {
		par := GetAccount(a.ParentAccount)
		if par == nil {
			return errors.New(fmt.Sprintf("Account %s: Parent account %s does not exist", a.Name, a.ParentAccount))
		} else if !par.IsGroup {
			return errors.New(fmt.Sprintf("Account %s: Parent account %s can not be a ledger", a.Name, a.ParentAccount))
		} else if par.Company != a.Company {
			return errors.New(fmt.Sprintf("Account %s: Parent account %s does not belong to company: %s", a.Name, a.ParentAccount, a.Company))
		}
	}
	return nil
}

func (a *Account) Validate() error {
	err := a.ValidateParent()
	if err != nil {
		return err
	}
	err = a.ValidateParentChildAccountType()
	if err != nil {
		return err
	}
	// validate other fields
	return nil
}

func GetAccountType(account string) string {
	// implementation of GetAccountType
	return ""
}

func GetAccount(account string) *Account {
	// implementation of GetAccount
	return nil
}

func main() {
	// example usage
	account := &Account{
		ParentAccount: "parent",
		AccountType:   "Direct Income",
		IsGroup:       false,
		Company:       "company",
	}
	err := account.Validate()
	if err != nil {
		fmt.Println(err)
	}
}
