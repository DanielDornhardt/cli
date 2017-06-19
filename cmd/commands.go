package cmd

import (
	"github.com/Scalingo/codegangsta-cli"
)

var (
	Commands = []cli.Command{
		// Apps
		AppsCommand,
		CreateCommand,
		DestroyCommand,

		// Apps Actions
		LogsCommand,
		RunCommand,

		// Apps Process Actions
		PsCommand,
		ScaleCommand,
		RestartCommand,

		// Events
		UserTimelineCommand,
		TimelineCommand,

		// Environment
		EnvCommand,
		EnvSetCommand,
		EnvUnsetCommand,

		// Domains
		DomainsListCommand,
		DomainsAddCommand,
		DomainsRemoveCommand,
		DomainsSSLCommand,

		// Deployments
		DeploymentsListCommand,
		DeploymentLogCommand,
		DeploymentFollowCommand,
		DeploymentDeployCommand,

		// Collaborators
		CollaboratorsListCommand,
		CollaboratorsAddCommand,
		CollaboratorsRemoveCommand,

		// Addons
		AddonProvidersListCommand,
		AddonProvidersPlansCommand,
		AddonsListCommand,
		AddonsAddCommand,
		AddonsRemoveCommand,
		AddonsUpgradeCommand,

		// Notifications
		NotificationsListCommand,
		NotificationsAddCommand,
		NotificationsUpdateCommand,
		NotificationsRemoveCommand,

		// Notifiers
		NotifiersListCommand,

		// DB Access
		DbTunnelCommand,
		RedisConsoleCommand,
		MongoConsoleCommand,
		MySQLConsoleCommand,
		PgSQLConsoleCommand,
		InfluxDBConsoleCommand,

		// Stats
		StatsCommand,

		// SSH keys
		ListSSHKeyCommand,
		AddSSHKeyCommand,
		RemoveSSHKeyCommand,

		// Sessions
		LoginCommand,
		LogoutCommand,
		SignUpCommand,

		// Version
		UpdateCommand,

		// Help
		HelpCommand,
	}
)
