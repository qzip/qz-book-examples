package main

import (
	"fmt"
)

func get() map[string]interface{} {
	return map[string]interface{}{
		"Application of Funds (Assets)": map[string]interface{}{
			"Current Assets": map[string]interface{}{
				"Accounts Receivable": map[string]interface{}{
					"Debtors": map[string]interface{}{
						"account_type": "Receivable",
					},
				},
				"Bank Accounts": map[string]interface{}{
					"account_type": "Bank",
					"is_group":     1,
				},
				"Cash In Hand": map[string]interface{}{
					"Cash": map[string]interface{}{
						"account_type": "Cash",
					},
					"account_type": "Cash",
				},
				"Loans and Advances (Assets)": map[string]interface{}{
					"Employee Advances": map[string]interface{}{},
				},
				"Securities and Deposits": map[string]interface{}{
					"Earnest Money": map[string]interface{}{},
				},
				"Stock Assets": map[string]interface{}{
					"Stock In Hand": map[string]interface{}{
						"account_type": "Stock",
					},
					"account_type": "Stock",
				},
				"Tax Assets": map[string]interface{}{
					"is_group": 1,
				},
			},
			"Fixed Assets": map[string]interface{}{
				"Capital Equipments": map[string]interface{}{
					"account_type": "Fixed Asset",
				},
				"Electronic Equipments": map[string]interface{}{
					"account_type": "Fixed Asset",
				},
				"Furnitures and Fixtures": map[string]interface{}{
					"account_type": "Fixed Asset",
				},
				"Office Equipments": map[string]interface{}{
					"account_type": "Fixed Asset",
				},
				"Plants and Machineries": map[string]interface{}{
					"account_type": "Fixed Asset",
				},
				"Buildings": map[string]interface{}{
					"account_type": "Fixed Asset",
				},
				"Softwares": map[string]interface{}{
					"account_type": "Fixed Asset",
				},
				"Accumulated Depreciation": map[string]interface{}{
					"account_type": "Accumulated Depreciation",
				},
				"CWIP Account": map[string]interface{}{
					"account_type": "Capital Work in Progress",
				},
			},
			"Investments": map[string]interface{}{
				"is_group": 1,
			},
			"Temporary Accounts": map[string]interface{}{
				"Temporary Opening": map[string]interface{}{
					"account_type": "Temporary",
				},
			},
			"root_type": "Asset",
		},
		"Expenses": map[string]interface{}{
			"Direct Expenses": map[string]interface{}{
				"Stock Expenses": map[string]interface{}{
					"Cost of Goods Sold": map[string]interface{}{
						"account_type": "Cost of Goods Sold",
					},
					"Expenses Included In Asset Valuation": map[string]interface{}{
						"account_type": "Expenses Included In Asset Valuation",
					},
					"Expenses Included In Valuation": map[string]interface{}{
						"account_type": "Expenses Included In Valuation",
					},
					"Stock Adjustment": map[string]interface{}{
						"account_type": "Stock Adjustment",
					},
				},
			},
			"Indirect Expenses": map[string]interface{}{
				"Administrative Expenses": map[string]interface{}{},
				"Commission on Sales":     map[string]interface{}{},
				"Depreciation": map[string]interface{}{
					"account_type": "Depreciation",
				},
				"Entertainment Expenses": map[string]interface{}{},
				"Freight and Forwarding Charges": map[string]interface{}{
					"account_type": "Chargeable",
				},
				"Legal Expenses": map[string]interface{}{},
				"Marketing Expenses": map[string]interface{}{
					"account_type": "Chargeable",
				},
				"Miscellaneous Expenses": map[string]interface{}{
					"account_type": "Chargeable",
				},
				"Office Maintenance Expenses": map[string]interface{}{},
				"Office Rent":                map[string]interface{}{},
				"Postal Expenses":            map[string]interface{}{},
				"Print and Stationery":       map[string]interface{}{},
				"Round Off": map[string]interface{}{
					"account_type": "Round Off",
				},
				"Salary":          map[string]interface{}{},
				"Sales Expenses":  map[string]interface{}{},
				"Telephone Expenses": map[string]interface{}{},
				"Travel Expenses":    map[string]interface{}{},
				"Utility Expenses":   map[string]interface{}{},
				"Write Off":          map[string]interface{}{},
				"Exchange Gain/Loss": map[string]interface{}{},
				"Gain/Loss on Asset Disposal": map[string]interface{}{},
			},
			"root_type": "Expense",
		},
		"Income": map[string]interface{}{
			"Direct Income": map[string]interface{}{
				"Sales":   map[string]interface{}{},
				"Service": map[string]interface{}{},
			},
			"Indirect Income": map[string]interface{}{
				"is_group": 1,
			},
			"root_type": "Income",
		},
		"Source of Funds (Liabilities)": map[string]interface{}{
			"Current Liabilities": map[string]interface{}{
				"Accounts Payable": map[string]interface{}{
					"Creditors": map[string]interface{}{
						"account_type": "Payable",
					},
					"Payroll Payable": map[string]interface{}{},
				},
				"Stock Liabilities": map[string]interface{}{
					"Stock Received But Not Billed": map[string]interface{}{
						"account_type": "Stock Received But Not Billed",
					},
					"Asset Received But Not Billed": map[string]interface{}{
						"account_type": "Asset Received But Not Billed",
					},
				},
				"Duties and Taxes": map[string]interface{}{
					"account_type": "Tax",
					"is_group":     1,
				},
				"Loans (Liabilities)": map[string]interface{}{
					"Secured Loans":        map[string]interface{}{},
					"Unsecured Loans":      map[string]interface{}{},
					"Bank Overdraft Account": map[string]interface{}{},
				},
			},
			"root_type": "Liability",
		},
		"Equity": map[string]interface{}{
			"Capital Stock":          map[string]interface{}{"account_type": "Equity"},
			"Dividends Paid":         map[string]interface{}{"account_type": "Equity"},
			"Opening Balance Equity": map[string]interface{}{"account_type": "Equity"},
			"Retained Earnings":      map[string]interface{}{"account_type": "Equity"},
			"root_type":              "Equity",
		},
	}
}

func main() {
	result := get()
	fmt.Println(result)
}


