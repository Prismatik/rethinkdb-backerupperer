## RethinkDB Backerupperer

It backs up RethinkDB. It puts those backups on S3.

It can just do it once, or, if you want, more than once.

### Usage

When run, it does a `rethinkdb dump` against the target database. It saves the output to a file called something like `2015-10-23T19:51:38+11:00.tar.gz`

It then uploads that file to the specified S3 bucket. Once the file is uploaded, it deletes the file.

If you like, you can specify an [encryption key for S3 to use](http://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html)

If you like, you can specify a [cron expression](https://godoc.org/github.com/robfig/cron) and it will run on that interval.

The required environment variables are:

* AWS_ACCESS_KEY_ID
* AWS_SECRET_ACCESS_KEY
* S3_BUCKET
* RETHINK_LOC (eg localhost:28015)

Optional environment variables are:

* SSE_KEY (In my testing, S3 barfs if this isn't 32 characters in length)
* CRON_STRING (Normal cron times, with extras defined [here](https://godoc.org/github.com/robfig/cron))

If SSE_KEY is not defined, the file will just be stored in the clear.
If CRON_STRING is not defined, the process will just run once then exit.

So an example invocation would be:

`$ AWS_ACCESS_KEY_ID={{KEY}} AWS_SECRET_ACCESS_KEY={{SECRET}} S3_BUCKET={{bucket}} RETHINK_LOC=localhost:28015 SSE_KEY=EnE50AzSFcm0k6iq0DGmBMUIjM2NozxS CRON_STRING="@every 1h" go run rethinkdb-backerupperer.go`

Every time it succeeds, it will log something similar to:

`2015/10/23 20:42:13 Successfully uploaded backup 2015-10-23T20:42:10+11:00.tar.gz`

Any failure will exit the process with an exit code of 1

### Download

You can download binaries for OSX, Linux and Windows [here](https://github.com/Prismatik/rethinkdb-backerupperer/releases/latest)

NB: I have not tested the Windows build. Please let me know if it works (or doesn't)
