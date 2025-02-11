// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package transfer

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a AWS Transfer User resource. Managing SSH keys can be accomplished with the [`transfer.SshKey` resource](https://www.terraform.io/docs/providers/aws/r/transfer_ssh_key.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/transfer_user.html.markdown.
type User struct {
	s *pulumi.ResourceState
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOpt) (*User, error) {
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.ServerId == nil {
		return nil, errors.New("missing required argument 'ServerId'")
	}
	if args == nil || args.UserName == nil {
		return nil, errors.New("missing required argument 'UserName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["homeDirectory"] = nil
		inputs["policy"] = nil
		inputs["role"] = nil
		inputs["serverId"] = nil
		inputs["tags"] = nil
		inputs["userName"] = nil
	} else {
		inputs["homeDirectory"] = args.HomeDirectory
		inputs["policy"] = args.Policy
		inputs["role"] = args.Role
		inputs["serverId"] = args.ServerId
		inputs["tags"] = args.Tags
		inputs["userName"] = args.UserName
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:transfer/user:User", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &User{s: s}, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.ID, state *UserState, opts ...pulumi.ResourceOpt) (*User, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["homeDirectory"] = state.HomeDirectory
		inputs["policy"] = state.Policy
		inputs["role"] = state.Role
		inputs["serverId"] = state.ServerId
		inputs["tags"] = state.Tags
		inputs["userName"] = state.UserName
	}
	s, err := ctx.ReadResource("aws:transfer/user:User", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &User{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *User) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *User) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Amazon Resource Name (ARN) of Transfer User
func (r *User) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// The landing directory (folder) for a user when they log in to the server using their SFTP client.  It should begin with a `/`.  The first item in the path is the name of the home bucket (accessible as `${Transfer:HomeBucket}` in the policy) and the rest is the home directory (accessible as `${Transfer:HomeDirectory}` in the policy). For example, `/example-bucket-1234/username` would set the home bucket to `example-bucket-1234` and the home directory to `username`.
func (r *User) HomeDirectory() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["homeDirectory"])
}

// An IAM JSON policy document that scopes down user access to portions of their Amazon S3 bucket. IAM variables you can use inside this policy include `${Transfer:UserName}`, `${Transfer:HomeDirectory}`, and `${Transfer:HomeBucket}`. These are evaluated on-the-fly when navigating the bucket.
func (r *User) Policy() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["policy"])
}

// Amazon Resource Name (ARN) of an IAM role that allows the service to controls your user’s access to your Amazon S3 bucket.
func (r *User) Role() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["role"])
}

// The Server ID of the Transfer Server (e.g. `s-12345678`)
func (r *User) ServerId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["serverId"])
}

// A mapping of tags to assign to the resource.
func (r *User) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// The name used for log in to your SFTP server.
func (r *User) UserName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["userName"])
}

// Input properties used for looking up and filtering User resources.
type UserState struct {
	// Amazon Resource Name (ARN) of Transfer User
	Arn interface{}
	// The landing directory (folder) for a user when they log in to the server using their SFTP client.  It should begin with a `/`.  The first item in the path is the name of the home bucket (accessible as `${Transfer:HomeBucket}` in the policy) and the rest is the home directory (accessible as `${Transfer:HomeDirectory}` in the policy). For example, `/example-bucket-1234/username` would set the home bucket to `example-bucket-1234` and the home directory to `username`.
	HomeDirectory interface{}
	// An IAM JSON policy document that scopes down user access to portions of their Amazon S3 bucket. IAM variables you can use inside this policy include `${Transfer:UserName}`, `${Transfer:HomeDirectory}`, and `${Transfer:HomeBucket}`. These are evaluated on-the-fly when navigating the bucket.
	Policy interface{}
	// Amazon Resource Name (ARN) of an IAM role that allows the service to controls your user’s access to your Amazon S3 bucket.
	Role interface{}
	// The Server ID of the Transfer Server (e.g. `s-12345678`)
	ServerId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The name used for log in to your SFTP server.
	UserName interface{}
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// The landing directory (folder) for a user when they log in to the server using their SFTP client.  It should begin with a `/`.  The first item in the path is the name of the home bucket (accessible as `${Transfer:HomeBucket}` in the policy) and the rest is the home directory (accessible as `${Transfer:HomeDirectory}` in the policy). For example, `/example-bucket-1234/username` would set the home bucket to `example-bucket-1234` and the home directory to `username`.
	HomeDirectory interface{}
	// An IAM JSON policy document that scopes down user access to portions of their Amazon S3 bucket. IAM variables you can use inside this policy include `${Transfer:UserName}`, `${Transfer:HomeDirectory}`, and `${Transfer:HomeBucket}`. These are evaluated on-the-fly when navigating the bucket.
	Policy interface{}
	// Amazon Resource Name (ARN) of an IAM role that allows the service to controls your user’s access to your Amazon S3 bucket.
	Role interface{}
	// The Server ID of the Transfer Server (e.g. `s-12345678`)
	ServerId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The name used for log in to your SFTP server.
	UserName interface{}
}
