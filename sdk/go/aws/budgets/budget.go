// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package budgets

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a budgets budget resource. Budgets use the cost visualisation provided by Cost Explorer to show you the status of your budgets, to provide forecasts of your estimated costs, and to track your AWS usage, including your free tier usage.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/budgets_budget.html.markdown.
type Budget struct {
	// URN is this resource's unique name assigned by Pulumi.
	URN pulumi.URNOutput `pulumi:"urn"`

	// ID is this resource's unique identifier assigned by its provider.
	ID pulumi.IDOutput `pulumi:"id"`

	// The ID of the target account for budget. Will use current user's accountId by default if omitted.
	AccountId pulumi.StringOutput `pulumi:"accountId"`

	// Whether this budget tracks monetary cost or usage.
	BudgetType pulumi.StringOutput `pulumi:"budgetType"`

	// Map of CostFilters key/value pairs to apply to the budget.
	CostFilters pulumi.MapOutput `pulumi:"costFilters"`

	// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions..
	CostTypes pulumi.AnyOutput `pulumi:"costTypes"`

	// The amount of cost or usage being measured for a budget.
	LimitAmount pulumi.StringOutput `pulumi:"limitAmount"`

	// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
	LimitUnit pulumi.StringOutput `pulumi:"limitUnit"`

	// The name of a budget. Unique within accounts.
	Name pulumi.StringOutput `pulumi:"name"`

	// The prefix of the name of a budget. Unique within accounts.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`

	// Object containing Budget Notifications. Can be used multiple times to define more than one budget notification
	Notifications pulumi.ArrayOutput `pulumi:"notifications"`

	// The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
	TimePeriodEnd pulumi.StringOutput `pulumi:"timePeriodEnd"`

	// The start of the time period covered by the budget. The start date must come before the end date. Format: `2017-01-01_12:00`.
	TimePeriodStart pulumi.StringOutput `pulumi:"timePeriodStart"`

	// The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`.
	TimeUnit pulumi.StringOutput `pulumi:"timeUnit"`
}

// NewBudget registers a new resource with the given unique name, arguments, and options.
func NewBudget(ctx *pulumi.Context,
	name string, args *BudgetArgs, opts ...pulumi.ResourceOpt) (*Budget, error) {
	if args == nil || args.BudgetType == nil {
		return nil, errors.New("missing required argument 'BudgetType'")
	}
	if args == nil || args.LimitAmount == nil {
		return nil, errors.New("missing required argument 'LimitAmount'")
	}
	if args == nil || args.LimitUnit == nil {
		return nil, errors.New("missing required argument 'LimitUnit'")
	}
	if args == nil || args.TimePeriodStart == nil {
		return nil, errors.New("missing required argument 'TimePeriodStart'")
	}
	if args == nil || args.TimeUnit == nil {
		return nil, errors.New("missing required argument 'TimeUnit'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["name"] = pulumi.Any()
	if args != nil {
		inputs["accountId"] = args.AccountId
		inputs["budgetType"] = args.BudgetType
		inputs["costFilters"] = args.CostFilters
		inputs["costTypes"] = args.CostTypes
		inputs["limitAmount"] = args.LimitAmount
		inputs["limitUnit"] = args.LimitUnit
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["notifications"] = args.Notifications
		inputs["timePeriodEnd"] = args.TimePeriodEnd
		inputs["timePeriodStart"] = args.TimePeriodStart
		inputs["timeUnit"] = args.TimeUnit
	}
	var resource Budget
	err := ctx.RegisterResource("aws:budgets/budget:Budget", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBudget gets an existing Budget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBudget(ctx *pulumi.Context,
	name string, id pulumi.ID, state *BudgetState, opts ...pulumi.ResourceOpt) (*Budget, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		inputs["accountId"] = state.AccountId
		inputs["budgetType"] = state.BudgetType
		inputs["costFilters"] = state.CostFilters
		inputs["costTypes"] = state.CostTypes
		inputs["limitAmount"] = state.LimitAmount
		inputs["limitUnit"] = state.LimitUnit
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["notifications"] = state.Notifications
		inputs["timePeriodEnd"] = state.TimePeriodEnd
		inputs["timePeriodStart"] = state.TimePeriodStart
		inputs["timeUnit"] = state.TimeUnit
	}
	var resource Budget
	err := ctx.ReadResource("aws:budgets/budget:Budget", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetURN returns this resource's unique name assigned by Pulumi.
func (r *Budget) GetURN() pulumi.URNOutput {
	return r.URN
}

// GetID returns this resource's unique identifier assigned by its provider.
func (r *Budget) GetID() pulumi.IDOutput {
	return r.ID
}
// Input properties used for looking up and filtering Budget resources.
type BudgetState struct {
	// The ID of the target account for budget. Will use current user's accountId by default if omitted.
	AccountId pulumi.StringInput `pulumi:"accountId"`
	// Whether this budget tracks monetary cost or usage.
	BudgetType pulumi.StringInput `pulumi:"budgetType"`
	// Map of CostFilters key/value pairs to apply to the budget.
	CostFilters pulumi.MapInput `pulumi:"costFilters"`
	// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions..
	CostTypes pulumi.AnyInput `pulumi:"costTypes"`
	// The amount of cost or usage being measured for a budget.
	LimitAmount pulumi.StringInput `pulumi:"limitAmount"`
	// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
	LimitUnit pulumi.StringInput `pulumi:"limitUnit"`
	// The name of a budget. Unique within accounts.
	Name pulumi.StringInput `pulumi:"name"`
	// The prefix of the name of a budget. Unique within accounts.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// Object containing Budget Notifications. Can be used multiple times to define more than one budget notification
	Notifications pulumi.ArrayInput `pulumi:"notifications"`
	// The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
	TimePeriodEnd pulumi.StringInput `pulumi:"timePeriodEnd"`
	// The start of the time period covered by the budget. The start date must come before the end date. Format: `2017-01-01_12:00`.
	TimePeriodStart pulumi.StringInput `pulumi:"timePeriodStart"`
	// The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`.
	TimeUnit pulumi.StringInput `pulumi:"timeUnit"`
}

// The set of arguments for constructing a Budget resource.
type BudgetArgs struct {
	// The ID of the target account for budget. Will use current user's accountId by default if omitted.
	AccountId pulumi.StringInput `pulumi:"accountId"`
	// Whether this budget tracks monetary cost or usage.
	BudgetType pulumi.StringInput `pulumi:"budgetType"`
	// Map of CostFilters key/value pairs to apply to the budget.
	CostFilters pulumi.MapInput `pulumi:"costFilters"`
	// Object containing CostTypes The types of cost included in a budget, such as tax and subscriptions..
	CostTypes pulumi.AnyInput `pulumi:"costTypes"`
	// The amount of cost or usage being measured for a budget.
	LimitAmount pulumi.StringInput `pulumi:"limitAmount"`
	// The unit of measurement used for the budget forecast, actual spend, or budget threshold, such as dollars or GB. See [Spend](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/data-type-spend.html) documentation.
	LimitUnit pulumi.StringInput `pulumi:"limitUnit"`
	// The name of a budget. Unique within accounts.
	Name pulumi.StringInput `pulumi:"name"`
	// The prefix of the name of a budget. Unique within accounts.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// Object containing Budget Notifications. Can be used multiple times to define more than one budget notification
	Notifications pulumi.ArrayInput `pulumi:"notifications"`
	// The end of the time period covered by the budget. There are no restrictions on the end date. Format: `2017-01-01_12:00`.
	TimePeriodEnd pulumi.StringInput `pulumi:"timePeriodEnd"`
	// The start of the time period covered by the budget. The start date must come before the end date. Format: `2017-01-01_12:00`.
	TimePeriodStart pulumi.StringInput `pulumi:"timePeriodStart"`
	// The length of time until a budget resets the actual and forecasted spend. Valid values: `MONTHLY`, `QUARTERLY`, `ANNUALLY`.
	TimeUnit pulumi.StringInput `pulumi:"timeUnit"`
}
