# cloudconvert
Golang CloudConvert client

Installation
------------

```go
  go get github.com/pinda/cloudconvert
```

Usage
-----
First, start a new process:

```go
cloudClient := cloudconvert.NewClient(<YOUR_API_KEY>)

process := cloudconvert.ProcessInput{
	InputFormat:  <INPUT>,
	OutputFormat: <OUTPUT>,
}
pr, err := cloudClient.Process.New(process)
if err != nil {
   return
}
```

Now start your new conversion:

```go
conversion := cloudconvert.NewConversion(...)

cloudClient.Conversion.New(pr.URL, conversion)
```

Check the status of the conversion using:

```go
status, err := cloudClient.Conversion.Status(pr.URL)
if err != nil {
   return
}
```

If the conversion is ready, you can download the file contents using:

```go
body, err := cloudClient.Conversion.Download(pr.Output.URL)
if err != nil {
   return
}
defer body.Close()

<< DO SOMETHING >>
```

Amazon S3
-----
If you want to convert using your s3 bucket I provided some simple methods for that.
Start with creating your process as described above and then start a `S3Conversion`:

```go
conversion := cloudconvert.NewS3Conversion(...)

cloudClient.Conversion.NewS3(pr.URL, conversion)
```


License
-------

The MIT License (MIT)

Contributed by [Joeri Djojosoeparto](https://github.com/pinda)
