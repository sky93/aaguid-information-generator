package aaguids

/*
Package aaguids provides Go data structures and enumerations for describing authenticator metadata
in alignment with the FIDO Alliance specifications:

 • FIDO Metadata Service (v3.0, May 18, 2021)
   http://fidoalliance.org/specs/mds/fido-metadata-service-v3.0-ps-20210518.html
 • FIDO Metadata Statement (v3.0, May 18, 2021)
   http://fidoalliance.org/specs/mds/fido-metadata-statement-v3.0-ps-20210518.html

These specifications define:
 • A format (the “Metadata BLOB”) for distributing authenticator metadata statements.
 • A normative structure for describing trusted attestation roots, authenticator capabilities,
   status information, and recommended reliance information.
 • Mechanisms for enumerating known issues or revocations, firmware updates, and other critical
   authenticator events.

The structs and constants below map directly to the data fields and enumerations described in the
FIDO Metadata specifications, providing Go-based relying-party or FIDO-server implementations with
a straightforward way to parse and consume these metadata objects.
*/

/*
AuthenticatorStatus is defined in the FIDO Metadata Service specification § 3.1.4 “AuthenticatorStatus enum”
and enumerates the possible status values an authenticator model can have. These status values help
Relying Parties evaluate risks or certifications associated with the model.
*/
type AuthenticatorStatus string

const (
	/*
	   NOT_FIDO_CERTIFIED
	   § 3.1.4.1 “Certification Related Statuses”
	   — Title: “NOT_FIDO_CERTIFIED”

	   Indicates the authenticator is not FIDO-certified.
	   Possible Metadata fields (see § 3.1.4.1):
	     • effectiveDate: when this status was achieved
	     • authenticatorVersion: minimum applicable version
	     • url: additional info or product page
	*/
	NOT_FIDO_CERTIFIED AuthenticatorStatus = "NOT_FIDO_CERTIFIED"

	/*
	   FIDO_CERTIFIED
	   § 3.1.4.1 “Certification Related Statuses”
	   — Title: “FIDO_CERTIFIED”

	   The authenticator has passed FIDO functional certification. This certification may be
	   superseded by FIDO_CERTIFIED_L1 in newer programs. Relevant fields (see § 3.1.4.1):
	     • effectiveDate: date of certification issuance
	     • authenticatorVersion: min version of the certified solution
	     • certificationDescriptor: e.g. "Munikey 7c Black Edition"
	     • certificateNumber: FIDO Alliance certificate number
	     • certificationPolicyVersion: policy version
	     • certificationRequirementsVersion: security requirements version
	     • url: link to certificate or relevant article
	*/
	FIDO_CERTIFIED AuthenticatorStatus = "FIDO_CERTIFIED"

	/*
	   USER_VERIFICATION_BYPASS
	   § 3.1.4.2 “Security Notification Statuses”
	   — Title: “USER_VERIFICATION_BYPASS”

	   Indicates malware or an exploit can bypass user verification, allowing use without user
	   knowledge or consent. Relevant fields (see § 3.1.4.2):
	     • effectiveDate: date of the incident
	     • authenticatorVersion: min affected version
	     • url: link to explanation or corporate article
	*/
	USER_VERIFICATION_BYPASS AuthenticatorStatus = "USER_VERIFICATION_BYPASS"

	/*
	   ATTESTATION_KEY_COMPROMISE
	   § 3.1.4.2 “Security Notification Statuses”
	   — Title: “ATTESTATION_KEY_COMPROMISE”

	   The authenticator’s attestation key is known to be compromised. Relying Parties should check the
	   "certificate" field to identify the compromised batch or reject new registrations if not set.
	   Relevant fields (see § 3.1.4.2):
	     • effectiveDate: date compromise occurred
	     • authenticatorVersion: min affected version
	     • certificate: base64 DER PKIX cert for compromised root
	     • url: link describing the incident
	*/
	ATTESTATION_KEY_COMPROMISE AuthenticatorStatus = "ATTESTATION_KEY_COMPROMISE"

	/*
	   USER_KEY_REMOTE_COMPROMISE
	   § 3.1.4.2 “Security Notification Statuses”
	   — Title: “USER_KEY_REMOTE_COMPROMISE”

	   Known weaknesses allow user credential keys to be compromised remotely (e.g. weak entropy,
	   side-channel vulnerabilities). Relevant fields (see § 3.1.4.2):
	     • effectiveDate: date of the reported incident
	     • authenticatorVersion: min affected version
	     • url: link describing the issue
	*/
	USER_KEY_REMOTE_COMPROMISE AuthenticatorStatus = "USER_KEY_REMOTE_COMPROMISE"

	/*
	   USER_KEY_PHYSICAL_COMPROMISE
	   § 3.1.4.2 “Security Notification Statuses”
	   — Title: “USER_KEY_PHYSICAL_COMPROMISE”

	   Attacker with physical possession can extract user keys due to weaknesses in key protection.
	   Relevant fields (see § 3.1.4.2):
	     • effectiveDate: date of the incident
	     • authenticatorVersion: min affected version
	     • url: link describing the issue
	*/
	USER_KEY_PHYSICAL_COMPROMISE AuthenticatorStatus = "USER_KEY_PHYSICAL_COMPROMISE"

	/*
	   UPDATE_AVAILABLE
	   § 3.1.4.3 “Info Statuses”
	   — Title: “UPDATE_AVAILABLE”

	   A software or firmware update is available for the authenticator. If it fixes severe issues,
	   authenticatorVersion must be incremented accordingly. Relevant fields (see § 3.1.4.3):
	     • effectiveDate: date of the update publication
	     • authenticatorVersion: new version matching metadata statement
	     • url: link for update info
	*/
	UPDATE_AVAILABLE AuthenticatorStatus = "UPDATE_AVAILABLE"

	/*
	   REVOKED
	   § 3.1.4.1 “Certification Related Statuses”
	   — Title: “REVOKED”

	   The FIDO Alliance determined this authenticator is not trustworthy (e.g. fraudulent or
	   backdoored). Relying Parties should reject its future registrations. Relevant fields:
	     • effectiveDate: date of revocation
	     • authenticatorVersion: if relevant
	     • url: link explaining the reason for revocation
	*/
	REVOKED AuthenticatorStatus = "REVOKED"

	/*
	   SELF_ASSERTION_SUBMITTED
	   § 3.1.4.1 “Certification Related Statuses”
	   — Title: “SELF_ASSERTION_SUBMITTED”

	   The vendor completed and submitted the self-certification checklist to the FIDO Alliance. If
	   publicly available, "url" contains the link. Fields:
	     • effectiveDate: submission date
	     • authenticatorVersion: if relevant
	*/
	SELF_ASSERTION_SUBMITTED AuthenticatorStatus = "SELF_ASSERTION_SUBMITTED"

	/*
	   FIDO_CERTIFIED_L1
	   § 3.1.4.1 “Certification Related Statuses”
	   — Title: “FIDO_CERTIFIED_L1”

	   Authenticator passed FIDO Certification at level 1, a more strict successor to FIDO_CERTIFIED.
	*/
	FIDO_CERTIFIED_L1 AuthenticatorStatus = "FIDO_CERTIFIED_L1"

	/*
	   FIDO_CERTIFIED_L1plus
	   § 3.1.4.1 “Certification Related Statuses”
	   — Title: “FIDO_CERTIFIED_L1plus”

	   Authenticator passed FIDO Certification at level 1+, stricter than level 1.
	*/
	FIDO_CERTIFIED_L1plus AuthenticatorStatus = "FIDO_CERTIFIED_L1plus"

	/*
	   FIDO_CERTIFIED_L2
	   § 3.1.4.1 “Certification Related Statuses”
	   — Title: “FIDO_CERTIFIED_L2”

	   Authenticator passed FIDO Certification at level 2, stricter than level 1+.
	*/
	FIDO_CERTIFIED_L2 AuthenticatorStatus = "FIDO_CERTIFIED_L2"

	/*
	   FIDO_CERTIFIED_L2plus
	   § 3.1.4.1 “Certification Related Statuses”
	   — Title: “FIDO_CERTIFIED_L2plus”

	   Authenticator passed certification at level 2+, stricter than level 2.
	*/
	FIDO_CERTIFIED_L2plus AuthenticatorStatus = "FIDO_CERTIFIED_L2plus"

	/*
	   FIDO_CERTIFIED_L3
	   § 3.1.4.1 “Certification Related Statuses”
	   — Title: “FIDO_CERTIFIED_L3”

	   Authenticator passed FIDO Certification at level 3, stricter than level 2+.
	*/
	FIDO_CERTIFIED_L3 AuthenticatorStatus = "FIDO_CERTIFIED_L3"

	/*
	   FIDO_CERTIFIED_L3plus
	   § 3.1.4.1 “Certification Related Statuses”
	   — Title: “FIDO_CERTIFIED_L3plus”

	   Authenticator passed FIDO Certification at level 3+, stricter than level 3.
	*/
	FIDO_CERTIFIED_L3plus AuthenticatorStatus = "FIDO_CERTIFIED_L3plus"
)

/*
StatusReport
§ 3.1.3 “StatusReport dictionary”
Title: “StatusReport”

Contains an AuthenticatorStatus plus additional data (such as when the status was set, which
versions are affected, any related certificate, etc.). The latest StatusReport typically reflects
the authenticator’s current state.

Fields reference:

  - status: AuthenticatorStatus
  - effectiveDate: If set, ISO-8601 date the status took effect
  - authenticatorVersion: If relevant, the firmware version at or above which this status applies
  - certificate: Base64 DER-encoded PKIX certificate identifying compromised batches, if applicable
  - url: https/URL to relevant details
  - certificationDescriptor / certificateNumber / certificationPolicyVersion / certificationRequirementsVersion:
    For certification-based statuses (FIDO_CERTIFIED*, etc.)
*/
type StatusReport struct {
	Status                           AuthenticatorStatus `json:"status"`
	EffectiveDate                    *string             `json:"effectiveDate"`
	AuthenticatorVersion             *uint64             `json:"authenticatorVersion"`
	Certificate                      *string             `json:"certificate"`
	URL                              *string             `json:"url"`
	CertificationDescriptor          *string             `json:"certificationDescriptor"`
	CertificateNumber                *string             `json:"certificateNumber"`
	CertificationPolicyVersion       *string             `json:"certificationPolicyVersion"`
	CertificationRequirementsVersion *string             `json:"certificationRequirementsVersion"`
}

/*
AlternativeDescription
§ 3.1.2 in “FIDO Metadata Statement” includes a mention of descriptive fields in alternative languages,
while in “FIDO Metadata Service” the concept is repeated in multiple places for localized descriptions.

This map is used to hold short, human-readable descriptions in various locales, keyed by IETF language
tags (e.g., "en-US", "fr-FR", "zh-CN"). For example:

	{
	  "ru-RU": "Пример U2F аутентификатора от FIDO Alliance",
	  "fr-FR": "Exemple U2F authenticator de FIDO Alliance"
	}
*/
type AlternativeDescription map[string]string

/*
MetadataStatement
Primarily defined in the “FIDO Metadata Statement” specification § 5 “Metadata Statement Format”,
Title: “MetadataStatement”, and cross-referenced by “FIDO Metadata Service” in § 3.1.1.

This structure describes, in detail, an authenticator’s capabilities, keys, user verification methods,
attestation roots, and other properties. Key highlights include:

  - legalHeader: Must be present as per the statement’s legal acceptance.
  - aaid, aaguid: Identify UAF or FIDO2 authenticators respectively.
  - attestationCertificateKeyIdentifiers: For U2F authenticators that rely on dedicated attestation certs.
  - description and alternativeDescriptions: short text in English and optional localized translations.
  - authenticatorVersion: earliest version that satisfies the statement’s security and functionality.
  - protocolFamily: "uaf", "u2f", or "fido2".
  - schema: metadata statement version (3 for v3.0).
  - icon: data: URL (PNG) representing the authenticator visually.
*/
type MetadataStatement struct {
	LegalHeader                          string                 `json:"legalHeader"`
	AAID                                 string                 `json:"aaid"`
	AAGUID                               string                 `json:"aaguid"`
	AttestationCertificateKeyIdentifiers []string               `json:"attestationCertificateKeyIdentifiers"`
	Description                          string                 // Typically ASCII-only short descriptor in English
	AlternativeDescriptions              AlternativeDescription `json:"alternativeDescriptions"`
	AuthenticatorVersion                 uint64                 `json:"authenticatorVersion"`
	ProtocolFamily                       string                 `json:"protocolFamily"`
	Schema                               uint16                 `json:"schema"`

	// The fields below are selectively included from the “FIDO Metadata Statement” specification.
	// They can be expanded further to include userVerificationDetails, etc. as needed.
	KeyProtection bool `json:"-"` // Example placeholder; real spec field is an array of strings
	// ... other fields ...

	// For demonstration here, we only show a subset. In a full implementation, all required
	// metadata statement fields from §5 FIDO Metadata Statement would appear.
	IsKeyRestricted                 bool   `json:"isKeyRestricted"`
	IsFreshUserVerificationRequired bool   `json:"isFreshUserVerificationRequired"`
	Icon                            string `json:"icon"`
	IconDark                        string `json:"icon_dark"`
}

/*
BiometricStatusReport
§ 3.1.2 “BiometricStatusReport dictionary” in the FIDO Metadata Service v3.0

Contains details about biometric certification status for a specific modality (e.g. fingerprint).
Relevant fields:

  - certLevel: The biometric certification level
  - modality: The user verification modality (e.g. "fingerprint_internal")
  - effectiveDate: When the certLevel took effect
  - certificationDescriptor, certificateNumber, policy version, requirements version, etc.
*/
type BiometricStatusReport struct {
	CertLevel                        uint8   `json:"certLevel"`
	Modality                         string  `json:"modality"`
	EffectiveDate                    *string `json:"effectiveDate"`
	CertificationDescriptor          *string `json:"certificationDescriptor"`
	CertificateNumber                *string `json:"certificateNumber"`
	CertificationPolicyVersion       *string `json:"certificationPolicyVersion"`
	CertificationRequirementsVersion *string `json:"certificationRequirementsVersion"`
}

/*
Entry
§ 3.1.1 “Metadata BLOB Payload Entry dictionary” (Title: “MetadataBLOBPayloadEntry”)
in FIDO Metadata Service v3.0

Each Entry corresponds to an authenticator model recognized in the Metadata BLOB, identified by
AAGUID (FIDO2), AAID (UAF), or attestationCertificateKeyIdentifiers (U2F).

  - aaid, aaguid: identify the model
  - metadataStatement: the FIDO “MetadataStatement” with more detailed info
  - biometricStatusReports: optional if there's biometric certification
  - statusReports: array describing status transitions, from earliest to latest
  - timeOfLastStatusChange: when this array last changed
  - rogueListURL, rogueListHash: optional for referencing a list of rogue individual authenticators
*/
type Entry struct {
	AAGUID                               string                  `json:"aaguid"`
	AAID                                 string                  `json:"aaid"`
	MetadataStatement                    MetadataStatement       `json:"metadataStatement"`
	AttestationCertificateKeyIdentifiers []string                `json:"attestationCertificateKeyIdentifiers"`
	BiometricStatusReports               []BiometricStatusReport `json:"biometricStatusReports"`
	StatusReports                        []StatusReport          `json:"statusReports"`
	TimeOfLastStatusChange               string                  `json:"timeOfLastStatusChange"`
	RogueListURL                         string                  `json:"rogueListURL"`
	RogueListHash                        string                  `json:"rogueListHash"`
}
