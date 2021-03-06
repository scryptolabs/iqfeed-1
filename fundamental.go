package iqfeed

import (
	"strings"
	"time"
)

// FundamentalMsg cannot be customized and is used to provide detail of a particular matched symbol.
type FundamentalMsg struct {
	Symbol             string    // The Symbol ID to match with watch request
	ExchaangeID        string    // Deprecated Use Listed Market (field 45 below) instead
	PE                 float64   // Price/earnings ratio
	AvgVolume          int       // Average daily volume (4 week avg)
	Fifty2WkHigh       float64   // Highest price of the last 52 weeks for futures, this is the contract High.
	Fifty2WkLow        float64   // Lowest price of the last 52 weeks. For futures, this is the contract Low.
	CalYearHigh        float64   // High price for the current calendar year.
	CalyearLow         float64   // Low price for the current calendar year.
	DivYield           float64   // The annual dividends per share paid by the company divided by the current market price per share of stock sent as a percentage.
	DivAmt             float64   // The current quarter actual dividend
	DivRate            float64   // The annualized amount at which a dividend is expected to be paid by a company.
	PayDate            time.Time // Date on which a company made its last dividend payment (MM/DD/YYYY).
	ExDivDate          time.Time // The actual date in which a stock goes ex-dividend, typically about 3 weeks before the dividend is paid to shareholders of record. Also the amount of the dividend is reflected in a reduction of the share price on this date. (MM/DD/YYYY).
	Reserved1          string    // Reserved field.
	Reserved2          string    // Reserved field.
	Reserved3          string    // Reserved field.
	ShortInterest      int       // The total number of shares of a security that have been sold short by customers and securities firms that have not been repurchased to settle outstanding short positions in the market.
	Reserved4          string    // Reserved field.
	CurrentYrEPS       float64   // The portion of a company's profit allocated to each outstanding share of common stock.
	NextYrEPS          float64   // The total amount of earnings per share a company is estimated to accumulate over the next four quarters of the current fiscal year.
	FiveYrGrowthPct    float64   // Earnings Per Share growth rate over a five year period.
	FiscalYrEnd        int       // The two digit month that the fiscal year ends for a company.
	Reserved5          string    // Reserved field.
	CompanyName        string    // Company name or contract description
	RootOptionSymbol   []string  // A list of root option symbols, there may be more than one.
	PctHeldByInst      float64   // A percentage of outstanding shares held by banks and institutions.
	Beta               float64   // A coefficient measuring a stock’s relative volatility. It is the covariance of a stock in relation to the rest of the stock market. 30 day historical volatility.
	Leaps              string    // Long term equity anticipation securities.
	CurrentAssets      float64   // The amount of total current assets held by a company as of a specific date in Millions (lastADate).
	CurrentLiabilities float64   // The amount of total current liabilities held by a company as of a specific date in Millions (lastADate).
	BalSheetDate       time.Time // Last date that a company issued their quarterly report. (MM/DD/YYYY).
	LongTermDebt       float64   // The amount of long term debt held by a company as of a specific date in Millions(lastADate).
	ComShrOutstanding  float64   // The amount of common shares outstanding.
	Reserved6          string    // Reserved field.
	SplitFactor1       string    // A float a space, then MM/DD/YYYY
	SplitFactor2       string    // A float a space, then MM/DD/YYYY
	Reserved7          string    // Reserved field.
	Reserved8          string    // Reserved field.
	FormatCode         string    // Display format code, See: Price Format Codes http://www.iqfeed.net/dev/api/docs/PriceFormatCodes.cfm.
	Precision          int       // Number of decimal digits.
	SIC                int       // Federally designed numbering system identifying companies by industry. This 4 digit number corresponds to a specific industry.
	HistVolatility     float64   // 30-trading day volatility that it is calculated using Black-Scholes (https://en.wikipedia.org/wiki/Black%E2%80%93Scholes_model).
	SecurityType       string    // The security type code, See: Security Types (http://www.iqfeed.net/dev/api/docs/SecurityTypes.cfm).
	ListedMarket       string    // The listing market ID, See: Listed Markets
	Fifty2WkHighDate   time.Time // The date of the lowest price of the last 52 weeks. For futures, this is the contract Low Date. (MM/DD/YYYY)
	Fifty2WkLowDate    time.Time // The date of the lowest price of the last 52 weeks. For futures, this is the contract Low Date. (MM/DD/YYYY)
	CalYearHighDate    time.Time // Date at which the High price for the current calendar year occurred. (MM/DD/YYYY)
	CalYearLowDate     time.Time // Date at which the Low price for the current calendar year occurred. (MM/DD/YYYY)
	YrEndClose         float64   // Price of Year End Close. (Equities Only)
	MaturityDate       time.Time // Date of maturity for a bond.
	CouponRate         float64   // Interest rate for a bond.
	ExpirationDate     time.Time // IEOptions, Futures, FutureOptions, and SSFutures only
	StrikePrice        float64   // IEOptions only
	NAICS              int       // North American Industry Classification System (http://www.census.gov/eos/www/naics/)
	ExchangeRoot       string    // The root symbol that you can find this symbol listed under at the exchange.
}

// UnMarshall sends the data into the usable struct for consumption by the application.
func (f *FundamentalMsg) UnMarshall(d []byte, loc *time.Location) {
	items := strings.Split(string(d), ",")
	f.Symbol = items[0]                                  // APL,
	f.ExchaangeID = items[1]                             // 5,
	f.PE = GetFloatFromStr(items[2])                     // 9.9,
	f.AvgVolume = GetIntFromStr(items[3])                // 53599000,
	f.Fifty2WkHigh = GetFloatFromStr(items[4])           // 134.5400,
	f.Fifty2WkLow = GetFloatFromStr(items[5])            // 92.0000,
	f.CalYearHigh = GetFloatFromStr(items[6])            // 105.8500,
	f.CalyearLow = GetFloatFromStr(items[7])             // 92.3900,
	f.DivYield = GetFloatFromStr(items[8])               // 2.2100,
	f.DivAmt = GetFloatFromStr(items[9])                 // 0.5200,
	f.DivRate = GetFloatFromStr(items[10])               // 2.0800,
	f.PayDate = GetDateMMDDCCYY(items[11], loc)          // 02/11/2016,
	f.ExDivDate = GetDateMMDDCCYY(items[12], loc)        // 02/04/2016,
	f.Reserved1 = items[13]                              // ,
	f.Reserved2 = items[14]                              // ,
	f.Reserved3 = items[15]                              // ,
	f.ShortInterest = GetIntFromStr(items[16])           // 63543520,
	f.Reserved4 = items[17]                              // ,
	f.CurrentYrEPS = GetFloatFromStr(items[18])          // 9.46,
	f.NextYrEPS = GetFloatFromStr(items[19])             // ,
	f.FiveYrGrowthPct = GetFloatFromStr(items[20])       // 0.34,
	f.FiscalYrEnd = GetIntFromStr(items[21])             // 09,
	f.Reserved5 = items[22]                              // ,
	f.CompanyName = items[23]                            // APPLE,
	f.RootOptionSymbol = strings.Split(items[24], " ")   // AAPL AAPL7,
	f.PctHeldByInst = GetFloatFromStr(items[25])         // 67.1,
	f.Beta = GetFloatFromStr(items[26])                  // 1.35,
	f.Leaps = items[27]                                  // ,
	f.CurrentAssets = GetFloatFromStr(items[28])         // 89378.0,
	f.CurrentLiabilities = GetFloatFromStr(items[29])    // 80610.0,
	f.BalSheetDate = GetDateMMDDCCYY(items[30], loc)     // 12/31/2015,
	f.LongTermDebt = GetFloatFromStr(items[31])          // 53463.0,
	f.ComShrOutstanding = GetFloatFromStr(items[32])     // 5544583,
	f.Reserved6 = items[33]                              // 334220,
	f.SplitFactor1 = items[34]                           // 0.14 06/09/2014,
	f.SplitFactor2 = items[35]                           // 0.50 02/28/2005,
	f.Reserved7 = items[36]                              // ,
	f.Reserved8 = items[37]                              // 0,
	f.FormatCode = items[38]                             // 14,
	f.Precision = GetIntFromStr(items[39])               // 4,
	f.SIC = GetIntFromStr(items[40])                     // 3571,
	f.HistVolatility = GetFloatFromStr(items[41])        // 36.98,
	f.SecurityType = items[42]                           // 1,
	f.ListedMarket = items[43]                           // 21,
	f.Fifty2WkHighDate = GetDateMMDDCCYY(items[44], loc) // 04/28/2015,
	f.Fifty2WkLowDate = GetDateMMDDCCYY(items[45], loc)  // 08/24/2015,
	f.CalYearHighDate = GetDateMMDDCCYY(items[46], loc)  // 01/05/2016,
	f.CalYearLowDate = GetDateMMDDCCYY(items[47], loc)   // 01/28/2016,
	f.YrEndClose = GetFloatFromStr(items[48])            // 105.26,
	f.MaturityDate = GetDateMMDDCCYY(items[49], loc)     // ,
	f.CouponRate = GetFloatFromStr(items[50])            // ,
	f.ExpirationDate = GetDateMMDDCCYY(items[51], loc)   // ,
	f.StrikePrice = GetFloatFromStr(items[52])           // ,
	f.NAICS = GetIntFromStr(items[53])                   // 334220,
	f.ExchangeRoot = items[54]                           // ,
}
