package main

type Config struct {
	AzureAccesToken      string `yaml:"AzureAccesToken" json:"AzureAccesToken"`
	AzureOrganizationUrl string `yaml:"AzureOrganizationUrl" json:"AzureOrganizationUrl"`
	AzureFeed            string `yaml:"AzureFeed" json:"AzureFeed"`
}
