# go-solid
Clean Architecture and SOLID


Install golangci-lint
```
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

Tambahkan Konfigurasi golangci-lint: Buat file .golangci.yml di root proyek Anda untuk mengonfigurasi golangci-lint.

```
linters-settings:
  errcheck:
    check-type-assertions: true
    check-blank: true

issues:
  exclude-use-default: false
  exclude:
    - "should have comment or be unexported"
  exclude-dirs:
    - vendor
```

jalankan golangci-lint
```
golangci-lint run
```

Contoh case yang ke-detect lewat linter

Sebelum 
```
func LoadConfig() (*Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
```
Setelah 
```
func LoadConfig() (*Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Failed to close file: %v", err)
		}
	}()

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
```
