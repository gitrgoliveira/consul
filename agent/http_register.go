package agent

func init() {
	registerEndpoint("/v1/acl/bootstrap", []string{"PUT"}, (*HTTPHandlers).ACLBootstrap)
	registerEndpoint("/v1/acl/login", []string{"POST"}, (*HTTPHandlers).ACLLogin)
	registerEndpoint("/v1/acl/logout", []string{"POST"}, (*HTTPHandlers).ACLLogout)
	registerEndpoint("/v1/acl/replication", []string{"GET"}, (*HTTPHandlers).ACLReplicationStatus)
	registerEndpoint("/v1/acl/policies", []string{"GET"}, (*HTTPHandlers).ACLPolicyList)
	registerEndpoint("/v1/acl/policy", []string{"PUT"}, (*HTTPHandlers).ACLPolicyCreate)
	registerEndpoint("/v1/acl/policy/", []string{"GET", "PUT", "DELETE"}, (*HTTPHandlers).ACLPolicyCRUD)
	registerEndpoint("/v1/acl/policy/name/", []string{"GET"}, (*HTTPHandlers).ACLPolicyReadByName)
	registerEndpoint("/v1/acl/roles", []string{"GET"}, (*HTTPHandlers).ACLRoleList)
	registerEndpoint("/v1/acl/role", []string{"PUT"}, (*HTTPHandlers).ACLRoleCreate)
	registerEndpoint("/v1/acl/role/name/", []string{"GET"}, (*HTTPHandlers).ACLRoleReadByName)
	registerEndpoint("/v1/acl/role/", []string{"GET", "PUT", "DELETE"}, (*HTTPHandlers).ACLRoleCRUD)
	registerEndpoint("/v1/acl/binding-rules", []string{"GET"}, (*HTTPHandlers).ACLBindingRuleList)
	registerEndpoint("/v1/acl/binding-rule", []string{"PUT"}, (*HTTPHandlers).ACLBindingRuleCreate)
	registerEndpoint("/v1/acl/binding-rule/", []string{"GET", "PUT", "DELETE"}, (*HTTPHandlers).ACLBindingRuleCRUD)
	registerEndpoint("/v1/acl/auth-methods", []string{"GET"}, (*HTTPHandlers).ACLAuthMethodList)
	registerEndpoint("/v1/acl/auth-method", []string{"PUT"}, (*HTTPHandlers).ACLAuthMethodCreate)
	registerEndpoint("/v1/acl/auth-method/", []string{"GET", "PUT", "DELETE"}, (*HTTPHandlers).ACLAuthMethodCRUD)
	registerEndpoint("/v1/acl/tokens", []string{"GET"}, (*HTTPHandlers).ACLTokenList)
	registerEndpoint("/v1/acl/token", []string{"PUT"}, (*HTTPHandlers).ACLTokenCreate)
	registerEndpoint("/v1/acl/token/self", []string{"GET"}, (*HTTPHandlers).ACLTokenSelf)
	registerEndpoint("/v1/acl/token/", []string{"GET", "PUT", "DELETE"}, (*HTTPHandlers).ACLTokenCRUD)
	registerEndpoint("/v1/agent/token/", []string{"PUT"}, (*HTTPHandlers).AgentToken)
	registerEndpoint("/v1/agent/self", []string{"GET"}, (*HTTPHandlers).AgentSelf)
	registerEndpoint("/v1/agent/host", []string{"GET"}, (*HTTPHandlers).AgentHost)
	registerEndpoint("/v1/agent/maintenance", []string{"PUT"}, (*HTTPHandlers).AgentNodeMaintenance)
	registerEndpoint("/v1/agent/reload", []string{"PUT"}, (*HTTPHandlers).AgentReload)
	registerEndpoint("/v1/agent/monitor", []string{"GET"}, (*HTTPHandlers).AgentMonitor)
	registerEndpoint("/v1/agent/metrics", []string{"GET"}, (*HTTPHandlers).AgentMetrics)
	registerEndpoint("/v1/agent/metrics/stream", []string{"GET"}, (*HTTPHandlers).AgentMetricsStream)
	registerEndpoint("/v1/agent/services", []string{"GET"}, (*HTTPHandlers).AgentServices)
	registerEndpoint("/v1/agent/service/", []string{"GET"}, (*HTTPHandlers).AgentService)
	registerEndpoint("/v1/agent/checks", []string{"GET"}, (*HTTPHandlers).AgentChecks)
	registerEndpoint("/v1/agent/members", []string{"GET"}, (*HTTPHandlers).AgentMembers)
	registerEndpoint("/v1/agent/join/", []string{"PUT"}, (*HTTPHandlers).AgentJoin)
	registerEndpoint("/v1/agent/leave", []string{"PUT"}, (*HTTPHandlers).AgentLeave)
	registerEndpoint("/v1/agent/force-leave/", []string{"PUT"}, (*HTTPHandlers).AgentForceLeave)
	registerEndpoint("/v1/agent/health/service/id/", []string{"GET"}, (*HTTPHandlers).AgentHealthServiceByID)
	registerEndpoint("/v1/agent/health/service/name/", []string{"GET"}, (*HTTPHandlers).AgentHealthServiceByName)
	registerEndpoint("/v1/agent/check/register", []string{"PUT"}, (*HTTPHandlers).AgentRegisterCheck)
	registerEndpoint("/v1/agent/check/deregister/", []string{"PUT"}, (*HTTPHandlers).AgentDeregisterCheck)
	registerEndpoint("/v1/agent/check/pass/", []string{"PUT"}, (*HTTPHandlers).AgentCheckPass)
	registerEndpoint("/v1/agent/check/warn/", []string{"PUT"}, (*HTTPHandlers).AgentCheckWarn)
	registerEndpoint("/v1/agent/check/fail/", []string{"PUT"}, (*HTTPHandlers).AgentCheckFail)
	registerEndpoint("/v1/agent/check/update/", []string{"PUT"}, (*HTTPHandlers).AgentCheckUpdate)
	registerEndpoint("/v1/agent/connect/authorize", []string{"POST"}, (*HTTPHandlers).AgentConnectAuthorize)
	registerEndpoint("/v1/agent/connect/ca/roots", []string{"GET"}, (*HTTPHandlers).AgentConnectCARoots)
	registerEndpoint("/v1/agent/connect/ca/leaf/", []string{"GET"}, (*HTTPHandlers).AgentConnectCALeafCert)
	registerEndpoint("/v1/agent/service/register", []string{"PUT"}, (*HTTPHandlers).AgentRegisterService)
	registerEndpoint("/v1/agent/service/deregister/", []string{"PUT"}, (*HTTPHandlers).AgentDeregisterService)
	registerEndpoint("/v1/agent/service/maintenance/", []string{"PUT"}, (*HTTPHandlers).AgentServiceMaintenance)
	registerEndpoint("/v1/catalog/register", []string{"PUT"}, (*HTTPHandlers).CatalogRegister)
	registerEndpoint("/v1/catalog/connect/", []string{"GET"}, (*HTTPHandlers).CatalogConnectServiceNodes)
	registerEndpoint("/v1/catalog/deregister", []string{"PUT"}, (*HTTPHandlers).CatalogDeregister)
	registerEndpoint("/v1/catalog/datacenters", []string{"GET"}, (*HTTPHandlers).CatalogDatacenters)
	registerEndpoint("/v1/catalog/nodes", []string{"GET"}, (*HTTPHandlers).CatalogNodes)
	registerEndpoint("/v1/catalog/services", []string{"GET"}, (*HTTPHandlers).CatalogServices)
	registerEndpoint("/v1/catalog/service/", []string{"GET"}, (*HTTPHandlers).CatalogServiceNodes)
	registerEndpoint("/v1/catalog/node/", []string{"GET"}, (*HTTPHandlers).CatalogNodeServices)
	registerEndpoint("/v1/catalog/node-services/", []string{"GET"}, (*HTTPHandlers).CatalogNodeServiceList)
	registerEndpoint("/v1/catalog/gateway-services/", []string{"GET"}, (*HTTPHandlers).CatalogGatewayServices)
	registerEndpoint("/v1/config/", []string{"GET", "DELETE"}, (*HTTPHandlers).Config)
	registerEndpoint("/v1/config", []string{"PUT"}, (*HTTPHandlers).ConfigApply)
	registerEndpoint("/v1/connect/ca/configuration", []string{"GET", "PUT"}, (*HTTPHandlers).ConnectCAConfiguration)
	registerEndpoint("/v1/connect/ca/roots", []string{"GET"}, (*HTTPHandlers).ConnectCARoots)
	registerEndpoint("/v1/connect/intentions", []string{"GET", "POST"}, (*HTTPHandlers).IntentionEndpoint) // POST is deprecated
	registerEndpoint("/v1/connect/intentions/match", []string{"GET"}, (*HTTPHandlers).IntentionMatch)
	registerEndpoint("/v1/connect/intentions/check", []string{"GET"}, (*HTTPHandlers).IntentionCheck)
	registerEndpoint("/v1/connect/intentions/exact", []string{"GET", "PUT", "DELETE"}, (*HTTPHandlers).IntentionExact)
	registerEndpoint("/v1/connect/intentions/", []string{"GET", "PUT", "DELETE"}, (*HTTPHandlers).IntentionSpecific) // deprecated
	registerEndpoint("/v1/coordinate/datacenters", []string{"GET"}, (*HTTPHandlers).CoordinateDatacenters)
	registerEndpoint("/v1/coordinate/nodes", []string{"GET"}, (*HTTPHandlers).CoordinateNodes)
	registerEndpoint("/v1/coordinate/node/", []string{"GET"}, (*HTTPHandlers).CoordinateNode)
	registerEndpoint("/v1/coordinate/update", []string{"PUT"}, (*HTTPHandlers).CoordinateUpdate)
	registerEndpoint("/v1/internal/federation-states", []string{"GET"}, (*HTTPHandlers).FederationStateList)
	registerEndpoint("/v1/internal/federation-states/mesh-gateways", []string{"GET"}, (*HTTPHandlers).FederationStateListMeshGateways)
	registerEndpoint("/v1/internal/federation-state/", []string{"GET"}, (*HTTPHandlers).FederationStateGet)
	registerEndpoint("/v1/discovery-chain/", []string{"GET", "POST"}, (*HTTPHandlers).DiscoveryChainRead)
	registerEndpoint("/v1/event/fire/", []string{"PUT"}, (*HTTPHandlers).EventFire)
	registerEndpoint("/v1/event/list", []string{"GET"}, (*HTTPHandlers).EventList)
	registerEndpoint("/v1/health/node/", []string{"GET"}, (*HTTPHandlers).HealthNodeChecks)
	registerEndpoint("/v1/health/checks/", []string{"GET"}, (*HTTPHandlers).HealthServiceChecks)
	registerEndpoint("/v1/health/state/", []string{"GET"}, (*HTTPHandlers).HealthChecksInState)
	registerEndpoint("/v1/health/service/", []string{"GET"}, (*HTTPHandlers).HealthServiceNodes)
	registerEndpoint("/v1/health/connect/", []string{"GET"}, (*HTTPHandlers).HealthConnectServiceNodes)
	registerEndpoint("/v1/health/ingress/", []string{"GET"}, (*HTTPHandlers).HealthIngressServiceNodes)
	registerEndpoint("/v1/internal/ui/metrics-proxy/", []string{"GET"}, (*HTTPHandlers).UIMetricsProxy)
	registerEndpoint("/v1/internal/ui/nodes", []string{"GET"}, (*HTTPHandlers).UINodes)
	registerEndpoint("/v1/internal/ui/node/", []string{"GET"}, (*HTTPHandlers).UINodeInfo)
	registerEndpoint("/v1/internal/ui/services", []string{"GET"}, (*HTTPHandlers).UIServices)
	registerEndpoint("/v1/internal/ui/gateway-services-nodes/", []string{"GET"}, (*HTTPHandlers).UIGatewayServicesNodes)
	registerEndpoint("/v1/internal/ui/gateway-intentions/", []string{"GET"}, (*HTTPHandlers).UIGatewayIntentions)
	registerEndpoint("/v1/internal/ui/service-topology/", []string{"GET"}, (*HTTPHandlers).UIServiceTopology)
	registerEndpoint("/v1/internal/acl/authorize", []string{"POST"}, (*HTTPHandlers).ACLAuthorize)
	registerEndpoint("/v1/kv/", []string{"GET", "PUT", "DELETE"}, (*HTTPHandlers).KVSEndpoint)
	registerEndpoint("/v1/operator/raft/configuration", []string{"GET"}, (*HTTPHandlers).OperatorRaftConfiguration)
	registerEndpoint("/v1/operator/raft/leader-transfer", []string{"POST"}, (*HTTPHandlers).OperatorRaftLeaderTransfer)
	registerEndpoint("/v1/operator/raft/peer", []string{"DELETE"}, (*HTTPHandlers).OperatorRaftPeer)
	registerEndpoint("/v1/operator/keyring", []string{"GET", "POST", "PUT", "DELETE"}, (*HTTPHandlers).OperatorKeyringEndpoint)
	registerEndpoint("/v1/operator/autopilot/configuration", []string{"GET", "PUT"}, (*HTTPHandlers).OperatorAutopilotConfiguration)
	registerEndpoint("/v1/operator/autopilot/health", []string{"GET"}, (*HTTPHandlers).OperatorServerHealth)
	registerEndpoint("/v1/operator/autopilot/state", []string{"GET"}, (*HTTPHandlers).OperatorAutopilotState)
	registerEndpoint("/v1/query", []string{"GET", "POST"}, (*HTTPHandlers).PreparedQueryGeneral)
	// specific prepared query endpoints have more complex rules for allowed methods, so
	// the prefix is registered with no methods.
	registerEndpoint("/v1/query/", []string{}, (*HTTPHandlers).PreparedQuerySpecific)
	registerEndpoint("/v1/session/create", []string{"PUT"}, (*HTTPHandlers).SessionCreate)
	registerEndpoint("/v1/session/destroy/", []string{"PUT"}, (*HTTPHandlers).SessionDestroy)
	registerEndpoint("/v1/session/renew/", []string{"PUT"}, (*HTTPHandlers).SessionRenew)
	registerEndpoint("/v1/session/info/", []string{"GET"}, (*HTTPHandlers).SessionGet)
	registerEndpoint("/v1/session/node/", []string{"GET"}, (*HTTPHandlers).SessionsForNode)
	registerEndpoint("/v1/session/list", []string{"GET"}, (*HTTPHandlers).SessionList)
	registerEndpoint("/v1/status/leader", []string{"GET"}, (*HTTPHandlers).StatusLeader)
	registerEndpoint("/v1/status/peers", []string{"GET"}, (*HTTPHandlers).StatusPeers)
	registerEndpoint("/v1/snapshot", []string{"GET", "PUT"}, (*HTTPHandlers).Snapshot)
	registerEndpoint("/v1/txn", []string{"PUT"}, (*HTTPHandlers).Txn)

	// Deprecated ACL endpoints, they do nothing but return an error
	registerEndpoint("/v1/acl/create", []string{"PUT"}, (*HTTPHandlers).ACLLegacy)
	registerEndpoint("/v1/acl/update", []string{"PUT"}, (*HTTPHandlers).ACLLegacy)
	registerEndpoint("/v1/acl/destroy/", []string{"PUT"}, (*HTTPHandlers).ACLLegacy)
	registerEndpoint("/v1/acl/info/", []string{"GET"}, (*HTTPHandlers).ACLLegacy)
	registerEndpoint("/v1/acl/clone/", []string{"PUT"}, (*HTTPHandlers).ACLLegacy)
	registerEndpoint("/v1/acl/list", []string{"GET"}, (*HTTPHandlers).ACLLegacy)
	registerEndpoint("/v1/acl/rules/translate", []string{"POST"}, (*HTTPHandlers).ACLLegacy)
	registerEndpoint("/v1/acl/rules/translate/", []string{"GET"}, (*HTTPHandlers).ACLLegacy)
}
