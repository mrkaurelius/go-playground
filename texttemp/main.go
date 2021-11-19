package main

import (
	"encoding/json"
	"fmt"
)

// TODO ethereum ile ilgili utiller importt edilebilri

type CliqueGenesis struct {
	CliqueGenesisConfig CliqueGenesisConfig `json:"config"`

	Nonce      string `json:"nonce"`
	Timestamp  string `json:"timestamp"`
	Extradata  string `json:"extraData"`
	GasLimit   string `json:"gasLimit"`
	Difficulty string `json:"difficulty"`
	MixHash    string `json:"mixHash"`
	Coinbase   string `json:"coinbase"`

	// Alloc icin struct tanimlamamak suan daha kolay
	Alloc map[string]map[string]string `json:"alloc"`

	Number  string `json:"number"`
	GasUsed string `json:"gasUsed"`
	// TODO burada kaldim
}

type CliqueGenesisConfig struct {
	HomesteadBlock      int    `json:"homesteadBlock"`
	Eip150Block         int    `json:"eip150Block"`
	Eip150Hash          string `json:"eip150Hash"`
	Eip155Block         int    `json:"eip155Block"`
	Eip158Block         int    `json:"eip158Block"`
	ByzantiumBlock      int    `json:"byzantiumBlock"`
	ConstantinopleBlock int    `json:"constantinopleBlock"`
	PetersburgBlock     int    `json:"petersburgBlock"`
	IstanbulBlock       int    `json:"istanbulBlock"`

	CliqueConfig CliqueConfig `json:"clique"`
}

type CliqueConfig struct {
	Period int `json:"period"`
	Epoch  int `json:"epoch"`
}

func main() {
	fmt.Println("deneme")

	alloc := make(map[string]map[string]string, 0)
	alloc["addr1"] = make(map[string]string, 0)
	alloc["addr1"]["balance"] = "1234"
	fmt.Println(alloc)

	cliqueGenesisConfig := CliqueGenesisConfig{}
	cliqueGenesis := CliqueGenesis{Alloc: alloc, CliqueGenesisConfig: cliqueGenesisConfig}

	jsonStr, _ := json.MarshalIndent(cliqueGenesis, "", "  ")
	fmt.Println(string(jsonStr))
}
