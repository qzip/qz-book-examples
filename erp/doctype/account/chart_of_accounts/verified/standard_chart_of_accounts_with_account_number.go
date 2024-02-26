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
						"account_type":   "Receivable",
						"account_number": "1310",
					},
					"account_number": "1300",
				},
				"Bank Accounts": map[string]interface{}{
					"account_type":   "Bank",
					"is_group":       1,
					"account_number": "1200",
				},
				"Cash In Hand": map[string]interface{}{
					"Cash": map[string]interface{}{
						"account_type":   "Cash",
						"account_number": "1110",
					},
					"account_type":   "Cash",
					"account_number": "1100",
				},
				"Loans and Advances (Assets)": map[string]interface{}{
					"Employee Advances": map[string]interface{}{
						"account_number": "1610",
					},
					"account_number": "1600",
				},
				"Securities and Deposits": map[string]interface{}{
					"Earnest Money": map[string]interface{}{
						"account_number": "1651",
					},
					"account_number": "1650",
				},
				"Stock Assets": map[string]interface{}{
					"Stock In Hand": map[string]interface{}{
						"account_type":   "Stock",
						"account_number": "1410",
					},
					"account_type":   "Stock",
					"account_number": "1400",
				},
				"Tax Assets": map[string]interface{}{
					"is_group":       1,
					"account_number": "1500",
				},
				"account_number": "1100-1600",
			},
			"Fixed Assets": map[string]interface{}{
				"Capital Equipments": map[string]interface{}{
					"account_type":   "Fixed Asset",
					"account_number": "1710",
				},
				"Electronic Equipments": map[string]interface{}{
					"account_type":   "Fixed Asset",
					"account_number": "1720",
				},
				"Furnitures and Fixtures": map[string]interface{}{
					"account_type":   "Fixed Asset",
					"account_number": "1730",
				},
				"Office Equipments": map[string]interface{}{
					"account_type":   "Fixed Asset",
					"account_number": "1740",
				},
				"Plants and Machineries": map[string]interface{}{
					"account_type":   "Fixed Asset",
					"account_number": "1750",
				},
				"Buildings": map[string]interface{}{
					"account_type":   "Fixed Asset",
					"account_number": "1760",
				},
				"Softwares": map[string]interface{}{
					"account_type":   "Fixed Asset",
					"account_number": "1770",
				},
				"Accumulated Depreciation": map[string]interface{}{
					"account_type":   "Accumulated Depreciation",
					"account_number": "1780",
				},
				"CWIP Account": map[string]interface{}{
					"account_type":   "Capital Work in Progress",
					"account_number": "1790",
				},
				"account_number": "1700",
			},
			"Investments": map[string]interface{}{
				"is_group":       1,
				"account_number": "1800",
			},
			"Temporary Accounts": map[string]interface{}{
				"Temporary Opening": map[string]interface{}{
					"account_type":   "Temporary",
					"account_number": "1910",
				},
				"account_number": "1900",
			},
			"root_type":      "Asset",
			"account_number": "1000",
		},
		"Expenses": map[string]interface{}{
			"Direct Expenses": map[string]interface{}{
				"Stock Expenses": map[string]interface{}{
					"Cost of Goods Sold": map[string]interface{}{
						"account_type":   "Cost of Goods Sold",
						"account_number": "5111",
					},
					"Expenses Included In Asset Valuation": map[string]interface{}{
						"account_type":   "Expenses Included In Asset Valuation",
						"account_number": "5112",
					},
					"Expenses Included In Valuation": map[string]interface{}{
						"account_type":   "Expenses Included In Valuation",
						"account_number": "5118",
					},
					"Stock Adjustment": map[string]interface{}{
						"account_type":   "Stock Adjustment",
						"account_number": "5119",
					},
					"account_number": "5110",
				},
				"account_number": "5100",
			},
			"Indirect Expenses": map[string]interface{}{
				"Administrative Expenses": map[string]interface{}{
					"account_number": "5201",
				},
				"Commission on Sales": map[string]interface{}{
					"account_number": "5202",
				},
				"Depreciation": map[string]interface{}{
					"account_type":   "Depreciation",
					"account_number": "5203",
				},
				"Entertainment Expenses": map[string]interface{}{
					"account_number": "5204",
				},
				"Freight and Forwarding Charges": map[string]interface{}{
					"account_type":   "Chargeable",
					"account_number": "5205",
				},
				"Legal Expenses": map[string]interface{}{
					"account_number": "5206",
				},
				"Marketing Expenses": map[string]interface{}{
					"account_type":   "Chargeable",
					"account_number": "5207",
				},
				"Office Maintenance Expenses": map[string]interface{}{
					"account_number": "5208",
				},
				"Office Rent": map[string]interface{}{
					"account_number": "5209",
				},
				"Postal Expenses": map[string]interface{}{
					"account_number": "5210",
				},
				"Print and Stationery": map[string]interface{}{
					"account_number": "5211",
				},
				"Round Off": map[string]interface{}{
					"account_type":   "Round Off",
					"account_number": "5212",
				},
				"Salary": map[string]interface{}{
					"account_number": "5213",
				},
				"Sales Expenses": map[string]interface{}{
					"account_number": "5214",
				},
				"Telephone Expenses": map[string]interface{}{
					"account_number": "5215",
				},
				"Travel Expenses": map[string]interface{}{
					"account_number": "5216",
				},
				"Utility Expenses": map[string]interface{}{
					"account_number": "5217",
				},
				"Write Off": map[string]interface{}{
					"account_number": "5218",
				},
				"Exchange Gain/Loss": map[string]interface{}{
					"account_number": "5219",
				},
				"Gain/Loss on Asset Disposal": map[string]interface{}{
					"account_number": "5220",
				},
				"Miscellaneous Expenses": map[string]interface{}{
					"account_type":   "Chargeable",
					"account_number": "5221",
				},
				"account_number": "5200",
			},
			"root_type":      "Expense",
			"account_number": "5000",
		},
		"Income": map[string]interface{}{
			"Direct Income": map[string]interface{}{
				"Sales": map[string]interface{}{
					"account_number": "4110",
				},
				"Service": map[string]interface{}{
					"account_number": "4120",
				},
				"account_number": "4100",
			},
			"Indirect Income": map[string]interface{}{
				"is_group":       1,
				"account_number": "4200",
			},
			"root_type":      "Income",
			"account_number": "4000",
		},
		"Source of Funds (Liabilities)": map[string]interface{}{
			"Current Liabilities": map[string]interface{}{
				"Accounts Payable": map[string]interface{}{
					"Creditors": map[string]interface{}{
						"account_type":   "Payable",
						"account_number": "2110",
					},
					"Payroll Payable": map[string]interface{}{
						"account_number": "2120",
					},
					"account_number": "2100",
				},
				"Stock Liabilities": map[string]interface{}{
					"Stock Received But Not Billed": map[string]interface{}{
						"account_type":   "Stock Received But Not Billed",
						"account_number": "2210",
					},
					"Asset Received But Not Billed": map[string]interface{}{
						"account_type":   "Asset Received But Not Billed",
						"account_number": "2211",
					},
					"account_number": "2200",
				},
				"Duties and Taxes": map[string]interface{}{
					"TDS Payable": map[string]interface{}{
						"account_number": "2310",
					},
					"account_type":   "Tax",
					"is_group":       1,
					"account_number": "2300",
				},
				"Loans (Liabilities)": map[string]interface{}{
					"Secured Loans": map[string]interface{}{
						"account_number": "2410",
					},
					"Unsecured Loans": map[string]interface{}{
						"account_number": "2420",
					},
					"Bank Overdraft Account": map[string]interface{}{
						"account_number": "2430",
					},
					"account_number": "2400",
				},
				"account_number": "2100-2400",
			},
			"root_type":      "Liability",
			"account_number": "2000",
		},
		"Equity": map[string]interface{}{
			"Capital Stock": map[string]interface{}{
				"account_type":   "Equity",
				"account_number": "3100",
			},
			"Dividends Paid": map[string]interface{}{
				"account_type":   "Equity",
				"account_number": "3200",
			},
			"Opening Balance Equity": map[string]interface{}{
				"account_type":   "Equity",
				"account_number": "3300",
			},
			"Retained Earnings": map[string]interface{}{
				"account_type":   "Equity",
				"account_number": "3400",
			},
			"root_type":      "Equity",
			"account_number": "3000",
		},
	}
}

func main() {
	result := get()
	fmt.Println(result)
}


