package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Account struct {
	AccountName     string `json:"account_name"`
	AccountNumber   string `json:"account_number"`
	AccountType     string `json:"account_type"`
	RootType        string `json:"root_type"`
	IsGroup         bool   `json:"is_group"`
	TaxRate         string `json:"tax_rate"`
	AccountCurrency string `json:"account_currency"`
}

type Chart struct {
	Name string            `json:"name"`
	Tree map[string]Account `json:"tree"`
}

func main() {
	company := ""
	chartTemplate := ""
	existingCompany := ""
	customChart := ""
	fromCoaImporter := false

	createCharts(company, chartTemplate, existingCompany, customChart, fromCoaImporter)
}

func createCharts(company, chartTemplate, existingCompany, customChart string, fromCoaImporter bool) {
	chart := getChart(chartTemplate, existingCompany)
	if chart != nil {
		accounts := make([]Account, 0)
		var importAccounts func(map[string]interface{}, string, string, bool)
		importAccounts = func(children map[string]interface{}, parent, rootType string, rootAccount bool) {
			for accountName, child := range children {
				if rootAccount {
					rootType = child.(map[string]interface{})["root_type"].(string)
				}
				if accountName != "account_name" && accountName != "account_number" && accountName != "account_type" && accountName != "root_type" && accountName != "is_group" && accountName != "tax_rate" && accountName != "account_currency" {
					accountNumber := strings.TrimSpace(fmt.Sprintf("%v", child.(map[string]interface{})["account_number"]))
					accountName, accountNameInDB := addSuffixIfDuplicate(accountName, accountNumber, accounts)
					isGroup := identifyIsGroup(child.(map[string]interface{}))
					reportType := ""
					if rootType == "Asset" || rootType == "Liability" || rootType == "Equity" {
						reportType = "Balance Sheet"
					} else {
						reportType = "Profit and Loss"
					}
					account := Account{
						AccountName:     child.(map[string]interface{})["account_name"].(string),
						AccountNumber:   accountNumber,
						AccountType:     child.(map[string]interface{})["account_type"].(string),
						RootType:        rootType,
						IsGroup:         isGroup,
						TaxRate:         child.(map[string]interface{})["tax_rate"].(string),
						AccountCurrency: child.(map[string]interface{})["account_currency"].(string),
					}
					if rootAccount || allowUnverifiedCharts {
						account.Flags.IgnoreMandatory = true
					}
					account.Flags.IgnorePermissions = true
					account.Insert()
					accounts = append(accounts, accountNameInDB)
					importAccounts(child.(map[string]interface{}), account.Name, rootType, false)
				}
			}
		}

		ignoreUpdateNSM := true
		importAccounts(chart.Tree, "", "", true)
		rebuildTree("Account", "parent_account")
		ignoreUpdateNSM = false
	}
}

func addSuffixIfDuplicate(accountName, accountNumber string, accounts []Account) (string, string) {
	accountNameInDB := ""
	if accountNumber != "" {
		accountNameInDB = unidecode(fmt.Sprintf("%s - %s", accountNumber, strings.ToLower(strings.TrimSpace(accountName)))))
	} else {
		accountNameInDB = unidecode(strings.ToLower(strings.TrimSpace(accountName)))
	}
	count := 0
	for _, acc := range accounts {
		if acc.AccountName == accountNameInDB {
			count++
		}
	}
	if count > 0 {
		accountName = fmt.Sprintf("%s %d", accountName, count)
	}
	return accountName, accountNameInDB
}

func identifyIsGroup(child map[string]interface{}) bool {
	isGroup := false
	if child["is_group"] != nil {
		isGroup = child["is_group"].(bool)
	} else if len(child) > 7 {
		isGroup = true
	}
	return isGroup
}

func getChart(chartTemplate, existingCompany string) *Chart {
	chart := &Chart{}
	if existingCompany != "" {
		return getAccountTreeFromExistingCompany(existingCompany)
	} else if chartTemplate == "Standard" {
		chart = getStandardChartOfAccounts()
	} else if chartTemplate == "Standard with Numbers" {
		chart = getStandardChartOfAccountsWithAccountNumber()
	} else {
		folders := []string{"verified"}
		if allowUnverifiedCharts {
			folders = []string{"verified", "unverified"}
		}
		for _, folder := range folders {
			path := fmt.Sprintf("%s/%s", os.Getenv("GOPATH"), folder)
			files, err := ioutil.ReadDir(path)
			if err != nil {
				continue
			}
			for _, file := range files {
				if strings.HasSuffix(file.Name(), ".json") {
					content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", path, file.Name()))
					if err != nil {
						continue
					}
					json.Unmarshal(content, &chart)
					if chart.Name == chartTemplate {
						return chart
					}
				}
			}
		}
	}
	return chart
}

func getStandardChartOfAccounts() *Chart {
	chart := &Chart{}
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", os.Getenv("GOPATH"), "erpnext/accounts/doctype/account/chart_of_accounts/verified/standard.json"))
	if err != nil {
		return chart
	}
	json.Unmarshal(content, &chart)
	return chart
}

func getStandardChartOfAccountsWithAccountNumber() *Chart {
	chart := &Chart{}
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", os.Getenv("GOPATH"), "erpnext/accounts/doctype/account/chart_of_accounts/verified/standard_with_account_number.json"))
	if err != nil {
		return chart
	}
	json.Unmarshal(content, &chart)
	return chart
}

func getAccountTreeFromExistingCompany(existingCompany string) *Chart {
	allAccounts := make([]Account, 0)
	// Get all accounts from existing company
	accountTree := make(map[string]interface{})
	if len(allAccounts) > 0 {
		buildAccountTree(accountTree, nil, allAccounts)
	}
	chart := &Chart{
		Tree: accountTree,
	}
	return chart
}

func buildAccountTree(tree map[string]interface{}, parent *Account, allAccounts []Account) {
	parentAccount := ""
	if parent != nil {
		parentAccount = parent.Name
	}
	children := make([]Account, 0)
	for _, acc := range allAccounts {
		if acc.ParentAccount == parentAccount {
			children = append(children, acc)
		}
	}
	if len(children) == 0 && parent.IsGroup {
		tree["is_group"] = true
		tree["account_number"] = parent.AccountNumber
	}
	for _, child := range children {
		tree[child.AccountName] = make(map[string]interface{})
		if child.AccountNumber != "" {
			tree[child.AccountName]["account_number"] = child.AccountNumber
		}
		if child.AccountType != "" {
			tree[child.AccountName]["account_type"] = child.AccountType
		}
		if child.TaxRate != "" {
			tree[child.AccountName]["tax_rate"] = child.TaxRate
		}
		if parent == nil {
			tree[child.AccountName]["root_type"] = child.RootType
		}
		buildAccountTree(tree[child.AccountName].(map[string]interface{}), &child, allAccounts)
	}
}

func validateBankAccount(coa, bankAccount string) bool {
	accounts := make([]Account, 0)
	chart := getChart(coa, "")
	if chart != nil {
		getAccountNames(chart.Tree, &accounts)
	}
	for _, acc := range accounts {
		if acc.AccountName == bankAccount {
			return true
		}
	}
	return false
}

func getAccountNames(accountMaster map[string]interface{}, accounts *[]Account) {
	for accountName, child := range accountMaster {
		if accountName != "account_number" && accountName != "account_type" && accountName != "root_type" && accountName != "is_group" && accountName != "tax_rate" {
			*accounts = append(*accounts, Account{
				AccountName: accountName,
			})
			getAccountNames(child.(map[string]interface{}), accounts)
		}
	}
}

func buildTreeFromJSON(chartTemplate string, chartData map[string]interface{}, fromCoaImporter bool) []Account {
	chart := chartData
	if chart == nil {
		chart = getChart(chartTemplate, "")
	}
	accounts := make([]Account, 0)
	var importAccounts func(map[string]interface{}, string)
	importAccounts = func(children map[string]interface{}, parent string) {
		for accountName, child := range children {
			account := Account{}
			if accountName == "account_name" || accountName == "account_number" || accountName == "account_type" || accountName == "root_type" || accountName == "is_group" || accountName == "tax_rate" || accountName == "account_currency" {
				continue
			}
			if fromCoaImporter {
				accountName = child.(map[string]interface{})["account_name"].(string)
			}
			account.ParentAccount = parent
			account.Expandable = identifyIsGroup(child.(map[string]interface{}))
			account.Value = fmt.Sprintf("%s - %s", strings.TrimSpace(fmt.Sprintf("%v", child.(map[string]interface{})["account_number"])), accountName)
			if child.(map[string]interface{})["account_number"] == nil {
				account.Value = accountName
			}
			accounts = append(accounts, account)
			importAccounts(child.(map[string]interface{}), account.Value)
		}
	}
	importAccounts(chart, "")
	return accounts
}


