# Yet another gotmpl processor

Flags: <br>
`--template` for specifying a template file, defaults to _./template.gotmpl_<br>
`--values` for specifying a file, containing keys and values for the above template, defaults to _./values.yaml_
  
### howto
```shell script
git get https://github.com/sigurdblueface/gibddAppeal
go build gibddappeal
```
Built binary is able to merge any template-values pair, In the attached examples it serves composing appeals to russian traffic police department.