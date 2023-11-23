package object

// AuthzSimple creates an authorization object that will control access to
// network services.
//
// The identity parameter identifies the user and its format depends on the
// network service that authorization object is associated with. For authorizing
// based on TLS x509 certificates, the identity must be the x509 distinguished
// name. Note that care must be taken to escape any commas in the distinguished name.
func AuthzSimple(id string, identity string) *cli.Option {
	props := []*cli.Property{
		cli.NewProperty("id", id),
		cli.NewProperty("identity", identity),
	}

	return cli.NewOption("object", "authz-simple", props...)
}

// AuthzListFile creates an authorization object that will control access to
// network services.
//
// The file parameter is the fully qualified path to a file containing the
// access control list rules in JSON format.
//
// An example set of rules that match against SASL usernames might look like:
//
//	{
//		"rules": [
//			{ "match": "fred", "policy": "allow", "format": "exact" },
//			{ "match": "bob", "policy": "allow", "format": "exact" },
//			{ "match": "danb", "policy": "deny", "format": "glob" },
//			{ "match": "dan*", "policy": "allow", "format": "exact" }
//		],
//		"policy": "deny"
//	}
//
// When checking access the object will iterate over all the rules and the first
// rule to match will have its policy value returned as the result. If no rules
// match, then the default policy value is returned.
//
// The rules can either be an exact string match, or they can use the simple Unix
// glob pattern matching to allow wildcards to be used.
//
// If the refresh parameter is set to true the file will be monitored and
// automatically reloaded whenever its content changes.
func AuthzListFile(id string, file string, refresh bool) *cli.Option {
	props := []*cli.Property{
		cli.NewProperty("id", id),
		cli.NewProperty("filename", file),
		cli.NewProperty("refresh", refresh),
	}

	return cli.NewOption("object", "authz-listfile", props...)
}

// AuthzPAM creates an authorization object that will control access to network
// services.
//
// The service parameter provides the name of a PAM service to use for authorization.
// It requires that a file `/etc/pam.d/service` exist to provide the configuration
// for the account subsystem.
func AuthzPAM(id string, service string) *cli.Option {
	props := []*cli.Property{
		cli.NewProperty("id", id),
		cli.NewProperty("service", service),
	}

	return cli.NewOption("object", "authz-pam", props...)
}
