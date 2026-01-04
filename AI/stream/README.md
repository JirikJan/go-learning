Stream tokenů, resp textů je klíčový pro relevantní UX v moderních aplikacích. 

Streamování textu je v Go zajištěno pomocí callback funnkce:

```go
_, err = llm.GenerateContent(ctx, messages,
	llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))
		return nil
	}),
)
if err != nil {
	log.Fatal(err)
}
```

```go
_, err = llm.GenerateContent
.
.
.
```
Promítneme zde jen error, neboť o "completions" nemáme zde zájem, a proto je použita skrytá hodnota pomocí "_"