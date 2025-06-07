// Copyright Jetstack Ltd. See LICENSE for details.
package cases

import (
	_ "github.com/TremoloSecurity/kube-oidc-proxy/test/e2e/suite/cases/audit"
	_ "github.com/TremoloSecurity/kube-oidc-proxy/test/e2e/suite/cases/headers"
	_ "github.com/TremoloSecurity/kube-oidc-proxy/test/e2e/suite/cases/impersonation"
	_ "github.com/TremoloSecurity/kube-oidc-proxy/test/e2e/suite/cases/passthrough"
	_ "github.com/TremoloSecurity/kube-oidc-proxy/test/e2e/suite/cases/probe"
	_ "github.com/TremoloSecurity/kube-oidc-proxy/test/e2e/suite/cases/rbac"
	_ "github.com/TremoloSecurity/kube-oidc-proxy/test/e2e/suite/cases/token"
	_ "github.com/TremoloSecurity/kube-oidc-proxy/test/e2e/suite/cases/upgrade"
)
