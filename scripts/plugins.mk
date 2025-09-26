# Install Pulumi and plugins required at build time.
install_plugins: .make/install_plugins
.make/install_plugins: export PULUMI_HOME := $(WORKING_DIR)/.pulumi
.make/install_plugins:
	pulumi plugin install converter terraform 1.0.19
	pulumi plugin install resource random 4.16.7
	pulumi plugin install resource aws 6.56.1
	pulumi plugin install resource gitlab 8.5.0
	pulumi plugin install resource local 0.1.5
	pulumi plugin install resource null 0.0.8
	@touch $@
.PHONY: install_plugins
