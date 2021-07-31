package confighelper

type Config struct {
	Database struct {
		Host string `yaml:"host" envconfig:"DATABASE_HOST"`
		Port string `yaml:"port" envconfig:"DATABASE_PORT"`
	} `yaml:"Database"`
	Organizations []struct {
		PAT             string `yaml:"PAT"`
		OrganizationUrl string `yaml:"OrganisationUrl"`
		Feeds           []struct {
			FeedName string   `yaml:"FeedName"`
			Packages []string `yaml:"Packages"`
		} `yaml:"Feeds"`
	} `yaml:"Organizations"`
}
