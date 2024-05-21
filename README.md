# file-go-openai

![](https://img.shields.io/github/stars/michalswi/file-go-openai)
![](https://img.shields.io/github/issues/michalswi/file-go-openai)
![](https://img.shields.io/github/forks/michalswi/file-go-openai)
![](https://img.shields.io/github/last-commit/michalswi/file-go-openai)
![](https://img.shields.io/github/release/michalswi/file-go-openai)

OpenAI in version **GPT-4.0** .

You need [OpenAI API key](https://platform.openai.com/api-keys) .

```
export API_KEY=<>

./file-go-openai -h
Options:
  -f, --file <path>/<file>  Path to the file to be reviewed [required]
  -m, --message <string>    Message to OpenAI model [required]
  -o, --out                 Save file's review output to a file [optional]
```

**IMPORTANT**  

If you encounter such error, it's because there are some limitations related to the number of lines that file can have. It's approximately **600** lines
```
Request too large for gpt-4 in organization <org> on tokens per min (TPM): Limit 10000, Requested 43034. The input or output tokens must be reduced in order to run successfully. Visit https://platform.openai.com/account/rate-limits to learn more.
```
you have to split your file to have ~**600** lines. More about **Rate limits** you can find [here](https://platform.openai.com/docs/guides/rate-limits?context=tier-free) .
