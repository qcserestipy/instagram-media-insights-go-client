package api

import "github.com/goharbor/go-client/pkg/sdk/v2.0/client/robot"

func GetRobot(robotID int64) (*robot.GetRobotByIDOK, error) {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}
	response, err := client.Robot.GetRobotByID(ctx, &robot.GetRobotByIDParams{RobotID: robotID})
	if err != nil {
		return nil, err
	}

	return response, nil
}
