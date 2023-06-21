package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"io"
	"log"
	"os"
	"strings"
)

func getFileByPath(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file, path: %s, err: %s", path, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file, path: %s, err: %s", path, err)
		}
	}(file)

	return file, nil
}

func GetContractABIByFile(path string) (*abi.ABI, error) {
	bytes, err := GetFileBytes(path)
	if err != nil {
		return nil, err
	}

	contractABI, err := abi.JSON(strings.NewReader(string(bytes)))
	if err != nil {
		log.Fatalf("failed to parse abi, path: %s, err: %s", path, err)
	}

	return &contractABI, nil
}

func GetFileBytes(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to read file, path: %s, err: %s", path, err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file, path: %s, err: %s", path, err)
		}
	}(file)

	return io.ReadAll(file)
}
