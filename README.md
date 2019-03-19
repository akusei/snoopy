# Readme coming soon

## Initialization
The sensitive files will need to be decrypted before compiling.
Run the following command to decrypt the required files.

Note: both the encrypt and decrypt functions can use an environment
variable (SNOOPY_PASSWORD) for the encryption/decryption password

```
docker-compose run --rm snoopy decrypt
```

## Building
This project includes building helper functionality. To build
windows, linux and mac variations in 64 and 32 bit just run

```
docker-compose run --rm snoopy build
```

## Encrypting Files
Certain files will need to be encrypted before pushing to the repo.
Run the following command to encrypt the files:

```
docker-compose run --rm snoopy encrypt
```