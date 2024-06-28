# file-go-openai

![](https://img.shields.io/github/stars/michalswi/file-go-openai)
![](https://img.shields.io/github/issues/michalswi/file-go-openai)
![](https://img.shields.io/github/forks/michalswi/file-go-openai)
![](https://img.shields.io/github/last-commit/michalswi/file-go-openai)
![](https://img.shields.io/github/release/michalswi/file-go-openai)

OpenAI model version used **gpt-3.5-turbo**

You need [OpenAI API key](https://platform.openai.com/api-keys) .

```
export API_KEY=<>

./file-go-openai -h
Options:
  -f, --file <path>/<file>  Path to the file to be reviewed [required]
  -m, --message <string>    Message to OpenAI model [required OR use '-p']
  -p, --pattern <string>    Pattern name [required OR use '-m']
  -o, --out                 Save file's review output to a file [optional]
  -v, --version             Display OpenAI model version
```

### **IMPORTANT**  

For example for OpenAI in version **GPT-4.0** .

If you encounter such error, it's because there are some API limitations.
```
Request too large for gpt-4 in organization <org> on tokens per min (TPM): Limit 10000, Requested 43034. The input or output tokens must be reduced in order to run successfully. Visit https://platform.openai.com/account/rate-limits to learn more.
```
More about **Rate limits** for **tier-1** you can find [here](https://platform.openai.com/docs/guides/rate-limits/usage-tiers?context=tier-one). In **tier-1** for GPT-4, TPM is 10,000. You might be using different tier than tier-1 e.g. free, tier-2 etc. where TPM values are different.

### \# pattern's list

[analyze_requests_init](./patterns/analyze_requests_init/README.md)


### \# example usage

#### > analyze and display review
```
./file-go-openai \
-f /tmp/input.log \
-m "please list all uniq Request lines in one section and uniq User-Agent lines in separate section."
```

#### > analyze and save review to a file
```
./file-go-openai \
-f /tmp/input.log \
-m "please list all uniq Request lines in one section and uniq User-Agent lines in separate section." \
-o
```

#### > analyze based on defined pattern and save review to a file

Patterns can be find [here](./patterns/) .

```
./file-go-openai \
-f /tmp/input.log \
-p analyze_requests_init \
-o
```