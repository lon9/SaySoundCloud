# tools

They are tools to prepare sounds for the application.

## Initializing sounds on local

1. Execute `initsounds`
1. Execute `mksounddb`

## Initializing sounds on Firebase

1. Execute `initsounds`
1. Execute `gsutil -m cp -r {output_dir} gs:{your-bucket}`
1. Execute `mksounddbfromfirebase`


## Adding aditional sounds to Firebase

1. Execute `addsounds`
1. Execute `gsutil -m cp -r {output_dir} gs:{your-bucket}`
1. Execute `mksounddbfromfirebase`
