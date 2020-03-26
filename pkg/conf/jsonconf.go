package conf

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// JSONConfig Loads configuration from a set of JSON files into the trg structure.
// Files are processed in the order specified, each file overriding, or augmenting, config info stored so far.
// Files that don't exist are silently skipped.
// JSON format errors are reported to stderr before proceeding with processing.
// The ordered list of files is specified in code, but can be overridden by CSV list contained in CONF_FILES environment variable,
// which, in turn, can be overridden by CLI flag -conf.files.
// The configuration information parsed from json files can be overridden by json stored in CONF_JSON environment variable, which, in turn,
// can be overridden by json passed through -conf.json CLI flag.
type JSONConfig struct {
	Files []string
}

const defFileName = "config.json"

var (
	cliFileList    string
	cliJSON        string
	cliPrintConfig bool
	cliConfig      bool
)

func init() {
	flag.StringVar(&cliFileList, "conf.files", "", "CSV list of json configuration files to load. Overrides env var CONF_FILES.")
	flag.StringVar(&cliJSON, "conf.json", "", "JSON config information override. Overrides env var CONF_JSON.")
	flag.BoolVar(&cliPrintConfig, "conf.print", false, "Print config info after loading all sources.")
	flag.BoolVar(&cliConfig, "conf.config", false, "Print config info after loading all sources and exit.")
}

// NewJSONConfig creates an instance of ConfigLoader to load config from the specified set of config files
func NewJSONConfig(confFiles ...string) JSONConfig {
	rv := JSONConfig{Files: confFiles}
	return rv
}

// DefaultJSONConfig creates an instance of ConfigLoader to load config from the standard set of config files
func DefaultJSONConfig() JSONConfig {
	cfName := computeFileName(os.Args[0], defFileName)
	rv := JSONConfig{Files: []string{cfName}}
	return rv
}

// LoadInto loads config and its overrides from all sources into the trg structure
func (loader *JSONConfig) LoadInto(trg interface{}) {

	// get the list of json files to parse
	jsonFiles := getListOfFiles(loader.Files)

	// load config from all files that exist and are parse-able
	for _, f := range jsonFiles {
		loadFromFile(trg, f)
	}

	// override with CONF_JSON env var, if exist
	if val, wasDefined := os.LookupEnv("CONF_JSON"); wasDefined {
		loadFromBytes(trg, []byte(val))
	}

	// override with -conf.json cli flag, if specified
	if cliJSON != "" {
		loadFromBytes(trg, []byte(cliJSON))
	}

	// if specified, print effective config json
	if cliPrintConfig || cliConfig {
		fmt.Fprintf(os.Stderr, "conf: Effective Config: %s\n", ToJSONString(trg))
	}

	// if specified, exit
	if cliConfig {
		os.Exit(0)
	}
}

// ToJSONString - Marshal trg config structure to JSON string
func ToJSONString(trg interface{}) string {
	buff, _ := json.MarshalIndent(trg, "", "  ")
	return string(buff)
}

// GetDefaultFilePath returns full path of the default config file in the default location
func GetDefaultFilePath(basename string) string {
	return computeFileName(os.Args[0], basename)
}

func getListOfFiles(defaultList []string) []string {
	// check cli flags
	if !flag.Parsed() {
		flag.Parse()
	}

	if cliFileList != "" {
		return strToFileList(cliFileList)
	}

	// check env var
	if val, wasDefined := os.LookupEnv("CONF_FILES"); wasDefined {
		return strToFileList(val)
	}

	// return default
	return defaultList
}

func strToFileList(str string) []string {
	return strings.FieldsFunc(str, func(r rune) bool { return r == ',' })
}

func loadFromFile(trg interface{}, filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "conf: File does not exist, skipping: %s\n", filename)
		return
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "conf: Error reading file %s, skipping: %s\n", filename, err)
		return
	}

	fmt.Fprintf(os.Stderr, "conf: Loading json from file: %s\n", filename)
	loadFromBytes(trg, data)

}

func loadFromBytes(trg interface{}, data []byte) {
	err := json.Unmarshal(data, trg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "conf: Error unmarshalling data: %s\n", err)
	}
}

func computeFileName(exeFile, confFile string) string {
	path := filepath.Dir(exeFile)
	return filepath.Join(path, confFile)
}
