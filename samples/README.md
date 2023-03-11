# Samples

## Description

This directory contains sample files used by the unit tests implemented in this
library. Some files may contain private keys that are not used in production
environment and thus, do not represent any real security risks.

## The sample files

### sample.pfx

The sample.pfx file is a PKCS #12 with the password `password` created with the
command:

```
$ openssl pkcs12 -export -out sample.pfx -inkey key.pem -in cert.pem -keysig -legacy
```

It is very important to notice that, in order to be compatible with MS Windows PFX
format, the flag `-legacy` must be used.
