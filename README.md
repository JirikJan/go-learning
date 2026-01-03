Tento repozitář obsahuje ukázku příkladů AI aplikace s využitím jazyka Go. Repozitář obsahuje jednotlivé složky s ukázkou základních aplikačních patternů, které lze dále rozvíjet.

Projekt se zaměřuje na:

- volání LLM (OpenAI) a streaming textu;
- použijí toolů (Tavily Search a Fire Crawl pro vyhledávání + scraping webů);
- vytvoření a použití custom toolu;
- napojení toolu pomocí MCP (Model Context Protocol);
- paměť (persistenci - krátkodobou, dlouhodobou (Postgres), vektorovou pro poskytnutí násobného znalostního kontextu).


Tento projekt lze použít jako startovací set pro postavení AI agentní aplikace. Jeho ambicí je ukázat základní koncepty AI aplikace v jazyce Go.

## Spuštění aplikace

### Požadavky
- Go 1.21 nebo novější
- .env soubor s API klíči

### Instalace a spuštění
```bash
# Naklonování repozitáře
git clone https://github.com/JirikJan/go-learning.git
cd Go-first

# Instalace závislostí
go mod download

# Spuštění aplikace
go run main.go
```

### Konfigurace
Před spuštěním vytvořte .env soubor s potřebnými API klíči podle vzoru v projektu.


Pro úspěšné spuštění nástrojů vyhledávání na internetu a scrapingu webovek je třeba si zajistit API klíče od [Tavily Search](https://www.tavily.com) a [Fire Crawl](https://www.firecrawl.dev)
