package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/picatz/hunter"
	"github.com/spf13/cobra"
)

func main() {
	// handle CTRL+C quit
	cleanup := func() {
		os.Exit(0)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			cleanup()
		}
	}()

	client := hunter.New(hunter.UseDefaultEnvVariable, hunter.UseDefaultHTTPClient)

	var cmdAccount = &cobra.Command{
		Use:   "account",
		Short: "Get information regarding your hunter.io account",
		Long:  "ACCOUNT\nDocumentation Taken From: https://hunter.io/api/v2/docs#account \n\nEnables you to get information regarding your Hunter account at any time. This API call is free.\n\n",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			result, err := client.Account()
			if err != nil {
				panic(err)
			}
			json, err := json.Marshal(result)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(json))
		},
	}

	var (
		cmdSearchDomainFlag     string
		cmdSearchCompanyFlag    string
		cmdSearchLimitFlag      string
		cmdSearchOffsetFlag     string
		cmdSearchTypeFlag       string
		cmdSearchSeniorityFlag  string
		cmdSearchDepartmentFlag string
	)

	var cmdSearch = &cobra.Command{
		Use:   "search",
		Short: "Search all the email addresses corresponding to one website or company",
		Long:  "SEARCH\nDocumentation Taken From: https://hunter.io/api/v2/docs#domain-search \n\nSearch all the email addresses corresponding to one website or compan.\n\nEach response will return up to 100 emails. Use the `--offset` flag to get all of them. A new query is counted for calls returning at least one result.\n\nThe number of sources is limited to 20 for each email address. The `extracted_on` attribute of a source contains the date it was found for the first time, whereas the `last_seen_on` attribute contains the date it was found for the last time.\n\ntype returns the value `personal` or `generic`. A `generic` email address is a role-based email address, like contact@hunter.io. On the contrary, a `personal` email address is the address of someone in the company.\n\n`confidence` is our estimation of the probability the email address returned is correct. It depends on several criteria such as the number and quality of sources.\n\nNote that this API call is rate limited to 15 requests per second.\n\n* You must send at least the domain name or the company name. You can also send both.\n\n",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			params := hunter.Params{
				"domain":     cmdSearchDomainFlag,
				"company":    cmdSearchCompanyFlag,
				"limit":      cmdSearchLimitFlag,
				"offset":     cmdSearchOffsetFlag,
				"type":       cmdSearchTypeFlag,
				"seniority":  cmdSearchSeniorityFlag,
				"department": cmdSearchDepartmentFlag,
			}
			if params["domain"] == "" && params["company"] == "" {
				fmt.Println("missing either the `--domain` or `--company` flag")
				os.Exit(1)
			}
			result, err := client.DomainSearch(params)
			if err != nil {
				panic(err)
			}
			json, err := json.Marshal(result)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(json))
		},
	}

	cmdSearch.Flags().StringVar(&cmdSearchDomainFlag, "domain", "", "Domain name from which you want to find the email addresses. For example, `stripe.com`.")
	cmdSearch.Flags().StringVar(&cmdSearchCompanyFlag, "company", "", "The company name from which you want to find the email addresses. For example, `stripe`. Note that you'll get better results by supplying the domain name as we won't have to find it. If you send a request with both the domain and the company name, we'll use the domain name. It doesn't need to be in lowercase.")
	cmdSearch.Flags().StringVar(&cmdSearchLimitFlag, "limit", "10", "Specifies the max number of email addresses to return.")
	cmdSearch.Flags().StringVar(&cmdSearchOffsetFlag, "offset", "0", "Specifies the number of email addresses to skip.")
	cmdSearch.Flags().StringVar(&cmdSearchSeniorityFlag, "seniority", "", "Get only email addresses for people with the selected seniority level. The possible values are junior, senior or executive. Several seniority levels can be selected (delimited by a comma).")
	cmdSearch.Flags().StringVar(&cmdSearchDepartmentFlag, "department", "", "Get only email addresses for people working in the selected department(s). The possible values are `executive`, `it`, `finance`, `management`, `sales`, `legal`, `support`, `hr`, `marketing` or `communication`. Several departments can be selected (comma-delimited).")

	var (
		cmdFindDomainFlag    string
		cmdFindCompanyFlag   string
		cmdFindFirstNameFlag string
		cmdFindLastNameFlag  string
		cmdFindFullNameFlag  string
	)

	var cmdFind = &cobra.Command{
		Use:   "find",
		Short: "Generates or retrieves the most likely email address from a domain name, a first name and a last name",
		Long:  "FIND\nDocumentation Taken From: https://hunter.io/api/v2/docs#email-finder \n\nGenerates or retrieves the most likely email address from a domain name, a first name and a last name.\n\nThe score returned is an estimation of the probability the email generated is correct.\n\nIf we have found the retrieved email address somewhere on the web, we display the sources here. The number of sources is limited to 20. The extracted_on attribute contains the date it was found for the first time, whereas the last_seen_on attribute contains the date it was found for the last time.\n\n* You must send at least the domain name or the company name. You can also send both.\n\n** You must send at least the first name and the last name or the full name.\nThe score returned is an estimation of the probability the email generated is correct.\n\nIf we have found the retrieved email address somewhere on the web, we display the sources here. The number of sources is limited to 20. The `extracted_on attribute` contains the date it was found for the first time, whereas the `last_seen_on` attribute contains the date it was found for the last time.\n\n",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			params := hunter.Params{
				"domain":     cmdFindDomainFlag,
				"company":    cmdFindCompanyFlag,
				"first_name": cmdFindFirstNameFlag,
				"last_name":  cmdFindLastNameFlag,
				"full_name":  cmdFindFullNameFlag,
			}
			if params["domain"] == "" && params["company"] == "" {
				fmt.Println("missing either the `--domain` or `--company` flag")
				os.Exit(1)
			}
			if params["first_name"] == "" || params["last_name"] == "" {
				if params["full_name"] == "" {
					fmt.Println("missing either the `--first-name` AND `--last-name` flags OR the `--full-name` flag")
					os.Exit(1)
				}
			}
			result, err := client.FindEmail(params)
			if err != nil {
				panic(err)
			}
			json, err := json.Marshal(result)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(json))
		},
	}

	cmdFind.Flags().StringVar(&cmdFindDomainFlag, "domain", "", "Domain name of the company, used for emails.")
	cmdFind.Flags().StringVar(&cmdFindCompanyFlag, "company", "", "The company name from which you want to find the email addresses.")
	cmdFind.Flags().StringVar(&cmdFindFirstNameFlag, "first-name", "", "The person's first name. It doesn't need to be in lowercase.")
	cmdFind.Flags().StringVar(&cmdFindLastNameFlag, "last-name", "", "The person's last name. It doesn't need to be in lowercase.")
	cmdFind.Flags().StringVar(&cmdFindFullNameFlag, "full-name", "", "The person's full name. Note that you'll get better results by supplying the person's first and last name if you can. It doesn't need to be in lowercase.")

	var (
		cmdVerifyEmailFlag string
	)

	var cmdVerify = &cobra.Command{
		Use:   "verify",
		Short: "Allows you to verify the deliverability of an email address",
		Long:  "VERIFY\nDocumentation Taken From: https://hunter.io/api/v2/docs#email-verifier \n\nHunter focuses on B2B. Therefore, webmails are not verified. We'll run every check but won't reach the remote SMTP server.\n\nThis endpoint is rate-limited by domain name. You can check up to 200 email addresses for a domain name every 24 hours. You can check the number of requests remaining using the X-RateLimit-Remaining header.\n\nThe request will run for 20 seconds. If it was not able to provide a response in time, we will return a 202 status code. You will then be able to poll the same endpoint to get the verification's result. Of course, all the requests in this case are counted only once.\n\n\nReading Results:\n`score` is the deliverability score we give to the email address.\n`regexp` is true if the email address passes our regular expression.\n`gibberish` is true if we find this is an automatically generated email address (for example `e65rc109q@company.com`).\n`disposable` is true if we find this is an email address from a disposable email service.\n`webmail` is true if we find this is an email from a webmail (for example Gmail).\n`mx_records` is true if we find MX records exist on the domain of the given email address.\n`smtp_server` is true if we connect to the SMTP server successfully.\n`smtp_check` is true if the email address doesn't bounce.\n`accept_all` is true if the SMTP server accepts all the email addresses. It means you can have have false positives on SMTP checks.\n`block` is true if the SMTP server prevented us to perform the STMP check.\n`sources` If we have found the given email address somewhere on the web, we display the sources here. The number of sources is limited to 20.\n`extracted_on` contains the date it was found for the first time.\n`last_seen_on` contains the date it was found for the last time.\n\n",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			params := hunter.Params{
				"email": cmdVerifyEmailFlag,
			}
			if params["email"] == "" {
				fmt.Println("missing the `--email` flag")
				os.Exit(1)
			}
			result, err := client.VerifyEmail(params)
			if err != nil {
				panic(err)
			}
			json, err := json.Marshal(result)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(json))
		},
	}

	cmdVerify.Flags().StringVar(&cmdVerifyEmailFlag, "email", "", "The email address you want to verify.")

	var rootCmd = &cobra.Command{Use: "hunter"}
	rootCmd.AddCommand(cmdAccount)
	rootCmd.AddCommand(cmdSearch)
	rootCmd.AddCommand(cmdFind)
	rootCmd.AddCommand(cmdVerify)
	rootCmd.Execute()
}
