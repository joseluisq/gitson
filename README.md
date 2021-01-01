# Gitson

> Get some Git commands output in JSON format.

WIP project under **active** development.

## Usage

```go
package main

import (
    "fmt"
    "os"

    git "github.com/joseluisq/gitson"
)

func main() {
    // Get `git log` output in JSON format
	jsonb, err := git.Log(os.Args, "/home/joseluisq/my-git-repo")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(jsonb))
}
```

#### JSON Output example

```json
[
    {
        "hash": "7cb004b7bb0da0e693a0d92b28832e18b39fc3c4",
        "hash_short": "7cb004b",
        "tree_hash": "1e86b094c645dd6fafbf831862b2a625463766d2",
        "tree_hash_short": "1e86b09",
        "parent_hashes": "",
        "parent_hashes_short": "",
        "author": "Jose Quintana",
        "author_mailmap": "Jose Quintana",
        "author_email": "joseluisq@localhost",
        "author_email_mailmap": "joseluisq@localhost",
        "author_email_local": "joseluisq",
        "author_local": "joseluisq",
        "author_date": "Thu Dec 31 02:40:56 2020 +0100",
        "author_date_rfc2822": "Thu, 31 Dec 2020 02:40:56 +0100",
        "author_date_relative": "2 days ago",
        "author_date_timestamp": "1609378856",
        "author_date_iso": "2020-12-31 02:40:56 +0100",
        "author_date_iso8601": "2020-12-31T02:40:56+01:00",
        "author_date_short": "2020-12-31",
        "committer": "Jose Quintana",
        "committer_mailmap": "Jose Quintana",
        "committer_email": "joseluisq@localhost",
        "committer_email_mailmap": "joseluisq@localhost",
        "committer_email_local": "joseluisq",
        "committer_email_local_mailmap": "joseluisq",
        "committer_date": "Thu Dec 31 02:40:56 2020 +0100",
        "committer_date_rfc2822": "Thu, 31 Dec 2020 02:40:56 +0100",
        "committer_date_relative": "2 days ago",
        "committer_date_timestamp": "1609378856",
        "committer_date_iso": "2020-12-31 02:40:56 +0100",
        "committer_date_iso8601": "2020-12-31T02:40:56+01:00",
        "committer_date_short": "2020-12-31",
        "ref_names": " (HEAD -> master, origin/master, origin/HEAD)",
        "ref_names_no_wrapping": "HEAD -> master, origin/master, origin/HEAD",
        "ref_names_given": "HEAD",
        "encoding": "",
        "subject": "feat-git-log-output-as-json",
        "body": "",
        "raw_body": "feat: git log output as json\n",
        "notes": "",
        "gpg_raw_verification": "gpg: Signature made Thu Dec 31 02:41:36 2020 CET\ngpg:                using RSA key 12E57137E61B7DE5395639909ED7C351C78E1F64\ngpg:                issuer \"joseluisq@localhost\"\ngpg: Good signature from \"Jose Quintana <joseluisq@localhost>\" [ultimate]\n",
        "gpg_verification_code": "G",
        "gpg_signer_name": "Jose Quintana <joseluisq@localhost>",
        "gpg_key": "9ED7C351C78E1F64",
        "gpg_key_fingerprint": "12E57137E61B7DE5395639909ED7C351C78E1F64",
        "gpg_key_fingerprint_subkey": "12E57137E61B7DE5395639909ED7C351C78E1F64",
        "gpg_key_trust_level": "ultimate",
        "reflog_selector": "",
        "reflog_selector_short": "",
        "reflog_identity_name": "",
        "reflog_identity_name_mailmap": "",
        "reflog_identity_email": "",
        "reflog_identity_email_mailmap": "",
        "reflog_subject": ""
    }
]
```

## Contributions

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in current work by you, as defined in the Apache-2.0 license, shall be dual licensed as described below, without any additional terms or conditions.

Feel free to send some [Pull request](https://github.com/joseluisq/gitson/pulls) or [issue](https://github.com/joseluisq/gitson/issues).

## License

This work is primarily distributed under the terms of both the [MIT license](LICENSE-MIT) and the [Apache License (Version 2.0)](LICENSE-APACHE).

© 2020-present [Jose Quintana](https://git.io/joseluisq)
