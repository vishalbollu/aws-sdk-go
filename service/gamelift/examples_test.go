// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package gamelift_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/gamelift"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleGameLift_CreateAlias() {
	svc := gamelift.New(session.New())

	params := &gamelift.CreateAliasInput{
		Name:        aws.String("NonZeroAndMaxString"), // Required
		Description: aws.String("NonZeroAndMaxString"),
		RoutingStrategy: &gamelift.RoutingStrategy{
			FleetId: aws.String("FleetId"),
			Message: aws.String("FreeText"),
			Type:    aws.String("RoutingStrategyType"),
		},
	}
	resp, err := svc.CreateAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_CreateBuild() {
	svc := gamelift.New(session.New())

	params := &gamelift.CreateBuildInput{
		Name:    aws.String("NonZeroAndMaxString"),
		Version: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.CreateBuild(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_CreateFleet() {
	svc := gamelift.New(session.New())

	params := &gamelift.CreateFleetInput{
		BuildId:          aws.String("BuildId"),             // Required
		EC2InstanceType:  aws.String("EC2InstanceType"),     // Required
		Name:             aws.String("NonZeroAndMaxString"), // Required
		ServerLaunchPath: aws.String("NonZeroAndMaxString"), // Required
		Description:      aws.String("NonZeroAndMaxString"),
		EC2InboundPermissions: []*gamelift.IpPermission{
			{ // Required
				FromPort: aws.Int64(1),                 // Required
				IpRange:  aws.String("NonBlankString"), // Required
				Protocol: aws.String("IpProtocol"),     // Required
				ToPort:   aws.Int64(1),                 // Required
			},
			// More values...
		},
		LogPaths: []*string{
			aws.String("NonZeroAndMaxString"), // Required
			// More values...
		},
		ServerLaunchParameters: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.CreateFleet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_CreateGameSession() {
	svc := gamelift.New(session.New())

	params := &gamelift.CreateGameSessionInput{
		MaximumPlayerSessionCount: aws.Int64(1), // Required
		AliasId:                   aws.String("AliasId"),
		FleetId:                   aws.String("FleetId"),
		GameProperties: []*gamelift.GameProperty{
			{ // Required
				Key:   aws.String("GamePropertyKey"),   // Required
				Value: aws.String("GamePropertyValue"), // Required
			},
			// More values...
		},
		Name: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.CreateGameSession(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_CreatePlayerSession() {
	svc := gamelift.New(session.New())

	params := &gamelift.CreatePlayerSessionInput{
		GameSessionId: aws.String("GameSessionId"),       // Required
		PlayerId:      aws.String("NonZeroAndMaxString"), // Required
	}
	resp, err := svc.CreatePlayerSession(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_CreatePlayerSessions() {
	svc := gamelift.New(session.New())

	params := &gamelift.CreatePlayerSessionsInput{
		GameSessionId: aws.String("GameSessionId"), // Required
		PlayerIds: []*string{ // Required
			aws.String("NonZeroAndMaxString"), // Required
			// More values...
		},
	}
	resp, err := svc.CreatePlayerSessions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DeleteAlias() {
	svc := gamelift.New(session.New())

	params := &gamelift.DeleteAliasInput{
		AliasId: aws.String("AliasId"), // Required
	}
	resp, err := svc.DeleteAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DeleteBuild() {
	svc := gamelift.New(session.New())

	params := &gamelift.DeleteBuildInput{
		BuildId: aws.String("BuildId"), // Required
	}
	resp, err := svc.DeleteBuild(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DeleteFleet() {
	svc := gamelift.New(session.New())

	params := &gamelift.DeleteFleetInput{
		FleetId: aws.String("FleetId"), // Required
	}
	resp, err := svc.DeleteFleet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeAlias() {
	svc := gamelift.New(session.New())

	params := &gamelift.DescribeAliasInput{
		AliasId: aws.String("AliasId"), // Required
	}
	resp, err := svc.DescribeAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeBuild() {
	svc := gamelift.New(session.New())

	params := &gamelift.DescribeBuildInput{
		BuildId: aws.String("BuildId"), // Required
	}
	resp, err := svc.DescribeBuild(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeEC2InstanceLimits() {
	svc := gamelift.New(session.New())

	params := &gamelift.DescribeEC2InstanceLimitsInput{
		EC2InstanceType: aws.String("EC2InstanceType"),
	}
	resp, err := svc.DescribeEC2InstanceLimits(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeFleetAttributes() {
	svc := gamelift.New(session.New())

	params := &gamelift.DescribeFleetAttributesInput{
		FleetIds: []*string{
			aws.String("FleetId"), // Required
			// More values...
		},
		Limit:     aws.Int64(1),
		NextToken: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.DescribeFleetAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeFleetCapacity() {
	svc := gamelift.New(session.New())

	params := &gamelift.DescribeFleetCapacityInput{
		FleetIds: []*string{
			aws.String("FleetId"), // Required
			// More values...
		},
		Limit:     aws.Int64(1),
		NextToken: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.DescribeFleetCapacity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeFleetEvents() {
	svc := gamelift.New(session.New())

	params := &gamelift.DescribeFleetEventsInput{
		FleetId:   aws.String("FleetId"), // Required
		EndTime:   aws.Time(time.Now()),
		Limit:     aws.Int64(1),
		NextToken: aws.String("NonZeroAndMaxString"),
		StartTime: aws.Time(time.Now()),
	}
	resp, err := svc.DescribeFleetEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeFleetPortSettings() {
	svc := gamelift.New(session.New())

	params := &gamelift.DescribeFleetPortSettingsInput{
		FleetId: aws.String("FleetId"), // Required
	}
	resp, err := svc.DescribeFleetPortSettings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeFleetUtilization() {
	svc := gamelift.New(session.New())

	params := &gamelift.DescribeFleetUtilizationInput{
		FleetIds: []*string{
			aws.String("FleetId"), // Required
			// More values...
		},
		Limit:     aws.Int64(1),
		NextToken: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.DescribeFleetUtilization(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribeGameSessions() {
	svc := gamelift.New(session.New())

	params := &gamelift.DescribeGameSessionsInput{
		AliasId:       aws.String("AliasId"),
		FleetId:       aws.String("FleetId"),
		GameSessionId: aws.String("GameSessionId"),
		Limit:         aws.Int64(1),
		NextToken:     aws.String("NonZeroAndMaxString"),
		StatusFilter:  aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.DescribeGameSessions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_DescribePlayerSessions() {
	svc := gamelift.New(session.New())

	params := &gamelift.DescribePlayerSessionsInput{
		GameSessionId:             aws.String("GameSessionId"),
		Limit:                     aws.Int64(1),
		NextToken:                 aws.String("NonZeroAndMaxString"),
		PlayerId:                  aws.String("NonZeroAndMaxString"),
		PlayerSessionId:           aws.String("PlayerSessionId"),
		PlayerSessionStatusFilter: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.DescribePlayerSessions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_GetGameSessionLogUrl() {
	svc := gamelift.New(session.New())

	params := &gamelift.GetGameSessionLogUrlInput{
		GameSessionId: aws.String("GameSessionId"), // Required
	}
	resp, err := svc.GetGameSessionLogUrl(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_ListAliases() {
	svc := gamelift.New(session.New())

	params := &gamelift.ListAliasesInput{
		Limit:               aws.Int64(1),
		Name:                aws.String("NonEmptyString"),
		NextToken:           aws.String("NonEmptyString"),
		RoutingStrategyType: aws.String("RoutingStrategyType"),
	}
	resp, err := svc.ListAliases(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_ListBuilds() {
	svc := gamelift.New(session.New())

	params := &gamelift.ListBuildsInput{
		Limit:     aws.Int64(1),
		NextToken: aws.String("NonEmptyString"),
		Status:    aws.String("BuildStatus"),
	}
	resp, err := svc.ListBuilds(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_ListFleets() {
	svc := gamelift.New(session.New())

	params := &gamelift.ListFleetsInput{
		BuildId:   aws.String("BuildId"),
		Limit:     aws.Int64(1),
		NextToken: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.ListFleets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_RequestUploadCredentials() {
	svc := gamelift.New(session.New())

	params := &gamelift.RequestUploadCredentialsInput{
		BuildId: aws.String("BuildId"), // Required
	}
	resp, err := svc.RequestUploadCredentials(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_ResolveAlias() {
	svc := gamelift.New(session.New())

	params := &gamelift.ResolveAliasInput{
		AliasId: aws.String("AliasId"), // Required
	}
	resp, err := svc.ResolveAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_UpdateAlias() {
	svc := gamelift.New(session.New())

	params := &gamelift.UpdateAliasInput{
		AliasId:     aws.String("AliasId"), // Required
		Description: aws.String("NonZeroAndMaxString"),
		Name:        aws.String("NonZeroAndMaxString"),
		RoutingStrategy: &gamelift.RoutingStrategy{
			FleetId: aws.String("FleetId"),
			Message: aws.String("FreeText"),
			Type:    aws.String("RoutingStrategyType"),
		},
	}
	resp, err := svc.UpdateAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_UpdateBuild() {
	svc := gamelift.New(session.New())

	params := &gamelift.UpdateBuildInput{
		BuildId: aws.String("BuildId"), // Required
		Name:    aws.String("NonZeroAndMaxString"),
		Version: aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.UpdateBuild(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_UpdateFleetAttributes() {
	svc := gamelift.New(session.New())

	params := &gamelift.UpdateFleetAttributesInput{
		FleetId:     aws.String("FleetId"), // Required
		Description: aws.String("NonZeroAndMaxString"),
		Name:        aws.String("NonZeroAndMaxString"),
	}
	resp, err := svc.UpdateFleetAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_UpdateFleetCapacity() {
	svc := gamelift.New(session.New())

	params := &gamelift.UpdateFleetCapacityInput{
		DesiredInstances: aws.Int64(1),          // Required
		FleetId:          aws.String("FleetId"), // Required
	}
	resp, err := svc.UpdateFleetCapacity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_UpdateFleetPortSettings() {
	svc := gamelift.New(session.New())

	params := &gamelift.UpdateFleetPortSettingsInput{
		FleetId: aws.String("FleetId"), // Required
		InboundPermissionAuthorizations: []*gamelift.IpPermission{
			{ // Required
				FromPort: aws.Int64(1),                 // Required
				IpRange:  aws.String("NonBlankString"), // Required
				Protocol: aws.String("IpProtocol"),     // Required
				ToPort:   aws.Int64(1),                 // Required
			},
			// More values...
		},
		InboundPermissionRevocations: []*gamelift.IpPermission{
			{ // Required
				FromPort: aws.Int64(1),                 // Required
				IpRange:  aws.String("NonBlankString"), // Required
				Protocol: aws.String("IpProtocol"),     // Required
				ToPort:   aws.Int64(1),                 // Required
			},
			// More values...
		},
	}
	resp, err := svc.UpdateFleetPortSettings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleGameLift_UpdateGameSession() {
	svc := gamelift.New(session.New())

	params := &gamelift.UpdateGameSessionInput{
		GameSessionId:             aws.String("GameSessionId"), // Required
		MaximumPlayerSessionCount: aws.Int64(1),
		Name: aws.String("NonZeroAndMaxString"),
		PlayerSessionCreationPolicy: aws.String("PlayerSessionCreationPolicy"),
	}
	resp, err := svc.UpdateGameSession(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
