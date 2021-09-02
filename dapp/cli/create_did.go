package main

import (
	"context"
	"fmt"
	"os"

	acc "github.com/dappley/go-dappley/core/account"
	"github.com/dappley/go-dappley/logic"
	"github.com/dappley/go-dappley/util"
	"github.com/dappley/go-dappley/wallet"
)

func createDIDCommandHandler(ctx context.Context, account interface{}, flags cmdFlags) {
	dm, err := logic.GetDIDManager(wallet.GetDIDFilePath())
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	prompter := util.NewTerminalPrompter()

	didSet := acc.NewDID()
	if !acc.CheckDIDFormat(didSet.DID) {
		fmt.Println("DID formatted incorrectly.")
		return
	}

	name, err := prompter.Prompt("Enter the name to be used for the new DID document: ")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	if _, err := os.Stat(name + ".txt"); err == nil {
		fmt.Println("Error: file already exists.")
		return
	}

	didDoc := acc.CreateDIDDocument(didSet, name)
	if didDoc == nil {
		fmt.Println("Could not create file.")
		return
	}
	docFile, err := os.OpenFile(didDoc.Name+".txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Error opening file.")
		return
	}
	defer docFile.Close()
	fmt.Println("Document created. New did is", didSet.DID)
	for {
		addItem, err := prompter.PromptConfirm("Add a value to the document?")
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}
		if !addItem {
			break
		}

		newKey, err := prompter.Prompt("Enter the key for the new value: ")
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}
		if _, found := didDoc.Values[newKey]; found {
			fmt.Println("Error: key already exists in document.")
			continue
		}

		newValue, err := prompter.Prompt("Enter the value to be added: ")
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}
		didDoc.Values[newKey] = newValue
		newEntry := ",\n" + newKey + ":" + newValue
		if _, err = docFile.WriteString(newEntry); err != nil {
			fmt.Println("Error writing to file.")
			return
		}
		fmt.Println("Value added successfully.")
	}

	dm.AddDID(didSet)
	dm.SaveDIDsToFile()
	fmt.Println("Operation complete!")
}
