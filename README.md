# file-go-openai

![](https://img.shields.io/github/stars/michalswi/file-go-openai)
![](https://img.shields.io/github/issues/michalswi/file-go-openai)
![](https://img.shields.io/github/forks/michalswi/file-go-openai)
![](https://img.shields.io/github/last-commit/michalswi/file-go-openai)
![](https://img.shields.io/github/release/michalswi/file-go-openai)

OpenAI in version **GPT-4.0** .

You need [OpenAI API key](https://platform.openai.com/api-keys) .

```
export API_KEYS=<>

./file-go-openai -h
Options:
  -f, --file <path>/<file>  Path to the file to be reviewed [required]
  -m, --message <string>    Message to OpenAI model [required]
  -o, --out                 Save file's review output to a file [optional]
```