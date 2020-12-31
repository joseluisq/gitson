package gitson

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"strings"
)

// LogCommit represents placeholders that expand to information extracted from the commit.
type LogCommit struct {
	// %H
	// commit hash
	Hash string `json:"hash"`

	// %h
	// abbreviated commit hash
	HashShort string `json:"hash_short"`

	// %T
	// tree hash
	TreeHash string `json:"tree_hash"`

	// %t
	// abbreviated tree hash
	TreeHashShort string `json:"tree_hash_short"`

	// %P
	// parent hashes
	ParentHashes string `json:"parent_hashes"`

	// %p
	// abbreviated parent hashes
	ParentHashesShort string `json:"parent_hashes_short"`

	// %an
	// author name
	Author string `json:"author"`

	// %aN
	// author name (respecting .mailmap, see git-shortlog[1] or git-blame[1])
	AuthorMailmap string `json:"author_mailmap"`

	// %ae
	// author email
	AuthorEMail string `json:"author_email"`

	// %aE
	// author email (respecting .mailmap, see git-shortlog[1] or git-blame[1])
	AuthorEMailMailmap string `json:"author_email_mailmap"`

	// %al
	// author email local-part (the part before the @ sign)
	AuthorEMailLocal string `json:"author_email_local"`

	// %aL
	// author local-part (see %al) respecting .mailmap, see git-shortlog[1] or git-blame[1])
	AuthorLocal string `json:"author_local"`

	// %ad
	// author date (format respects --date= option)
	AuthorDate string `json:"author_date"`

	// %aD
	// author date, RFC2822 style
	AuthorDateRFC2822 string `json:"author_date_rfc2822"`

	// %ar
	// author date, relative
	AuthorDateRelative string `json:"author_date_relative"`

	// %at
	// author date, UNIX timestamp
	AuthorDateTimestamp string `json:"author_date_timestamp"`

	// %ai
	// author date, ISO 8601-like format
	AuthorDateISO string `json:"author_date_iso"`

	// %aI
	// author date, strict ISO 8601 format
	AuthorDateISO8601 string `json:"author_date_iso8601"`

	// %as
	// author date, short format (YYYY-MM-DD)
	AuthorDateShort string `json:"author_date_short"`

	// %cn
	// committer name
	Committer string `json:"committer"`

	// %cN
	// committer name (respecting .mailmap, see git-shortlog[1] or git-blame[1])
	CommitterMailmap string `json:"committer_mailmap"`

	// %ce
	// committer email
	CommitterEMail string `json:"committer_email"`

	// %cE
	// committer email (respecting .mailmap, see git-shortlog[1] or git-blame[1])
	CommitterEMailMailmap string `json:"committer_email_mailmap"`

	// %cl
	// committer email local-part (the part before the @ sign)
	CommitterEMailLocal string `json:"committer_email_local"`

	// %cL
	// committer local-part (see %cl) respecting .mailmap, see git-shortlog[1] or git-blame[1])
	CommitterEMailLocalMailmap string `json:"committer_email_local_mailmap"`

	// %cd
	// committer date (format respects --date= option)
	CommitterDate string `json:"committer_date"`

	// %cD
	// committer date, RFC2822 style
	CommitterDateRFC2822 string `json:"committer_date_rfc2822"`

	// %cr
	// committer date, relative
	CommitterDateRelative string `json:"committer_date_relative"`

	// %ct
	// committer date, UNIX timestamp
	CommitterDateTimestamp string `json:"committer_date_timestamp"`

	// %ci
	// committer date, ISO 8601-like format
	CommitterDateISO string `json:"committer_date_iso"`

	// %cI
	// committer date, strict ISO 8601 format
	CommitterDateISO8601 string `json:"committer_date_iso8601"`

	// %cs
	// committer date, short format (YYYY-MM-DD)
	CommitterDateShort string `json:"committer_date_short"`

	// %d
	// ref names, like the --decorate option of git-log[1]
	RefNames string `json:"ref_names"`

	// %D
	// ref names without the " (", ")" wrapping.
	RefNamesNoWrapping string `json:"ref_names_no_wrapping"`

	// %S
	// ref name given on the command line by which the commit was reached (like git log --source), only works with git log
	RefNamesGiven string `json:"ref_names_given"`

	// %e
	// encoding
	Encoding string `json:"encoding"`

	// %f
	// sanitized subject line, suitable for a filename
	Subject string `json:"subject"`

	// %b
	// body
	Body string `json:"body"`

	// %B
	// raw body (unwrapped subject and body)
	RawBody string `json:"raw_body"`

	// %N
	// commit notes
	Notes string `json:"notes"`

	// %GG
	// raw verification message from GPG for a signed commit
	GPGRawVerification string `json:"gpg_raw_verification"`

	// %G?
	// show "G" for a good (valid) signature, "B" for a bad signature, "U" for a good signature with unknown validity, "X" for a good signature that has expired, "Y" for a good signature made by an expired key, "R" for a good signature made by a revoked key, "E" if the signature cannot be checked (e.g. missing key) and "N" for no signature
	GPGVerificationCode string `json:"gpg_verification_code"`

	// %GS
	// show the name of the signer for a signed commit
	GPGSignerName string `json:"gpg_signer_name"`

	// %GK
	// show the key used to sign a signed commit
	GPGKey string `json:"gpg_key"`

	// %GF
	// show the fingerprint of the key used to sign a signed commit
	GPGKeyFingerprint string `json:"gpg_key_fingerprint"`

	// %GP
	// show the fingerprint of the primary key whose subkey was used to sign a signed commit
	GPGKeyFingerprintSubkey string `json:"gpg_key_fingerprint_subkey"`

	// %GT
	// show the trust level for the key used to sign a signed commit
	GPGKeyTrustLevel string `json:"gpg_key_trust_level"`

	// %gD
	// reflog selector, e.g., refs/stash@{1} or refs/stash@{2 minutes ago}; the format follows the rules described for the -g option. The portion before the @ is the refname as given on the command line (so git log -g refs/heads/master would yield refs/heads/master@{0}).
	ReflogSelector string `json:"reflog_selector"`

	// %gd
	// shortened reflog selector; same as %gD, but the refname portion is shortened for human readability (so refs/heads/master becomes just master).
	ReflogSelectorShort string `json:"reflog_selector_short"`

	// %gn
	// reflog identity name
	ReflogIdentityName string `json:"reflog_identity_name"`

	// %gN
	// reflog identity name (respecting .mailmap, see git-shortlog[1] or git-blame[1])
	ReflogIdentityNameMailmap string `json:"reflog_identity_name_mailmap"`

	// %ge
	// reflog identity email
	ReflogIdentityEmail string `json:"reflog_identity_email"`

	// %gE
	// reflog identity email (respecting .mailmap, see git-shortlog[1] or git-blame[1])
	ReflogIdentityEmailMailmap string `json:"reflog_identity_email_mailmap"`

	// %gs
	// reflog subject
	ReflogSubject string `json:"reflog_subject"`
}

var logFormatPlaceholders = [][]string{
	// %H
	// commit hash
	{"hash", "%H"},

	// %h
	// abbreviated commit hash
	{"hash_short", "%h"},

	// %T
	// tree hash
	{"tree_hash", "%T"},

	// %t
	// abbreviated tree hash
	{"tree_hash_short", "%t"},

	// %P
	// parent hashes
	{"parent_hashes", "%P"},

	// %p
	// abbreviated parent hashes
	{"parent_hashes_short", "%p"},

	// %an
	// author name
	{"author", "%an"},

	// %aN
	// author name (respecting .mailmap, see git-shortlog[1] or git-blame[1])
	{"author_mailmap", "%aN"},

	// %ae
	// author email
	{"author_email", "%ae"},

	// %aE
	// author email (respecting .mailmap, see git-shortlog[1] or git-blame[1])
	{"author_email_mailmap", "%aE"},

	// %al
	// author email local-part (the part before the @ sign)
	{"author_email_local", "%al"},

	// %aL
	// author local-part (see %al) respecting .mailmap, see git-shortlog[1] or git-blame[1])
	{"author_local", "%aL"},

	// %ad
	// author date (format respects --date= option)
	{"author_date", "%ad"},

	// %aD
	// author date, RFC2822 style
	{"author_date_rfc2822", "%aD"},

	// %ar
	// author date, relative
	{"author_date_relative", "%ar"},

	// %at
	// author date, UNIX timestamp
	{"author_date_timestamp", "%at"},

	// %ai
	// author date, ISO 8601-like format
	{"author_date_iso", "%ai"},

	// %aI
	// author date, strict ISO 8601 format
	{"author_date_iso8601", "%aI"},

	// %as
	// author date, short format (YYYY-MM-DD)
	{"author_date_short", "%as"},

	// %cn
	// committer name
	{"committer", "%cn"},

	// %cN
	// committer name (respecting .mailmap, see git-shortlog[1] or git-blame[1])
	{"committer_mailmap", "%cN"},

	// %ce
	// committer email
	{"committer_email", "%ce"},

	// %cE
	// committer email (respecting .mailmap, see git-shortlog[1] or git-blame[1])
	{"committer_email_mailmap", "%cE"},

	// %cl
	// committer email local-part (the part before the @ sign)
	{"committer_email_local", "%cl"},

	// %cL
	// committer local-part (see %cl) respecting .mailmap, see git-shortlog[1] or git-blame[1])
	{"committer_email_local_mailmap", "%cL"},

	// %cd
	// committer date (format respects --date= option)
	{"committer_date", "%cd"},

	// %cD
	// committer date, RFC2822 style
	{"committer_date_rfc2822", "%cD"},

	// %cr
	// committer date, relative
	{"committer_date_relative", "%cr"},

	// %ct
	// committer date, UNIX timestamp
	{"committer_date_timestamp", "%ct"},

	// %ci
	// committer date, ISO 8601-like format
	{"committer_date_iso", "%ci"},

	// %cI
	// committer date, strict ISO 8601 format
	{"committer_date_iso8601", "%cI"},

	// %cs
	// committer date, short format (YYYY-MM-DD)
	{"committer_date_short", "%cs"},

	// %d
	// ref names, like the --decorate option of git-log[1]
	{"ref_names", "%d"},

	// %D
	// ref names without the " (", ")" wrapping.
	{"ref_names_no_wrapping", "%D"},

	// %S
	// ref name given on the command line by which the commit was reached (like git log --source), only works with git log
	{"ref_names_given", "%S"},

	// %e
	// encoding
	{"encoding", "%e"},

	// %f
	// sanitized subject line, suitable for a filename
	{"subject", "%f"},

	// %b
	// body
	{"body", "%b"},

	// %B
	// raw body (unwrapped subject and body)
	{"raw_body", "%B"},

	// %N
	// commit notes
	{"notes", "%N"},

	// %GG
	// raw verification message from GPG for a signed commit
	{"gpg_raw_verification", "%GG"},

	// %G?
	// show "G" for a good (valid) signature, "B" for a bad signature, "U" for a good signature with unknown validity, "X" for a good signature that has expired, "Y" for a good signature made by an expired key, "R" for a good signature made by a revoked key, "E" if the signature cannot be checked (e.g. missing key) and "N" for no signature
	{"gpg_verification_code", "%G?"},

	// %GS
	// show the name of the signer for a signed commit
	{"gpg_signer_name", "%GS"},

	// %GK
	// show the key used to sign a signed commit
	{"gpg_key", "%GK"},

	// %GF
	// show the fingerprint of the key used to sign a signed commit
	{"gpg_key_fingerprint", "%GF"},

	// %GP
	// show the fingerprint of the primary key whose subkey was used to sign a signed commit
	{"gpg_key_fingerprint_subkey", "%GP"},

	// %GT
	// show the trust level for the key used to sign a signed commit
	{"gpg_key_trust_level", "%GT"},

	// %gD
	// reflog selector, e.g., refs/stash@{1} or refs/stash@{2 minutes ago}; the format follows the rules described for the -g option. The portion before the @ is the refname as given on the command line (so git log -g refs/heads/master would yield refs/heads/master@{0}).
	{"reflog_selector", "%gD"},

	// %gd
	// shortened reflog selector; same as %gD, but the refname portion is shortened for human readability (so refs/heads/master becomes just master).
	{"reflog_selector_short", "%gd"},

	// %gn
	// reflog identity name
	{"reflog_identity_name", "%gn"},

	// %gN
	// reflog identity name (respecting .mailmap, see git-shortlog[1] or git-blame[1])
	{"reflog_identity_name_mailmap", "%gN"},

	// %ge
	// reflog identity email
	{"reflog_identity_email", "%ge"},

	// %gE
	// reflog identity email (respecting .mailmap, see git-shortlog[1] or git-blame[1])
	{"reflog_identity_email_mailmap", "%gE"},

	// %gs
	// reflog subject
	{"reflog_subject", "%gs"},
}

const (
	// EOL defines an End Of Line char
	EOL string = "@_EOL_@"
	// SEPVAL defines char separator beetween a value-pair
	SEPVAL string = "@_SEP_VAL_@"
	// SEPPAIR defines a char separator per every pair
	SEPPAIR string = "@_SEP_PAIR_@"
)

// Log gets Git log information in JSON bytes.
func Log(args []string, repoPath string) ([]byte, error) {
	var argsv []string
	var pretty bool = false
	// Skip always the first element of args slice
	for _, v := range args[1:] {
		if strings.HasPrefix(v, "--help") || strings.HasPrefix(v, "-h") {
			continue
		}
		if strings.HasPrefix(v, "--pretty") {
			if !strings.HasPrefix(v, "--pretty=") {
				pretty = true
			}
			continue
		}
		if pretty {
			continue
		}
		argsv = append(argsv, v)
	}

	csep := SEPPAIR
	format := ""
	for i, f := range logFormatPlaceholders {
		if i == len(logFormatPlaceholders)-1 {
			csep = ""
		}
		format = format + f[0] + SEPVAL + f[1] + csep
	}
	argsv = append(
		[]string{"log", "--no-decorate", "--pretty=format:" + format + EOL},
		argsv...,
	)

	var out bytes.Buffer
	cmd := exec.Command("git", argsv...)
	if repoPath != "" {
		cmd.Dir = repoPath
	}
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	EOLB := []byte(EOL)
	SEPB := []byte(SEPPAIR)
	LFB := []byte("\n")
	r := bufio.NewReader(&out)
	scanner := bufio.NewScanner(r)
	scanner.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.Index(data, EOLB); i >= 0 {
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	})

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	for scanner.Scan() {
		b := scanner.Bytes()
		if bytes.HasPrefix(b, EOLB[1:]) {
			b = b[len(EOLB[1:]):]
		}
		if bytes.HasPrefix(b, LFB) {
			b = b[1:]
		}
		if len(b) == 0 {
			continue
		}

		p := []byte(",")
		pairs := bytes.Split(b, SEPB)
		var pairsr [][]byte
		for n, v := range pairs {
			r := bytes.Split(v, []byte(SEPVAL))
			s := bytes.ReplaceAll(r[1], []byte("\n"), []byte(`\n`))
			s = bytes.ReplaceAll(s, []byte("\r"), []byte(`\r`))
			s = bytes.ReplaceAll(s, []byte("\t"), []byte(`\t`))
			s = bytes.ReplaceAll(s, []byte("\""), []byte(`\"`))
			if n == len(pairs)-1 {
				p = []byte("")
			}
			pairsr = append(pairsr, []byte("\""))
			pairsr = append(pairsr, []byte(r[0]))
			pairsr = append(pairsr, []byte("\""))
			pairsr = append(pairsr, []byte(":"))
			pairsr = append(pairsr, []byte("\""))
			pairsr = append(pairsr, []byte(s))
			pairsr = append(pairsr, []byte("\""))
			pairsr = append(pairsr, []byte(p))
		}
		rs := append([]byte("{"), bytes.Join(pairsr, []byte(""))...)
		rs = append(rs, []byte("},")...)
		_, err := w.Write(rs)
		if err != nil {
			return nil, err
		}
	}
	w.Flush()

	z := buf.Bytes()
	// remove trailing comma
	z = z[:len(z)-1]
	z = append([]byte("["), z...)
	z = append(z, []byte("]")...)

	return z, nil
}
