package core

import (
	// "log"
	"time"

	platforms "github.com/YaleOpenLab/openx/platforms"
)

// Project defines the project investment structure in opensolar
type Project struct {
	// The project is split into two parts - parts which are used in the smart contract and parts which are not
	// we define them as critparams and noncritparams

	// Index is the project index
	Index int

	// TotalValue is the value of value of the advertised project
	TotalValue float64

	// Lock locks investments in order to wait for the recipient's confirmation
	Lock bool

	// LockPwd is the recipient's seedpwd. Will be set to null after used once
	LockPwd string

	// Chain is the blockchain the smart contract is based on
	Chain string

	// OneTimeUnlock is a one time unlock password where the recipient stores their. Set to null after single use
	OneTimeUnlock string

	// AmountOwed is the amount owed to investors.
	AmountOwed float64

	// Reputation is the positive reputation associated with a project
	Reputation float64

	// Votes is the number of votes towards a proposed contract by investors
	Votes float64

	// OwnershipShift is the percentage of the project that the recipient owns
	OwnershipShift float64

	// StageData is the data associated with stage migrations
	StageData []string

	// StageChecklist is the checklist that has to be completed before moving on to the next stage
	StageChecklist []map[string]bool

	// InvestorMap publicKey: %investment map
	InvestorMap map[string]float64

	// SeedInvestorMap publicKey: %investment map for seed invstors
	SeedInvestorMap map[string]float64

	// WaterfallMap publickey:amount map used to pay project stakeholders
	WaterfallMap map[string]float64

	// RecipientIndex is the index of the project's main recipient
	RecipientIndex int

	// OriginatorIndex is the originator's index
	OriginatorIndex int

	// GuarantorIndex is the person guaranteeing the project
	GuarantorIndex int

	// ContractorIndex is the person who proposed the contract
	ContractorIndex int

	// InvestorIndices contains the various investors who have invested
	InvestorIndices []int

	// SeedInvestorIndices contains investors who took part before the contract was at stage 3
	SeedInvestorIndices []int

	// DateInitiated contains the date when the project was created
	DateInitiated string

	// DateFunded contains the date that the project completed the stage 4-5 migration
	DateFunded string

	// DateLastPaid contains the int64 ie unix time of last payment
	DateLastPaid int64

	// AuctionType is the type of the auction the recipient has chosen (if they have)
	AuctionType string

	// InvestmentType is the type of investment - equity crowdfunding, municipal bond, normal crowdfunding, etc
	InvestmentType string

	// PaybackPeriod is the frequency in number of weeks that the recipient has to pay back the platform
	PaybackPeriod time.Duration

	// Stage is the stage at which the contract is at
	Stage int

	// InvestorAssetCode the code of the asset given to investors on investment in the project
	InvestorAssetCode string

	// DebtAssetCode is the code of the asset given to recipients on receiving a project
	DebtAssetCode string

	// PaybackAssetCode is the code of the asset given to recipients on receiving a project
	PaybackAssetCode string

	// SeedAssetCode is the code of the asset given to seed investors on seed investment in the project
	SeedAssetCode string

	// SeedInvestmentFactor is the factor that a seed investor's investment is multiplied by in case they do invest at the seed stage
	SeedInvestmentFactor float64

	// SeedInvestmentCap is the max amount that a seed investor can put in a project when it is the seed stage
	SeedInvestmentCap float64

	// EscrowPubkey is the publickey of the escrow we setup after project investment
	EscrowPubkey string

	// EscrowLock is used to lock the escrow in case someting goes wrong
	EscrowLock bool

	// MoneyRaised is total money that has been raised until now
	MoneyRaised float64

	// SeedMoneyRaised is the total seed money that has been raised until now
	SeedMoneyRaised float64

	// EstimatedAcquisition is the year by which the recipient is expected to repay the initial investment amount
	EstimatedAcquisition int

	// BalLeft is the balance left against the original investment
	BalLeft float64

	// AdminFlagged is set if someone reports the project
	AdminFlagged bool

	// FlaggedBy is the index of the admin who flagged the project
	FlaggedBy int

	// UserFlaggedBy contains the indices of users who flagged the project
	UserFlaggedBy []int

	// Reports is the total number of reports against this particular project
	Reports int

	// TellerUrl isthe url of the teller
	TellerUrl string

	// BrokerUrl isthe url of the MQTT broker
	BrokerUrl string

	// TellerPublishTopic is the topic using which the publisher / subscriber must post / subscribe messages from
	TellerPublishTopic string

	// Metadata contains other metadata and is used to derive project asset ids.
	Metadata string

	// below are all the non critical params only used on the frontend
	Name                  string    // the name of the project / the identifier by which its referred to
	City                  string    // the city in which the project is located at
	State                 string    // the state in which the project has been installed in
	Country               string    // the country in which the project has been installed in
	PanelSize             string    // size of the given panel, for diplsaying to the user who wants to bid stuff
	Image                 string    // a URL to an image of the project that will be displayed on the frontendf
	InterestRate          float64   // the rate of return for investors
	ProposedInvestmentCap float64   // the max amount that an investor can invest in when the project is in its proposed stage (stage 2)
	SelfFund              float64   // the amount that a beneficiary / recipient puts in a project without asking from other investors. This is not included as a seed investment because this would mean the recipient pays his own investment back in the project
	MainDeveloperIndex    int       // the main developer of the project
	DeveloperIndices      []int     // the indices of the developers involved in the project`
	ContractorFee         float64   // fee paid to the contractor from the total fee of the project
	OriginatorFee         float64   // fee paid to the originator included in the total value of the project
	DeveloperFee          []float64 // the fees charged by the developers
	Location              string    `json:"Location"`
	DonationType          string    `json:"Donation Type"`
	Originator            string    `json:"Originator"`
	Description           string    `json:"Description"`
	Bullet1               string    `json:"Bullet1"`
	Bullet2               string    `json:"Bullet2"`
	Bullet3               string    `json:"Bullet3"`
	Solar                 string    `json:"Solar"`
	InverterSize          string    `json:"Inverter Size"`
	Battery               string    `json:"Battery"`
	Storage               string    `json:"Storage"`
	Tariff                string    `json:"Tariff"`
	Return                string    `json:"Return"`
	Rating                string    `json:"Rating"`
	Tax                   string    `json:"Tax"`
	DailyAvgGeneration    string    `json:"Daily Average Generation"`
	Maturity              string    `json:"Maturity"`
	Acquisition           string    `json:"Acquisition"`
	AmountFunded          float64   `json:"Amount Funded"`
}

// Feedback defines a structure that is used for providing feedback
type Feedback struct {
	// Content is the content of the feedback
	Content string

	// From denotes who gave the feedback
	From Entity

	// To denotes to whom the feedback was targeted towards
	To Entity

	// Date contains the data when the feedback was given
	Date string

	// Contract is the project for which the feedback was given for.
	Contract []Project
}

// Stage is the evolution of the erstwhile static stage integer construction
type Stage struct {
	Number          int
	FriendlyName    string   // the informal name that one can use while referring to the stage
	Name            string   // this is a more formal name to give to the given stage
	Activities      []string // the activities that are covered in this particular stage and need to be fulfilled in order to move to the next stage.
	StateTrigger    []string // trigger state change from n to n+1
	BreachCondition []string // define breach conditions for a particular stage
}

// ContractAuction is an auction struct
type ContractAuction struct {
	AllContracts    []Project
	AllContractors  []Entity
	WinningContract Project
}

const (
	// InvestorWeight is the percentage weight of the project's total reputation assigned to the investor
	InvestorWeight = 0.1

	// OriginatorWeight is the percentage weight of the project's total reputation assigned to the originator
	OriginatorWeight = 0.1

	// ContractorWeight is the percentage weight of the project's total reputation assigned to the contractor
	ContractorWeight = 0.3

	// DeveloperWeight is the percentage weight of the project's total reputation assigned to the developer
	DeveloperWeight = 0.2

	// RecipientWeight is the percentage weight of the project's total reputation assigned to the recipient
	RecipientWeight = 0.3

	// NormalThreshold is the normal payback interval of 1 payback period. Regular notifications are sent regardless of whether the user has paid back towards the project.
	NormalThreshold = 1

	// AlertThreshold is the threshold above which the user gets a nice email requesting a quick payback whenever possible
	AlertThreshold = 2

	// SternAlertThreshold is the threshold above when the user gets a warning that services will be disconnected if the user doesn't payback soon.
	SternAlertThreshold = 4

	// DisconnectionThreshold is the threshold above which the user gets a notification telling that services have been disconnected.
	DisconnectionThreshold = 6
)

// InitializePlatform imports handlers from the main platform struct that are necessary for starting the platform
func InitializePlatform() error {
	return platforms.InitializePlatform()
}

// RefillPlatform checks whether the platform has any xlm and if its balance
// is less than 21 XLM, it proceeds to ask friendbot for more test xlm
func RefillPlatform(publicKey string) error {
	return platforms.RefillPlatform(publicKey)
}
