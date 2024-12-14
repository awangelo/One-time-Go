# One-Time Pad Encryption Tool
Allows you to encrypt and decrypt files using the One-Time Pad encryption method.

## How It Works
The program reads an input file and determines whether to encrypt or decrypt based on the existence of a pad file:

- Encryption: If the pad file does not exist, the program encrypts the input file and generates a new pad file.
- Decryption: If the pad file exists, the program decrypts the input file using the existing pad file.
The encryption and decryption processes use the Vernam cipher (XOR operation) to combine the input data with the pad.

## Usage
Build the binary
```bash
go build -o otp
```

### Encrypting a File
To encrypt a file, run the program with the `-input` flag specifying the input file. Optionally, you can specify output and pad file names.

```bash
./otp -input <input_file> [-output <output_file>] [-pad <pad_file>]
```

- `<input_file>`: Path to the file you want to encrypt.
- `<output_file>`: (Optional) Name of the encrypted output file (default is result.out).
- `<pad_file>`: (Optional) Name of the pad file to generate (default is `pad.out`).

#### Example:

```bash
./otp -i message.txt
```

This command will encrypt `message.txt`, output the encrypted data to `result.out`, and create a pad file named `pad.out`.

### Decrypting a File
To decrypt a file, ensure the pad file used during encryption is present.

```bash
./otp -input <encrypted_file> -pad <pad_file> [-output <output_file>]
```

- `<encrypted_file>`: Path to the encrypted file.
- `<pad_file>`: Path to the pad file used during encryption.
- `<output_file>`: (Optional) Name of the decrypted output file (default is `result.out`).

#### Example:

```bash
./otp -i result.out -p pad.out -o decrypted.txt
```

This command will decrypt `result.out` using `pad.out` and save the decrypted data to `decrypted.txt`.

## Flags
- `-input` or `-i`: Specifies the input file (required).
- `-output` or `-o`: Specifies the output file name.
- `-pad` or `-p`: Specifies the pad file name (required for decryption).