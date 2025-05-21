# Zendesk Terraform Provider

> [!IMPORTANT]  
> Read the contribution guideline before adding a pull request.

<!-- TOC -->
* [Zendesk Terraform Provider](#zendesk-terraform-provider)
  * [Developing the Provider](#developing-the-provider)
    * [Adding and Updating Dependencies](#adding-and-updating-dependencies)
    * [Documentation Generation](#documentation-generation)
    * [Testing](#testing)
      * [Acceptance Testing](#acceptance-testing)
<!-- TOC -->

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (
see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin`
directory.

### Adding and Updating Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

To update an existing dependency:

```shell
go get -u github.com/author/dependency
go mod tidy
```

In either case, commit the changes to `go.mod` and `go.sum`.

### Documentation Generation

> [!IMPORTANT]
> This is required when making any schema changes, as the pipeline will fail if docs are not updated in PR
> To generate or update documentation, run `go generate`.

### Testing

#### Acceptance Testing

> [!CAUTION]
> In order to run acceptance tests, the following environment variables must be set in a .env file:

| Name              | Description                                                                  |
|-------------------|------------------------------------------------------------------------------|
| ZENDESK_SUBDOMAIN | Zendesk subdomain, ex: dynatrace                                             |
| ZENDESK_USERNAME  | Zendesk username used for basic token auth.                                  |
| ZENDESK_API_TOKEN | API token used for Zendesk basic token auth.                                 |

Once these variables are set, please run following command:

```shell
  task acctest
```


