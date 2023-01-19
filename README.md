# pgen
password & passphrase generator written in go

## Usage
```shell
$ pgen
password & passphrase generator                                                                                                                                                                                                                           121.141.26.61

Usage:
  pgen [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  passphrase  Generate passphrase (Alias: pp)
  password    Generate password (Alias: pw)

Flags:
  -h, --help      help for pgen
  -v, --version   version for pgen

Use "pgen [command] --help" for more information about a command.
```
#### Example w/ default options
```shell
$ pgen pp
mucosogranular-corrected-sawdusty 
```
#### Example w/ options
```shell
$ pgen pp -n=true -s _ -c 5
scratchboard652_capybara_ophthalaiater_taxidea_naturopathy
```
### Generate passphrase
```shell
$ pgen passphrase --help
generate passphrase --help                                                                                                                                                                                                                                121.141.26.61

Usage:
  pgen passphrase [flags]

Flags:
  -h, --help               help for passphrase
  -n, --num                whether include numbers or not
  -s, --separator string   passphrase separator (default "-")
  -c, --word.count int     words count (default 3)
```
### Generate password
```shell
$ pgen password --help
generate password                                                                                                                                                                                                                                 121.141.26.61

Usage:
  pgen password [flags]

Flags:
  -h, --help              help for password
  -L, --length int        password length (default 16)
  -l, --lower             whether include lowercase or not (default true)
  -n, --num               whether include numbers or not (default true)
  -N, --num.max int       maximum numbers (default 3)
  -s, --special           whether include special characters or not (default true)
  -S, --special.max int   maximum special characters (default 3)
  -u, --upper             whether include uppercase or not (default true)
```
#### Example w/ default options
```shell
pgen pw
!Abokn^cnpGAPBca$H 
```
#### Example w/ options
```shell
$ pgen pw -n=false -l=false -s=true -S 10 -L 48
HQ#JF*QN%NZMVJ$FMXC&OBO&BVYGNIDJRST&CK@#FIIKWR*ME 
```