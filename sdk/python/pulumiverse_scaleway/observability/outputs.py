# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = [
    'AlertManagerContactPoint',
    'CockpitEndpoint',
    'CockpitPushUrl',
    'TokenScopes',
    'GetInstanceEndpointResult',
    'GetInstancePushUrlResult',
]

@pulumi.output_type
class AlertManagerContactPoint(dict):
    def __init__(__self__, *,
                 email: Optional[builtins.str] = None):
        """
        :param builtins.str email: Email addresses for the alert receivers
        """
        if email is not None:
            pulumi.set(__self__, "email", email)

    @property
    @pulumi.getter
    def email(self) -> Optional[builtins.str]:
        """
        Email addresses for the alert receivers
        """
        return pulumi.get(self, "email")


@pulumi.output_type
class CockpitEndpoint(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "alertmanagerUrl":
            suggest = "alertmanager_url"
        elif key == "grafanaUrl":
            suggest = "grafana_url"
        elif key == "logsUrl":
            suggest = "logs_url"
        elif key == "metricsUrl":
            suggest = "metrics_url"
        elif key == "tracesUrl":
            suggest = "traces_url"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CockpitEndpoint. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CockpitEndpoint.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CockpitEndpoint.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 alertmanager_url: Optional[builtins.str] = None,
                 grafana_url: Optional[builtins.str] = None,
                 logs_url: Optional[builtins.str] = None,
                 metrics_url: Optional[builtins.str] = None,
                 traces_url: Optional[builtins.str] = None):
        """
        :param builtins.str alertmanager_url: (Deprecated) URL for the [Alert manager](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#alert-manager).
        :param builtins.str grafana_url: (Deprecated) URL for Grafana.
        :param builtins.str logs_url: (Deprecated) URL for [logs](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#logs) to retrieve in the [Data sources tab](https://console.scaleway.com/cockpit/dataSource) of the Scaleway console.
        :param builtins.str metrics_url: (Deprecated) URL for [metrics](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#metric) to retrieve in the [Data sources tab](https://console.scaleway.com/cockpit/dataSource) of the Scaleway console.
        :param builtins.str traces_url: (Deprecated) URL for [traces](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#traces) to retrieve in the [Data sources tab](https://console.scaleway.com/cockpit/dataSource) of the Scaleway console.
        """
        if alertmanager_url is not None:
            pulumi.set(__self__, "alertmanager_url", alertmanager_url)
        if grafana_url is not None:
            pulumi.set(__self__, "grafana_url", grafana_url)
        if logs_url is not None:
            pulumi.set(__self__, "logs_url", logs_url)
        if metrics_url is not None:
            pulumi.set(__self__, "metrics_url", metrics_url)
        if traces_url is not None:
            pulumi.set(__self__, "traces_url", traces_url)

    @property
    @pulumi.getter(name="alertmanagerUrl")
    def alertmanager_url(self) -> Optional[builtins.str]:
        """
        (Deprecated) URL for the [Alert manager](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#alert-manager).
        """
        return pulumi.get(self, "alertmanager_url")

    @property
    @pulumi.getter(name="grafanaUrl")
    def grafana_url(self) -> Optional[builtins.str]:
        """
        (Deprecated) URL for Grafana.
        """
        return pulumi.get(self, "grafana_url")

    @property
    @pulumi.getter(name="logsUrl")
    def logs_url(self) -> Optional[builtins.str]:
        """
        (Deprecated) URL for [logs](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#logs) to retrieve in the [Data sources tab](https://console.scaleway.com/cockpit/dataSource) of the Scaleway console.
        """
        return pulumi.get(self, "logs_url")

    @property
    @pulumi.getter(name="metricsUrl")
    def metrics_url(self) -> Optional[builtins.str]:
        """
        (Deprecated) URL for [metrics](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#metric) to retrieve in the [Data sources tab](https://console.scaleway.com/cockpit/dataSource) of the Scaleway console.
        """
        return pulumi.get(self, "metrics_url")

    @property
    @pulumi.getter(name="tracesUrl")
    def traces_url(self) -> Optional[builtins.str]:
        """
        (Deprecated) URL for [traces](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#traces) to retrieve in the [Data sources tab](https://console.scaleway.com/cockpit/dataSource) of the Scaleway console.
        """
        return pulumi.get(self, "traces_url")


@pulumi.output_type
class CockpitPushUrl(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "pushLogsUrl":
            suggest = "push_logs_url"
        elif key == "pushMetricsUrl":
            suggest = "push_metrics_url"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CockpitPushUrl. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CockpitPushUrl.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CockpitPushUrl.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 push_logs_url: Optional[builtins.str] = None,
                 push_metrics_url: Optional[builtins.str] = None):
        """
        :param builtins.str push_logs_url: Push URL for logs (Grafana Loki)
        :param builtins.str push_metrics_url: Push URL for metrics (Grafana Mimir)
        """
        if push_logs_url is not None:
            pulumi.set(__self__, "push_logs_url", push_logs_url)
        if push_metrics_url is not None:
            pulumi.set(__self__, "push_metrics_url", push_metrics_url)

    @property
    @pulumi.getter(name="pushLogsUrl")
    def push_logs_url(self) -> Optional[builtins.str]:
        """
        Push URL for logs (Grafana Loki)
        """
        return pulumi.get(self, "push_logs_url")

    @property
    @pulumi.getter(name="pushMetricsUrl")
    def push_metrics_url(self) -> Optional[builtins.str]:
        """
        Push URL for metrics (Grafana Mimir)
        """
        return pulumi.get(self, "push_metrics_url")


@pulumi.output_type
class TokenScopes(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "queryLogs":
            suggest = "query_logs"
        elif key == "queryMetrics":
            suggest = "query_metrics"
        elif key == "queryTraces":
            suggest = "query_traces"
        elif key == "setupAlerts":
            suggest = "setup_alerts"
        elif key == "setupLogsRules":
            suggest = "setup_logs_rules"
        elif key == "setupMetricsRules":
            suggest = "setup_metrics_rules"
        elif key == "writeLogs":
            suggest = "write_logs"
        elif key == "writeMetrics":
            suggest = "write_metrics"
        elif key == "writeTraces":
            suggest = "write_traces"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TokenScopes. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TokenScopes.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TokenScopes.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 query_logs: Optional[builtins.bool] = None,
                 query_metrics: Optional[builtins.bool] = None,
                 query_traces: Optional[builtins.bool] = None,
                 setup_alerts: Optional[builtins.bool] = None,
                 setup_logs_rules: Optional[builtins.bool] = None,
                 setup_metrics_rules: Optional[builtins.bool] = None,
                 write_logs: Optional[builtins.bool] = None,
                 write_metrics: Optional[builtins.bool] = None,
                 write_traces: Optional[builtins.bool] = None):
        """
        :param builtins.bool query_logs: Permission to query logs.
        :param builtins.bool query_metrics: Permission to query metrics.
        :param builtins.bool query_traces: Permission to query traces.
        :param builtins.bool setup_alerts: Permission to set up alerts.
        :param builtins.bool setup_logs_rules: Permission to set up logs rules.
        :param builtins.bool setup_metrics_rules: Permission to set up metrics rules.
        :param builtins.bool write_logs: Permission to write logs.
        :param builtins.bool write_metrics: Permission to write metrics.
        :param builtins.bool write_traces: Permission to write traces.
        """
        if query_logs is not None:
            pulumi.set(__self__, "query_logs", query_logs)
        if query_metrics is not None:
            pulumi.set(__self__, "query_metrics", query_metrics)
        if query_traces is not None:
            pulumi.set(__self__, "query_traces", query_traces)
        if setup_alerts is not None:
            pulumi.set(__self__, "setup_alerts", setup_alerts)
        if setup_logs_rules is not None:
            pulumi.set(__self__, "setup_logs_rules", setup_logs_rules)
        if setup_metrics_rules is not None:
            pulumi.set(__self__, "setup_metrics_rules", setup_metrics_rules)
        if write_logs is not None:
            pulumi.set(__self__, "write_logs", write_logs)
        if write_metrics is not None:
            pulumi.set(__self__, "write_metrics", write_metrics)
        if write_traces is not None:
            pulumi.set(__self__, "write_traces", write_traces)

    @property
    @pulumi.getter(name="queryLogs")
    def query_logs(self) -> Optional[builtins.bool]:
        """
        Permission to query logs.
        """
        return pulumi.get(self, "query_logs")

    @property
    @pulumi.getter(name="queryMetrics")
    def query_metrics(self) -> Optional[builtins.bool]:
        """
        Permission to query metrics.
        """
        return pulumi.get(self, "query_metrics")

    @property
    @pulumi.getter(name="queryTraces")
    def query_traces(self) -> Optional[builtins.bool]:
        """
        Permission to query traces.
        """
        return pulumi.get(self, "query_traces")

    @property
    @pulumi.getter(name="setupAlerts")
    def setup_alerts(self) -> Optional[builtins.bool]:
        """
        Permission to set up alerts.
        """
        return pulumi.get(self, "setup_alerts")

    @property
    @pulumi.getter(name="setupLogsRules")
    def setup_logs_rules(self) -> Optional[builtins.bool]:
        """
        Permission to set up logs rules.
        """
        return pulumi.get(self, "setup_logs_rules")

    @property
    @pulumi.getter(name="setupMetricsRules")
    def setup_metrics_rules(self) -> Optional[builtins.bool]:
        """
        Permission to set up metrics rules.
        """
        return pulumi.get(self, "setup_metrics_rules")

    @property
    @pulumi.getter(name="writeLogs")
    def write_logs(self) -> Optional[builtins.bool]:
        """
        Permission to write logs.
        """
        return pulumi.get(self, "write_logs")

    @property
    @pulumi.getter(name="writeMetrics")
    def write_metrics(self) -> Optional[builtins.bool]:
        """
        Permission to write metrics.
        """
        return pulumi.get(self, "write_metrics")

    @property
    @pulumi.getter(name="writeTraces")
    def write_traces(self) -> Optional[builtins.bool]:
        """
        Permission to write traces.
        """
        return pulumi.get(self, "write_traces")


@pulumi.output_type
class GetInstanceEndpointResult(dict):
    def __init__(__self__, *,
                 alertmanager_url: builtins.str,
                 grafana_url: builtins.str,
                 logs_url: builtins.str,
                 metrics_url: builtins.str,
                 traces_url: builtins.str):
        """
        :param builtins.str alertmanager_url: (Deprecated) URL for the [Alert manager](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#alert-manager).
        :param builtins.str grafana_url: (Deprecated) URL for Grafana.
        :param builtins.str logs_url: (Deprecated) URL for [logs](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#logs) to retrieve in the [Data sources tab](https://console.scaleway.com/cockpit/dataSource) of the Scaleway console.
        :param builtins.str metrics_url: (Deprecated) URL for [metrics](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#metric) to retrieve in the [Data sources tab](https://console.scaleway.com/cockpit/dataSource) of the Scaleway console.
        :param builtins.str traces_url: The traces URL.
        """
        pulumi.set(__self__, "alertmanager_url", alertmanager_url)
        pulumi.set(__self__, "grafana_url", grafana_url)
        pulumi.set(__self__, "logs_url", logs_url)
        pulumi.set(__self__, "metrics_url", metrics_url)
        pulumi.set(__self__, "traces_url", traces_url)

    @property
    @pulumi.getter(name="alertmanagerUrl")
    def alertmanager_url(self) -> builtins.str:
        """
        (Deprecated) URL for the [Alert manager](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#alert-manager).
        """
        return pulumi.get(self, "alertmanager_url")

    @property
    @pulumi.getter(name="grafanaUrl")
    def grafana_url(self) -> builtins.str:
        """
        (Deprecated) URL for Grafana.
        """
        return pulumi.get(self, "grafana_url")

    @property
    @pulumi.getter(name="logsUrl")
    def logs_url(self) -> builtins.str:
        """
        (Deprecated) URL for [logs](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#logs) to retrieve in the [Data sources tab](https://console.scaleway.com/cockpit/dataSource) of the Scaleway console.
        """
        return pulumi.get(self, "logs_url")

    @property
    @pulumi.getter(name="metricsUrl")
    def metrics_url(self) -> builtins.str:
        """
        (Deprecated) URL for [metrics](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#metric) to retrieve in the [Data sources tab](https://console.scaleway.com/cockpit/dataSource) of the Scaleway console.
        """
        return pulumi.get(self, "metrics_url")

    @property
    @pulumi.getter(name="tracesUrl")
    def traces_url(self) -> builtins.str:
        """
        The traces URL.
        """
        return pulumi.get(self, "traces_url")


@pulumi.output_type
class GetInstancePushUrlResult(dict):
    def __init__(__self__, *,
                 push_logs_url: builtins.str,
                 push_metrics_url: builtins.str):
        """
        :param builtins.str push_logs_url: Push URL for logs (Grafana Loki)
        :param builtins.str push_metrics_url: Push URL for metrics (Grafana Mimir)
        """
        pulumi.set(__self__, "push_logs_url", push_logs_url)
        pulumi.set(__self__, "push_metrics_url", push_metrics_url)

    @property
    @pulumi.getter(name="pushLogsUrl")
    def push_logs_url(self) -> builtins.str:
        """
        Push URL for logs (Grafana Loki)
        """
        return pulumi.get(self, "push_logs_url")

    @property
    @pulumi.getter(name="pushMetricsUrl")
    def push_metrics_url(self) -> builtins.str:
        """
        Push URL for metrics (Grafana Mimir)
        """
        return pulumi.get(self, "push_metrics_url")


