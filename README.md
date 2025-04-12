# AAGUID Information Generator

This repository contains a command-line utility that fetches the [FIDO Alliance Metadata Service (MDS3)](https://mds.fidoalliance.org/) and [Passkey Provider AAGUIDs](https://github.com/passkeydeveloper/passkey-authenticator-aaguids) and generates a static Go map of **AAGUID → Entry** objects. These objects capture various authenticator metadata entries as defined by the [FIDO specifications](https://fidoalliance.org/specifications/).

## Overview

- **`main.go`** — The generator tool that:
    1. Downloads the MDS3 JWT and Passkey Provider AAGUIDs
    2. Verifies the JWT signature (using x5c cert chain) for MDS3
    3. Extracts the JSON payload and unmarshals it
    4. Builds a static map (`map[string]Entry`)
    5. Writes two files (`types.go`, `metadata.go`) under user provided location. By default `internal/aaguids/`.

- **`internal/aaguids/types.go`** — Contains the Go types for describing authenticator metadata, enumerations, and status objects.
- **`internal/aaguids/metadata.go`** — Contains the `metadata` map literal of **AAGUID → Entry**, generated automatically by the tool. Also includes helper functions (`GetEntry`) to retrieve metadata for a particular AAGUID.

## Installation

```bash
go install github.com/sky93/aaguid-information-generator@latest
```

This installs the `aaguid-information-generator` command into your Go toolchain.

## Usage

You can simply run the command below in your project root directory:

```bash
aaguid-information-generator -o=internal/
```

And you will have:

- `internal/aaguids/types.go` (if not already present)
- `internal/aaguids/metadata.go` updated with the latest data from MDS3

Then you can use it like below:

```go
data, exists := aaguids.GetEntry("AUTHENTICATOR_AAGUID")
```

## Security Considerations

1. **MDS Trust**  
   The FIDO MDS root of trust is the FIDO Alliance certificate authority. By default, this tool verifies the JWT’s x5c chain using system trust. You can add custom logic if you have stricter pinning requirements.

2. **Updates**  
   FIDO MDS is updated over time. (Typically once per month.) By running `aaguid-information-generator` again, you ensure you have the newest data. Check `BLOBPayload.NextUpdate` in the code if you want an automatic refresh schedule.

3. **AAID vs AAGUID**  
   This generator focuses on FIDO2 AAGUIDs. UAF-based entries with `AAID` only (no AAGUID) are generally skipped. If you need to handle UAF or U2F certificate key identifiers, you can customize the code.

## License

[MIT License](LICENSE)

---

**Enjoy generating and embedding up-to-date authenticator metadata for your FIDO2/WebAuthn ecosystem!**