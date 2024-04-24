package config

const (
	BasicAuth    AuthType = "basic"
	URLParamAuth AuthType = "urlParam"
)

type Config struct {
	M3uURL         string
	Auth           *AuthenticationInformation
	OutputFilePath string
	FullMatch      MatchCriteria `yaml:"fullmatch"`
	SoftMatch      MatchCriteria `yaml:"softmatch"`
}

type AuthenticationInformation struct {
	AuthType AuthType
	User     string
	Password string
}

type AuthType string

type MatchCriteria struct {
	Group []string `yaml:"group"` // group-title / category
	Name  []string `yaml:"name"`  // tvg-name
}
