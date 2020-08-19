# hugo-ts

`hugo-ts` updates the `date:` line in YAML front matter with the current timestamp. This is useful as a replacement for the old `hugo undraft` command, which used to do this for you.

**CAUTION**: This tool uses a simple regular expression and is thus fairly fragile. It does its edit in place, so it has the potential to destroy data. Please make sure your content is safely committed to version control (or otherwise backed up) before running this command.

## Compiling

`go build` should be sufficient. This tool has no dependencies outside of Go itself.

## Usage

```bash
$ hugo-ts <full-path>
```

You must supply the full path to the file to be edited, e.g. `content/posts/hello-world/index.md`.
